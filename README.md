# mqtt-projects

This project uses paho.mqtt.golang as MQTT client library, install:
```console
go get github.com/eclipse/paho.mqtt.golang
```

If you're encountering a red import error in Go when trying to import a package from GitHub
Clear your Go module cache: If you're using Go modules, try clearing your module cache by running the following command:
```console
go clean -modcache
```
This will remove all cached modules and force Go to re-download them.

You may set GO111MODULE env var to auto
```console
go env -w GO111MODULE=auto
```