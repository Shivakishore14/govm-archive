# Govm [ Archive ]

## please find the new repo at [golang-vm/govm](https://github.com/golang-vm/govm)

Govm is a fast and flexible version manager for Go built with Go.

Govm lets you work with multiple versions of Go, actively.


# installation

```
go get github.com/Shivakishore14/govm
```
### configure govm
```
govm configure
```

# List versions
List all versions available for download
```
govm ls-remote
```

List versions installed locally.
```
govm ls
```

# Download / Install a Go version
```
govm install 1.10
```

# Use a version
```
govm use 1.10
```

# uninstall a version
```
govm uninstall 1.10
```

# Execute a command with specific version
```
govm exec 1.10 go env
```