package main

import ( "fmt"
         "math/rand" )

const HomeTown string = "Libercourt"

func prln(s string) {
  fmt.Println(s)
}

func main()  {
  prln($HomeTown)
  prln(rand.Int(100))
}
