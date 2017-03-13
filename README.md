# gotestsample
sample the test code for golang

## execute test

### case succeed
```
GOPATH=`pwd` go test ./...
ok  	_/Users/suzujun/workspace/gotestsample/pkg1	0.004s
ok  	_/Users/suzujun/workspace/gotestsample/pkg2	0.004s
GOPATH=`pwd` go test ./... -cover
GOPATH=`pwd` go test ./... -cover -coverpkg ./...
```

### 

### case error
cannot use test profile flag with multiple packages.
```
GOPATH=`pwd` go test ./... -cover -coverprofile=cover.out
ok  	_/Users/suzujun/workspace/gotestsample/pkg1	0.003s	coverage: 75.0% of statements
ok  	_/Users/suzujun/workspace/gotestsample/pkg2	0.003s	coverage: 100.0% of statements
GOPATH=`pwd` go test ./... -cover -coverprofile=cover.out -coverpkg ./...
ok  	_/Users/suzujun/workspace/gotestsample/pkg1	0.005s	coverage: 37.5% of statements in ./...
ok  	_/Users/suzujun/workspace/gotestsample/pkg2	0.003s	coverage: 50.0% of statements in ./...
```


