# Linear Data Structure

Various applications, such as Facebook, Twitter, and Google, use lists and linear data structures.
As We have discussed previously, data structures allow us to organize vast swathes of data in a sequential and organized manner, 
thereby reducing time and effect in working with such a data. Lists, stacks, sets, and tuples are some of the commonly used
linear data structures.

Let's discuss the operations related to add, update, remove, and lookup on linked list and doubly linked list in the following section.

## Lists
> A list is a collection of ordered elements that are used to store list of items. Unlike array lists, these can expand and shrink dynamically.

Lists also be used as a base for other data structures, such as stack and queue.

Lists can be used to store lists of users, car parts, ingredients, to-do items, and various other such elements.

Lists are the most commonly used linear data structures.

## LinkedList
> LinkedList is a sequence of nodes that have properties and a reference to the next new node in the sequence.

It is a linear data structure that is used to store data.

The data structure permits the addition and deletion of components from any node next to another node.

They are not stored contiguously in memory, which makes them different arrays.

## The Node class
> The Node class has an integer typed variable with the name property.

The class has another variable with the name nextNode, which is a node pointer.

Linked list will have a set of nodes with integer properties, as follows:

    // Node class 
    type Node struct {
        property int
        nextNode *Node
    }
    
**The LinkedList class**
> The LinkedList class has the headNode pointer as its property.
By traversing to nextNode from headNode, you can iterate through the linked list, as shown in the following 

    // LinkedList class
    type LinkedList struct {
        headNode *Node
    }

The different methods of the LinkedList class, such as AddToHead, IterateList, LastNode, AddToEnd,
NodeWithValue, AddAfter, and the main method, are discussed in the following sections.

**The AddToHead method**
> The AddToHead method adds the node to the start of the linked list.

# Sets
> A Set is a linear data structure that has a collection of values that are not repeated.

A set can store unique values without any particular order.

Dynamic and mutable sets allow for the insertion and deletion of elements.

Algebraic operations such as union, intersection, difference, and subset can be defined on the sets.

# Tuples
> Tuples are finite ordered sequences of objects. They can contain a mixture of other data types 
and are used to group related data into a data structure.

# Queues
> A queue consists of elements to be processed in a particular order or based on priority

A priority-based queue of orders is shown in the following code, structured as a heap.

# Tree
> A Tree is a non-linear data structure.