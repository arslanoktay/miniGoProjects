package main

import (
	handlers "Deneme/httpServer"
	"Deneme/middleware"
	gorillaRouting "Deneme/routing"
)

func main() {
	handlers.HttpServerHandlers()
	gorillaRouting.MainRouting()
	middleware.MiddlewareMain()
}
