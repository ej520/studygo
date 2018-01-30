//Go提供了另一个重要的数据类型：映射(Map)，它将唯一键映射到值。 键是用于在检索值的对象。
//给定一个键和一个值就可以在Map对象中设置值。设置存储值后，就可以使用其键检索它对应的值了

package main

import (
	"fmt"
)

func main() {
	//declare map
	var m map[string]string
	//define the map as nil map can not be assigned any value
	//create a map
	m = make(map[string]string)

	//insert key-value pairs in the map
	m["this"] = "that"
	m["it"] = "is"
	m["game"] = "over"

	//print map using keys
	for k := range m {
		fmt.Println("key is", k, "value is", m[k])
	}
	fmt.Println("----------------------------------------")

	// print map use key-value
	for k, v := range m {
		fmt.Println("key is", k, "value is", v)
	}
	fmt.Println("----------------------------------------")

	//test key "wang" is present in the map or not
	v, ok := m["wang"]
	//if ok is true, present; otherise, absent
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("wang is not present")
	}
}
