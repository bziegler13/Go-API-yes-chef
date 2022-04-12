package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/schollz/ingredients"
)

func returnJSON(w http.ResponseWriter, r *http.Request) {
	url, ok := r.URL.Query()["url"]

	if !ok || len(url[0]) < 1 {
		log.Println("URL parameter is missing")
		return
	}

	log.Println("url is: " + url[0])

	recipe, _ := ingredients.NewFromURL(url[0])

	json.NewEncoder(w).Encode(recipe.IngredientList().Ingredients)

	b, err := json.Marshal(recipe.IngredientList().Ingredients)
	if err != nil {
		fmt.Printf("Error :%s", err)
		return
	}
	fmt.Println(string(b))

}

func main() {
	http.HandleFunc("/", returnJSON)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
