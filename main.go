package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	baseUlrCountries := "https://restcountries.eu"
	// baseUlrGoogle := "https://translation.googleapis.com"

	response, err := http.Get(baseUlrCountries + "/rest/v2/alpha/col")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}
