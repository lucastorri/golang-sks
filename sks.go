package main

import (
    "errors"
    "flag"
    "strings"
    "github.com/lucastorri/sks/server"
    "github.com/lucastorri/sks/store"
)

type config struct {
    port int
    store store.Store
}

func main() {
    cfg := parseArgs()
    server.New(cfg.port, cfg.store)
}

func parseArgs() config {
    port := flag.Int("port", 12121, "server port")
    storeArg := flag.String("store", "mem", "storage method")
    flag.Parse()

    storeParts := strings.Split(*storeArg, ":")
    var s store.Store
    switch storeParts[0] {
        case "mem": s = store.NewMemStore()
        case "dir": s = store.NewFileStore(storeParts[1])
        default: panic(errors.New("Invalid store param"))
    }

    return config { *port, s }
}

