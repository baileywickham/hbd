package main

import (
	"fmt"

	"github.com/plivo/plivo-go"
)

var client *plivo.Client

func sendText(dst, message string) {
	resp, err := client.Messages.Create(plivo.MessageCreateParams{
		Src:  "19712063040",
		Dst:  dst,
		Text: message,
	})
	fmt.Printf("%+v\n", resp)
	if err != nil {
		panic(err)
	}

}

func init() {
	var err error
	client, err = plivo.NewClient("", "", &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}
}
