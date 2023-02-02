package main

import "fmt"

func main() {
	var str = "74NtLAFqOEeAEgJ2abdNmNnIEiIcujwnAhxAdrVa3CjsUxgrEhSLcnblXT54AA/+RuNjL6pv21WLb297aPk7Ng=="
	fmt.Println(len(str))

	var val interface{} = "ddd"
	val = val.(string)

}

func call(val string) string {
	fmt.Printf("%v , %T", val, val)
	return ""
}
