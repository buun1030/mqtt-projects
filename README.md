# mqtt-projects

This project uses paho.mqtt.golang as MQTT client library, install:
```console
go get github.com/eclipse/paho.mqtt.golang
```

If you're encountering a red import error in Go when trying to import a package from GitHub.

> **Solution 1:** You may set GO111MODULE env var to auto
> ```console
> go env -w GO111MODULE=auto
> ```
>
> **Solution 2:** Clear your Go module cache: If you're using Go modules, try clearing your module cache by running the following command:
> ```console
> go clean -modcache
> ```
> This will remove all cached modules and force Go to re-download them.
>
> **Solution 3:** Check your workspace settings: If you are using Visual Studio Code, make sure that your workspace settings are configured to use the correct GOPATH. GOPATH check by
>  ```console
> go env
> ```
> Configuring GOPATH by opening your workspace settings (Ctrl+Shift+P, then search for "Preferences: Open Workspace Settings"), and then setting the "go.gopath" property to the correct directory.
