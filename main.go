package main
 
import (
    "errors"
    "fmt"
)
 
// 使用されない定数
const Url = "http://example.com"
 
type Person struct {
    name string
    // 使用されない
    age int
}
 
func main() {
    // 意味のない代入
    ineffassign := "hoge"
    ineffassign  = "ineffassign"
 
    fmt.Println(ineffassign)
 
    person := &Person{name:"hoge"}
    fmt.Println(person.name)
 
    // エラーチェックしない
    returnError()
}
 
func returnError() error {
    fmt.Println("return error")
    return errors.New("new error")
}
 
// 使用されない関数
func unused() {
    fmt.Println("unused")
}
