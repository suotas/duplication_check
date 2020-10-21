package main

import (
    "os"
    "bufio"
	"fmt"
	"regexp"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("data/*")
	results := make(map[string][]string)

	for _, v := range files {
		rep := regexp.MustCompile(`.c$`)
		filename := filepath.Base(rep.ReplaceAllString(v, ""))
		results[filename] = readFile(v)
	}

	finalResult := make(map[string][]string)

	for k, v := range results {
		// fmt.Println(k)
		for _, v2 := range v {
			// fmt.Println(v)
			for k3, v3 := range results {
				if k != k3 {
					for _, v4 := range v3 {
						if v2 == v4 {
							finalResult[v4] = append(finalResult[v4], k3)
						}
					}
				}
			}
		}
	}

	for k, v := range finalResult {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func readFile(filename string) []string{
	var result []string
    fp, err := os.Open(filename)
    if err != nil {
		panic(0)
    }
    defer fp.Close()

    scanner := bufio.NewScanner(fp)

    for scanner.Scan() {
		result = append(result, scanner.Text())
    }

    if err = scanner.Err(); err != nil {
        panic(0)
	}
	
	return result
}