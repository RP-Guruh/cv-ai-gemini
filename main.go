package main

import (
	"bytes"
	"fmt"
	"goweb-cv-ai/internal/ai"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/phpdave11/gofpdf"
)

// --- STRUCT & FUNGSI PEMBANTU UNTUK PARSING CV ---

// StructuredEntryParts: Struct untuk menampung hasil parsing entri terstruktur
type StructuredEntryParts struct {
	Title    string // Judul (misalnya, Asisten Apoteker)
	Location string // Lokasi atau Perusahaan
	Date     string // Rentang waktu (misalnya, Oktober 2022 - Sekarang)
}

// isSection: Mendeteksi apakah baris adalah judul bagian utama (misalnya PENDIDIKAN)
func isSection(line string) bool {
	// Memastikan semua karakter adalah huruf besar dan memiliki lebih dari 4 karakter
	if len(line) < 5 {
		return false
	}
	// Asumsi: Judul bagian biasanya terdiri dari 1-3 kata (e.g., PENGALAMAN KERJA, KEMAMPUAN)
	return strings.ToUpper(line) == line && strings.Count(line, " ") <= 3
}

// isStructuredEntry: Mencoba mendeteksi baris yang berisi judul pekerjaan/pendidikan yang terstruktur
func isStructuredEntry(line string) bool {
	// Kriteria: mengandung range tahun/bulan (implied structure)
	// Kita cari pola yang mengandung '-' dan salah satu indikator waktu
	isDateRange := strings.Contains(line, "-") && (strings.Contains(line, "20") || strings.Contains(line, "Sekarang") || strings.Contains(line, "Januari"))

	// Cek apakah bukan bullet point dan bukan section title
	isNotBullet := !strings.HasPrefix(line, "•") && !strings.HasPrefix(line, "- ")
	isNotSection := !isSection(line)

	return isNotBullet && isNotSection && isDateRange
}

// splitEntry: Membagi baris entri terstruktur menjadi Role, Location, dan Date
func splitEntry(line string) StructuredEntryParts {
	parts := StructuredEntryParts{
		Title:    line,
		Location: "",
		Date:     "",
	}

	// 1. Mencari Tanggal di Akhir
	dateEndIndex := -1
	words := strings.Fields(line)
	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]
		// Mencari kata yang terlihat seperti tahun (20xx) atau "Sekarang"
		if len(word) >= 4 && (strings.HasPrefix(word, "20") || strings.Contains(word, "Sekarang")) {
			dateEndIndex = i
			break
		}
	}

	if dateEndIndex != -1 {
		dateStartIndex := dateEndIndex
		// Cari pemisah range (e.g. '-') atau awal bulan untuk menemukan awal range
		for i := dateEndIndex - 1; i >= 0; i-- {
			// Jika menemukan dash ('-') atau kata yang terlihat seperti nama bulan
			monthNames := "Januari Februari Maret April Mei Juni Juli Agustus September Oktober November Desember"
			isMonth := len(words[i]) >= 3 && strings.Contains(monthNames, words[i])

			if strings.Contains(words[i], "-") || isMonth {
				dateStartIndex = i
				break
			}
		}

		parts.Date = strings.Join(words[dateStartIndex:], " ")
		remainingLine := strings.Join(words[:dateStartIndex], " ")

		// 2. Mencari Location/Company dari sisa baris
		// Mencari pemisah ' - ' yang memisahkan Location/Company dari Title
		if idx := strings.LastIndex(remainingLine, " - "); idx != -1 {
			// Asumsi: Location di kanan ' - ' terakhir, Title di kiri
			parts.Location = strings.TrimSpace(remainingLine[idx+3:])
			parts.Title = strings.TrimSpace(remainingLine[:idx])
		} else if idx := strings.LastIndex(remainingLine, ","); idx != -1 {
			// Jika hanya ada koma
			parts.Location = strings.TrimSpace(remainingLine[idx+1:])
			parts.Title = strings.TrimSpace(remainingLine[:idx])
		} else {
			// Tidak ada pemisah, biarkan Location kosong
			parts.Title = remainingLine
			parts.Location = ""
		}
	} else {
		// Jika tidak ada tanggal yang terdeteksi
		parts.Title = line
		parts.Location = ""
		parts.Date = ""
	}

	return parts
}

// --- FUNGSI UTAMA SERVER (MAIN) ---

func main() {
	// Inisialisasi template engine untuk HTML
	engine := html.New("./public", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Menyajikan file statis dari direktori 'public/assets'
	app.Static("/assets", "./public/assets")

	// Routing
	app.Get("/", landing_page)                       // Handler untuk halaman utama
	app.Get("/cv", input_cv)                         // Handler untuk halaman input CV
	app.Post("/generate", func(c *fiber.Ctx) error { // Handler untuk generasi CV via AI
		var input ai.CVInput

		if err := c.BodyParser(&input); err != nil {
			fmt.Println("Error parsing JSON:", err)
			body := c.Body()
			fmt.Println("🧾 Raw body:", string(body))

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		fmt.Printf("Parsed input: %+v\n", input)

		// Memanggil fungsi AI untuk generate CV (asumsi goweb-cv-ai/internal/ai sudah didefinisikan)
		result, err := ai.GenerateCV(input)
		if err != nil {
			fmt.Println("❌ AI error:", err)
			return c.Status(500).JSON(fiber.Map{"error": "AI generation failed"})
		}

		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "CV generated successfully",
			"cv":      result,
		})
	})

	app.Post("/download-pdf", downloadPDFHandler) // Handler untuk generate PDF

	// Menjalankan server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}

// --- HANDLER VIEWS ---

func landing_page(c *fiber.Ctx) error {
	// Render halaman index.html (dengan layout app.html)
	return c.Render("index", fiber.Map{
		"Title": "Welcome to Fiber with HTML Templates",
	}, "app")
}

func input_cv(c *fiber.Ctx) error {
	// Render halaman cv/index.html
	return c.Render("cv/index", fiber.Map{
		"Title": "Input CV Anda",
	})
}

func downloadPDFHandler(c *fiber.Ctx) error {
	var req struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(20, 20, 20)
	pdf.AddPage()

	pageWidth, _ := pdf.GetPageSize()
	leftMargin, _, rightMargin, _ := pdf.GetMargins()
	contentWidth := pageWidth - leftMargin - rightMargin

	// Font UTF-8
	fontFile := "DejaVuSans.ttf"
	boldFontFile := "DejaVuSans-Bold.ttf"
	italicFontFile := "DejaVuSans-Oblique.ttf"
	if _, err := os.Stat(fontFile); err == nil {
		pdf.AddUTF8Font("DejaVu", "", fontFile)
		pdf.AddUTF8Font("DejaVu", "B", boldFontFile)
		pdf.AddUTF8Font("DejaVu", "I", italicFontFile)
		pdf.SetFont("DejaVu", "", 12)
	} else {
		pdf.SetFont("Arial", "", 12)
	}

	// Bersihkan isi dari baris kosong
	lines := strings.Split(req.Content, "\n")
	var contentLines []string
	for _, line := range lines {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			contentLines = append(contentLines, trimmed)
		}
	}
	if len(contentLines) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Content cannot be empty")
	}

	// Header nama
	if strings.HasPrefix(contentLines[0], "**") && strings.HasSuffix(contentLines[0], "**") {
		header := strings.TrimSuffix(strings.TrimPrefix(contentLines[0], "**"), "**")
		pdf.SetFont("DejaVu", "B", 20)
		pdf.CellFormat(0, 10, strings.ToUpper(header), "", 1, "C", false, 0, "")
		pdf.Ln(3)
		contentLines = contentLines[1:]
	}

	// Kontak
	if len(contentLines) > 0 && strings.Contains(contentLines[0], "|") {
		pdf.SetFont("DejaVu", "", 10)
		pdf.SetTextColor(80, 80, 80)
		contacts := strings.Split(contentLines[0], "|")
		cellWidth := contentWidth / float64(len(contacts))
		for _, contact := range contacts {
			pdf.CellFormat(cellWidth, 6, strings.TrimSpace(contact), "", 0, "C", false, 0, "")
		}
		pdf.Ln(8)
		contentLines = contentLines[1:]
	}
	pdf.SetTextColor(0, 0, 0)

	pdf.SetFont("DejaVu", "", 11)

	for _, rawLine := range contentLines {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			continue
		}

		// Judul section
		if strings.HasPrefix(line, "**") && strings.HasSuffix(line, "**") {
			title := strings.TrimSuffix(strings.TrimPrefix(line, "**"), "**")
			pdf.Ln(6)
			pdf.SetFont("DejaVu", "B", 13)
			pdf.CellFormat(0, 8, strings.TrimSpace(title), "", 1, "L", false, 0, "")
			pdf.Ln(2)
			pdf.SetFont("DejaVu", "", 11)
			continue
		}

		// Subjudul pekerjaan
		if strings.Contains(line, "di ") && strings.Contains(line, "(") && strings.Contains(line, ")") {
			pdf.Ln(3)
			pdf.SetFont("DejaVu", "B", 11)
			pdf.MultiCell(0, 6, line, "", "J", false)
			pdf.SetFont("DejaVu", "", 11)
			pdf.Ln(1)
			continue
		}

		// Bullet point
		if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "•") || strings.HasPrefix(line, "●") || strings.HasPrefix(line, "·") {
			text := strings.TrimSpace(line)
			text = strings.TrimPrefix(text, "- ")
			text = strings.TrimPrefix(text, "•")
			text = strings.TrimPrefix(text, "●")
			text = strings.TrimPrefix(text, "·")
			pdf.SetX(pdf.GetX() + 2)
			pdf.SetFont("DejaVu", "", 10)
			pdf.CellFormat(2, 4, "•", "", 0, "L", false, 0, "")
			pdf.MultiCell(contentWidth-6, 5, text, "", "J", false)
			pdf.SetFont("DejaVu", "", 11)
			pdf.Ln(1)
			continue
		}

		// Baris biasa
		cleanLine := strings.ReplaceAll(line, "*", "")
		pdf.MultiCell(0, 6, cleanLine, "", "J", false)
		pdf.Ln(2)
	}

	// Output PDF
	buf := new(bytes.Buffer)
	if err := pdf.Output(buf); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate PDF: " + err.Error())
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=cv_%s.pdf", strings.ReplaceAll(req.Name, " ", "_")))
	return c.Send(buf.Bytes())
}
