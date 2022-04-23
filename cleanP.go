package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	cont := []string{}
	uniqPath := []string{}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {

		line := sc.Text()
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
