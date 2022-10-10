package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
)

func main() {
	client_id := "xx"
	client_secret_id := "xx"
	tenant_id := "xx"
	authority := "https://login.microsoftonline.com/" + tenant_id
	userAssertion := "xx"

	cred, err := confidential.NewCredFromSecret(client_secret_id)
	if err != nil {
		log.Fatal(err)
	}
	app, err := confidential.New(client_id, cred, confidential.WithAuthority(authority))
	if err != nil {
		log.Fatal(err)
	}
	result, err := app.AcquireTokenOnBehalfOf(context.Background(), userAssertion, []string{"user.read"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("On behalf of token is " + result.AccessToken)
}
