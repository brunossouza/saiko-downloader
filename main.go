package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// Anime type creation
type Anime struct {
	name    string
	epsodio string
	link    string
}

var (
	animesList = []Anime{}
)

func addAnimeToList(element *colly.HTMLElement) {
	name := element.ChildText("div[class='post-name']")

	if name != "" {
		epsodio := strings.Split(element.ChildText("div[class='post-ep']"), " ")[1]
		link := element.ChildAttr("a", "href")

		animesList = append(animesList, Anime{name: name, epsodio: epsodio, link: link})
	}
}

func getCollyConfig() *colly.Collector {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	return c
}

func readOpcao() (opcao int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nPara sair digite: sair(s)\n\nSelecione uma opção: ")
	scanner.Scan()
	var opcaoTxt = scanner.Text()
	if opcaoTxt == "s" || opcaoTxt == "sair" {
		fmt.Println("saindo...")
		os.Exit(1)
	}
	opcao, err := strconv.Atoi(opcaoTxt)
	if err != nil {
		log.Fatalln("Opção inválida, error:", err)
	}
	return
}

func getAnimeLastUpdates() {
	c := getCollyConfig()

	c.OnHTML("#content", func(e *colly.HTMLElement) {
		addAnimeToList(e)
	})
	c.Visit("https://saikoanimes.net/")
}

func getAnimePage(anime Anime) {
	c := getCollyConfig()

	c.OnHTML("div[class='ani-titulo']", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println(string(r.Body))
	// })

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(anime.link)
}

func menuHandler() {
	fmt.Print("---------------------------------------------------\n")
	fmt.Print("|Anime\t|Episodio\t|Nome\t\n")
	fmt.Print("---------------------------------------------------\n")
	for idx, a := range animesList {
		fmt.Print("|", idx, "\t|", a.epsodio, "\t\t|", a.name, "\n")
	}

	var opcao = readOpcao()

	var anime = animesList[opcao]

	getAnimePage(anime)
}

func main() {
	getAnimeLastUpdates()
	menuHandler()
}
