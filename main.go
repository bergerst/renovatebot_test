package main

import (
 "fmt"
 "github.com/antlr/antlr4/runtime/Go/antlr"
 "github.com/go-http-utils/headers"
 "google.golang.org/genproto/googleapis/type/month"
)


func main() {
 unspecified := month.Month_MONTH_UNSPECIFIED
 language := headers.AcceptLanguage
 invalidType := antlr.TokenInvalidType
 fmt.Printf("%v-%v-%v", unspecified, language, invalidType)
}
