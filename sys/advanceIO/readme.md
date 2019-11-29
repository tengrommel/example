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
There are two other interfaces that are related to readers: io.Closer and io.Seeker:
    
    type Closer interface {
        Close() error
    }
    
    type Seeker interface {
        Seek(offset int64, whence int) (int64, error)
    }

There are usually combined with io.Reader, and the resulting interfaces are as follows:

    type ReadCloser interface {
        Reader
        Closer
    }
    
    type ReadSeeker interface {
        Reader
        Seeker
    }

> The Close Method ensures that the resource gets released and avoids leaks, while
the Seek method make it possible to move the cursor of the current object(for example,a Writer) 
to the desired offset from the start/end of the file, or from its current position.

# Writing to file
> As we have seen for reading, there are different ways to write files, each one with its own flaws
and strengths.

# Writer interface
> The same principle that is valid for reading also applies for writing - there's an 
interface in the io package that determines writing behaviors,as shown in the following code:

    type Writer interface {
        Write(p []byte) (n int, err error)
    }

#  Buffer and format
> In the previous section, we saw how bytes.Buffer can be used to store data temporarily and how 
it handles its own growth by appending the underlying slice.

# Efficient writing
> Each time the os.File method, that is, Write, is executed, this translates to a system call, 
which is an operation that comes with some overhead.

# Handling Stream

# Stream
> Writers and readers are not just for files; they are interfaces that abstract flows of data in one direction or another.

# Input and readers
> Incoming streams of data are considered the io.Reader interface if the application has no control over the data flow, 
and will wait for an error to end the process, receiving the in.EOF vale in the best 
case scenario, which is a special error that signals that there is no more content to read,
or another error otherwise.

# The bytes reader
> The bytes package contains a useful structure that treats a slice of bytes as an io.Reader interface, and it implements many more I/O interfaces

# The strings reader
> The strings package contains another structure that is very similar to the io.Reader interface, called strings.Reader.
