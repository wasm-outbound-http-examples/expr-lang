package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/expr-lang/expr"
)

func main() {
	env := map[string]interface{}{
		"httpget": func(url string) string {
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			return string(body)
		},
		"println": fmt.Println,
	}

	code := `println(httpget('https://httpbin.org/anything'))`

	_, err := expr.Eval(code, env)
	if err != nil {
		panic(err)
	}
}
