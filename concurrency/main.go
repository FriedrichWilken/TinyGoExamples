package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// The morining routine: drink coffee and watch websites

	urls := []string{
		"https://www.tagesspiegel.de/",
		"https://www.theguardian.com/",
		"https://www.newyorker.com/",
		"https://github.com/",
		"https://twitter.com/",
		"https://www.golem.de/",
		"https://www.youtube.com/",
		"https://en.wikipedia.org/",
		"https://golang.org/",
		"https://www.qwant.com/",
		"https://dev.to/",
	}
	var wg sync.WaitGroup

	// lets open up all the websites at the same time ...
	for _, url := range urls {
		wg.Add(1)
		go getNews(&wg, url)
	}
	fmt.Println("Waiting for the websites...")

	// ...and at the same time, lets make coffee
	c1 := make(chan string)
	c2 := make(chan string)

	wg.Add(1)
	go drinkCoffee(c2, &wg) // really we only want to drink coffee...
	go makeCoffee(c1)       // ...unfortunetly we have to wait until its done ...
	go coffeeToCup(c1, c2)  // ...and poured to the cup.

	wg.Wait()
	fmt.Println("Done.")
}

func getNews(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Printf("page '%v', status code: %v\n", response.Request.URL, response.StatusCode)
}

func makeCoffee(c chan string) {
	fmt.Println("let's make coffee!")
	time.Sleep(time.Second * 1)
	c <- "coffee is done"
}
func coffeeToCup(c1, c2 chan string) {
	fmt.Println(<-c1)
	fmt.Println("pouring coffee in")
	time.Sleep(time.Second * 1)
	c2 <- "coffee is in the cup"
}
func drinkCoffee(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-c)
	fmt.Println("drinking coffee..")
	time.Sleep(time.Second * 1)
	fmt.Println("yummy")
}
