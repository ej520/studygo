//map增加对象 和 删除对象
//add: m["ddd"] = "444ddd"
//del: delete(m,"aaa")
//update: m["ddd"] = "444444"
//qry: v, ok := m["ddd"]
package main

import (
	"fmt"
)

func main() {
	//create a map
	m := map[string]string{"aaa": "111aaa", "bbb": "222bbb", "ccc": "333ccc"}
	m["ddd"] = "444ddd"
	fmt.Println("Original map:")
	//print map
	for k := range m {
		fmt.Println("key is", k, "value is", m[k])
	}

	//delete an entry
	delete(m, "aaa")
	fmt.Println("Entry for aaa is deleted")
	fmt.Println("Update map")
	//print map
	for k := range m {
		fmt.Println("key is", k, "value is", m[k])
	}

	//update an entry
	m["ddd"] = "444444"
	//print the entry ddd
	v, ok := m["ddd"]
	if ok {
		fmt.Println("the entry ddd value is ", v)
	} else {
		fmt.Println("the entry ddd is not present")
	}
}
