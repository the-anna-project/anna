# cheat sheet
Here are we collect some useful snippets to ease some pains that may occur
during development.

### fix misspells
There may be the project check failing like follows.
```
$ make projectcheck
go vet    succeeded
golint    succeeded
misspell  failed
make: *** [projectcheck] Error 1
```

The following snippet can be used to easily check misspells as reported by
https://github.com/client9/misspell.
```
find . -not -path "./.git/*" -not -path "./.workspace/*" | xargs misspell
```

You will see output similar to the following depending on the current issues.
```
file:388:4:found "wrong." a misspelling of "right."
file:404:21:found "wrong" a misspelling of "right"
file:4:found "wrong." a misspelling of "right."
file:21:found "wrong" a misspelling of "right"
```

The following snippet can be used to easily correct misspells. Simply add `-w`
to the command we used to check the misspells. The output will be the same as
already seen and the corrections will be written to the related files.
```
find . -not -path "./.git/*" -not -path "./.workspace/*" | xargs misspell -w
```

### run single tests
The following snippet can be used to easily run unit tests of single packages.
This is unlike to the makefile target, that runs the tests of all packages.
Here all tests of the package `pkg` will be run. For more options consider
reading `go test -h`. Note that it is a good idea to run tests with having the
race detector enabled. For more information on this one see
https://blog.golang.org/race-detector.
```
GOPATH=$(pwd)/.workspace/ go test -race ./pkg
```

### search and replace
The following snippet can be used to easily search and replace strings
recursively within a given directory. Here `search` is replaced by `replace`
within all files which names match the expression `*.go` within the current
directory given by `.`. To accomplish that we make use of the command line
tools `sed` and `find`.
```
sed -i 's/search/replace/g' $(find . -name *.go)
```

### list storage keys
The following snippet can be used to list all used storage keys. Note that this
requires the [key package
interface](https://godoc.org/github.com/xh3b4sd/anna/key) to be aligned with the
command below.
```
ack-grep "key\.NewNetworkKey\(" --nobreak --noheading --no-filename --type=go | sed -u 's/^.*key.NewNetworkKey("\(.*\)".*$/\1/p' | uniq | sort
```

The generated output will somehow look like this.
```
activate:configuration:behaviour-id:%s:behaviour-ids
activate:queue:behaviour-id:%s:network-payload
behaviour-id:%s:separator
clg-tree-id:%s:first-behaviour-id
clg-tree-id:%s:first-behaviour-id
event:network-payload
event:network-payload
events:network-payload
events:network-payload
feature:%s:positions
forward:configuration:behaviour-id:%s:behaviour-ids
information-id:%s:information-sequence
information-id:%s:information-sequence
information-id:%s:information-sequence
information-id:%s:information-sequence
information-sequence:%s:information-id
information-sequence:%s:information-id
information-sequence:%s:information-id
pair:syntactic:feature:%s:pair-id
```
