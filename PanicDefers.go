package gosystem

import ( "fmt" )

func DeferHello(){
    defer fmt.Println("World 1\n")
    defer fmt.Println("World 2\n")
    defer fmt.Println("World 3\n")
    defer fmt.Println("World 4\n")
          fmt.Println("Hello")
          testPanics()
          fmt.Println("World")
}


func testPanics(){
  defer func() {
    fmt.Println("This is how we catch panic")
    if err := recover(); err!=nil {
      fmt.Println(err)
      fmt.Println("We recovered from panic")
    }
  }()
  panic("Oh my god zombie apocalypse")
}
