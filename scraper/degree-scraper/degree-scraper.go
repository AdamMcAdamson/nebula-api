package main

import (
	"context"
	"fmt"
	//"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"log"
	"time"
	"strings"
)

func main() {

	
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	
	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	
	start := time.Now()
	var res []string
	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("Degree-Scraper 0.1"),

		// @TODO: read catelog-uris.txt (or programmatically navigate site)
		chromedp.Navigate(`https://catalog.utdallas.edu/2022/undergraduate/programs/ah/literature`),
			
		// Get inner text of all elements meeting the selector 'p > a'
		chromedp.Evaluate(`[...document.querySelectorAll('p > a')].map((e) => e.innerText)`, &res),
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, val := range res {
		temp := strings.Split(val," ")
		if len(temp) == 2 && len(temp[1]) == 4 && temp[1] != "Core" {
			// Course
			fmt.Printf("%s is a course\n", temp)
		} else if temp[len(temp)-1] == "Core" {
			// Core
			fmt.Printf("CORE: %s is a core requirement\n", temp)
		} else {
			// Garbage
			fmt.Printf("------'%s' is NOT a course\n", val)
		}
	}

	//fmt.Printf("res:\n'%s'\n", res)
	fmt.Printf("\nTook: %f secs\n", time.Since(start).Seconds())
}