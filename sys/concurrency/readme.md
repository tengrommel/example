# concurrency in Go
> Goroutines, Channels, and Pipelines

The previous chapter discussed systems programming in Go, including the Go functions and techniques that allow you to
communicate with your operating system.

Go offers its own unique and innovative way of achieving concurrency, which
comes in the form of goroutines and channels.

*Everything in Go is executed using goroutine, which makes perfect sense
since Go is concurrent programming language by design. Therefore, when a Go
program starts its execution, its single goroutine calls the main() function,
which starts the actual program execution*

## The differences between processes, threads, and goroutines
- A process is an execution environment that contains instructions, user data, and
system data parts, as well as other types of resources that are obtained during
runtime, whereas a program is a file that contains instructions and data that
are used for initializing the instruction and user-data parts of a process.

- A thread is a smaller and lighter entity than a process or a program.
Threads are created by processes and have their own flow of control and stack.
A quick and simplistic way to differentiate a thread from a process is to
consider a process as the running binary file and a thread as a subset of process.

- A goroutine is the minimum Go entity that can be executed concurrently. The
use of the word "minimum" is very important here, as goroutines are not autonomous
entities like UNIX processes - goroutines live in UNIX threads that live
UNIX processes. The main advantage of goroutines is that they are extremely
lightweight and running thousands or hundreds of thousands of them on a single
machine is not a problem

## The Go scheduler
> The UNIX kernel scheduler is responsible for the execution of the threads of
a program. On the other hand, the Go runtime has own scheduler, which is responsible
for the execution of the goroutines using a technique known as m:n scheduling,
where m goroutines are executed using n operating system threads using multiplexing
The Go scheduler is the Go component responsible for the way and order in which
the goroutines of a Go program get executed. This makes the Go scheduler a really
important part of the Go programming language, as everything in a Go program
is executed as a goroutine.

*Be aware that as the Go scheduler only deals with the goroutines of a single
program, its operation is much simpler, cheaper, and faster than the operation
of the kernel scheduler*

## Concurrency versus parallelism
> It is a very common misconception that concurrency is the same thing as parallelism - this is just not true! Parallelism is the simultaneous execution of multiple entities
of some kind, whereas concurrency is a way of structuring your components so that they
can be executed independently when possible.

It is only when you build software components concurrently that you can safely execute
them in parallel, when and if your operating system and your hardware permit it. The Erlang
programming language did this a long time ago - long before CPUs had multiple cores
and computers had lots of RAM.

In a valid concurrent design, adding concurrent entities makes the whole system run faster
because more thing can be executed in parallel. So, the desired parallelism comes from a better
concurrent expression and implementation of the problem. The developer is responsible for
taking concurrency into account during the design phase of a system and will benefit from
a potential parallel execution of the components of the system. So, the developer should
not think about parallelism but about breaking things into independent components that
solve the initial problem when combined.

## The concurrency models of Erlang and Rust

## Creating goroutines

## Creating channels

## Reading or receiving data from a channel
>You can read a single value from a channel named c by executing <-c
In this case, the direction is from the channel to the outer world.

## Writing or sending data to a channel

## Creating pipelines
A pipeline is a virtual method for connecting goroutines and channels
so that the output of one goroutine becomes the input of another goroutine
using channels to transfer your data

## Waiting for your goroutines to finish

# Concurrency in Go - Advanced Topics
> A scheduler is responsible for distributing the amount of work that needs to be done
over the available resources in an efficient

## The select keyword

As you will learn in a short while, the select keyword is pretty powerful and can do
many things in a variety of situations. The select statement in Go looks like a switch
statement but for channels.

In practice, this means that select allows a goroutine to wait on multiple communication operations


## How the Go scheduler works
Go uses the fork-join concurrency model The fork part of the model states that a child
branch can be created at any point of a program. Analogously, the join part of the Go
concurrency model is where the child branch ends and joins with its parent. Among other
 things, both sync.Wait() statements and channels that collect the result s of goroutines
are join points, whereas each new goroutine creates a child branch.

The fair scheduling strategy, which is pretty straightforward and has a simple implementation,
shares all load evenly among the available processors.
- At first, this might look like the perfect strategy because it does not have to take many
things into consideration while keeping all processors equally occupied.

- However, it turns out that this is not exactly the case because most distributed tasks
usually depend on other tasks.

- Therefore, some processors are underutilized, or equivalently, some processors are utilized
more than others.

A goroutine in Go is a task, whereas everything after the calling statement of a goroutine is a continuation.

## Two techniques that allow you to time out a goroutine that takes longer than expected to finish

- Timing out a goroutine

## Signal channels

## Buffered channels

## Nil channels

## Monitor goroutines

## Shared memory and mutexes

## The sync.Mutex and sync.RWMutex types

## The context package and its advanced functionality

## The atomic package

## Worker pools

## Detecting race conditions
