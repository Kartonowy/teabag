package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"teabag/v2/form_component"
)


func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		k := [3]formcomponent.Component{
				formcomponent.Checkbox, 
				formcomponent.Range,
				formcomponent.Textarea,
			}
		b := formcomponent.InputComponent(k[:])
		w.Write([]byte(b))
	})

	err := http.ListenAndServe(":3333", router)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Listening on port 3333!")
}
