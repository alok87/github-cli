# github-cli

Perform github tasks from command line

## Prerequisite

* [Setup golang environment](https://golang.org/doc/install)
* [How to organize go code](https://golang.org/doc/code.html)

## Installation

Install the cli with the following command
```bash
> go get github.com/alok87/github-cli/cmd/ghi
```

## Usage

![Help](http://i.imgur.com/owpwAbc.png)

## Features

### Login
* Setup Login to access Github
```bash
> ghi login my_github_oauth_token
```

###  Get
* Get Repos
```bash
> ghi get repos
```

###  Create
* Create Repo
```bash
> ghi create repo luna
```

###  Delete
* Delete Repo
```bash
> ghi delete repo luna
```

## Contribution

### Resources:
`ghi` can easily be extended.

* [go-github docs](https://godoc.org/github.com/google/go-github/github)
* [viper](https://github.com/spf13/viper)
* [cobra](https://github.com/spf13/cobra)

### Build and Install
```bash
> go install github.com/alok87/github-cli/cmd/ghi
```

### Pull requests are welcome !
* [How to give a PR for a Golang project?](http://blog.campoy.cat/2014/03/github-and-go-forking-pull-requests-and.html)

### Other CLIs made in similar fashion
* [Kubernetes](http://kubernetes.io/)
* [rkt](https://github.com/coreos/rkt)
* [etcd](https://github.com/coreos/etcd)
* [Docker (distribution)](https://github.com/docker/distribution)
* [GopherJS](http://www.gopherjs.org/)
