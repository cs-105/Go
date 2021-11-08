package main

/*
func main() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	// Find and visit all links
	/*c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	}) */ /*

	c.OnHTML("head", func(e *colly.HTMLElement) {
		text := e.Text
		fmt.Print(text)
	})

	c.OnHTML("section", func(e *colly.HTMLElement) {
		text := e.Text
		fmt.Print(text)
	})

	c.OnHTML("p", func(e *colly.HTMLElement) {
		text := e.Text
		fmt.Println(text)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		myString := string(response.Body[:])
		fmt.Println(myString)
	})

	c.Visit("https://noah-de.github.io")
}
*/
