# Clipboard
A simple Golang library for reading and writing to the clipboard for Linux.

The only dependency is `xclip`, which is likely installed if you're using Linux.

## Use
Import the library:
```
import(
    "github.com/ZacJoffe/clipboard/xclip"
)
```

Use the `xclip.Read` and `xclip.Write` functions to read and write to/from the clipboard. 

See `example.go` in the root folder for actual use.