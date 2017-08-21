package main

import (
	"net/http"
	"html/template"
	"github.com/DandyDev/data-file-viewer/parsers"
	"github.com/DandyDev/data-file-viewer/templates"
	"strings"
	"os"
	"fmt"
)

type UnsupportedError struct {}

func (e UnsupportedError) Error() string {
	return "Unsupported file type"
}

func handler(w http.ResponseWriter, r *http.Request) {
	filepath := r.URL.Path[len("/view/"):]
	ext := filepath[strings.LastIndex(filepath, ".") + 1:]
	var parser parsers.Parser
	switch ext {
	case "csv":
		parser = parsers.CSVParser{}
	case "prn":
		var columns []string
		columnStr := r.URL.Query().Get("columns")
		if strings.Trim(columnStr, " ") != "" {
			columns = strings.Split(columnStr, ",")
		}
		parser = parsers.FixedWidthParser{Columns: columns}
	default:
		handleError(w, UnsupportedError{})
		return
	}
	reader, e := parsers.DecodeFile(filepath)
	if e != nil {
		handleError(w, e)
		return
	}
	table, e := parser.Parse(reader)
	table.Filename = filepath
	if e != nil {
		handleError(w, e)
		return
	}
	t, _ := template.New("table").Parse(templates.Table)
	t.Execute(w, table)
}

func handleError(w http.ResponseWriter, e error) {
	err := e.Error()
	var status int
	switch t := e.(type) {
	case *os.PathError:
		err = fmt.Sprintf("File '%s' not found", t.Path)
		status = 404
	case UnsupportedError:
		status = 400
	case parsers.EmptyFileError:
		status = 400
	case parsers.ColNotFound:
		status = 400
	default:
		err = fmt.Sprintf("Something went wrong that was unforseen: %s", e.Error())
		status = 500
	}
	http.Error(w, err, status)
}

func main() {
	http.HandleFunc("/view/", handler)
	http.ListenAndServe(":8080", nil)
}
