package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

// TODO(wperron) parameterize domain and GH user
const hostPath = "go.wperron.io/%s"
const sourcePath = "https://github.com/wperron/%s"

//go:embed templates/index.html.tpl
var index string

//go:embed templates/package.html.tpl
var pkg string

type Package struct {
	Title  string
	Href   string
	Source string
}

func main() {
	// TODO(wperron) parameterize output directory
	if err := os.Mkdir("./public", 0o700); err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Fatalf("failed to create public dir: %s", err)
		}
	}

	indexTpl := must(template.New("index").Parse(index))
	pkgTpl := must(template.New("index").Parse(pkg))

	f := must(os.Open("packages.csv"))
	scan := bufio.NewScanner(f)
	// Consume title line of CSV
	_ = scan.Scan()

	packages := ReadPackages(scan)

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

func ReadPackages(s *bufio.Scanner) []Package {
	packages := make([]Package, 0)
	for s.Scan() {
		l := s.Text()
		parts := strings.Split(l, ",")
		p := Package{
			Title:  parts[1],
			Href:   fmt.Sprintf(hostPath, parts[1]),
			Source: fmt.Sprintf(sourcePath, parts[0]),
		}
		packages = append(packages, p)
	}
	return packages
}

func must[T any](v T, err error) T {
	if err != nil {
		log.Fatalln(err)
	}
	return v
}
