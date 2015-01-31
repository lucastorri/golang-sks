package main

import "github.com/lucastorri/sks/server"
import "github.com/lucastorri/sks/store"

func main() {
    port := 12121
    dir := "/Users/lucastorri/tmp/data"
    server.New(port, store.New(dir))
}
