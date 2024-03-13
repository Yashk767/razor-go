package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"razor/cmd"
)

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	cmd.InitializeInterfaces()
	cmd.Execute()
}
