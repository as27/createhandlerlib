package main

const tmpl = `package {{Packagename}}

import (
	"bytes"
    "net/http"
)


// Handler can be used from http Server
func Handler(w http.ResponseWriter, r *http.Request){
    w.Write(libBytes)
}

// vueBytes contains vue.js v2.1.6 as []byte
var libBytes = {{LibBytes}}`

func main() {

}
