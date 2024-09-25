package main

import "github.com/scottjab/bplog"

func main() {
	b := bplog.NewBPLog()
	b.DBMigrate()
	//b.LoadCSVIntoDB("/Users/scottjab/Downloads/history.csv")
	b.Plot()
}
