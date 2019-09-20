package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/olivere/elastic"
)

type Product struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	client, err := elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"))

	if err != nil {
		log.Fatalf("Fail to connect to elasticsearch")
	}

	Product := Product{
		Name:  "Babaloo",
		Price: 3.9,
	}

	client.Index().Index("siani").Type("products").BodyJson(Product).Do(context.Background())

	res, err := client.Search().
		Index("siani").
		Do(context.Background())

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)
}
