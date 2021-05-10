// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"gopl.io/ch4/github"
)

//!+
func main() {
	var dateSorted = flag.Bool("s", false, "Sort issues by creation date.")
	flag.Parse()
	query := flag.Args()
	var now, dayAgo, weekAgo, monthAgo, yearAgo time.Time
	if *dateSorted {
		query = append(query, "&sort=created")
		now = time.Now()
		dayAgo = now.AddDate(0, 0, -1)
		weekAgo = now.AddDate(0, 0, -7)
		monthAgo = now.AddDate(0, -1, 0)
		yearAgo = now.AddDate(-1, 0, 0)
	}
	result, err := github.SearchIssues(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	timeCategories := make(map[time.Time]bool)
	for _, item := range result.Items {
		if *dateSorted {
			switch {
			case item.CreatedAt.Before(yearAgo):
				_, ok := timeCategories[yearAgo]
				if !ok {
					timeCategories[yearAgo] = true
					fmt.Println("### Older than a year")
				}

			case item.CreatedAt.Before(monthAgo):
				_, ok := timeCategories[monthAgo]
				if !ok {
					timeCategories[monthAgo] = true
					fmt.Println("### Older than a month")
				}
			case item.CreatedAt.Before(weekAgo):
				_, ok := timeCategories[weekAgo]
				if !ok {
					timeCategories[weekAgo] = true
					fmt.Println("### Older than a week")
				}
			case item.CreatedAt.Before(dayAgo):
				_, ok := timeCategories[dayAgo]
				if !ok {
					timeCategories[dayAgo] = true
					fmt.Println("### Older than a day")
				}
			default:
				_, ok := timeCategories[now]
				if !ok {
					timeCategories[now] = true
					fmt.Println("### Todays events")
				}
			}
		}
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
