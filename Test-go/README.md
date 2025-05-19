so first set up a module
```
go mod init test
```
then run
```
go get github.com/ozontech/allure-go/pkg/framework
```
then run 
```
 go test ./... -v
```
if there was any problem run
```
go mod tidy
```