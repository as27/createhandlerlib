package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

const tmpl = `package {{.Packagename}}

import (
	"bytes"
    "net/http"
)


// Handler can be used from http Server
func Handler(w http.ResponseWriter, r *http.Request){
    w.Write(libBytes)
}

var libBytes = {{.LibBytes}}`

const version = "v1.0.0"

type tmplStruct struct {
	Packagename string
	LibBytes    string
}

func main() {
	url := flag.String("url", "", "When set, the lib is loaded from a url")
	name := flag.String("pkg", "", "Package name")
	outDir := flag.String("dst", "", "Subfolder where the package is stored.")
	versFlag := flag.Bool("v", false, "Prints the version")
	flag.Parse()
	if *versFlag {
		fmt.Println("Version:", version)
		os.Exit(-1)
	}
	var r io.Reader
	var err error
	switch {
	case *url != "":
		log.Println("Loading ", *url)
		r, err = loadLib(*url)
		if err != nil {
			log.Println(err)
		}
	default:
		log.Println("Please set flags for source and package name!")
		os.Exit(-1)
	}
	libBytes, err := createLibBytes(r)
	tmplVars := tmplStruct{
		Packagename: *name,
		LibBytes:    string(libBytes),
	}
	t := template.Must(template.New("lib").Parse(tmpl))
	buf := &bytes.Buffer{}
	t.Execute(buf, tmplVars)
	filePath := path.Join(
		*outDir,
		*name,
		fmt.Sprintf("%s.go", *name),
	)
	log.Println("Writing ", filePath)
	os.MkdirAll(filepath.Dir(filePath), 0777)
	err = ioutil.WriteFile(filePath, buf.Bytes(), 0777)
}

func createLibBytes(r io.Reader) ([]byte, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return b, err
	}
	lb := bytes.NewBufferString(fmt.Sprintf("%#v", b))
	return lb.Bytes(), nil
}

func loadLib(url string) (io.Reader, error) {
	r, err := http.Get(url)
	return r.Body, err
}
