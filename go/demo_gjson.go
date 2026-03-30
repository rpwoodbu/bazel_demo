package main

import "github.com/tidwall/gjson"

func compute_hello(json string) string {
	greeting := gjson.Get(json, "greet").String()
	return "Hello, " + greeting + "!"
}
