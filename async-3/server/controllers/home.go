package controllers

import (
	"async-3/server/utils"
	"html/template"
	"log"
	"net/http"
	"path"
)

type HomeController interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type homeController struct{}

func NewHomeController() HomeController {
	return &homeController{}
}

func (*homeController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("static", "pages/home/index.html"), utils.LayoutMaster)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
