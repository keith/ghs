# ghs

ghs is a simple tool for searching Github repos using their [Search
API](http://developer.github.com/v3/search/)

[![Build Status](https://travis-ci.org/Keithbsmiley/ghs.png?branch=master)](https://travis-ci.org/Keithbsmiley/ghs)

## Installation

```
go get github.com/Keithbsmiley/ghs
go install github.com/Keithbsmiley/ghs
```

## Usage

Search for repos matching `AFNetworking`

```
$ ghs AFNetworking

AFNetworking/AFNetworking                                                      8560 Objective-C
AFNetworking/AFIncrementalStore                                                1414 Objective-C
AFNetworking/AFOAuth2Client                                                     401 Objective-C
steipete/AFDownloadRequestOperation                                             387 Objective-C
jnjosh/JJAFAcceleratedDownloadRequestOperation                                  232 Objective-C
AFNetworking/Xcode-Project-Templates                                            169 Objective-C
AFNetworking/AFHTTPRequestOperationLogger                                       154 Objective-C
AFNetworking/AFAmazonS3Client                                                   130 Objective-C
AFNetworking/AFOAuth1Client                                                     105 Objective-C
jaminguy/JGAFImageCache                                                          96 Objective-C
```

Search for a maximum of 2 repos matching `AFNetworking`

```
$ ghs -count=2 AFNetworking

AFNetworking/AFNetworking                                                      8560 Objective-C
AFNetworking/AFIncrementalStore                                                1413 Objective-C
```

Search for a repo written in [Go](http://golang.org/) matching
`postgres`

```
$ ghs -lang=Go postgres
lib/pq                                                                                   368 Go
deafbybeheading/dog                                                                       11 Go
jgallagher/go-libpq                                                                        4 Go
fdr/pg_logplexcollector                                                                    2 Go
darkhelmet/musclecar                                                                       1 Go
w4g3n3r/gopg                                                                               0 Go
levicook/tusk                                                                              0 Go
metakeule/pgsql                                                                            0 Go
jcoene/pgorp                                                                               0 Go
```


