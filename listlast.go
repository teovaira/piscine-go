package piscine

func ListLast(l *List) interface{} {
	if l == nil || l.Head == nil {
		return nil
	}

	if l.Tail != nil {
		return l.Tail.Data
	}

	it := l.Head
	for it.Next != nil {
		it = it.Next
	}
	return it.Data
}
