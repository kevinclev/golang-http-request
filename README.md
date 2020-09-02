# golang-http-request

## Purpose
Pulls the HTML Body from a URL and writes it to a file call `html.body`. Computes an MD5 sum of the file and writes that to a file called `body.md5`

## Installation
I'm still a newbie at Go so this is the way I would run this.

Pull down the repo to your GOPATH

```shell
go get github.com/kevinclev/golang-http-request
```
Compile code
```shell
go install <path to code>
```

Run the Binary (In the bin folder within your $GOPATH)
```shell
./golang-http-request
```
## TODO
- Not default the file to body.html since it's really just doing a GET request and you can pull down different types (ex: json responses) with this if you wanted.
