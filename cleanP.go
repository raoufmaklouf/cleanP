package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var static_exts = []string{"js", "css", "png", "jpg", "jpeg", "svg", "ico", "webp", "ttf", "otf", "woff", "gif", "pdf", "bmp", "eot", "mp3", "woff2", "mp4", "avi"}

func main() {
	cont := []string{}
	uniqPath := []string{}
	deepcheklist := []string{}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {

		line := sc.Text()
		regex1, _ := regexp.MatchString("[?]", line)
		regex2, _ := regexp.MatchString("=", line)
		if regex1 == true && regex2 == true {
			if isStatic(line) == false {
				urlInjected := injectparam(line)
				pathKey := uniqurl(line)
				deepChekey := deepchek(line)
				if contains(cont, urlInjected) == false {
					if contains(uniqPath, pathKey) == false {
						if contains(deepcheklist, deepChekey) == false {
							deepcheklist = append(deepcheklist, deepChekey)
							uniqPath = append(uniqPath, pathKey)
							cont = append(cont, urlInjected)
							fmt.Println(line)
						}
					}

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
func deepchek(url string) (finalkey string) {
	var key string
	baseul := strings.Split(url, "?")[0]
	domain := strings.Split(url, "/")[2]
	parameters := strings.Split(url, "?")[1]
	regex2, _ := regexp.MatchString("&", parameters)
	key = domain + ":"
	if regex2 == true {
		for _, i := range strings.Split(parameters, "&") {
			paramkey := strings.Split(i, "=")[0]
			key = key + paramkey + ":"

		}
	} else {
		paramkey := strings.Split(parameters, "=")[0]
		key = key + paramkey + ":"

	}
	for _, Path := range strings.Split(baseul, "/") {
		key = key + strconv.Itoa(len(Path)) + ":"

	}
	finalkey = key
	return

}

func isStatic(url string) bool {
	Url := strings.Split(url, "?")[0]
	var static bool = false
	for _, ext := range static_exts {
		if strings.HasSuffix(Url, ext) == true {
			static = true
			break
		}
	}
	return static
}
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
			break
		}
	}

	return false
}
