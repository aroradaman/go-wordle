package container

import (
	"fmt"
	"reflect"
)

type Container struct {
	data map[rune]int
}

func NewContainer(members ...rune) *Container {
	data := make(map[rune]int)
	for _, m := range members {
		_, ok := data[m]
		if !ok {
			data[m] = 1
		} else {
			data[m]++
		}
	}
	return &Container{data: data}
}

func (c Container) Copy() *Container {
	data := make(map[rune]int)
	for r, count := range c.data {
		data[r] = count
	}
	return &Container{data: data}
}

func (c Container) IsEmpty() bool {
	return len(c.data) == 0
}

func (c *Container) String() string {
	data := make(map[string]int)
	for r, count := range c.data {
		data[string(r)] = count
	}
	return fmt.Sprintf("Container(%v)", data)

}
func (c *Container) Equals(d *Container) bool {
	return reflect.DeepEqual(c, d)
}

func (c Container) Contains(r rune) bool {
	_, ok := c.data[r]
	return ok
}

func (c *Container) Add(runes ...rune) {
	for _, r := range runes {
		_, ok := c.data[r]
		if ok {
			c.data[r]++
		} else {
			c.data[r] = 1
		}
	}
}

func (c *Container) Pop(runes ...rune) {
	for _, r := range runes {
		_, ok := c.data[r]
		if ok {
			c.data[r]--
			if c.data[r] == 0 {
				delete(c.data, r)
			}
		}
	}
}

func (c *Container) Remove(runes ...rune) {
	for _, r := range runes {
		_, ok := c.data[r]
		if ok {
			delete(c.data, r)
		}
	}
}

func (c Container) GetCount(r rune) int {
	for ru, count := range c.data {
		if r == ru {
			return count
		}
	}
	return 0
}

func (c Container) UpdateCount(r rune, count int) {
	_, ok := c.data[r]
	if ok {
		c.data[r] = count
	}
}
