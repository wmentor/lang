package mcounter

import (
	"sort"
)

type Counter map[string]uint64

func New() Counter {
	return map[string]uint64{}
}

func (c Counter) Inc(name string, val uint64) uint64 {
	if val == 0 {
		return 0
	}

	v := c[name]
	v += val
	c[name] = v
	return v
}

func (c Counter) Dec(name string, val uint64) uint64 {
	v := c[name]
	if v <= val {
		delete(c, name)
		return 0
	}

	v -= val
	c[name] = v
	return v
}

func (c Counter) Set(name string, value uint64) {
	if value == 0 {
		delete(c, name)
	} else {
		c[name] = value
	}
}

func (c Counter) Get(name string) uint64 {
	return c[name]
}

func (c Counter) Slice(minVal uint64, orderDesc bool) []string {
	list := make([]string, 0, len(c))

	for k, v := range c {
		if v >= minVal {
			list = append(list, k)
		}
	}

	if orderDesc {
		sort.Slice(list, func(i, j int) bool {
			k1 := list[i]
			k2 := list[j]
			return c[k1] > c[k2] || c[k1] == c[k2] && k1 < k2
		})
	} else {
		sort.Slice(list, func(i, j int) bool {
			k1 := list[i]
			k2 := list[j]
			return c[k1] < c[k2] || c[k1] == c[k2] && k1 < k2
		})
	}

	return list
}

func (c Counter) Max() uint64 {
	res := uint64(0)

	for _, v := range c {
		if v > res {
			res = v
		}
	}

	return res
}

func (c Counter) Sum() uint64 {
	res := uint64(0)

	for _, v := range c {
		res += v
	}

	return res
}
