package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/dyatlov/go-oembed/oembed"
	"github.com/dyatlov/go-url2oembed/url2oembed"
)

func main() {
	providersData, err := ioutil.ReadFile("providers.json")

	if err != nil {
		panic(err)
	}

	oe := oembed.NewOembed()
	oe.ParseProviders(bytes.NewReader(providersData))

	parser := url2oembed.NewParser(oe)

	data := parser.Parse("http://youtube.com")

	fmt.Printf("%s\n", data)
}
