# A server that serves http requests to POST and GET images

### build steps
```
go mod vendor
# LINUX:
go build -m vendor -o server
# WINDOWS: 
go build -m vendor -o server.exe
```

# Run server:
```
# Linux:
./server
# Windows:
.\server.exe
```