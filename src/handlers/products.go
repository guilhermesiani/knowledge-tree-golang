package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/olivere/elastic"
)

type Product struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	client, err := elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"))

	if err != nil {
		log.Fatalf(err.Error())
		http.Error(
			w,
			"Some error occurred",
			http.StatusInternalServerError,
		)
	}

	res, err := client.Search().
		Index("siani").
		Do(context.Background())

	log.Println(res)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatalf(err.Error())
		http.Error(w, "Some error occurred", http.StatusInternalServerError)
	}

	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		log.Fatalf(err.Error())
		http.Error(w, "Some error occurred", http.StatusInternalServerError)
	}

	client, err := elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"))

	if err != nil {
		log.Fatalf(err.Error())
		http.Error(
			w,
			"Some error occurred",
			http.StatusInternalServerError,
		)
	}

	client.Index().
		Index("siani").
		Type("products").
		BodyJson(product).
		Do(context.Background())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
