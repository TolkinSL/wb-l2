L2.5  
Что выведет программа?  

Объяснить вывод программы.  

```
package main

// Создание структуры кастомной ошибки со свойством msg  

type customError struct {
  msg string
}

// Создание метода Error() через указатель на структуру  
// кастомной ошибки, который возвращает значение ошибки e.msg  

func (e *customError) Error() string {
  return e.msg
}

// Обычная функция которая в случае ошибки возвращает указатель  
// на структуру с кастомной ошибкой, и которая содержит саму ошибку

func test() *customError {
  // ... do something
  return nil
}

func main() {
    // Создание переменной err типа error, error это интерфейс
  var err error
    // Присваиваниение переменной err ошибки возвращаемой из функции test()  
    // так как *customError реализует интерфейс error имеет метод Error  
    // Возврат функции test() упаковывается в интерфейс error  

    // test() возвращает nil типа *customError со значением (*customError)(nil)
    // при err = test() присваивании в интерфейс error происходит упаковка  
    // интерфейс error хранит: (type, value)
    // type = *customError, value = nil     
    // Интерфейс не равен nil, потому что у него есть тип.
    // и функция напечатает println("error")


  err = test()
  if err != nil {
    println("error")
    return
  }
  println("ok")
}
```