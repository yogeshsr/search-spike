package elasticsearch_api

import (
	"github.com/olivere/elastic"
	"fmt"
	"reflect"
	"context"
	"log"
	"os"
	"errors"
)

type Voucher struct {
	Sponsor string `json:"sponsor"`
	Title   string `json:"title"`
	Desc    string `json:"description"`
}

func GetElasticClient(elasticUrl string) (*elastic.Client, error) {
	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)
	// see details- https://docs.docker.com/compose/networking/
	client, err := elastic.NewSimpleClient(elastic.SetErrorLog(errorlog), elastic.SetURL(elasticUrl))
	if err != nil {
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(elasticUrl).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion(elasticUrl)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("voucher").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if exists {
		fmt.Printf("connection ok\n")
		return client, nil
	}
	fmt.Printf("connection failed\n")

	return nil, errors.New("Failed to establish Elasticsearch connection")
}

func VoucherSearch(client *elastic.Client, searchText string) ([]Voucher) {
	// Search with a term query
	multiMatchQuery := elastic.NewMultiMatchQuery(searchText, "sponsor^3", "title")
	searchResult, err := client.Search().
		Index("voucher").        // search in index "twitter"
		Query(multiMatchQuery). // specify the query
		From(0).Size(10).        // take documents 0-9
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	var voucher Voucher
	var vouchers[] Voucher
	for _, item := range searchResult.Each(reflect.TypeOf(voucher)) {
		t := item.(Voucher)
		fmt.Printf("Voucher by %s | %s\n", t.Sponsor, t.Title)
		vouchers = append(vouchers, t)
	}
	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d vouchers\n", searchResult.TotalHits())

	return vouchers
}

