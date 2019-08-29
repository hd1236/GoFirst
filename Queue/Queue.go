package Queue

type Queue []int

func (this *Queue) Push(val int) {
	*this = append(*this, val)
}

func (this *Queue) Pop() int {
	if len(*this) > 0 {
		ret := (*this)[0]
		*this = (*this)[1:]
		return ret
	} else {
		return 0
	}
}

func (this *Queue) IsEmpty() bool {
	return len(*this) == 0
}
