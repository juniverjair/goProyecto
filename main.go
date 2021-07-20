package main

import (
	"cloud.google.com/go/translate"
	"context"
	"fmt"
	"golang.org/x/text/language"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func translateText(targetLanguage, text string) (string, error) {
	// text := "The Go Gopher is cute"
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}

func main() {

	q := `So let us begin anew--remembering on both sides that civility is not a
		sign of weakness, and sincerity is always subject to proof. Let us never
		negotiate out of fear. But let us never fear to negotiate`

	lan := `de`

	//{
	//	"q": "So let us begin anew--remembering on both sides that civility is not a
	//	sign of weakness, and sincerity is always subject to proof. Let us never
	//	negotiate out of fear. But let us never fear to negotiate.",
	//	"target": "de"
	//}

	translateText(q, lan)

	menu := `Bienvenido al traductor por ciudades`
	fmt.Println(menu)
	fmt.Println(`INGRESA TU CIUDAD ORIGEN: `)

	var origen string
	fmt.Scanln(&origen)

	fmt.Print(`INGRESA TU CIUDAD DESTINO: `)

	var destino string
	fmt.Scanln(&destino)

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
