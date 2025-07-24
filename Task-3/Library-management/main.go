package main

import (
	"Library-management/services"
	"Library-management/controllers"
)

func main() {
	library := services.NewLibrary()
	controller := &controllers.LibraryController{Manager: library}
	controller.Run()
}