package opcode

import (
	"fmt"
	"strconv"
	"strings"
)

type memory struct {
	Initial []int
	Current []int
	Size    int
}

func (m *memory) Reset() {
	m.Current = make([]int, len(m.Initial))
	m.Size = len(m.Current)
	copy(m.Current, m.Initial)
}

func (m *memory) Get(pointer int) int {
	log(fmt.Sprintf("[Memory.Get] Fetching value from %v", pointer))
	if pointer < 0 {
		return -1
	} else if pointer >= m.Size {
		m.Expand(pointer - m.Size + 1)
	}
	return m.Current[pointer]
}

func (m *memory) Set(pointer int, val int) {
	log(fmt.Sprintf("[Memory.Set] Writing value %v to position %v", val, pointer))
	if pointer < m.Size && pointer >= 0 {
		m.Current[pointer] = val
	}
}

func (m *memory) Expand(size int) {
	expanded := make([]int, size)
	for i := 0; i < size; i++ {
		expanded[i] = 0
	}
	m.Current = append(m.Current, expanded...)
	m.Size += size
}

func (m *memory) Initialize(input []string) {
	for _, line := range input {
		for _, val := range strings.Split(line, ",") {
			cleaned, err := strconv.Atoi(val)
			if err != nil {
				// For now we will skip unrecognized values
				continue
			}
			m.Initial = append(m.Initial, cleaned)
			m.Size += 1
		}
	}
}

func NewMemory(input []string) *memory {
	var mem memory
	mem.Initialize(input)
	mem.Reset()
	return &mem
}
