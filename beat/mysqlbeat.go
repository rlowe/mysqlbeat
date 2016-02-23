package beat

import (
  "time"

  "github.com/elastic/beats/libbeat/cfgfile"
  "github.com/elastic/beats/libbeat/publisher"

  _ "github.com/go-sql-driver/mysql"
)

type MySQLbeat struct {
  period   time.Duration
  host     string
  port     int
  network  string
  maxConn  int
  auth     bool
  pass     string
  RbConfig ConfigSettings
  events   publisher.Client

  statusStats    bool
  variablesStats bool
  sysschemaStats bool
  replStats      bool

  done chan struct{}
}

func New() *MySQLbeat {
  return &MySQLbeat{}
}


