package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// =====================
// Struct untuk Input
// =====================

type CVInput struct {
	Personal struct {
		FullName  string `json:"fullName"`
		JobTitle  string `json:"jobTitle"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		LinkedIn  string `json:"linkedin"`
		Portfolio string `json:"portfolio"`
		Location  string `json:"location"`
		Summary   string `json:"summary"`
	} `json:"personal"`

	Preferences struct {
		Template string `json:"template"`
		ATS      string `json:"ats"`
		Tone     string `json:"tone"`
		Language string `json:"language"`
	} `json:"preferences"`

	Education []struct {
		Institution string `json:"institution"`
		Degree      string `json:"degree"`
		StartYear   string `json:"startYear"`
		EndYear     string `json:"endYear"`
		GPA         string `json:"gpa"`
	} `json:"education"`

	Experience []struct {
		Position    string `json:"position"`
		Company     string `json:"company"`
		StartDate   string `json:"startDate"`
		EndDate     string `json:"endDate"`
		Description string `json:"description"`
		Projects    []struct {
			URL string `json:"url"`
		} `json:"projects"`
	} `json:"experience"`

	Skills []string `json:"skills"`

	PortfolioLinks []struct {
		Platform string `json:"platform"`
		URL      string `json:"url"`
	} `json:"portfolioLinks"`

	Certificates []struct {
		Name   string `json:"name"`
		Issuer string `json:"issuer"`
		Year   string `json:"year"`
	} `json:"certificates"`
}

// =====================
// Struct untuk Response dari Gemini
// =====================

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

// =====================
// Fungsi Utama GenerateCV
// =====================

func GenerateCV(input CVInput) (string, error) {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	apiKey := os.Getenv("GEMINI_API_KEY")

	prompt := fmt.Sprintf(`
Kamu adalah asisten profesional pembuat CV standar Indonesia modern.
Tugasmu adalah menghasilkan teks CV yang siap diubah ke PDF dengan gaya profesional dan menonjolkan keunggulan AI CV Builder:

- Nama besar di bagian atas, semua huruf kapital.
- Baris berikutnya berisi kontak (telepon | email | LinkedIn) dalam satu baris.
- Ringkasan profesional dibuat sebagai paragraf panjang, deskriptif, dan persuasif, menonjolkan keunggulan kandidat dan keahlian unik. Gunakan AI untuk memperluas kata-kata agar terlihat lebih menarik bagi HRD.
- Setiap bagian (Pendidikan, Pengalaman Kerja, Keahlian, Portfolio/Proyek) ditulis tebal, diikuti isi.
- Deskripsi pekerjaan (Pengalaman Kerja) diperluas secara profesional dengan bullet ●, menekankan tanggung jawab, pencapaian, dan hasil yang nyata. Gunakan bahasa yang menjual, formal, dan mudah dibaca.
- Gunakan bullet ● hanya untuk tanggung jawab atau pencapaian pekerjaan.
- Gunakan bahasa formal, profesional, mudah dibaca, dan ATS friendly.
- Tekankan keunggulan: AI Smart Content, Instan & Mudah, Export Fleksibel, Multi Bahasa, Aman & Private jika relevan.
- Jangan gunakan simbol *, #, garis, atau format markdown.
- Jangan tulis label "Curriculum Vitae".
- Jika data kosong, lewati bagian itu sepenuhnya.

Berikut data untuk CV:

Nama Lengkap        : %s
Profesi             : %s
Email               : %s
Telepon             : %s
Tempat / Tanggal Lahir : %s
LinkedIn            : %s
Portfolio           : %s

Ringkasan Profesional:
%s

Pendidikan:
%s

Pengalaman Kerja:
%s

Keahlian:
%s

Preferensi:
Template    : %s
ATS Friendly: %s
Gaya Bahasa : %s
Bahasa      : %s

Hasil CV harus terlihat profesional, siap dicetak atau dikirim ke HRD, menonjolkan keunggulan AI CV Builder, dengan ringkasan profesional panjang dan persuasif, deskripsi pekerjaan diperluas dan menjual, tata bahasa yang benar, paragraf deskriptif, bullet yang jelas untuk pencapaian/tanggung jawab, dan rapi.
`,
		input.Personal.FullName,
		input.Personal.JobTitle,
		input.Personal.Email,
		input.Personal.Phone,
		input.Personal.Location,
		input.Personal.LinkedIn,
		input.Personal.Portfolio,
		input.Personal.Summary,
		formatEducation(input.Education),
		formatExperience(input.Experience),
		strings.Join(input.Skills, ", "),
		input.Preferences.Template,
		input.Preferences.ATS,
		input.Preferences.Tone,
		input.Preferences.Language,
	)

	// 🔹 Siapkan request body untuk Gemini API
	body := map[string]interface{}{
		"contents": []map[string]interface{}{
			{"parts": []map[string]string{{"text": prompt}}},
		},
	}

	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST",
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent",
		bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-goog-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Println("❌ Gemini API error:", string(respBody))
		return "", fmt.Errorf("Gemini API returned status %d", resp.StatusCode)
	}

	var gemResp GeminiResponse
	if err := json.Unmarshal(respBody, &gemResp); err != nil {
		fmt.Println("❌ Error decode Gemini response:", err)
		fmt.Println("🧾 Raw:", string(respBody))
		return "", err
	}

	if len(gemResp.Candidates) == 0 || len(gemResp.Candidates[0].Content.Parts) == 0 {
		fmt.Println("⚠️ Gemini Response kosong:", string(respBody))
		return "", fmt.Errorf("no response from Gemini")
	}

	return gemResp.Candidates[0].Content.Parts[0].Text, nil
}

// =====================
// Helper Functions
// =====================

func formatEducation(list []struct {
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	StartYear   string `json:"startYear"`
	EndYear     string `json:"endYear"`
	GPA         string `json:"gpa"`
}) string {
	if len(list) == 0 {
		return "-"
	}
	var sb strings.Builder
	for _, e := range list {
		sb.WriteString(fmt.Sprintf("- %s, %s (%s - %s) GPA: %s\n",
			e.Institution, e.Degree, e.StartYear, e.EndYear, e.GPA))
	}
	return sb.String()
}

func formatExperience(list []struct {
	Position    string `json:"position"`
	Company     string `json:"company"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Description string `json:"description"`
	Projects    []struct {
		URL string `json:"url"`
	} `json:"projects"`
}) string {
	if len(list) == 0 {
		return "-"
	}
	var sb strings.Builder
	for _, e := range list {
		sb.WriteString(fmt.Sprintf("- %s di %s (%s - %s)\n  %s\n",
			e.Position, e.Company, e.StartDate, e.EndDate, e.Description))
		for _, p := range e.Projects {
			sb.WriteString(fmt.Sprintf("    🔗 %s\n", p.URL))
		}
	}
	return sb.String()
}

func formatPortfolio(list []struct {
	Platform string `json:"platform"`
	URL      string `json:"url"`
}) string {
	if len(list) == 0 {
		return "-"
	}
	var sb strings.Builder
	for _, e := range list {
		sb.WriteString(fmt.Sprintf("- %s: %s\n", e.Platform, e.URL))
	}
	return sb.String()
}
