package BinomialHeap

import "fmt"


const (
    PRINT_LEVEL_INCR = 4
)


type BinomialHeapNode struct {
    value 			int

    parent 			*BinomialHeapNode
    children_head	*BinomialHeapNode
    rightsibling	*BinomialHeapNode

    order			int
}


func newBinomialHeapNode(value int) *BinomialHeapNode {
    return &BinomialHeapNode {
        value: value,
        parent: nil,
        children_head: nil,
        rightsibling: nil,
        order: 0,
    }
}


/*
A node adopts another node, becoming its parent.
 */
func (bn* BinomialHeapNode) adopt(other *BinomialHeapNode) {
    // Assumes "other" has no parent and is not present within a linked list.

    // Sibling relations
    insertIntoLinkedList(&bn.children_head, other)

    // Parent relations
    other.parent = bn
}

/*
A node goes rogue, severing its ties with its parent and siblings.
 */
func (bn* BinomialHeapNode) rogue() {
    // Assumes this node has a parent already.
    // Note: Topmost nodes of forests should not be isolated using this function.

    // Sibling relations
    removeFromLinkedList(&bn.parent.children_head, bn)

    // Parent relations
    bn.parent = nil
}


/*
Link two given nodes; which is simply done with the principle:
    "The node with smaller value becomes the parent of the other one"
 */
func linkNodes(n1 *BinomialHeapNode, n2 *BinomialHeapNode) *BinomialHeapNode {
    if n1.value < n2.value {
        n1.order += 1
        n1.adopt(n2)
        return n1
    } else {
        n2.order += 1
        n2.adopt(n1)
        return n2
    }
}


// === Printing Utility ===
func (bn *BinomialHeapNode) print_single() {
    /*fmt.Printf("Value: %d Order: %d addr: %p ch_head: %p right: %p parent: %p\n",
                 bn.value, bn.order, bn, bn.children_head, bn.rightsibling, bn.parent)
    */
    fmt.Printf("Value: %d Order: %d\n", bn.value, bn.order)
}


func (bn *BinomialHeapNode) print_recursive(level int) {
    printSpaces(level)
    bn.print_single()

    for _, child := range nodeIterator(bn.children_head) {
        child.print_recursive(level + PRINT_LEVEL_INCR)
    }
}
