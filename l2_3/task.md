L2.3  
Что выведет программа? 

Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.  

```
package main

import (
  "fmt"
  "os"
)

// Тип интерфейса error  
// type error interface {  
//    Error() string  
// }  

// *os.PathError реализует интерфейс error,  
// у него есть метод Error() string  
// type PathError struct {  
// 	...  
// }  
// func (e *PathError) Error() string {  

// Функция Foo() объявлена с возвращаемым типом error  
// Внутри создаётся указатель типа *os.PathError, присваивается значение nil  
// При возврате этот указатель упаковывается в интерфейс error  
// В результате возвращается не nil-интерфейс,  
// а интерфейс с ненулевым типом и nil-значением

func Foo() error {
  var err *os.PathError = nil
  return err
}

func main() {
  err := Foo()
  fmt.Println(err) // <nil>
  fmt.Println(err == nil) // false
}

```