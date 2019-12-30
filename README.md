# go-readerer
Convert io.ReaderAt to io.Reader.

## Usage

```go
f, _ := os.Open("/path/to/file")
r := readerer.FromReaderAt(f, 42) // offset == 42
b, _ := ioutil.ReadAll(r)
```