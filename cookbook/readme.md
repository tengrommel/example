# I/O and Filesystems
> Go provides excellent support for both basic and complex I/O

The Go standard library frequently uses these interface, and they will be used
by recipes throughout the book.

You'll learn how to work with data in memory and in the form of streams. You'll see
examples of working with files, directories, and the CSV format.

- Using the common I/O interfaces
> The Go language provides a number of I/O interface that are used throughout the standard library

It is best practice to make use of these interfaces wherever possible rather than
passing structures or other types directly

    type Reader interface {
        Read(p []byte) (n int, err error)
    }

    type Writer interface {
        Write(p []byte) (n int, err error)
    }

Go also makes it easy to combine interfaces:

    type Seeker interface {
        Seek(offset int64, whence int) (int64, error)
    }
    
    type ReadSeeker interface {
        Reader
        Seeker
    }

- Using the bytes and strings packages
> The bytes and strings packages have a number of useful helpers to work with and
convert the data from string to byte types, and vice versa.
- Working with directories and files
- Working with the CSV format
> CSV is a common format that is used to manipulate data
- Working with temporary files
- Working with text/template and html/template

# Command-Line Tools
> Command-line applications are among the easiest ways to handle user input and output

## Using command-line flags
## Using command-line arguments
## Reading and setting environment variables
## Configuration using TOML, YAML, and JSON
## Working with Unix pipes
> Unix pipes are useful when we are passing the output of one program to the
input of another
## Catching and handling signals
## An ANSI coloring application