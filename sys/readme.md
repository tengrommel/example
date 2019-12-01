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

One of the main advantages of using a string instead of the byte reader, when dealing with strings that need to read, 
is the avoidance of copying the data when initializing it. This subtle difference helps with both performance and memory
usage because it does fewer allocations and requires the Garbage Collector(GC) to clean up the copy.

# Defining a reader
> Any Go application can define a custom implementing of the io.reader interface. A good general rule when implementing 
interfaces is to accept interfaces and return concrete types, avoiding unnecessary abstraction.

# Output and writers
> The reasoning that applies to incoming streams also applies to outgoing ones.

# The bytes writer
> We already saw that the bytes package offers Buffer, which has both reading and writing capabilities.

- io.Writer: This can act as a regular writer
- io.WriterAt: This makes it possible to write from a certain position onward
- io.ByteWriter: This makes it possible to write single bytes

bytes.Buffer is a very flexible structure considering that it works for both, Writer and ByteWriter and works  best if reused, 
thanks to the Reset and Truncate methods. Instead of leaving a used buffer to be recycled by the GC and make a new buffer,
 it is better to reset the existing one, keeping the underlying array for the buffer and setting the slice length to 0.

# The string writer
> A byte buffer executes a copy of the bytes in order to produce a string.

The only way of obtaining the final string is with the String method, which uses the 
unsafe package under the hood to convert the slice to a string without copying the underlying data.

# Define a writer
> Any custom implementation of any writer can be defined in the application.

# The Uses of Composite Types

## Structures
> Although arrays, slices, and maps are all very useful, they cannot group and hold multiple
values in the same place.When you need to group various types of variable types of variables and 
create a new handy type, you can use a structure.
