package main

import (
	"bufio"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)


type TemplateData struct{
	Path string
	Name string
}

type Data struct{
	Templates string
}

func main() {

	var files []string
	
	err := filepath.Walk("./components", func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			ext := filepath.Ext(info.Name())
			
			if ext == ".html" {
				files = append(files, path)
			}
		}
		return nil
	});
	
	if err != nil {
		log.Fatal(err)
	}

	var output string
	for _, file := range files {
		dat, err := ioutil.ReadFile(file)

		if err != nil {
			log.Fatal(err)
		}

		output += "\n" + string(dat)
	}

	var templates []TemplateData
	err = filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			ext := filepath.Ext(info.Name())

			if ext == ".html" {
				templates = append(templates, TemplateData{Path:path,Name:info.Name()})
			}
		}
		return nil
	});


	for _, t := range templates {
		tmpl := template.Must(template.New(t.Name).ParseFiles(t.Path))

		if err != nil {
			log.Fatal(err)
		}

		if _, err := os.Stat(t.Path); os.IsExist(err) {
			err := os.Remove(t.Name)

			if err != nil {
				log.Fatal(err)
			}
		}

		file, err := os.Create(t.Name)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		writer := bufio.NewWriter(file)

		err = tmpl.Execute(writer, Data{Templates: output})

		if err != nil {
			log.Fatal(err)
		}


		err = writer.Flush()

		if err != nil {
			log.Fatal(err)
		}

	}

}