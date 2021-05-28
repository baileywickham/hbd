package main

import (
	"fmt"

	"github.com/plivo/plivo-go"
)

const authId = "auth_id"
const authToken = "auth_token"

func main() {
	testPhloGetter()
}

func testPhloGetter() {
	phloClient, err := plivo.NewPhloClient(authId, authToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
	response, err := phloClient.Phlos.Get("phlo_id")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %#v\n", response)
}
