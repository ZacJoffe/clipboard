# Clipboard
Simple Golang libraries for reading and writing text to/from the clipboard in Linux.

Libraries for both [`xclip`](https://github.com/astrand/xclip) and [`xsel`](https://github.com/kfish/xsel) are included in separate packages. See the section below and installation and use.

As such, the only dependency for each package is its respective clipboard service (`xclip` or `xsel`).

Note that both libraries have the exact same functionality, just using two different backends.

## Installation and Use
Install the library you choose to use. For `xclip`:
```
go get -u github.com/ZacJoffe/clipboard/xclip
```

For `xsel`:
```
go get -u github.com/ZacJoffe/clipboard/xsel
```

Then import either of the libraries of your choice in your code:
```
import(
    "github.com/ZacJoffe/clipboard/xclip"
    "github.com/ZacJoffe/clipboard/xsel"
)
```

Use the `Read` and `Write` functions to read and write text to/from the clipboard.

See `example.go` in the root folder for actual use.