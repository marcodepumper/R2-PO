package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

type BenchmarkResult struct {
	SubscriptionsCallTimeMs float64 `json:"SubscriptionsCallTimeMs"`
	SubscriptionCount       int     `json:"SubscriptionCount"`
}

func main() {
	// Authenticate using DefaultAzureCredential
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to obtain a credential: %v", err)
	}

	// Create Subscriptions Client
	client, err := armsubscriptions.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("Failed to create subscriptions client: %v", err)
	}

	// Start timing the API call
	startTime := time.Now()

	// List subscriptions
	pager := client.NewListPager(nil)
	var subscriptionCount int

	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			log.Fatalf("Failed to get next page of subscriptions: %v", err)
		}

		for _, subscription := range page.Value {
			fmt.Printf("Display Name: %s\n", *subscription.DisplayName)
			subscriptionCount++
		}
	}

	// Stop timing the API call
	apiDuration := time.Since(startTime).Seconds() * 1000 // Convert to milliseconds

	// Create structured result
	result := BenchmarkResult{
		SubscriptionsCallTimeMs: apiDuration,
		SubscriptionCount:       subscriptionCount,
	}

	// Convert result to JSON for structured output
	jsonOutput, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Failed to encode result as JSON: %v", err)
	}

	fmt.Println("\n--- Timing Results ---")
	fmt.Println(string(jsonOutput))
}
