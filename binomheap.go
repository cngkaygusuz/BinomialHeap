package BinomialHeap
import "fmt"


/*
This is the struct that serves as entrypoint to the heap and implements the priority queue interface.
 */
type BinomialHeap struct {
    // In binomial heaps, there are no such thing as global root. It is a collection of trees, called a "forest"
    // forest_head points to the head of a descending "order"-ordered linked list. The linking mechanism is implemented
    // in "BinomialHeapNode"
    forest_head		*BinomialHeapNode

    // Amount of elements in the heap.
    size 			int
}


func NewBinomialHeap() *BinomialHeap {
    return &BinomialHeap {
        forest_head: nil,
        size: 0,
    }
}


/*
Insertion function that satisfies priorityqueue interface. Head over to "insert" for the actual insertion.
 */
func (bh* BinomialHeap) Insert(value int) {
    bh.size += 1

    newnode := newBinomialHeapNode(value)
    bh.insert(newnode)
}


/*
Popping mechanism is as follows;
    * Get the minimum node among the forest heads
    * Remove that node among forest heads.
    * Insert all of the minimum node's children to the queue again.
 */
func (bh* BinomialHeap) Pop() int {
    // Assume the queue is not empty.
    bh.size -= 1

    minnode := getMinimumNode(bh.forest_head)
    removeFromLinkedList(&bh.forest_head, minnode)

    for _, child := range nodeIterator(minnode.children_head) {
        removeFromLinkedList(&minnode.children_head, child)
        bh.insert(child)
    }

    return minnode.value
}


/*
Return the minimum valued node, but do not remove.
This operation can be made O(1) trivially. Do it.
 */
func (bh* BinomialHeap) Peek() int {
    return getMinimumNode(bh.forest_head).value
}


/*
Return the element count in the heap.
 */
func (bh* BinomialHeap) Size() int {
    return bh.size
}


/*
Merge operations is akin to "Pop";
    * For every forest head present on the other heap:
        * remove that node from the other heap
        * insert this node to this heap using the insertion procedure.
 */
func (bh* BinomialHeap) Merge(other *BinomialHeap) {
    bh.size += other.size

    for _, child := range nodeIterator(other.forest_head) {
        removeFromLinkedList(&other.forest_head, child)
        bh.insert(child)
    }
}


/*
Here is the actual magic. Suppose that we are attempting to insert a node with order N
    * If there is already a node exists in the forest with order N
        * Link these two nodes, yielding a node with order N+1  (read "linkNodes" for details)
        * Re-insert this newly made node
    * If there is not, simply put this node somewhere in the linked list of forest heads.

Recall that, there cannot exists two nodes with the same rank on a linked list (they can exist on different ones though)
Linking operation is done to address this, but this same conflict may happend for the node with order N+1.
This means we should insert and reinsert until we can make an insertion order-conflict-free.
 */
func (bh *BinomialHeap) insert(newnode *BinomialHeapNode) {
    srnode := getNodeWithOrder(bh.forest_head, newnode.order)

    if srnode == nil {
        insertIntoLinkedList(&bh.forest_head, newnode)
    } else {
        removeFromLinkedList(&bh.forest_head, srnode)
        linkednode := linkNodes(srnode, newnode)
        bh.insert(linkednode)
    }
}


/*
Print the elements of the heap in a depth-first fashion.
 */
func (bh *BinomialHeap) print() {
    if bh.forest_head == nil {
        fmt.Print("heap is empty.")
    }

    for _, node := range nodeIterator(bh.forest_head) {
        node.print_recursive(0)
    }
}

