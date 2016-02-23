package main

import (
  mysqlbeat "github.com/rlowe/mysqlbeat/beat"

  "github.com/elastic/beats/libbeat/beat"
)

var Name = "mysqlbeat"
var Version = "0.0.1"

func main() {
  beat.Run(Name, Version, mysqlbeat.New())
}
