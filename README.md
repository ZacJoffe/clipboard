# Clipboard
Simple Golang libraries for reading and writing text to/from the clipboard in Linux.

Libraries for both [`xclip`](https://github.com/astrand/xclip) and [`xsel`](https://github.com/kfish/xsel) are included in separate packages. See the section below and installation and use.

As such, the only dependency for each package is its respective clipboard service (`xclip` or `xsel`).

## Use
Import the library:
```
import(
    "github.com/ZacJoffe/clipboard/xclip"
)
```

Use the `xclip.Read` and `xclip.Write` functions to read and write to/from the clipboard.

See `example.go` in the root folder for actual use.