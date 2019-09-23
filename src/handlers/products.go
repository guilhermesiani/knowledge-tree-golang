package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/olivere/elastic/v7"
)

type Product struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		AddProduct(w, r)
		return
	}
	GetAllProducts(w, r)
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

	if err != nil {
		log.Fatalf(err.Error())
		http.Error(
			w,
			"Some error occurred",
			http.StatusInternalServerError,
		)
	}

	var products []Product

	for _, hit := range res.Hits.Hits {
		var product Product
		err := json.Unmarshal(hit.Source, &product)
		if err != nil {
			log.Fatalf(err.Error())
			http.Error(
				w,
				"Some error occurred",
				http.StatusInternalServerError,
			)
		}
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
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
