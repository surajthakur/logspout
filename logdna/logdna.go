package logdna

import (
    "errors"
    "log"
    "os"

    "github.com/gliderlabs/logspout/router"
    "github.com/logdna/logspout/logdna/adapter"
)

const (
    endpointVar         = "LOGDNA_URL"
    tokenVar            = "LOGDNA_KEY"
    tagsVar             = "TAGS"
    filterNameVar       = "FILTER_NAME"
)

func init() {
    router.AdapterFactories.Register(NewLogDNAAdapter, "logdna")

    r := &router.Route{
        Adapter:    "logdna",
        FilterName: os.Getenv(filterNameVar),
    }

    err := router.Routes.Add(r)
    if err != nil {
        log.Fatal("could not add route: ", err.Error())
    }
}

func NewLogDNAAdapter(route *router.Route) (router.LogAdapter, error) {
    endpoint := os.Getenv(endpointVar)
    token := os.Getenv(tokenVar)
    tags := os.Getenv(tagsVar)

    if endpoint == "" {
        endpoint = "logs.logdna.com/logs/ingest"
    }

    if token == "" {
        return nil, errors.New(
            "could not find environment variable LOGDNA_KEY",
        )
    }

    return adapter.New(
        endpoint,
        token,
        tags,
    ), nil
}