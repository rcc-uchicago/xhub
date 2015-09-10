package main

import (
    "log"
    "net/http"

    . "github.com/joyrexus/xhub"
    "github.com/julienschmidt/httprouter"
)

const (
	host = "localhost"
	port = 8080
}

func main() {
    db, err := NewDB("bolt.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    h := NewHandles(host, port, db)
    mux := httprouter.New() 
    mux.GET("/", Index)

	// study handles
    mux.GET("/studies", h.studies.List)
    mux.POST("/studies", h.studies.Create)
    mux.GET("/studies/:id", h.studies.Get)
    mux.PUT("/studies/:id", h.studies.Replace)
    mux.DELETE("/studies/:id", h.studies.Delete)

	// file handles (study level)
	mux.GET("/studies/:study/files", h.files.List)
	mux.POST("/studies/:study/files", h.files.Create)
	mux.GET("/studies/:study/files/:id", h.files.Get)
	mux.PUT("/studies/:study/files/:id", h.files.Replace)
	mux.DELETE("/studies/:study/files/:id", h.files.Delete)

	// trial handles (verbose routes)
	mux.GET("/studies/:study/trials", h.trials.List)
	mux.POST("/studies/:study/trials", h.trials.Create)
	mux.GET("/studies/:study/trials/:id", h.trials.Get)
	mux.PUT("/studies/:study/trials/:id", h.trials.Replace)
	mux.DELETE("/studies/:study/trials/:id", h.trials.Delete)

	// trial handles (convenience routes)
	mux.GET("/trials/:study", h.trials.List)
	mux.POST("/trials/:study", h.trials.Create)
	mux.GET("/trials/:study/:id", h.trials.Get)
	mux.PUT("/trials/:study/:id", h.trials.Replace)
	mux.DELETE("/trials/:study/:id", h.trials.Delete)

	// file handles (trial level)
	mux.GET("/files/:study/:trial", h.files.List)
	mux.POST("/files/:study/:trial", h.files.Create)
	mux.GET("/files/:study/:trial/:id", h.files.Get)
	mux.PUT("/files/:study/:trial/:id", h.files.Replace)
	mux.DELETE("/files/:study/:trial/:id", h.files.Delete)

	// serve it!
    log.Fatal(http.ListenAndServe(":8080", mux))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    http.Redirect(w, r, "/studies", http.StatusFound)
}
