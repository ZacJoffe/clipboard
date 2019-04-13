# Clipboard
A simple Golang library for reading and writing the clipboard in Linux.

The only dependency is [`xclip`](https://github.com/astrand/xclip), which is likely already installed on a Linux system.

## Use
Import the library:
```
import(
    "github.com/ZacJoffe/clipboard/xclip"
)
```

Use the `xclip.Read` and `xclip.Write` functions to read and write to/from the clipboard.

See `example.go` in the root folder for actual use.