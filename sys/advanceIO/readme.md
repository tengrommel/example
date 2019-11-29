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

# Using buffers
> A data buffer, or just a buffer, is a part of memory that is used to store temporary data while it is moved.

Byte buffers are implemented in the bytes package, and they are implemented by an underlying slice that is capable of growing every time 
the amount of data that needs to be stored will not fit.

#Peeking content
> Peeking is the ability to read content without advancing the reader cursor.

# Closer and seeker
> There are two other interfaces that are related to readers: io.Closer and io.Seeker:
    
    type Closer interface {
        Close() error
    }
    
    type Seeker interface {
        Seek(offset int64, whence int) (int64, error)
    }

