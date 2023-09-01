package cmd

import (
	"flag"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/quocbang/restful/config"
	"github.com/quocbang/restful/server/delivery/http"
)

func Run() {
	e := echo.New()
	flag.StringVar(&config.C.Host, "host", "localhost", "server host")
	flag.IntVar(&config.C.Port, "port", 7878, "server port")
	flag.Parse()

	e.Use(middleware.CORS())
	http.NewStationHandler(e)
	if err := e.Start(fmt.Sprintf("%s:%d", config.C.Host, config.C.Port)); err != nil {
		log.Fatalf("start failed, error: %v", err)
	}
}
