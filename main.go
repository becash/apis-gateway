package main

import (
	"apis_gateway/channels/get_your_guide"
	"context"
	"fmt"
	"log"
	"net/http"
	//"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

func main() {

	// custom HTTP client
	hc := http.Client{}

	client, err := get_your_guide.NewClientWithResponses("http://get-yourguide:1234", get_your_guide.WithHTTPClient(&hc))
	if err != nil {
		log.Fatal(err)
	}

	// Example query
	resp, err := client.GetPricingCategories(context.TODO(), "30000")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Expected HTTP 200 but received %d", resp.StatusCode)
	}

	fmt.Printf("resp.JSON200: %v\n", resp)
}
