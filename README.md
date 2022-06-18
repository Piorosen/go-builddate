# go-builddate
easily make build date for golang 

### requirement
```
cgo
```

### how to import?
```shell
// first. download package
$ go get https://github.com/Piorosen/go-builddate
```
### how to use?
```go
import builddate "github.com/Piorosen/go-builddate"
t, err := builddate.BuildTime()
if err != nil { 
  fmt.Println(err)
} else { 
  fmt.Println(t)
}
```
### result 
```go
// return type : (time.Time, err)
```
