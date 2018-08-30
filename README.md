# Wastebasket

| OS | master | develop |
| - | - | - |
| darwin / linux | [![Travis CI - Darwin and Linux - master](https://travis-ci.org/Bios-Marcel/wastebasket.svg?branch=master)](https://travis-ci.org/Bios-Marcel/wastebasket) | [![Travis CI - Darwin and Linux - develop](https://travis-ci.org/Bios-Marcel/wastebasket.svg?branch=develop)](https://travis-ci.org/Bios-Marcel/wastebasket) |
| windows | [![AppVeyor - Windows - master](https://ci.appveyor.com/api/projects/status/8tsgphvg9jn3mms2?svg=true)](https://ci.appveyor.com/project/Bios-Marcel/wastebasket) | [![AppVeyor - Windows- develop](https://ci.appveyor.com/api/projects/status/8tsgphvg9jn3mms2/branch/develop?svg=true)](https://ci.appveyor.com/project/Bios-Marcel/wastebasket) |

Wastebasket is a go library allowing you to move files into your trashbin.

## Dependencies

### Windows

The only dependency is PowerShell.

### Linux

Your either need to have `gio`, `gvfs-trash` or `trash-cli` installed.

### Mac OS

```bash
brew install trash
```

## How do i use it

Run

```bash
go get github.com/Bios-Marcel/wastebasket
```

Example usage in Go:

```GO
package main

import (
    "fmt"
    "io/ioutil"
    "os"

    "github.com/Bios-Marcel/wastebasket"
)

func main() {
    ioutil.WriteFile("test.txt", []byte("Test"), os.ModePerm)
    fmt.Println(wastebasket.Trash("test.txt"))
    wastebasket.Empty()
}

```