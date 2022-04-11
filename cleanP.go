package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	cont := []string{}
	uniqPath := []string{}
	file := os.Args[1]

	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		regex, _ := regexp.MatchString("[?]", line)
		if regex == true {
			urlInjected := injectparam(line)
			pathInjected := uniqurl(line)
			if contains(cont, urlInjected) == false {
				if contains(uniqPath, pathInjected) == false {
					uniqPath = append(uniqPath, pathInjected)
					cont = append(cont, urlInjected)
					fmt.Println(line)
				}

			}

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func uniqurl(url string) (finalKey string) {
	url = injectparam(url)
	paramLink := strings.Split(url, "?")[1]
	baseul := strings.Split(url, "?")[0]
	key11 := strings.Split(baseul, "/")
	key1 := key11[len(key11)-2]
	rootUrl := strings.Split(baseul, "/")[2]
	finalKey = rootUrl + ":" + key1 + ":" + paramLink
	return
}

func injectparam(url string) (finalResolt string) {
	paramLink := strings.Split(url, "?")[1]
	baseul := strings.Split(url, "?")[0] + "?"

	regex2, _ := regexp.MatchString("&", paramLink)
	if regex2 == true {
		paramslinks := strings.Split(paramLink, "&")
		for _, prmt := range paramslinks {
			key := strings.Split(prmt, "=")[0]
			injectedParam := key + "=FuZZrAouF&"
			baseul += injectedParam

		}
		finalResolt = strings.TrimSuffix(baseul, "&")

	} else {
		key2 := strings.Split(paramLink, "=")[0]
		injectedParam2 := key2 + "=FuZZrAouF"
		baseul += injectedParam2
		finalResolt = baseul
	}
	return
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
