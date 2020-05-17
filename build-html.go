package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type jsonData []struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Value       string   `json:"value"`
	Img         string   `json:"img"`
	Notables    []string `json:"notables"`
}

func main() {

	file, err := ioutil.ReadFile("data.json")
	data := jsonData{}

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		log.Fatal(err)
	}

	middleHTML, err := os.OpenFile("blocks/index-middle.tmpl",
		os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer middleHTML.Close()

	for i, card := range data {

		notableHTML := "<ul>"
		for _, notable := range card.Notables {
			notableHTML += fmt.Sprintf("<li>%s</li>", notable)
		}
		notableHTML += "</ul>"

		if len(card.Notables) == 0 {
			notableHTML = ""
		}

		mod := i % 2

		html := ""
		if mod == 0 {
			html += `<div class="row">`
		}

		html += fmt.Sprintf(`		
		<div class="col-sm-6">
			<div class="card h-100">
				<img class="mx-auto d-block" height="150px" src="%s">    
				<div class="card-body">
					<h5 class="card-title">%s</h5>
					%s												
					<p class="card-text">%s</p> 
					<p><i>%s</i></p>  
				</div>
			</div>
		</div>`, card.Img, card.Title, notableHTML, card.Description, card.Value)

		if mod == 1 || len(data) == i {
			html += `</div>`
		}

		if _, err := middleHTML.WriteString(html); err != nil {
			log.Fatal(err)
		}

	}

	writeOutput()

}

func writeOutput() {
	files := []string{"blocks/index-top.tmpl", "blocks/index-middle.tmpl", "blocks/index-bottom.tmpl"}
	var buf bytes.Buffer
	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		buf.Write(b)
	}

	err := ioutil.WriteFile("public/index.html", buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Writen index.html")
}
