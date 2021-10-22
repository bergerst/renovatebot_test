package main

import (
 "fmt"
 "github.com/go-http-utils/headers"
 "github.com/tidwall/gjson"
 "google.golang.org/genproto/googleapis/type/month"
)


func main() {
 unspecified := month.Month_MONTH_UNSPECIFIED
 language := headers.AcceptLanguage
 json := gjson.JSON
 fmt.Printf("%v-%v-%v", unspecified, language, json)
}
