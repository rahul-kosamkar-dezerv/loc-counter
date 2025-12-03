package counter

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	locClassifier "loc-counter/classifier"
	"loc-counter/syntax"
)

type Result struct {
	Files        int
	Blank        int
	Comments     int
	Code         int
	Imports      int
	Declarations int
	Total        int
}

func CountPath(path string, classifier locClassifier.LineClassifier, ext string) (Result, error) {
	r := Result{Imports: -1, Declarations: -1}
	provider := classifier.Provider
	var s syntax.Syntax
	if ext != "" {
		s = provider.GetByExt(ext)
	}
	if s == nil {
		s = provider.Default()
	}
	if s != nil {
		r.Imports = 0
		r.Declarations = 0
	}

	f, err := os.Open(path)
	if err != nil {
		return r, err
	}
	defer f.Close()

	stat, _ := f.Stat()
	if stat.IsDir() {
		return r, fmt.Errorf("path is directory")
	}

	// read file content
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return r, err
	}
	lines := strings.Split(string(contentBytes), "\n")
	inBlock := false
	for _, line := range lines {
		r.Total += 1
		lt := classifier.Classify(line, s, &inBlock)
		switch lt {
		case locClassifier.Blank:
			r.Blank++
		case locClassifier.Comment:
			r.Comments++
		case locClassifier.Code:
			r.Code++
		case locClassifier.Import:
			r.Imports++
		case locClassifier.Declaration:
			r.Declarations++
		}
	}
	r.Files = 1
	return r, nil
}

func CountDir(root string, classifier locClassifier.LineClassifier) (Result, error) {
	var total Result
	total.Imports = 0
	total.Declarations = 0
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		// check if extension is supported
		if classifier.Provider.GetByExt(ext) == nil {
			return nil
		}
		res, err := CountPath(path, classifier, ext)
		if err != nil {
			return err
		}
		total.Files += res.Files
		total.Blank += res.Blank
		total.Comments += res.Comments
		total.Code += res.Code
		total.Imports += res.Imports
		total.Declarations += res.Declarations
		total.Total += res.Total
		return nil
	})
	return total, err
}
