package main

type Set []string

func (s *Set) add(v string) bool {
	if s.contains(v) {
		return false
	}

	*s = append(*s, v)

	return true
}

func (s *Set) remove(v string) bool {
	if s.contains(v) {
		var removedIdx int = s.indexOf(v)

		for i := removedIdx; i < s.size()-1; i++ {
			(*s)[i] = (*s)[i+1]
		}

		*s = (*s)[:s.size()-1]
		
		return true
	}

	return false
}

func (s *Set) contains(v string) bool {
	for _, val := range *s {
		if val == v {
			return true
		}
	}

	return false
}

func (s *Set) indexOf(v string) int {
	for i, val := range *s {
		if val == v {
			return i
		}
	}

	return -1
}

func (s *Set) size() int {
	return len(*s)
}
