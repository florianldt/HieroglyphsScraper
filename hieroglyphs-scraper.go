package main

import (
	"fmt"
	"regexp"
	"encoding/json"
	"io"
   	"os"

	"github.com/gocolly/colly"
)

type Category struct {
	Id 		string	`json:"id"`
	Name	string	`json:"name"`
}

type Hieroglyph struct {
	Id 				string		`json:"id"`
	CategoryId 		string 		`json:"category_id"`
	Unicode			string		`json:"unicode"`
	Description 	string 		`json:"description"`
	Transliteration string 		`json:"transliteration"`
	Phonetic 		string 		`json:"phonetic"`
	Note 			string 		`json:"note"`
}

func writeCategoriesJson(categories []Category) {

	jsonFile, err := os.Create("./json/categories.json")

   	if err != nil {
    	fmt.Println("Error creating JSON file:", err)
      	return
	}

	jsonWriter := io.Writer(jsonFile)
	encoder := json.NewEncoder(jsonWriter)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&categories)
	
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}

func writeHieroglyphsJson(hieroglyphs []Hieroglyph) {

	jsonFile, err := os.Create("./json/hieroglyphs.json")

   	if err != nil {
    	fmt.Println("Error creating JSON file:", err)
      	return
	   }
	   
	jsonWriter := io.Writer(jsonFile)
	encoder := json.NewEncoder(jsonWriter)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&hieroglyphs)

	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}

func getGardinerCategories() {

	var categories []Category

	c := colly.NewCollector()

	c.OnHTML("li.tocsection-1 li .toctext", func(e *colly.HTMLElement) {
		re := regexp.MustCompile("([A-z]+)")
		category := Category {
			Id: re.FindStringSubmatch(e.Text)[1],
			Name: e.Text,
		}
		categories = append(categories, category)
	})

	c.OnScraped(func(r *colly.Response) {
		writeCategoriesJson(categories)
	})

	c.Visit("https://en.wikipedia.org/wiki/Gardiner%27s_sign_list")
}

func getHieroglyphs() {
	
	var hieroglyphs []Hieroglyph

	c := colly.NewCollector()

	c.OnXML("//*[@id='mw-content-text']/div/table[2]/tbody/tr[position()>1]", func(e *colly.XMLElement) {
		re := regexp.MustCompile("([A-z]+)")
		hieroglyph := Hieroglyph {
			Id: e.ChildText("/td[2]"),
			CategoryId: re.FindStringSubmatch(e.ChildText("/td[2]"))[1],
			Unicode: e.ChildText("/td[3]"),
			Description: e.ChildText("/td[4]"),
			Transliteration: e.ChildText("/td[5]"),
			Phonetic: e.ChildText("/td[6]"),
			Note: e.ChildText("/td[7]"),
		}
		hieroglyphs = append(hieroglyphs, hieroglyph)
	})

	c.OnScraped(func(r *colly.Response) {
		writeHieroglyphsJson(hieroglyphs)
	})

	c.Visit("https://en.wikipedia.org/wiki/List_of_Egyptian_hieroglyphs")
}

func main() {

	getGardinerCategories()
	getHieroglyphs()
}