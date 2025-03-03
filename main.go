package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	sqlWithParams, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Could not read clipboard", err)
		os.Exit(1)
	}

	fmt.Print("\n Is it you sql you would like to substitute?\n\n")
	fmt.Println(sqlWithParams)
	fmt.Print("\n")

	if AskQuestion("Y/n:") {
		fmt.Println("Ok, let's go")
	} else {
		fmt.Println("Then copy another stuff and run me again")
		os.Exit(0)
	}

	fmt.Println("\n Got your params, thanks", sqlWithParams)
	fmt.Print("\n\n\n\n")

	paramsMarker := "-- PARAMETERS: "
	paramsStart := strings.Index(sqlWithParams, paramsMarker)
	paramsString := sqlWithParams[paramsStart+len(paramsMarker):]

	var params []interface{}

	if err := json.Unmarshal([]byte(paramsString), &params); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	fmt.Println("Params found", params)

	sqlWithParams = sqlWithParams[:paramsStart]

	for i, param := range params {
		if parsed, ok := param.(string); ok {
			params[i] = `'` + parsed + `'`
		}

		sqlWithParams = strings.Replace(sqlWithParams, fmt.Sprintf("$%d", i+1), fmt.Sprint(params[i]), -1)
	}

	fmt.Print("\n\n", sqlWithParams, "\n\n")
}
