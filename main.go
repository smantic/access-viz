package main

import (
	"html/template"

	"github.com/labstack/echo/middleware"
	"github.com/namsral/flag"
	"go.smantic.dev/mux"
)

func main() {

	var (
		maxmindDB     string = flag.String("maxmind", "", "Location of the maxmind DB")
		mapboxToken   string = flag.String("mapbox", "", "MapBox Access Token")
		proxyLogsPath string = flag.String("logs", "", "path to logs to read")
	)

	flag.Parse()

	t := &Template{
		templates: template.Must(template.ParseGlob("web/index.tmpl")),
	}

	r := mux.New()
	r.Use(middleware.Recover())
	//e.Renderer = t
	r.Static("/public", "web/public")
}
