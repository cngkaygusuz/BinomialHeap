package BinomialHeap
import "fmt"

const (
    INIT_ARRAYSIZE = 4
)


/*
Given a node and head of a linked list, insert this node to the list keeping the "order"-ordering.

"head" is a double pointer, because we may have to modify the head as a result of this operation.
 */
func insertIntoLinkedList(head **BinomialHeapNode, node *BinomialHeapNode) {
    var prev *BinomialHeapNode
    var next *BinomialHeapNode

    prev = nil
    next = *head

    for next != nil && node.order < next.order {
        prev = next
        next = next.rightsibling
    }

    if prev == nil && next == nil {  // linked list is empty
        *head = node
    } else if prev == nil && next != nil {	// linked list is not empty and our new node has higher rank than the node pointed by head.
        node.rightsibling = *head
        *head = node
    } else if prev != nil && next == nil {  // We got to the end of the list, and our newnode has the smallest rank.
        prev.rightsibling = node
    } else if prev != nil && next != nil { 	// our node has found a place for itself somewhere in the list.
        prev.rightsibling = node
        node.rightsibling = next
    }
}


/*
Given a head and node of the list, remove the node.

"head" is a double pointer because it might be modified as a result of this operation.

This function takes O(logn), but if you use a doubly linked list it would take O(1). It would complicate the hell
out of the functions present here though. Do it.
 */
func removeFromLinkedList(head **BinomialHeapNode, node *BinomialHeapNode) {
    // Assume the node is present in the list.
    leftsib := getLeftsibling(*head, node)

    if leftsib == nil {
        // We are removing the head of this list.
        *head = node.rightsibling  // this can set to nil.
    } else {
        leftsib.rightsibling = node.rightsibling
    }
    node.rightsibling = nil
}


/*
Get the previous element of the given node in the list,
 */
func getLeftsibling(head *BinomialHeapNode, node *BinomialHeapNode) *BinomialHeapNode {
    // Assume the node is present in the list.

    if head == node {
        return nil
    }

    checknode := head

    for checknode.rightsibling != node {
        checknode = checknode.rightsibling
    }

    return checknode
}


/*
Get a node with the given order.
 */
func getNodeWithOrder(head* BinomialHeapNode, order int) *BinomialHeapNode {
    checknode := head

    for checknode != nil {
        if checknode.order == order {
            return checknode
        }
        checknode = checknode.rightsibling
    }
    return nil
}


/*
Get the minimum valued node within the queue.
 */
func getMinimumNode(head* BinomialHeapNode) *BinomialHeapNode {
    // Assume there exists at least 1 node.
    minnode := head
    checknode := head.rightsibling

    for checknode != nil {
        if checknode.value < minnode.value {
            minnode = checknode
        }
        checknode = checknode.rightsibling
    }
    return minnode
}


/*
Copy the linked list to an array, and return it.

This is used to iterate over all of the elements, traversing the linked list is not reliable because during one of
such traversals linked list is modified.
 */
func nodeIterator(head* BinomialHeapNode) []*BinomialHeapNode {
    arr := make([]*BinomialHeapNode, 0, INIT_ARRAYSIZE)

    trnode := head
    for trnode != nil {
        arr = append(arr, trnode)
        trnode = trnode.rightsibling
    }
    return arr
}


func printSpaces(cnt int) {
    for i:=0; i<cnt; i++ {
        fmt.Print(" ")
    }
}
