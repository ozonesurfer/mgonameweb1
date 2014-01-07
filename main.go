// main
package main

import (
	"fmt"
	"mgonameweb1/controllers"
	"net/http"
)

func main() {
	//	fmt.Println("Hello World!")
	fmt.Println("Starting server at localhost:8080")
	http.HandleFunc("/", controllers.HomeIndexController)
	http.HandleFunc("/home/validate", controllers.HomeValidateController)
	http.ListenAndServe("localhost:8080", nil)
}
