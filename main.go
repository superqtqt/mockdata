package main

import (
	"fmt"
	"github.com/superqtqt/mockdata/internal"
	"regexp"
)

type MockData interface {
	Mock(v interface{}) interface{}
}

var Init = func() MockData {
	return &internal.Imp{}
}

func main() {
	////var a *int
	////v:=reflect.ValueOf(a)
	////fmt.Printf("%s\t%s",v.Type(),v.Kind())
	////println(int(^uint(0)>>1))
	////rand.Seed(time.Now().UnixNano())
	////for i := 0; i < 3100; i++  {
	////	println(int(rand.Float64()*1000))
	////}
	//st:="a1d1f1"
	////for _,v := range strings.Split("a1d1f1", "1") {
	////	fmt.Println(v)
	////}
	//mather:=regexp.MustCompile("\\d")
	//mather.
	//for _,v := range mather.FindAllString(st, -1) {
	//	fmt.Println(v)
	//
	//}
	//for _,v := range regexp.MustCompile("\\d").Split(st, -1) {
	//	fmt.Println(v)
	//}
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < 1000; i++ {
	//	fmt.Println(rand.Intn(10))
	//}
	reg := regexp.MustCompile("\\d")

	str:="1a1b1c1d"
	for _,v := range reg.FindAllStringSubmatch(str, -1) {
		println(v)
	}
	newArr:=reg.Split(str,-1)
	for i := range newArr {
		fmt.Println(newArr[i])
	}

}
