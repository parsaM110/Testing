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
after creating the allure-results now run:
```
allure serve allure-results --host 0.0.0.0 --port 9999
```

# Sources
- https://github.com/ozontech/allure-go/tree/master