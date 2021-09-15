package main

import (
    "html/template"
    "net/http"
    "os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tpl.Execute(w, nil)
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    mux := http.NewServeMux()

    fs := http.FileServer(http.Dir("assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

    mux.HandleFunc("/", indexHandler)
    http.ListenAndServe(":"+port, mux)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    u, err := url.Parse(r.URL.String())
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }

    params := u.Query()
    searchKey := params.Get("q")
    page := params.Get("page")
    if page == "" {
        page = "1"
    }

    fmt.Println("Search Query is: ", searchKey)
    fmt.Println("Results page is: ", page)
}
