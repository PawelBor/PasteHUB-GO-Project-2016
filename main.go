package main

import "gopkg.in/macaron.v1"

func main() {
  m := macaron.Classic()

  m.Run(8080)
}
