package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

const data = `fizzbuzz-codegen`
const hostPath = "go.wperron.io/%s"
const sourcePath = "https://github.com/wperron/%s"

type Package struct {
	Title  string
	Href   string
	Source string
}

func main() {
	if err := os.Mkdir("./public", 0o600); err != nil {
		log.Fatalf("failed to create public dir: %s", err)
	}

	bs, err := os.ReadFile("./templates/index.html.tpl")
	if err != nil {
		log.Fatalf("failed to read index tpl file: %s", err)
	}

	indexTpl, err := template.New("index").Parse(string(bs))
	if err != nil {
		log.Fatalf("failed to parse template: %s", err)
	}

	bs, err = os.ReadFile("./templates/package.html.tpl")
	if err != nil {
		log.Fatalf("failed to read index tpl file: %s", err)
	}

	pkgTpl, err := template.New("index").Parse(string(bs))
	if err != nil {
		log.Fatalf("failed to parse template: %s", err)
	}

	scan := bufio.NewScanner(strings.NewReader(data))
	packages := make([]Package, 0)
	for scan.Scan() {
		l := scan.Text()
		p := Package{
			Title:  l,
			Href:   fmt.Sprintf(hostPath, l),
			Source: fmt.Sprintf(sourcePath, l),
		}
		packages = append(packages, p)
	}

	var buf bytes.Buffer
	if err := indexTpl.Execute(&buf, packages); err != nil {
		log.Fatalf("failed to execute index template: %s", err)
	}

	if err := os.WriteFile("./public/index.html", buf.Bytes(), 0o600); err != nil {
		log.Fatalf("failed to write index output file: %s", err)
	}

	for _, p := range packages {
		buf.Reset()
		if err := pkgTpl.Execute(&buf, p); err != nil {
			log.Fatalf("failed to execute package template for %s: %s", p.Title, err)
		}

		if err := os.WriteFile(fmt.Sprintf("./public/%s.html", p.Title), buf.Bytes(), 0o600); err != nil {
			log.Fatalf("failed to write package output file fot %s: %s", p.Title, err)
		}
	}
}
