package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Welcome to home"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w, "Oops, Error is happening. Keep calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "i'm learning golang web",
	// 	"content": "i'm learning golang web with afung setiawan",
	// }

	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 3}
	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 11},
		{ID: 2, Name: "Xpander", Price: 240000000, Stock: 4},
		{ID: 3, Name: "Avanza", Price: 210000000, Stock: 1},
	}

	// w.Write([]byte("Home"))
	// err = tmpl.Execute(w, nil)
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w, "Oops, Error is happening. Keep calm", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, saya sedang belajar golang web"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario hai hello"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNum, err := strconv.Atoi(id) // String to int

	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Product Page "))
	// fmt.Fprintf(w, "Product Page : %d", idNum)

	data := map[string]interface{}{
		"content": idNum,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Oops, Error is happening. Keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Oops, Error is happening. Keep calm", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini adalah GET"))
	case "POST":
		w.Write([]byte("Ini adalah POST"))
	default:
		http.Error(w, "Oops, Error is happening. Keep calm", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Oops, Error is happening. Keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Oops, Error is happening. Keep calm", http.StatusInternalServerError)
			return
		}

		return
	}
	http.Error(w, "Oops, Error is happening. Keep calm", http.StatusBadRequest)

}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm() //mengolah form yang ada
		if err != nil {
			log.Println(err)
			http.Error(w, "Oops, Error is happening. Keep calm", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Oops, Error is happening. Keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Oops, Error is happening. Keep calm", http.StatusInternalServerError)
			return
		}

		return
	}
	http.Error(w, "Oops, Error is happening. Keep calm", http.StatusBadRequest)
}
