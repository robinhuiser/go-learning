# Movie scanner

Private project for my father to create an inventory for his movies ;-)

## Build for Windows

~~~bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -o scan-movies.exe .
~~~