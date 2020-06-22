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

func getAnimeLastUpdates() {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)

	c.OnHTML("#content", func(e *colly.HTMLElement) {
		addAnimeToList(e)
	})
	c.Visit("https://saikoanimes.net/")
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

func showMenu() {
	fmt.Print("---------------------------------------------------\n")
	fmt.Print("|Anime\t|Episodio\t|Nome\t\n")
	fmt.Print("---------------------------------------------------\n")
	for idx, a := range animesList {
		fmt.Print("|", idx, "\t|", a.epsodio, "\t\t|", a.name, "\n")
	}

	var opcao = readOpcao()

	var anime = animesList[opcao]
	fmt.Println(anime)
}

func main() {
	getAnimeLastUpdates()
	showMenu()
}
