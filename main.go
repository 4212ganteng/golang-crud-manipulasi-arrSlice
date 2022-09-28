package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// struct
type Struktur struct {
	Name string
	Start_date string
	End_date string
	Deskripsi string
	Node string
	React string
	Laravel string
	Golang string
	Gambar string
	Duration string
	Id int
}

var iniArray = []Struktur{}
func main() {
	route := mux.NewRouter()

	// path prefix
	
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// routing

	route.HandleFunc("/",home).Methods("GET")
	route.HandleFunc("/add-blog", addProject).Methods("GET")
	route.HandleFunc("/store-blog", storeProject).Methods("POST")
	route.HandleFunc("/detail-blog/{id}", detailProject).Methods("GET")
	route.HandleFunc("/edit/{id}", editProject).Methods("GET")
	route.HandleFunc("/update-blog/{id}", updateProject).Methods("POST")
	route.HandleFunc("/delete/{id}", deleteProject).Methods("GET")


	// contact Route

	// route.HandleFunc("/contact", contact).Methods("GET")

	// server
	fmt.Println("server is runing on 127.0.0.1:5000")
	http.ListenAndServe("127.0.0.1:5000",route)
}

func home(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type","text/html; charset=utf-8")
	theme, err := template.ParseFiles("views/blog/index.html")

	if err != nil {
		res.Write([]byte("massage : HACKER JANGAN MENYERANG !" + err.Error()))
	}

	responPenampung := map[string]interface{}{
		"Penampungdata" : iniArray,

	}
	theme.Execute(res, responPenampung)
}

func addProject(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type","text/html; charset=utf-8")
	theme, err := template.ParseFiles("views/blog/addproject.html")

	if err != nil {
		res.Write([]byte("massage : HACKER JANGAN MENYERANG !" + err.Error()))
	}

	theme.Execute(res, nil)
}
func storeProject(res http.ResponseWriter, req *http.Request)  {
	err := req.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	title := req.PostForm.Get("title")
	start_date := req.PostForm.Get("start-date")
	end_date := req.PostForm.Get("end-date")
	desc := req.PostForm.Get("desc")
	node := req.PostForm.Get("node")
	laravel := req.PostForm.Get("laravel")
	react := req.PostForm.Get("react")
	golang := req.PostForm.Get("golang")

	layouts := "2006-01-02"
	convStartDate, _ := time.Parse(layouts, start_date)  
	convEndtDate, _ := time.Parse(layouts, end_date)  

	hourse := convEndtDate.Sub(convStartDate).Hours()
	days := hourse/24
	weeks := days/7
	months := days/30
	years := months/12

	var duration string
	if days >= 1 && days <= 6 {
        duration = strconv.Itoa(int(days)) + " days"
    } else if days >= 7 && days <= 29 {
        duration = strconv.Itoa(int(weeks)) + " weeks"
    } else if days >= 30 && days <= 364 {
        duration = strconv.Itoa(int(months)) + " months"
    } else if days >= 365 {
        duration = strconv.Itoa(int(years)) + " years"
    }


	var newProject = Struktur{
		Name : title,
		Start_date : start_date,
		End_date :  end_date,
		Deskripsi : desc,
		Node : node,
		Laravel : laravel,
		React : react,
		Golang : golang,
		Duration : duration,
		Id : len(iniArray),
	}

	iniArray = append(iniArray, newProject)
	fmt.Println(iniArray)

	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}

func detailProject(res http.ResponseWriter, req *http.Request)  {

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	theme, err := template.ParseFiles("views/blog/detail.html")

	if err != nil {
		res.Write([]byte("Hacker jangan menyerang! :" + err.Error()))
		return
	}
	var blogDetail = Struktur{}

	index, _ := strconv.Atoi(mux.Vars(req)["id"])

	for i, data := range iniArray{
		if index == i{
			blogDetail = Struktur{
				Name : data.Name,
				Deskripsi : data.Deskripsi,
				Duration: data.Duration,
				Start_date: data.Start_date,
				End_date: data.End_date,
				Node: data.Node,
				React: data.React,
				Laravel: data.Laravel,
				Golang: data.Golang,
			}
		}
	}

	data := map[string]interface{}{
		"detail" : blogDetail,
	}
	fmt.Println(data)
	theme.Execute(res, data)
}
func editProject(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/blog/edit-project.html")

	if err != nil {
		res.Write([]byte("message : "+ err.Error()))
		return
	}

	var editProject = Struktur{}

	index, _ := strconv.Atoi(mux.Vars(req)["id"])

	for i, project := range iniArray {
		if index == i {
			editProject = Struktur{
				Name: project.Name,
				Deskripsi: project.Deskripsi,
				Start_date: project.Start_date,
				End_date: project.End_date,
				Node: project.Node,
				Golang: project.Golang,
				React: project.React,
				Laravel: project.Laravel,
				Id: project.Id,
			}
		}

	}

	data := map[string]interface{}{
		"EditProject": editProject,
	}

	tmpl.Execute(res, data)
}

func updateProject(res http.ResponseWriter, req *http.Request){
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	title := req.PostForm.Get("nameProject")
	description := req.PostForm.Get("description")
		EditdataProject := Struktur{
		Name: title,
		Deskripsi: description,
		Id: id,
	}

	iniArray[id] = EditdataProject

	http.Redirect(res, req, "/", http.StatusFound)
}

func deleteProject(res http.ResponseWriter, req *http.Request){
	index, _ := strconv.Atoi(mux.Vars(req)["id"])

	iniArray = append(iniArray[:index], iniArray[index+1:]...)

	http.Redirect(res, req, "/", http.StatusFound)
}