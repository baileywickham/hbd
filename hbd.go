package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/emersion/go-vcard"
)

func readVCF(card string) {
	f, err := os.Open(card)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := vcard.NewDecoder(f)
	for {
		card, err := dec.Decode()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		handleVCF(card.PreferredValue(vcard.FieldFormattedName),
			card.PreferredValue(vcard.FieldTelephone),
			card.PreferredValue(vcard.FieldBirthday))
	}
}

func handleVCF(name, phone, birthday string) {
	phone = cleanPhone(phone)
	today := time.Now().Local().Format("01-02")
	if today == birthday[5:] {
		message := generateMessage(name)
		sendText(phone, message)
		println("sending: \""+message+"\" to:", phone)
	}
}

func cleanPhone(phone string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		panic(err)
	}
	phone = reg.ReplaceAllString(phone, "")
	if len(phone) == 10 {
		return "1" + phone
	}
	return phone

}
func generateMessage(name string) string {
	adj := []string{"Nice", "Funny", "Neat", "Kind", "Organized"}
	rand.Seed(time.Now().Unix())
	s := adj[rand.Intn(len(adj))]
	return fmt.Sprintf("hbd %s! you are very %s!", name, s)
}

func main() {
	if len(os.Args) != 2 {
		println("usage: hbd contact.vcf")
		os.Exit(1)
	}
	readVCF(os.Args[1])
}
