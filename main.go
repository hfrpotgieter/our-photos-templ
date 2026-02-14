package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/henro47/hfr-go-templ-cv/components"
)

//go:generate templ generate

const (
	runArgGen = "generate"
	runArgRun = "run"
)

// go:generate templ generate

func runDevelopmentServer() {
	homePage := components.Index()
	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/", templ.Handler(homePage))
	log.Println("Starting development server on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", pagesHandler))
}

func generateHTML() {
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = components.Index().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func main() {
	args := os.Args[1]
	if len(args) == 0 {
		log.Fatalf("You must provide arguments to genrate / run server")
	}
	switch args {
	case runArgGen:
		generateHTML()
	case runArgRun:
		runDevelopmentServer()
	default:
		log.Fatalf("Invalid argument: %s. Use either %s or %s", args, runArgGen, runArgRun)
	}
}
