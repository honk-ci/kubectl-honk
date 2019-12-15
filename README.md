# honk-plugin

## Purpose

This is a honk plugin for kubectl. Why? why not? honk!

## Running

```sh
$ GO111MODULE="on" go build cmd/kubectl-honk.go
# place the built binary somewhere in your PATH
$ cp ./kubectl-honk /usr/local/bin


$ kubectl honk
# or
$ kubectl honk --goose PATH_GOOSE_IMAGE

```

## Use Cases

Just for fun!

## Cleanup

You can "uninstall" this plugin from kubectl by simply removing it from your PATH:

    $ rm /usr/local/bin/kubectl-honk


Honk the Planet!
