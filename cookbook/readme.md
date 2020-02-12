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
- Working with directories and files
- Working with the CSV format
- Working with temporary files
- Working with text/template and html/template