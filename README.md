# ghs

ghs is a simple tool for searching Github repos using their [Search
API](http://developer.github.com/v3/search/)

[![Build Status](https://travis-ci.org/Keithbsmiley/ghs.png?branch=master)](https://travis-ci.org/Keithbsmiley/ghs)

## Installation

```
go get github.com/Keithbsmiley/ghs
```

## Usage

Search for repos matching `AFNetworking`

```
$ ghs AFNetworking

AFNetworking/AFNetworking                                                   11634 Objective-C
AFNetworking/AFIncrementalStore                                              1662 Objective-C
AFNetworking/AFOAuth2Client                                                   566 Objective-C
steipete/AFDownloadRequestOperation                                           519 Objective-C
chroman/Doppio                                                                299 Objective-C
subdigital/AFProgressiveImageDownload                                         253 Objective-C
jnjosh/JJAFAcceleratedDownloadRequestOperation                                239 Objective-C
xmartlabs/XLRemoteImageView                                                   211 Objective-C
AFNetworking/Xcode-Project-Templates                                          208 Objective-C
AFNetworking/AFHTTPRequestOperationLogger                                     183 Objective-C
```

Search for a maximum of 2 repos matching `AFNetworking`

```
$ ghs -count=2 AFNetworking

AFNetworking/AFNetworking                                                   11634 Objective-C
AFNetworking/AFIncrementalStore                                              1662 Objective-C
```

Search for a repo written in [Go](http://golang.org/) matching
`postgres`

```
$ ghs -lang=Go postgres

lib/pq                                                                         623 Go
gosexy/db                                                                      160 Go
vmihailenco/pg                                                                  96 Go
lxn/go-pgsql                                                                    82 Go
jbarham/gopgsqldriver                                                           40 Go
jbarham/pgsql.go                                                                39 Go
replicon/pgreplicaproxy                                                         18 Go
deafbybeheading/dog                                                             12 Go
JackC/pgx                                                                       11 Go
jgallagher/go-libpq                                                              9 Go
```
