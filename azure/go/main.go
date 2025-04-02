package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

func main() {
	// Start timing the entire operation
	startTime := time.Now()

	// Create a new DefaultAzureCredential which will use az login if available
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to obtain a credential: %v", err)
	}

	// Create a new client for the Subscriptions API
	client, err := armsubscriptions.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("Failed to create subscriptions client: %v", err)
	}

	// Start timing the API call
	apiStartTime := time.Now()

	// List subscriptions
	pager := client.NewListPager(nil)

	// Iterate through the subscription results
	var subscriptionCount int
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			log.Fatalf("Failed to get next page of subscriptions: %v", err)
		}

		// Print the subscription details and count
		for _, subscription := range page.Value {
			fmt.Printf("Subscription ID: %s, Display Name: %s\n", *subscription.SubscriptionID, *subscription.DisplayName)
			subscriptionCount++
		}
	}

	// Stop timing the API call
	apiEndTime := time.Now()
	apiDuration := apiEndTime.Sub(apiStartTime)

	// Stop timing the entire operation
	endTime := time.Now()
	totalDuration := endTime.Sub(startTime)

	// Output the results
	fmt.Printf("\n--- Timing Results ---\n")
	fmt.Printf("Total Execution Time: %v\n", totalDuration)
	fmt.Printf("API Call Time: %v\n", apiDuration)
	fmt.Printf("Subscription Count: %d\n", subscriptionCount)
}
