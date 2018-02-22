package main

import (
	"flag"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"sync"
	"strconv"
)

func main() {

	var wg sync.WaitGroup
	inputString := flag.String("input", "http://golang.org\nhttps://google.com", "url's")
	flag.Parse()

	k := strings.Split(*inputString,`\n`)
	wg.Add(len(k)) //ождаем все рутины

	var result = 0
	for _, v := range k {
		go func(s string) { //можно и вызывать без передаваемого параметра
			defer wg.Done()
			var url = ""
			var total = 0
			url, total = callGet(s)
			result = result + total
			fmt.Printf("Count for " + url + ":" + strconv.Itoa(total) + "\n")
		}(v)
	}

	wg.Wait()
	fmt.Println("Total ", result)
}


func callGet(s string) (name string, value int)  {
	r, _ := http.Get(s) //без обработки ошибок, что не соотв стандарту языка, но экономит место

	defer r.Body.Close()

	body, _ := ioutil.ReadAll(r.Body)

	bodyStr := string(body)
	name = s
	value = strings.Count(bodyStr, "Go")

	return
}
