## Build

```bash
env GOOS=windows GOARCH=amd64 go build . # for Windows
env GOOS=linux GOARCH=amd64 go build .   # for Linux
```

## Run
-c flag set a count of lines, default 1000

```bash
./csv-generator -c=00
```