package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning Map")
	lang := make(map[string]string)
	lang["JS"] = "javascript"
	lang["TS"] = "Typescript"
	lang["py"] = "python"

	fmt.Println(len(lang), lang, "type of lang map %T d", lang)

	for index, val := range lang {
		fmt.Printf("index is %v and value is %v  ", index, val)
	}

	fmt.Printf("Type %T", lang)

}



