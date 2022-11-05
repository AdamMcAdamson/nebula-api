package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"log"
	"strings"
	"time"

	. "github.com/UTDNebula/nebula-api/scraper/degree-scraper/model"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
)

// write to file in pretty JSON
func PrettyStruct(data interface{}, filepath string) (string, error) {
	// marshal to json
	val, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return "", err
	}

	// get string equivalent
	str := string(val)

	// fix core requirement formatting
	regexp_arr, _ := regexp.Compile(`\n\s*"\^",*\n\s*`)
	str2 := regexp_arr.ReplaceAllString(str, "")
	str = strings.Replace(str2, ",]", "]", -1)

	// write to file
	file, err := os.Create("./output/" + filepath + ".json")
	if err != nil {
		return "", err
	}

	if _, err := file.Write([]byte(str)); err != nil {
		return "", err
	}

	return str, nil
}

// get core flag
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

// read catalog_uris.json
func getCatelogUris(year string) []string {
	// @TODO error handling
	content, _ := os.ReadFile("./catalog_uris.json")

	var uris_og []interface{}
	var end_uris []string
	json.Unmarshal(content, &uris_og)

	uris := uris_og[0].([]interface{})

	for _, val := range uris {
		end_uris = append(end_uris, val.(string))
	}

	// handle edge-cases
	for _, uri := range uris_og[1].([]interface{}) {
		if val, ok := uri.(map[string]interface{})[year]; ok {
			end_uris = append(end_uris, val.(string))
			//fmt.Printf("year: %s, val: %s\n", year, val.(string))
		} else if val, ok := uri.(map[string]interface{})[year]; ok {
			end_uris = append(end_uris, val.(string))
			//fmt.Printf("year: else, val: %s\n", uri.(map[string]interface{})["else"].(string))
		}
	}

	return end_uris
}

// get school from uri shorthand
func getSchool(in string) string {
	switch in {
	case "ah":
		return "School of Arts and Humanities"
	case "atec":
		return "School of Arts, Technology, and Emerging Communication"
	case "bbs":
		return "School of Behavioral and Brain Sciences"
	case "epps":
		return "School of Economic, Political and Policy Sciences"
	case "ecs":
		return "Erik Jonsson School of Engineering and Computer Science"
	case "is":
		return "School of Interdisciplinary Studies"
	case "jsom":
		return "Naveen Jindal School of Management"
	case "nsm":
		return "School of Natural Sciences and Mathematics"
	default:
		return "unknown school"
	}
}

func main() {
	// list of years to scrape
	years := []int{2022, 2021, 2020, 2019}

	// course/core requirements (and some garbage)
	var res []string

	// degree name and abbreviation
	var degree_res []string

	var degree_res_verification []string

	// minimum credit hours
	var hours_res string

	// selector for coursesJS
	sel1 := "p > a"

	// JavaScript for getting inner text of all courses/core entries (and some extra garbage)
	coursesJS := fmt.Sprintf(`[...document.querySelectorAll('%s')].map((e) => e.innerText)`, sel1)

	// selector for degreeJS
	sel2 := "h2"

	// JavaScript for getting inner text of degree name and abbreviation
	degreeJS := fmt.Sprintf(`[...document.querySelectorAll('%s')].map((e) => e.innerText)`, sel2)

	// Get innerHTML of h2 elements
	degree_valJS := fmt.Sprintf(`[...document.querySelectorAll('%s')].map((e) => e.innerHTML)`, sel2)

	// selector for element whose text includes the minimum credit hours
	hours_sel := "p.cat-degh"

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 600*time.Second)
	defer cancel()

	for _, yearInt := range years {
		year := fmt.Sprintf("%d", yearInt)
		uris := getCatelogUris(year)

		// create folder for given year as needed
		os.Mkdir("./output/"+year, 0666)

		for _, uri := range uris {

			start := time.Now()

			var splitUri = strings.Split(uri, "/")

			// file path
			var filepath string = (year + "/" + splitUri[len(splitUri)-1])

			err := chromedp.Run(ctx,
				emulation.SetUserAgentOverride("Degree-Scraper v1.0"),

				// Navigate calalog site for given year and uri
				chromedp.Navigate(`https://catalog.utdallas.edu/`+year+"/"+uri),

				// Evaluate coursesJS
				chromedp.Evaluate(coursesJS, &res),

				// Evaluate degreeJS
				chromedp.Evaluate(degreeJS, &degree_res),

				// Evaluate degree_valJS
				chromedp.Evaluate(degree_valJS, &degree_res_verification),

				// Get credit hours element text
				chromedp.TextContent(hours_sel, &hours_res),
			)

			if err != nil {
				log.Fatal(err)
			}

			//	var requirements []Requirement
			var requirements []interface{}

			// extract minimum credit hours
			var re = regexp.MustCompile(`\d+`)
			hours, _ := strconv.Atoi(string(re.Find([]byte(hours_res))))

			// extract degree name and abbreviation
			var name string
			var abbreviation string
			if len(degree_res) >= 3 {
				if strings.Contains(degree_res_verification[0], "<em>") {
					// first h2 is not the abbreviation
					name = degree_res[2]
					abbreviation = degree_res[1]
				} else {
					name = degree_res[1]
					abbreviation = degree_res[0]
				}
			} else if len(degree_res) == 2 {
				name = degree_res[1]
				abbreviation = degree_res[0]
			} else if len(degree_res) == 1 {
				name = degree_res[0]
				abbreviation = degree_res[0]
			} else {
				name = "MISSING"
				abbreviation = "MISSING"
			}

			// degree requirement heading
			requirements = append(requirements, map[string]interface{}{
				"subtype":              "Major",
				"school":               getSchool(splitUri[len(splitUri)-2]),
				"year":                 year + "-" + fmt.Sprintf("%d", yearInt+1),
				"abbreviation":         abbreviation,
				"name":                 name,
				"minimum_credit_hours": hours,
				"catalog_uri":          `https://catalog.utdallas.edu/` + year + "/" + uri,
			})

			// parse and include course/core requirements
			for _, val := range res {
				temp := strings.Split(val, " ")
				if len(temp) == 2 && len(temp[1]) == 4 && temp[1] != "Core" {
					// Course
					fmt.Printf("%s is a course\n", temp)

					// append course to requirements array (in ndt format)
					requirements = append(requirements, temp[0]+" "+temp[1])
				} else if temp[len(temp)-1] == "Core" {
					// Core
					fmt.Printf("CORE: %s is a core requirement\n", temp)

					coreReq := GetCoreRequirement(strings.Join(temp, " "))
					coreArr := []interface{}{"^", "core", "^", coreReq.CoreFlag, "^", coreReq.Hours, "^"}

					// append core requirement to requirements array (in ndt format ("^" are used to adjust formatting))
					requirements = append(requirements, coreArr)
				} else {
					// Garbage
					fmt.Printf("------'%s' is NOT a course\n", val)
				}

			}

			out, _ := PrettyStruct(requirements, filepath)

			fmt.Printf("Degree Requirement (%s(%s)): \n%s\n", year, degree_res[0], out)

			fmt.Printf("\nTook: %f secs\n", time.Since(start).Seconds())
		}
	}
}
