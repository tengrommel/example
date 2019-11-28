# Reading from files

Getting the contents of a file can be done with an auxiliary function in the io/ioutil package, as well as with the ReadFile function,
which opens reads and closes the file at once.

# Reader interface
> For all operations that read from a disk, there's an interface that is paramount:

    type Reader interface {
        Read(p []byte) (n int, err error)
    }

# The file structure
> The os.File type satisfies the reader interface and is the main actor that's used to interact with file contents.    