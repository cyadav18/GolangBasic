package main

import "fmt"

func main() {
	fmt.Println("Main")
	prop := map[string]interface{}{"ID": []string{"some-ID"}, "Name": []string{"some-name"}, "Users": []string{"some-user"},
		"Applications": []string{"some-app"}}
	fmt.Println(prop)
	arryString := []string{}
	for i,v := range prop{
		fmt.Println(i,v)
		fmt.Printf("%T,%T\n",i,v)
		v1 := v.([]string)
		arryString = append(arryString,v1[0])
	}
	fmt.Println("\n\nlast",arryString)
}
