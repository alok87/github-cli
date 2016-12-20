# github-cli

[![Build Status](https://travis-ci.org/alok87/github-cli.svg?branch=master)](https://travis-ci.org/alok87/github-cli)

Perform github tasks from command line.

## Prerequisite

* [Setup golang environment](https://golang.org/doc/install)
* [How to organize go code](https://golang.org/doc/code.html)

## Installation

Install the cli with the following command
```bash
> go get -u github.com/alok87/github-cli
```

## Usage

### Login
* Setup Login to access Github
```bash
> github-cli login my_github_oauth_token
```

###  Get
* Get Repos
```bash
> github-cli get repos
```

###  Create
* Create Repo
```bash
> github-cli create repo luna
```

###  Delete
* Delete Repo
```bash
> github-cli delete repo luna
```


## Contribution

### Development Setup

1. Clone the repo & `cd` into it.
2. Ensure that [`govendor`](https://github.com/kardianos/govendor) is installed
and install all the vendor dependencies with `govendor sync`.
3. Install with:
```
go install github.com/alok87/github-cli
```

#### Dependencies

Add new dependencies with `govendor fetch <packagename>`. This would install
the dependencies under `vendor/` and add them to `vendor/vendor.json`, which
should be checked-in.


### Resources:
`github-cli` can easily be extended.

* [go-github docs](https://godoc.org/github.com/google/go-github/github)
* [viper](https://github.com/spf13/viper)
* [cobra](https://github.com/spf13/cobra)


### Pull requests are welcome !
* [How to give a PR for a Golang project?](http://blog.campoy.cat/2014/03/github-and-go-forking-pull-requests-and.html)

### Other CLIs made in similar fashion
* [Kubernetes](http://kubernetes.io/)
* [rkt](https://github.com/coreos/rkt)
* [etcd](https://github.com/coreos/etcd)
* [Docker (distribution)](https://github.com/docker/distribution)
* [GopherJS](http://www.gopherjs.org/)
