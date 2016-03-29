# bytesize
The mutual conversion between digital and human-readable for byte size.

## Usage
```go
// convert human-readable string to digital byte size.
// supported the unit: B,KB,MB,GB,TB,PB
// supported format: 1K, 1KB, 1kB, 1Kb, 10KB, 1.3KB, 13, etc, they are flexible!
n, err := bytesize.Parse("1.1MB")
if err != nil{
    fmt.Println(err)
}

// convert digital byte size to human-readable string
s := bytesize.ByteSizeToString(1*1024*1024)
fmt.Println(s) // 1MB
```
