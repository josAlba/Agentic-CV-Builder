package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/chromedp/cdproto/page"

	"github.com/chromedp/chromedp"
)

type CVData struct {
	PersonalInfo struct {
		Name  string `json:"name"`
		Title string `json:"title"`
		Email string `json:"email"`
		Phone string `json:"phone"`
		Links []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"links"`
		Bio   string `json:"bio"`
		Photo    string `json:"photo"`
		Location string `json:"location"`
	} `json:"personal_info"`
	Experience []struct {
		Company      string   `json:"company"`
		Role         string   `json:"role"`
		Period       string   `json:"period"`
		Description  string   `json:"description"`
		Achievements []string `json:"achievements"`
		Technologies []string `json:"technologies"`
	} `json:"experience"`
	Projects []struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		TechStack   []string `json:"tech_stack"`
	} `json:"projects"`
	Skills struct {
		Technical []struct {
			Category string   `json:"category"`
			Items    []string `json:"items"`
		} `json:"technical"`
		Soft  []string `json:"soft"`
		Tools []string `json:"tools"`
	} `json:"skills"`
	Languages []struct {
		Name  string `json:"name"`
		Level string `json:"level"`
	} `json:"languages"`
	Others struct {
		Licenses       []string `json:"licenses"`
		Certifications []string `json:"certifications"`
	} `json:"others"`
}

func main() {
	tmplName := flag.String("template", "modern", "Nombre de la plantilla a usar")
	flag.Parse()

	// 1. Leer JSON de datos
	jsonData, err := ioutil.ReadFile("../cv_data.json")
	if err != nil {
		log.Fatalf("Error leyendo cv_data.json: %v", err)
	}

	var cv CVData
	if err := json.Unmarshal(jsonData, &cv); err != nil {
		log.Fatalf("Error parseando JSON: %v", err)
	}

	// 2. Aplicar Template HTML con funciones personalizadas
	funcMap := template.FuncMap{
		"lower": strings.ToLower,
	}

	tmplPath := filepath.Join("templates", *tmplName, "template.html")
	tmpl, err := template.New("template.html").Funcs(funcMap).ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("Error cargando template: %v", err)
	}

	var finalHTML bytes.Buffer
	if err := tmpl.Execute(&finalHTML, cv); err != nil {
		log.Fatalf("Error ejecutando template: %v", err)
	}


	// Guardar HTML temporal
	tmpHTMLPath := "temp_cv.html"
	ioutil.WriteFile(tmpHTMLPath, finalHTML.Bytes(), 0644)
	defer os.Remove(tmpHTMLPath)

	// 3. Generar PDF
	absPath, _ := filepath.Abs(tmpHTMLPath)
	pdfPath := fmt.Sprintf("../curriculum_%s.pdf", *tmplName)

	if err := generatePDF("file://"+absPath, pdfPath); err != nil {
		log.Fatalf("Error generando PDF: %v", err)
	}

	fmt.Printf("¡Éxito! PDF generado con plantilla '%s' en: %s\n", *tmplName, pdfPath)
}

func generatePDF(url string, outputPath string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPaperWidth(8.27).
				WithPaperHeight(11.69).
				WithMarginTop(0).
				WithMarginBottom(0).
				WithMarginLeft(0).
				WithMarginRight(0).
				Do(ctx)
			return err
		}),
	); err != nil {
		return err
	}
	return ioutil.WriteFile(outputPath, buf, 0644)
}
