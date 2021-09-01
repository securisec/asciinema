# asciinema
**This is not a complete project** This repo is a fork of the [asciinema](https://github.com/asciinema/asciinema) repo under the `golang` branch. It has been refactored a bit so that it can be used as a lib. As the originating branch is quite old, this lib will be behind some of the latest features and improvements that has been made to asciinema overall. 

### Implemented
- Record
- Play

The for was made to use with a local project, and all PR's are welcome. 

## Usage
### Install
```sh
go get -u github.com/securisec/asciinema # for v1 asciinema format
go get -u github.com/securisec/asciinema/v2 # for v2 asciinema format
```

```go
package main

import "github.com/securisec/asciinema"

func main() {
    cli := asciinema.New()
    cast, err := cli.Rec()
    ...
}
```
