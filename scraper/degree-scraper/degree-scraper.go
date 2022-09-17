package main

import (
	"encoding/json"
	"context"
	"fmt"
	//"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	"log"
	"time"
	"strings"
	. "github.com/UTDNebula/nebula-api/scraper/degree-scraper/model"
)

func PrettyStruct(data interface{}) (string, error) {
    val, err := json.MarshalIndent(data, "", "    ")
    if err != nil {
        return "", err
    }
    return string(val), nil
}

func GetCoreRequirement(in string) CoreRequirement {
	switch in {
		case "Communication Core":
			return CoreRequirement{"010", 6}
		case "Mathematics Core":
			return CoreRequirement{"020", 3}
		case "Life and Physical Sciences Core":
			return CoreRequirement{"030", 6}
		case "Language, Philosophy and Culture Core":
			return CoreRequirement{"040", 3}
		case "Creative Arts Core":
			return CoreRequirement{"050", 3}
		case "American History Core":
			return CoreRequirement{"060", 6}
		case "Government/Political Science Core":
			return CoreRequirement{"070", 6}
		case "Social and Behavioral Sciences Core":
			return CoreRequirement{"080", 3}
		case "Component Area Option Core":
			return CoreRequirement{"090", 6}
		default:
			return CoreRequirement{in, -1}  
	}
}

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
	
	// selector for js
	sel1 := "p > a"
	
	// JavaScript for getting inner text of all elements meeting the selector sel1
	js := fmt.Sprintf(`[...document.querySelectorAll('%s')].map((e) => e.innerText)`, sel1)


	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride("Degree-Scraper 0.1"),

		// @TODO: read catelog-uris.txt (or programmatically navigate site)
		chromedp.Navigate(`https://catalog.utdallas.edu/2022/undergraduate/programs/ah/literature`),
			
		// Evaluate js
		chromedp.Evaluate(js, &res),
	)

	if err != nil {
		log.Fatal(err)
	}


	
	var requirements []Requirement
	
	
	for _, val := range res {
		temp := strings.Split(val," ")
		if len(temp) == 2 && len(temp[1]) == 4 && temp[1] != "Core" {
			// Course
			fmt.Printf("%s is a course\n", temp)
			requirements = append(requirements, TempCourse{temp[0], temp[1], "D-"})
		} else if temp[len(temp)-1] == "Core" {
			// Core
			fmt.Printf("CORE: %s is a core requirement\n", temp)
			requirements = append(requirements, GetCoreRequirement(strings.Join(temp, " ")))
		} else {
			// Garbage
			fmt.Printf("------'%s' is NOT a course\n", val)
		}
	}
	
	degRequirement := CollectionRequirement{"Degree Requirement", -1, requirements}
	
	out, err := PrettyStruct(degRequirement)

	fmt.Printf("Degree Requirement: \n%s\n", out)
	
	fmt.Printf("\nTook: %f secs\n", time.Since(start).Seconds())

}