package main

import(
    "strings"
	"net/http"
    "log"
    "fmt"
    "html/template"
    "github.com/ochload/core/static"
)

type HttpController struct {
}

func (this *HttpController) App(r http.ResponseWriter, req *http.Request) {
    t, err := template.ParseFS(static.Assets, "pages/App.tmpl")
    if err != nil {
        log.Fatal(err)
    }
    if err := t.Execute(r, nil); err != nil {
        log.Fatal(err)
    }
}

func (this *HttpController) Assets(r http.ResponseWriter, req *http.Request) {
    fileName := strings.TrimPrefix(req.URL.Path, "/assets/")
    file,err := static.Assets.ReadFile(fmt.Sprintf("pages/assets/%v", fileName))
    if err != nil {
        log.Fatal(err)
        return
    }
    r.Write(file)
}