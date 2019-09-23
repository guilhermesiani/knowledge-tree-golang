package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/olivere/elastic/v7"
)

func init() {
	client, err := elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"))

	if err != nil {
		log.Fatalf(err.Error())
	}

	client.DeleteIndex("siani").
		Do(context.Background())

	var product = Product{"TShirt", 35.10}

	client.Index().
		Index("siani").
		Type("products").
		BodyJson(product).
		Do(context.Background())
	time.Sleep(1 * time.Second)
}

func TestGetAllProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllProducts)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	expected := `{"took":5,"hits":{"total":{"value":1,"relation":"eq"},"max_score":1,"hits":[{"_score":1,"_index":"siani","_type":"products","_id":"z_d0W20BlI3uLLSTnOtb","_seq_no":null,"_primary_term":null,"_source":{"name":"TShirt","price":35.1}}]},"_shards":{"total":1,"successful":1,"failed":0}}`
	if rr.Body.String() != expected {
		t.Errorf(
			"handler returned unexpected body: got %v want %v",
			rr.Body.String(),
			expected,
		)
	}
}
