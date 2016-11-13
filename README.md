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

![Help](http://i.imgur.com/mRGLoGS.png)

## Features

### Login
* Setup Login to access Github
```bash
> ghi login my_github_oauth_token
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
