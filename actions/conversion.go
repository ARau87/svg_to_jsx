package actions

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
)

type SvgData struct {
	Name    string
	XMLData string
}

func ConvertFiles(ctx *cli.Context) error {

	currentDir, _ := os.Getwd()

	files, err := getFiles(ctx.String("dir"))
	if err != nil {
		log.Fatal("Cannot open file")
	}

	if _, err := os.Stat(ctx.String("out")); os.IsNotExist(err) {
		os.Mkdir(ctx.String("out"), os.ModePerm)
	}

	for _, file := range files {
		if !file.IsDir() {
			iconName := strings.ReplaceAll(strings.Title(strings.Split(file.Name(), ".")[0]+"Icon"), "-", "")
			filename := strings.ReplaceAll(file.Name(), ".svg", ".js")
			fileData := readFile(path.Join(ctx.String("dir"), file.Name()))
			if strings.Contains(file.Name(), ".svg") {
				svgData := &SvgData{Name: iconName, XMLData: fileData}
				writeDataToFile(path.Join(currentDir, ctx.String("out"), filename), svgData)
			}
		}
	}

	return nil

}

func writeDataToFile(path string, data *SvgData) error {

	file, err := os.Create(path)
	defer file.Close()

	if err != nil {
		return err
	}

	tmp := setupTemplate()
	tmp.Execute(file, data)

	return nil
}

func setupTemplate() *template.Template {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	templateFile := path.Join(dir, "template.tmpl")
	tmpl := template.Must(template.ParseFiles(templateFile))
	return tmpl

}

func readFile(filename string) string {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Cannot open file")
	}

	return string(data)
}

func getFiles(dir string) ([]os.FileInfo, error) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return files, nil
}
