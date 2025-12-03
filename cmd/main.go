package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"loc-counter/classifier"
	"loc-counter/counter"
	"loc-counter/syntax"
)

func main() {
	mode := flag.String("mode", "auto", "mode: auto | file | dir")
	lang := flag.String("lang", "auto", "language override: auto | java | c | cpp | js | python")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: loc-counter [options] <path>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	path := args[0]

	var provider *syntax.Provider = syntax.NewProvider()
	if *lang != "auto" {
		if s := provider.GetByName(*lang); s != nil {
			provider.SetDefault(s)
		}
	}

	classifier := classifier.NewLineClassifier(provider)

	var res counter.Result
	var err error

	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if info.IsDir() || *mode == "dir" {
		res, err = counter.CountDir(path, classifier)
	} else {
		ext := filepath.Ext(path)
		res, err = counter.CountPath(path, classifier, ext)
	}

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Print summary
	fmt.Printf("Files scanned: %d", res.Files)
	fmt.Printf("Blank: %d", res.Blank)
	fmt.Printf("Comments: %d", res.Comments)
	fmt.Printf("Imports: %d\n", res.Imports)
	fmt.Printf("Code: %d", res.Code)
	if res.Imports >= 0 {
		fmt.Printf("Imports: %d", res.Imports)
	}
	if res.Declarations >= 0 {
		fmt.Printf("Declarations: %d", res.Declarations)
	}
	fmt.Printf("Total lines: %d", res.Total)
}
