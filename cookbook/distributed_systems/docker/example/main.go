package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/MrGru/GruGo/cookbook/distributed_systems/docker"
)

var (
	version   string
	builddate string
)

var versioninfo docker.VersionInfo

func init() {
	versioninfo.Version = version
	i, err := strconv.ParseInt(builddate, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	versioninfo.BuildDate = tm
}

func main() {
	http.HandleFunc("/version", docker.VersionHandler(&versioninfo))
	fmt.Printf("version %s listenng on :8000\n", versioninfo.Version)
	panic(http.ListenAndServe(":8000", nil))
}
