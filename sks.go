package main

import "sks/server"
import "sks/store"

func main() {
    port := 12121
    dir := "/Users/lucastorri/tmp/data"
    server.New(port, store.New(dir))
}
