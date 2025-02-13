# dir2txt

`dir2txt` is a command-line tool written in Go that reads all text files in a specified directory and prints their contents to the console. Binary files are automatically excluded from the output.

Output is formatted like: 

```md
---
File: /Users/name/code/frontend/hello.js
---
// the hello world program
console.log('Hello World');
```

## Features

- Recursively traverse directories
- Print contents of text files to the console
- Exclude binary files based on null byte detection

## Installation

### Download Pre-Compiled Binaries

Pre-compiled binaries are available for various operating systems.  from the Releases page.

### Building from Source

To build the `dir2txt` binary from the source code, follow these steps:

1. Ensure you have Go installed on your system. You can download it from [the official Go website](https://golang.org/dl/).

2. Clone the repository or download the source code.

3. Navigate to the directory containing the `dir2txt.go` file and run:

   ```bash
   go build -o dir2txt main.go
   ```
