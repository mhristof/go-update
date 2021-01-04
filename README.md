go-update
=========

Library to self update golang applications from Github.

Example
-------

You can add this snippet in your code to check/get for new github releases.

Make sure you update the github URL.

```go
url := fmt.Sprintf("https://github.com/mhristof/germ/releases/latest/download/germ.%s", runtime.GOOS)

updates, updateFunc, err := update.Check(url)
if err != nil {
    panic(err)
}

if updates {
    log.Info("New version is available")
}

if dryRun {
    return
}

updateFunc()
```
