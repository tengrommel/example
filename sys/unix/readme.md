# Unix processes
> The subject of this chapter is systems programming because, after all,
Go is a mature systems programming language that was born out of frustration.
Its spiritual fathers were unsatisfied with the programming language choices
they had for creating systems software, so they decided to create a new programming language.

# About UNIX processes
> Strictly speaking, a process is an execution environment that contains instructions,
 user data and system data parts, and other types of resources that are obtained during runtime

 On the other hand, a program is a binary file that contains instructions and data
 that are used for initializing the instruction and user data parts of a process.

 There are three categories of processes:
 - user processes
 > User processes run in user space and usually have no special access rights
 - daemon processes
 > Daemon processes are programs that can be found in the user space and
 run in the background without the need for a terminal.
 - kernel processes
 > Kernel processes are executed in kernel space only and can fully access all kernel data structures.

 ## The flag package
 > Flags are specially-formatted strings that are passed into a program to control its behavior

The flag package makes no assumptions about the order of command-line arguments and options,
and it prints helpful messages in case there was an error in the way the command-line utility was executed.

## The viper package
> viper is a powerful Go package that supports a plethora of options


