package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cast"
	"github.com/thoas/go-funk"
)

var str = `{
        "tyAccount": "17720206785",
        "bookList": [
            {
                "id": 6328899,
                "name": "361度 光谷天地店",
                "mobile": "86643525"
            },
            {
                "id": 6334878,
                "name": "陈鹏",
                "mobile": "02783613231"
            },
            {
                "id": 6323397,
                "name": "陈鹏",
                "mobile": "13669099006"
            },
            {
                "id": 6326610,
                "name": "邓先生回家",
                "mobile": "10000"
            },
            {
                "id": 6327468,
                "name": "邓先生",
                "mobile": "10000"
            },
            {
                "id": 6328906,
                "name": "889898898988988",
                "mobile": "889898898988988"
            },
            {
                "id": 6328908,
                "name": "9",
                "mobile": "9"
            },
            {
                "id": 6328907,
                "name": "9",
                "mobile": "5585868686"
            },
            {
                "id": 6328909,
                "name": "91",
                "mobile": "44577575"
            },
            {
                "id": 6334870,
                "name": "摸摸你",
                "mobile": "+.---.---."
            }
        ]
}`

var idStr = "6328909,6327468,6327468"

type ContactQueryResult struct {
	TyAccount string            `json:"tyAccount"`
	BookList  []ContactAddrBook `json:"bookList"`
}

type ContactAddrBook struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

func main() {
	s1 := []string{"a", "b", "c", "d"}
	s2 := s1

	s2[0] = "Sun"

	fmt.Println(s1) // Output: [Sun Tue Wed Thu Fri Sat Sun]
	fmt.Println(s2) // Output: [Sun Tue]

	var list []string = strings.Split(idStr, ",")
	temp := funk.Uniq(list)
	fmt.Println(temp.([]string))

	// deal()
}

func deal() {

	var data ContactQueryResult
	err := json.Unmarshal([]byte(str), &data)
	fmt.Println(err, data)

	temp := removeElement(data.BookList, idStr)

	data.BookList = temp
	buf, err := json.Marshal(data)
	fmt.Println(string(buf))
}

func removeElement(s []ContactAddrBook, ids string) []ContactAddrBook {
	var result []ContactAddrBook
	var idArr []string = strings.Split(ids, ",")
	for _, v := range s {
		if !contains(idArr, cast.ToString(v.ID)) {
			result = append(result, v)
		}
	}
	return result
}

func contains(s []string, elem string) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}
