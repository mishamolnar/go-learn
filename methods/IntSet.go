package methods

import (
	"bytes"
	"fmt"
	"math/bits"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint64(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) addAll(nums ...int) {
	for _, val := range nums {
		s.Add(val)
	}
}

func (s *IntSet) Contains(x int) bool {
	word, bit := x/64, uint64(x%64)
	return word < len(s.words) && (s.words[word]&(1<<bit)) != 0
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i := range t.words {
		if i < len(s.words) {
			s.words[i] |= t.words[i]
		} else {
			s.words = append(s.words, t.words[i])
		}
	}
}

func (s *IntSet) String() string {
	var buff bytes.Buffer
	buff.WriteByte('[')
	for ind, word := range s.words {
		if word != 0 {
			base := ind * 64
			for i := 0; i < 64; i++ {
				if word&(1<<i) != 0 {
					if buff.Len() > 1 {
						buff.WriteByte(' ')
					}
					fmt.Fprintf(&buff, "%d", base+i)
				}
			}
		}
	}
	buff.WriteByte(']')
	return buff.String()
}

func Test() {
	set := &IntSet{words: make([]uint64, 10)}
	set.Add(1)
	set.Add(2)
	set.Add(3000)
	set.Add(123)
	set.addAll(999, 1000, 1001)
	fmt.Println(set.Contains(3))
	fmt.Println(set.Contains(1))
	fmt.Println(set)
}

func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}
