package main

import (
	"cloud.google.com/go/translate"
	"context"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)
type Format string

type Language struct {
	// Name is the human-readable name of the language.
	Name string

	// Tag is a standard code for the language.
	Tag language.Tag
}

type Detection struct {
	// Language is the code of the language detected.
	Language language.Tag

	// Confidence is a number from 0 to 1, with higher numbers indicating more
	// confidence in the detection.
	Confidence float64

	// IsReliable indicates whether the language detection result is reliable.
	IsReliable bool
}
type Options struct {
	// Source is the language of the input strings. If empty, the service will
	// attempt to identify the source language automatically and return it within
	// the response.
	Source language.Tag

	// Format describes the format of the input texts. The choices are HTML or
	// Text. The default is HTML.
	Format Format

	// The model to use for translation. The choices are "nmt" or "base". The
	// default is "base".
	Model string
}
type Translation struct {
	// Text is the input text translated into the target language.
	Text string

	// Source is the detected language of the input text, if source was
	// not supplied to Client.Translate. If source was supplied, this field
	// will be empty.
	Source language.Tag

	// Model is the model that was used for translation.
	// It may not match the model provided as an option to Client.Translate.
	Model string
}


type Country struct {
	Code string `json:"English"`
	Name string `json:"alpha2"`
}

type Traduccion struct {
	Code string
	Text string
}

func createClientWithKey() {
	ctx := context.Background()

	const apiKey = "AIzaSyCkA6lmPa0eATWDgBS4FYDSD0CPMaL60nI"
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Translate(ctx, []string{"Hello, world!"}, language.Russian, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", resp)
}

func translateToText(targetLanguage, text string) string {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "Lenguaje no soportado !"
	}

	const apiKey = "AIzaSyCkA6lmPa0eATWDgBS4FYDSD0CPMaL60nI"
	client, err := translate.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "Lenguaje no soportado !"
	}

	return resp[0].Text
}

func detectLanguage(text string) (*translate.Detection, error) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	lang, err := client.DetectLanguage(ctx, []string{text})
	if err != nil {
		return nil, err
	}

	return &lang[0][0], nil
}

func listSupportedLanguages(w io.Writer, targetLanguage string) error {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return err
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	langs, err := client.SupportedLanguages(ctx, lang)
	if err != nil {
		return err
	}

	for _, lang := range langs {
		fmt.Fprintf(w, "%q: %s\n", lang.Tag, lang.Name)
	}

	return nil
}

func translateTextWithModel(targetLanguage, text, model string) (string, error) {
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", err
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, &translate.Options{
		Model: model, // Either "mnt" or "base".
	})
	if err != nil {
		return "", err
	}
	return resp[0].Text, nil
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENTER HOME")
	t1, _ := template.ParseFiles("templates/home.html")
	t1.Execute(w, nil)
}

func getContries(w http.ResponseWriter, r *http.Request) {

	response, err := http.Get("https://pkgstore.datahub.io/core/language-codes/language-codes_json/data/97607046542b532c395cf83df5185246/language-codes_json.json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data []Country
	json.Unmarshal([]byte(responseData), &data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}


func getTranslate(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query().Get("text")
	l := strings.ToLower(r.URL.Query().Get("lang"))

	var t Traduccion
	t.Code = l
	t.Text = translateToText(l, q)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func handleRequests() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/getContries", getContries)
	http.HandleFunc("/getTranslate", getTranslate)

	port := ":9000"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, nil)
}

func main() {
	handleRequests()
}
