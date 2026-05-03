# CV-Engine-Go: Professional Curriculum Vitae Generator

Automated and modern Curriculum Vitae generation engine. Transform your professional experience from Markdown into high-quality PDF Curriculum Vitae using Go and AI-based transformation skills.

## Template Preview

| Modern Template | Tech Template |
| :---: | :---: |
| ![Modern Preview](assets/modern_preview.png) | ![Tech Preview](assets/tech_preview.png) |

## Project Structure

```text
.
├── cv_data.json         # Structured source of truth (JSON)
├── jos/                 # Source Markdown files (Modular)
├── skills/              # AI Protocols (Skills)
│   └── CV_Transformer/  # Transformation and generation skill
│       ├── skill.md           # Master index and workflow
│       ├── data_extraction.md # MD -> JSON logic
│       └── pdf_generation.md  # JSON -> PDF logic
├── references/          # Reference documents (Git ignored)
├── src/                 # Go Generator
│   ├── main.go          # Rendering and PDF logic
│   └── templates/       # HTML/CSS Templates
│       ├── default/     # Classic design
│       ├── modern/      # Premium two-column design
│       └── tech/        # Tech-focused design (Violet & Orange)
└── curriculum_*.pdf      # Generated PDF files
```

## Workflow

### 1. Update Data (Markdown)
Data is organized by sections in the `jos/` folder. Edit these files to add experience or skills.

### 2. Sync with JSON (The Skill)
Use the skill in `skills/CV_Transformer/` to update `cv_data.json`.
- Suggested command: *"Use the CV_Transformer skill to update cv_data.json"*.

### 3. Generate the PDF Curriculum Vitae
```bash
cd src
# Modern template (default)
go run main.go -template modern
# Tech template
go run main.go -template tech
```

## Technical Requirements
- **Go** 1.20+
- **Google Chrome/Edge** (for Chromedp)

---
*Developed with ❤️ and AI.*
