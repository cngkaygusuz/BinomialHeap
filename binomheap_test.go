package BinomialHeap


import (
    "testing"
    "math/rand"
)


func Test_insert_1(t *testing.T) {
    heap := NewBinomialHeap()

    node := newBinomialHeapNode(1)

    heap.insert(node)

    if heap.forest_head != node {t.Errorf("something went wrong.")}
}


func Test_insert_2(t *testing.T) {
    const RANGE = 16

    heap := NewBinomialHeap()
    insert_mult(heap, interval(0, RANGE))

    if heap.forest_head.order != 4 {t.Errorf("something went wrong")}
}


func Test_pop_1(t *testing.T) {
    const RANGE = 16

    heap := NewBinomialHeap()
    insert_mult(heap, reverse(interval(0, RANGE)))

    for i:=0; i<RANGE; i++ {
        pval := heap.Pop()
        if pval != i {t.Errorf("expected %d, got %d", i, pval)}
    }
}


func Test_pop_shuffle_1(t *testing.T) {
    const RANGE = 10000

    heap := NewBinomialHeap()

    elems := shuffle(interval(0, RANGE))
    insert_mult(heap, elems)

    for i:=0; i<RANGE; i++ {
        pval :=	heap.Pop()
        if pval != i {t.Errorf("expected %d, got %d", i, pval)}
    }

}


func Test_merge_1(t *testing.T) {
    const (
        SIZE1 = 1000
        SIZE2 = 2000
    )

    h1 := NewBinomialHeap()
    h2 := NewBinomialHeap()

    insert_mult(h1, shuffle(interval(0, SIZE1)))
    insert_mult(h2, shuffle(interval(SIZE1, SIZE2)))

    h1.Merge(h2)

    if h1.size != SIZE2 {t.Errorf("size error, expected %d, got %d,", SIZE2, h1.size)}

    for i:=0; i<SIZE2; i++ {
        pval := h1.Pop()
        if pval != i {t.Errorf("expected %d, got %d", i, pval)}
    }
}






// Helpers
func interval(start int, end int) []int {
    // [start, end)
    slice := make([]int, 0, end-start)

    for i:=start; i<end; i++ {
        slice = append(slice, i)
    }
    return slice
}


func reverse(values []int) []int {
    // reverse the array
    reversed := make([]int, len(values), len(values))
    for i:=0; i<len(values); i++ {
        reversed[len(values)-i-1] = values[i]
    }
    return reversed
}


func shuffle(slice []int) []int {
    // shuffle the elements of the array
    shuffled := make([]int, len(slice), len(slice))
    copy(shuffled, slice)

    for i:=0; i<len(shuffled); i++ {
        index := rand.Intn(len(shuffled)-i)
        index += i

        // swap i'th and index'th element
        tmp := shuffled[i]
        shuffled[i] = shuffled[index]
        shuffled[index] = tmp
    }
    return shuffled
}


func insert_mult(bh* BinomialHeap, values []int) {
    // wrapper for inserting multiple elements into a heap.
    for _, elem := range values {
        bh.Insert(elem)
    }
}
