# go-ipfs-routing

[![](https://img.shields.io/badge/made%20by-Protocol%20Labs-blue.svg?style=flat-square)](http://ipn.io)
[![](https://img.shields.io/badge/project-IPFS-blue.svg?style=flat-square)](http://ipfs.io/)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![Build Status](https://travis-ci.org/ipfs/go-ipfs-routing.svg?branch=master)](https://travis-ci.org/ipfs/go-ipfs-routing)

> go-ipfs-routing provides go-libp2p-routing implementations used in go-ipfs.

## ❗ This repo is no longer maintained.
👉 We highly recommend switching to the maintained version at https://github.com/ipfs/boxo/tree/main/routing.
🏎️ Good news!  There is [tooling and documentation](https://github.com/ipfs/boxo#migrating-to-boxo) to expedite a switch in your repo. 

⚠️ If you continue using this repo, please note that security fixes will not be provided (unless someone steps in to maintain it).

📚 Learn more, including how to take the maintainership mantle or ask questions, [here](https://github.com/ipfs/boxo/wiki/Copied-or-Migrated-Repos-FAQ).


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [License](#license)

## Install

`go-ipfs-routing` works like a set of regular Go packages:

```
> go get github.com/ipfs/go-ipfs-routing/...
```

This module uses [Gx](https://github.com/whyrusleeping/gx) to manage
dependencies. You can use `make all` to build it with the `gx` dependencies.

## Usage

This repo contains 3 different packages.

### Mock

[![GoDoc](https://godoc.org/github.com/ipfs/go-ipfs-routing/mock?status.svg)](https://godoc.org/github.com/ipfs/go-ipfs-routing/mock)

```
import "github.com/ipfs/go-ipfs-routing/mock"
```

Mock is a fake router useful for tests. It provides a mock client that
implements the `IpfsRouting` interface and a mock server from which the client
retrieves routing records.


### Offline

[![GoDoc](https://godoc.org/github.com/ipfs/go-ipfs-routing/offline?status.svg)](https://godoc.org/github.com/ipfs/go-ipfs-routing/offline)

```
import "github.com/ipfs/go-ipfs-routing/offline"
```

Offline is an offline router that can put and get records to and from a local
`Datastore` but can't retrieve them from the network.

### None

[![GoDoc](https://godoc.org/github.com/ipfs/go-ipfs-routing/none?status.svg)](https://godoc.org/github.com/ipfs/go-ipfs-routing/none)

```
import "github.com/ipfs/go-ipfs-routing/none"
```

None is a router no-op router that doesn't do anything. Puts always succeed and
lookups always fail.

## License

MIT © Protocol Labs, Inc.
