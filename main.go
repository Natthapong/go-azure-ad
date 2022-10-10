package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/confidential"
)

func main() {
	client_id := "47d06252-7fb3-463c-a841-b3be6ecd87f4"
	client_secret_id := "ZvQ8Q~Ca8p3DQUXgR2u4VA1kKcbFgaJNTjhulbko"
	tenant_id := "49b5155e-6e7f-4868-a760-df12b09855fc"
	authority := "https://login.microsoftonline.com/" + tenant_id
	userAssertion := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IjJaUXBKM1VwYmpBWVhZR2FYRUpsOGxWMFRPSSJ9.eyJhdWQiOiI0N2QwNjI1Mi03ZmIzLTQ2M2MtYTg0MS1iM2JlNmVjZDg3ZjQiLCJpc3MiOiJodHRwczovL2xvZ2luLm1pY3Jvc29mdG9ubGluZS5jb20vNDliNTE1NWUtNmU3Zi00ODY4LWE3NjAtZGYxMmIwOTg1NWZjL3YyLjAiLCJpYXQiOjE2NjU0MTYzNjUsIm5iZiI6MTY2NTQxNjM2NSwiZXhwIjoxNjY1NDIwODA0LCJhaW8iOiJBVFFBeS84VEFBQUF4NXpSc1diZ2tPY3Nqbms5RC8ydWdYMVVOL0pMYWpCVWU5amM2VXFDWXdHUUIyR1JJNEs1MTJUa1VINFhtMm9WIiwiYXpwIjoiZGQyYWFiMWUtY2IxZS00Mzk5LTlhZmUtYmZhMDMyYzAzMWM3IiwiYXpwYWNyIjoiMCIsIm5hbWUiOiJOYXR0aGFwb25nIENob29zdXAiLCJvaWQiOiIzYTQwMDZmZC01NjFiLTQ4M2EtYmQ4My1lNDM3MDU4NGZlMzgiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJuYXR0aGFwb25nLmNob0Bra3B0ZXN0Lm9ubWljcm9zb2Z0LmNvbSIsInJoIjoiMC5BVllBWGhXMVNYOXVhRWluWU44U3NKaFZfRkppMEVlemZ6eEdxRUd6dm03TmhfUldBS0UuIiwic2NwIjoiYWNjZXNzX2FzX3VzZXIiLCJzdWIiOiJhTHBNaDl5cng0UmFpN0M5RDUyenVVemhFaWpYNmhlNjNrUzR1ZkViNWdzIiwidGlkIjoiNDliNTE1NWUtNmU3Zi00ODY4LWE3NjAtZGYxMmIwOTg1NWZjIiwidXRpIjoiTU92SHBNeGN6VWloT0ZTUHY4WTJBQSIsInZlciI6IjIuMCJ9.V2gvns2DGqJqTwnrPqq6c0ugLt2ZaHAHkXqnJUAwDNpeH5wl5bwOcY-RU_ehGOf_Uc53Tig_TI7PLk71Zl0-M65zC3dKHA2Ln6xKE5BqNCiDd5bhwLvM2gAeKi1TUkQVbfRrQ8I4lseJoLIMf8-fZUn8vCe04JrJkPz2aGnFQNNCn4vdjB8Y1a_0ts7zosJVG--OTMlaHF28XG-8NPAzdDwBjfwXE-Ps_6J3h4-QMsfum5TV6mxUHlvBreWREYkJflLkPrTGFwF11cMMfOs8etggQl-uc89td-NBP72F9aCPlGyhTpjGiG01s0JMnHWLlQ725Q-o0IVn9xS6YPkPkQ"

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
