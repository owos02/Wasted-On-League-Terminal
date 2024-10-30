package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("---Wasted-On-League-Terminal---")

	accountLink := checkArgs()
	fmt.Println(accountLink)

	html_content, ok := getWebsiteInformation(accountLink)

	if !ok {
		fmt.Println("Could not establish connection. Exiting...")
		return
	}

	time, ok := getTimeFromInformation(html_content)

	if !ok {
		fmt.Println("Could not parse out the time. Exiting...")
		return
	}

	printResults(time)
}

func checkArgs() string {
	programArguments := os.Args[1:]

	if len(programArguments) == 0 || os.Args[1] == "--help" {
		fmt.Println("\nNo link or username entered\n")
		fmt.Println("Usage: Wasted-On-League-Terminal [Username#GameTag|WoL-LINK] [Server]")
		fmt.Println("Fetches time played in hours from WoL.gg.\n")
		fmt.Println("Examples:")
		fmt.Println("./Wasted-On-League-Terminal \"the inescapable#EUW\" \"EUW\"")
		fmt.Println("./Wasted-On-League-Terminal \"https://wol.gg/stats/euw/theinescapable-euw/\"")
		os.Exit(0)
	}

	url, _ := url.Parse(os.Args[1])

	if url.Scheme != "" {
		return os.Args[1]
	}

	newUsername := strings.Replace(os.Args[1], "#", "-", -1)

	toJoin := []string{"https://wol.gg/stats/", os.Args[2], "/", strings.ReplaceAll(newUsername, " ", ""), "/"}

	userLink := strings.Join(toJoin, "")

	return userLink
}

func getWebsiteInformation(accountLink string) (io.ReadCloser, bool) {

	type return_type struct {
		data io.ReadCloser
		ok   bool
	}

	var result return_type
	result.data = nil
	result.ok = true

	response, _ := http.Get(accountLink)

	if response.Status != "200 OK" {
		result.ok = false
	}

	result.data = response.Body

	return result.data, result.ok
}

func getTimeFromInformation(html_context io.ReadCloser) (time string, err bool) {

	type return_type struct {
		time string
		ok   bool
	}

	var result return_type
	result.time = ""
	result.ok = true

	data := html.NewTokenizer(html_context)

	for {
		if data.Next() == html.ErrorToken {

			if result.time == "" {
				result.ok = !result.ok
			}
			break
		}

		node := data.Token()
		toProcess := node.Attr

		for _, attribute := range toProcess {
			if attribute.Key == "id" && attribute.Val == "time-hours" {

				data.Next()
				data.Next()
				node := data.Token()

				result.time = node.Data
				break
			}
		}
	}

	return result.time, result.ok
}

func printResults(time string) {
	fmt.Println("Time (h): ", time)
}
