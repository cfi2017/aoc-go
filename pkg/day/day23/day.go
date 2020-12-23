package day23

import "strconv"

type Day23 struct {
}

func (d Day23) Solve(input string) (a, b int) {
	cups := parse(input)
	a = int(solve(cups, uint32(len(cups)), 100, false))
	b = int(solve(cups, 1000000, 10000000, true))
	return
}

func parse(input string) []uint32 {
	res := make([]uint32, len(input))
	for i, c := range input {
		num, _ := strconv.Atoi(string(c))
		res[i] = uint32(num)
	}
	return res
}

type Node struct {
	val  uint32
	next *Node
	prev *Node
}

func max(v []uint32) uint32 {
	var m uint32
	for _, e := range v {
		if e > m {
			m = e
		}
	}
	return m
}

func makeLinkedList(cups []uint32, length uint32) *Node {
	head := &Node{val: cups[0]}
	prev := head
	for _, i := range cups[1:] {
		node := &Node{val: i, prev: prev}
		prev.next = node
		prev = node
	}
	for i := max(cups) + 1; i <= length; i++ {
		node := &Node{val: i, prev: prev}
		prev.next = node
		prev = node
	}
	head.prev = prev
	prev.next = head
	return head
}

func makeNodeMap(head *Node, length int) map[uint32]*Node {
	nodes := make(map[uint32]*Node)
	current := head
	for i := 0; i < length; i++ {
		nodes[current.val] = current
		current = current.next
	}
	return nodes
}

func in(arr []uint32, v uint32) bool {
	for _, i := range arr {
		if v == i {
			return true
		}
	}
	return false
}

func solve(cups []uint32, length uint32, iterations int, mul bool) uint64 {
	head := makeLinkedList(cups, length)
	nodes := makeNodeMap(head, int(length))
	current := head

	for i := 0; i < iterations; i++ {
		left := current.next
		right := left.next.next
		picked := []uint32{left.val, left.next.val, right.val}

		// unhook next three elements from list
		current.next = right.next
		current.next.prev = current

		dest := (current.val - 1) % length
		if dest == 0 {
			dest = length
		}
		for in(picked, dest) {
			dest = (dest - 1) % length
			if dest == 0 {
				dest = length
			}
		}
		dnode := nodes[dest]

		// hook next three elements back into to list
		right.next = dnode.next
		right.next.prev = right
		dnode.next = left
		left.prev = dnode

		// increment current
		current = current.next

	}

	if mul {
		return uint64(nodes[1].next.val) * uint64(nodes[1].next.next.val)
	} else {
		var n uint64 = 0
		one := nodes[1]
		for i := 0; i < len(cups)-1; i++ {
			one = one.next
			n *= 10
			n += uint64(one.val)
		}
		return n
	}

}
