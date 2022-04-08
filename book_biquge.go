package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func StartBookOfBiQuGe() {
	getChapter()
	// getBook()
}

//获取章节信息
func getChapter() {
	url := "http://www.beqiku.com/book/176/"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	docDetail, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	docDetail.Find("#list > dl").Each(func(i int, s *goquery.Selection) {

		a := s.Find("dd:nth-child(2) > a")
		h, ok := a.Attr("href")
		if ok {
			fmt.Println(h)
		}
	})

	fmt.Println("END")
}

//根据某一章的链接，获取对应章节的内容
func getBook() {
	url := "http://www.beqiku.com/book/176/2856557.html"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	docDetail, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	text := docDetail.Find("#content").Text()
	fmt.Println(text)
}
