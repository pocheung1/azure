package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}

	clientId, ok := os.LookupEnv("AZURE_CLIENT_ID")
	if !ok {
		panic("AZURE_CLIENT_ID could not be found")
	}

	options := azidentity.ManagedIdentityCredentialOptions{ID: azidentity.ClientID(clientId)}
	cred, err := azidentity.NewManagedIdentityCredential(&options)
	handleError(err)

	// URL of your storage account blob service
	url := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	// Create a blob service client using the ManagedIdentityCredential
	client, err := azblob.NewClient(url, cred, nil)
	handleError(err)

	fmt.Println("Client created")

	pager := client.NewListContainersPager(&azblob.ListContainersOptions{
		Include: azblob.ListContainersInclude{Metadata: true, Deleted: true},
	})

	fmt.Println("Pager created")

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		fmt.Printf("Page response: %v\n", resp)
		handleError(err)
		for _, container := range resp.ContainerItems {
			fmt.Printf("%v\n", container)
		}
	}
}
