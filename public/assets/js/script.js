
        // Data storage
        let formData = {
            education: [],
            experience: [],
            skills: [],
            portfolioLinks: [],
            certificates: []
        };

        // Initialize form with one education and experience entry
        window.addEventListener('DOMContentLoaded', () => {
            addEducation();
            addExperience();
            addSkill('JavaScript');
            addSkill('React');
            addSkill('Node.js');
            addPortfolioLink();
            setupPhotoUpload();
            setupCertificateUpload();
        });

        // Photo upload handler
        function setupPhotoUpload() {
            document.getElementById('photoInput').addEventListener('change', function(e) {
                const file = e.target.files[0];
                if (file) {
                    if (file.size > 2 * 1024 * 1024) {
                        alert('Ukuran file terlalu besar! Maksimal 2MB');
                        return;
                    }
                    const reader = new FileReader();
                    reader.onload = function(event) {
                        document.getElementById('photoPreview').innerHTML = 
                            `<img src="${event.target.result}" class="w-full h-full object-cover">`;
                    };
                    reader.readAsDataURL(file);
                }
            });
        }

        // Add Education
        function addEducation() {
            const id = Date.now();
            const educationHTML = `
                <div class="border border-gray-200 rounded-xl p-6 hover:border-purple-300 transition" data-id="${id}">
                    <div class="grid md:grid-cols-2 gap-4">
                        <div>
                            <label class="block text-sm font-semibold text-gray-700 mb-2">Institusi *</label>
                            <input type="text" placeholder="Universitas Indonesia" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent">
                        </div>
                        <div>
                            <label class="block text-sm font-semibold text-gray-700 mb-2">Gelar *</label>
                            <input type="text" placeholder="S1 Teknik Informatika" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent">
                        </div>
                        <div>
                            <label class="block text-sm font-semibold text-gray-700 mb-2">Tahun Mulai *</label>
                            <input type="text" placeholder="2018" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent">
                        </div>
                        <div>
                            <label class="block text-sm font-semibold text-gray-700 mb-2">Tahun Selesai *</label>
                            <input type="text" placeholder="2022 atau 'Sekarang'" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent">
                        </div>
                        <div class="md:col-span-2">
                            <label class="block text-sm font-semibold text-gray-700 mb-2">IPK (Opsional)</label>
                            <input type="text" placeholder="3.75 / 4.00" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent">
                        </div>
                    </div>
                    <button type="button" onclick="removeEducation(${id})" class="text-red-500 hover:text-red-700 text-sm mt-4">
                        <i class="fas fa-trash-alt mr-1"></i> Hapus
                    </button>
                </div>
            `;
            document.getElementById('educationList').insertAdjacentHTML('beforeend', educationHTML);
        }

        function removeEducation(id) {
            const element = document.querySelector(`#educationList [data-id="${id}"]`);
            element.remove();
        }

        // Add Experience
        function addExperience() {
            const id = Date.now();
            const experienceHTML = `
                <div class="border border-gray-200 rounded-xl p-6 hover:border-purple-300 transition" data-id="${id}">
                    <div class="grid md:grid-cols-2 gap-4">
                        <div>
                            <label class="block text-sm font-semibold text-gray-700 mb-2">Posisi *</label>
                            <input type="text" placeholder="Senior Software Engineer" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent">
                        </div>
                        <div>
                            <label class="block text-sm font-semibold text-gray-700 mb-2">Perusahaan *</label>
                            <input type="text" placeholder="Google Indonesia" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent">
                        </div>
                        <div>
                            <label class="block text-sm font-semibold text-gray-700 mb-2">Tahun Mulai *</label>
                            <input type="text" placeholder="Jan 2022" class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent">
                        </div>
                        <div>
                            <label class="block text-sm font-semibold text-gray-700 mb-2">Tahun Selesai *</label>
                            <div class="flex gap-2">
                                <input type="text" placeholder="Des 2023 atau 'Sekarang'" class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent" id="endDate${id}">
                                <label class="flex items-center gap-2 bg-gray-100 px-3 rounded-lg">
                                    <input type="checkbox" class="w-4 h-4 text-purple-600" onchange="toggleCurrentJob(${id})">
                                    <span class="text-sm">Sekarang</span>
                                </label>
                            </div>
                        </div>
                        <div class="md:col-span-2">
                            <label class="block text-sm font-semibold text-gray-700 mb-2">Deskripsi Pekerjaan *</label>
                            <textarea rows="4" placeholder="Tuliskan tanggung jawab dan pencapaian Anda. AI akan memformat menjadi bullet points profesional." class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent"></textarea>
                            <p class="text-xs text-gray-500 mt-1"><i class="fas fa-magic text-purple-500"></i> AI akan mengoptimalkan deskripsi ini</p>
                        </div>
                    </div>

                    <!-- Portfolio/Karya untuk Pengalaman -->
                    <div class="mt-6 pt-6 border-t border-gray-200">
                        <label class="block text-sm font-semibold text-gray-700 mb-3">Portfolio / Karya (Opsional)</label>
                        <div class="space-y-3">
                            <div id="projectLinks${id}" class="space-y-2"></div>
                            <button type="button" onclick="addProjectLink(${id})" class="text-purple-600 hover:text-purple-700 text-sm font-semibold">
                                <i class="fas fa-plus-circle mr-1"></i> Tambah Link Project
                            </button>
                            
                            <div class="mt-4">
                                <label class="block text-xs font-semibold text-gray-700 mb-2">Upload Foto Karya</label>
                                <div class="grid grid-cols-2 md:grid-cols-4 gap-3" id="projectImages${id}">
                                    <div class="relative group">
                                        <input type="file" id="projectImageUpload${id}" accept="image/*" multiple class="hidden">
                                        <div onclick="document.getElementById('projectImageUpload${id}').click()" class="aspect-square bg-gray-100 rounded-lg flex items-center justify-center border-2 border-dashed border-gray-300 hover:border-purple-500 cursor-pointer transition">
                                            <i class="fas fa-image text-gray-400 text-2xl"></i>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <button type="button" onclick="removeExperience(${id})" class="text-red-500 hover:text-red-700 text-sm mt-4">
                        <i class="fas fa-trash-alt mr-1"></i> Hapus Pengalaman
                    </button>
                </div>
            `;
            document.getElementById('experienceList').insertAdjacentHTML('beforeend', experienceHTML);
            
            // Setup project image upload
            document.getElementById(`projectImageUpload${id}`).addEventListener('change', function(e) {
                handleProjectImages(e, id);
            });
        }

        function removeExperience(id) {
            const element = document.querySelector(`#experienceList [data-id="${id}"]`);
            element.remove();
        }

        function toggleCurrentJob(id) {
            const endDateInput = document.getElementById(`endDate${id}`);
            if (event.target.checked) {
                endDateInput.value = 'Sekarang';
                endDateInput.disabled = true;
                endDateInput.classList.add('bg-gray-100');
            } else {
                endDateInput.value = '';
                endDateInput.disabled = false;
                endDateInput.classList.remove('bg-gray-100');
            }
        }

        // Add project link for experience
        function addProjectLink(expId) {
            const linkId = Date.now();
            const linkHTML = `
                <div class="flex gap-2" data-link-id="${linkId}">
                    <input type="text" placeholder="https://github.com/project atau link demo" class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent text-sm">
                    <button type="button" onclick="removeProjectLink(${expId}, ${linkId})" class="px-3 py-2 bg-red-100 text-red-600 rounded-lg hover:bg-red-200 transition">
                        <i class="fas fa-times"></i>
                    </button>
                </div>
            `;
            document.getElementById(`projectLinks${expId}`).insertAdjacentHTML('beforeend', linkHTML);
        }

        function removeProjectLink(expId, linkId) {
            const element = document.querySelector(`#projectLinks${expId} [data-link-id="${linkId}"]`);
            element.remove();
        }

        // Handle project images upload
        function handleProjectImages(e, expId) {
            const files = e.target.files;
            const container = document.getElementById(`projectImages${expId}`);
            
            Array.from(files).forEach(file => {
                if (file.size > 5 * 1024 * 1024) {
                    alert('Ukuran file terlalu besar! Maksimal 5MB per file');
                    return;
                }
                
                const reader = new FileReader();
                reader.onload = function(event) {
                    const imageId = Date.now() + Math.random();
                    const imageHTML = `
                        <div class="relative group" data-image-id="${imageId}">
                            <img src="${event.target.result}" class="aspect-square object-cover rounded-lg">
                            <button type="button" onclick="removeProjectImage(${expId}, ${imageId})" class="absolute top-1 right-1 w-6 h-6 bg-red-500 text-white rounded-full opacity-0 group-hover:opacity-100 transition flex items-center justify-center text-xs">
                                <i class="fas fa-times"></i>
                            </button>
                        </div>
                    `;
                    container.insertAdjacentHTML('beforeend', imageHTML);
                };
                reader.readAsDataURL(file);
            });
        }

        function removeProjectImage(expId, imageId) {
            const element = document.querySelector(`#projectImages${expId} [data-image-id="${imageId}"]`);
            element.remove();
        }

        // Skills management
        function addSkill(skillName = null) {
            const input = document.getElementById('skillInput');
            const skills = skillName || input.value.trim();
            
            if (!skills) return;
            
            // Split by comma for multiple skills
            const skillArray = skills.split(',').map(s => s.trim()).filter(s => s);
            
            skillArray.forEach(skill => {
                const skillId = Date.now() + Math.random();
                const colors = ['purple', 'blue', 'green', 'orange', 'pink', 'indigo', 'cyan', 'red'];
                const color = colors[Math.floor(Math.random() * colors.length)];
                
                const skillHTML = `
                    <span class="px-4 py-2 bg-${color}-100 text-${color}-700 rounded-full text-sm font-semibold flex items-center gap-2" data-skill-id="${skillId}">
                        ${skill}
                        <button type="button" onclick="removeSkill(${skillId})" class="hover:text-${color}-900">
                            <i class="fas fa-times"></i>
                        </button>
                    </span>
                `;
                document.getElementById('skillsList').insertAdjacentHTML('beforeend', skillHTML);
            });
            
            if (!skillName) input.value = '';
        }

        function removeSkill(skillId) {
            const element = document.querySelector(`#skillsList [data-skill-id="${skillId}"]`);
            element.remove();
        }

        // Portfolio links management
        function addPortfolioLink() {
            const linkId = Date.now();
            const linkHTML = `
                <div class="flex gap-2" data-portfolio-id="${linkId}">
                    <input type="text" placeholder="https://github.com/username" class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent">
                    <select class="px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent">
                        <option>GitHub</option>
                        <option>Behance</option>
                        <option>Dribbble</option>
                        <option>Medium</option>
                        <option>YouTube</option>
                        <option>Lainnya</option>
                    </select>
                    <button type="button" onclick="removePortfolioLink(${linkId})" class="px-4 py-2 bg-red-100 text-red-600 rounded-lg hover:bg-red-200 transition">
                        <i class="fas fa-times"></i>
                    </button>
                </div>
            `;
            document.getElementById('portfolioLinksList').insertAdjacentHTML('beforeend', linkHTML);
        }

        function removePortfolioLink(linkId) {
            const element = document.querySelector(`#portfolioLinksList [data-portfolio-id="${linkId}"]`);
            element.remove();
        }

        // Certificate upload handler
        function setupCertificateUpload() {
            document.getElementById('certificateUpload').addEventListener('change', function(e) {
                const files = e.target.files;
                const container = document.getElementById('certificatesList');
                
                Array.from(files).forEach(file => {
                    if (file.size > 5 * 1024 * 1024) {
                        alert('Ukuran file terlalu besar! Maksimal 5MB per file');
                        return;
                    }
                    
                    const certId = Date.now() + Math.random();
                    const certHTML = `
                        <div class="flex items-center justify-between bg-gray-50 p-3 rounded-lg" data-cert-id="${certId}">
                            <div class="flex items-center gap-3">
                                <i class="fas fa-file-pdf text-red-500 text-2xl"></i>
                                <div>
                                    <p class="font-semibold text-gray-900 text-sm">${file.name}</p>
                                    <p class="text-xs text-gray-500">${(file.size / 1024).toFixed(2)} KB</p>
                                </div>
                            </div>
                            <button type="button" onclick="removeCertificate(${certId})" class="text-red-500 hover:text-red-700">
                                <i class="fas fa-trash-alt"></i>
                            </button>
                        </div>
                    `;
                    container.insertAdjacentHTML('beforeend', certHTML);
                });
            });
        }

        function removeCertificate(certId) {
            const element = document.querySelector(`#certificatesList [data-cert-id="${certId}"]`);
            element.remove();
        }

        // Import functions
        function importLinkedIn() {
            alert('Fitur impor LinkedIn akan segera hadir! Untuk demo, Anda akan diarahkan untuk authorize LinkedIn.');
            // Implement LinkedIn OAuth here
        }

        function uploadOldCV() {
            document.getElementById('oldCVUpload').click();
            document.getElementById('oldCVUpload').addEventListener('change', function(e) {
                const file = e.target.files[0];
                if (file) {
                    alert(`File "${file.name}" berhasil diupload! AI sedang mengekstrak data...`);
                    // Implement CV parsing here
                }
            });
        }

        function startManual() {
            window.scrollTo({ top: document.getElementById('cvForm').offsetTop - 100, behavior: 'smooth' });
        }

        // Save draft
        function saveDraft() {
            const data = collectFormData();
            localStorage.setItem('cvDraft', JSON.stringify(data));
            
            // Show success notification
            const notification = document.createElement('div');
            notification.className = 'fixed top-20 right-4 bg-green-500 text-white px-6 py-3 rounded-lg shadow-lg z-50 animate-bounce';
            notification.innerHTML = '<i class="fas fa-check-circle mr-2"></i>Draft berhasil disimpan!';
            document.body.appendChild(notification);
            
            setTimeout(() => {
                notification.remove();
            }, 3000);
        }

        // Preview CV
        function previewCV() {
            const data = collectFormData();
            console.log('Preview CV Data:', data);
            alert('Preview CV akan ditampilkan di modal. Data CV:\n\n' + JSON.stringify(data, null, 2).substring(0, 200) + '...');
            // Implement preview modal here
        }

        // 🔧 Ambil data Education dari DOM
function getEducationData() {
    const educations = [];
    document.querySelectorAll('#educationList > [data-id]').forEach(block => {
        const inputs = block.querySelectorAll('input');
        educations.push({
            institution: inputs[0]?.value.trim() || "",
            degree: inputs[1]?.value.trim() || "",
            startYear: inputs[2]?.value.trim() || "",
            endYear: inputs[3]?.value.trim() || "",
            gpa: inputs[4]?.value.trim() || ""
        });
    });
    return educations;
}

// 🔧 Ambil data Experience dari DOM
function getExperienceData() {
    const experiences = [];
    document.querySelectorAll('#experienceList > [data-id]').forEach(block => {
        const inputs = block.querySelectorAll('input[type="text"]');
        const textarea = block.querySelector('textarea');
        const projectLinks = Array.from(block.querySelectorAll('[id^="projectLinks"] input'))
            .map(i => i.value.trim()).filter(Boolean);

        experiences.push({
            position: inputs[0]?.value.trim() || "",
            company: inputs[1]?.value.trim() || "",
            startDate: inputs[2]?.value.trim() || "",
            endDate: inputs[3]?.value.trim() || "",
            description: textarea?.value.trim() || "",
            projects: projectLinks
        });
    });
    return experiences;
}

// 🔧 Ambil data Skills dari DOM
function getSkillsData() {
    return Array.from(document.querySelectorAll('#skillsList [data-skill-id]'))
        .map(s => s.textContent.replace('×', '').trim());
}

// 🔧 Ambil data Portfolio links dari DOM
function getPortfolioData() {
    const links = [];
    document.querySelectorAll('#portfolioLinksList > [data-portfolio-id]').forEach(div => {
        const input = div.querySelector('input');
        const select = div.querySelector('select');
        if (input && input.value.trim()) {
            links.push({
                platform: select.value,
                url: input.value.trim()
            });
        }
    });
    return links;
}

// 🔧 Ambil data Certificates dari DOM
function getCertificatesData() {
    return Array.from(document.querySelectorAll('#certificatesList [data-cert-id]')).map(div => {
        const name = div.querySelector('p.font-semibold')?.textContent.trim();
        const size = div.querySelector('p.text-xs')?.textContent.trim();
        return { name, size };
    });
}


        // Smooth scroll for skill input on Enter
        document.getElementById('skillInput').addEventListener('keypress', function(e) {
            if (e.key === 'Enter') {
                e.preventDefault();
                addSkill();
            }
        });
    