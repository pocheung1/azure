package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	clientId, ok := os.LookupEnv("AZURE_CLIENT_ID")
	if !ok {
		panic("AZURE_CLIENT_ID could not be found")
	}

	options := azidentity.ManagedIdentityCredentialOptions{ID: azidentity.ClientID(clientId)}
	cred, err := azidentity.NewManagedIdentityCredential(&options)
	handleError(err)

	scopes := []string{
		"https://storage.azure.com/.default",
	}
	token, err := cred.GetToken(context.Background(), policy.TokenRequestOptions{Scopes: scopes})
	handleError(err)

	fmt.Printf("Access Token: %s\n", token.Token)
}
