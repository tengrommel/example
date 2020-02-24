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

# Working with numberic data types using math and math/big
> The math and math/big packages focus on exposing more complex mathematical operations to the
Go language, such as Pow, Sqrt, and Cos.

# Currency conversions and float64 considerations
> Working with currency is always a tricky process.
It can be tempting to represent money as a float64, but this can result in some pretty tricky (rand wrong)
rounding errors when doing calculations. For this reason, it's preferable to think of money in terms of cents
and store the figure as an int64 instance.

When collection user input froms, the command line, or other sources,
money is usually represented in dollar form. For this reason, it's best to treat
it as a string and convert that string directly to cents without floating-point conversions.
This recipe will present ways to convert a string representation of currency into an int64(cents)
instance and back again.

# Using pointers and SQL NullTypes for encoding and decoding
> When you encode or decode into an object in Go, types that are not explicitly set will be set to their default values.

Strings will default to empty string("") and integers will default to 0, as an example.

Normally, this is fine, unless 0 means something for your API or service that is consuming the user input or returning it.

In addition, if you use struct tags such as json omitempty, the 0 value will be ignored even if they're valid.

Another example of this is Null, which returns from SQL What value best represents Null for an Int?