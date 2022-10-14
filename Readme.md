## Build

```bash
env GOOS=windows GOARCH=amd64 go build  .  # for Windows
env GOOS=linux GOARCH=amd64 go build .     # for Linux
```


## Usage
```bash
Usage of ./sheet-generator:
  -format string
        Format of file(xlsx, csv) (default "xlsx")
  -lang string
        Faker Language(ru, en) (default "ru")
  -line int
        Count of lines (default 10000)
  -size string
        Size of file(low, middle, big) (default "middle")  
```
## Example 

```bash
./sheet-generator -line=2000 -lang=en -size=low -format=csv

```