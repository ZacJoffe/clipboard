# Clipboard
Simple Golang libraries for reading and writing text to/from the clipboard in Linux.

Libraries for both [`xclip`](https://github.com/astrand/xclip) and [`xsel`](https://github.com/kfish/xsel) are included in separate packages. See the section below and installation and use.

As such, the only dependency for each package is its respective clipboard service (`xclip` or `xsel`).

The library for `xclip` can also be used to copy images to the clipboard.

## Installation and Use
Install the library of your choise. For `xclip`:
```
go get -u github.com/ZacJoffe/clipboard/xclip
```

For `xsel`:
```
go get -u github.com/ZacJoffe/clipboard/xsel
```

Then import either of the libraries of your choice in your code:
```
import (
    "github.com/ZacJoffe/clipboard/xclip"
    "github.com/ZacJoffe/clipboard/xsel"
)
```

Use the `Read` and `WriteText` functions to read and write text to/from the clipboard.

For xclip, use the `WriteImage` function to copy an image file to the clipboard.

See `example.go` in the root folder for actual use.