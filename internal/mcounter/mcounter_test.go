package mcounter_test

import (
	"strings"
	"testing"

	"github.com/wmentor/lang/internal/mcounter"
)

func TestMCounter(t *testing.T) {
	c := mcounter.New()
	if c == nil {
		t.Fatal("New failed")
	}

	tSet := func(name string, val uint64) {
		c.Set(name, val)
		if c.Get(name) != val {
			t.Fatalf("Set/Get failed for key=%s value=%d", name, val)
		}
	}

	tSlice := func(min uint64, desc bool, wait []string) {
		if strings.Join(c.Slice(min, desc), ", ") != strings.Join(wait, ", ") {
			t.Fatalf("Slice failed for min=%d desc=%t wait=%s", min, desc, strings.Join(wait, ", "))
		}
	}

	tMax := func(wait uint64) {
		if c.Max() != wait {
			t.Fatal("Max failed")
		}
	}

	tSum := func(wait uint64) {
		if c.Sum() != wait {
			t.Fatal("Sum failed")
		}
	}

	tInc := func(name string, inc uint64, wait uint64) {
		if c.Inc(name, inc) != wait {
			t.Fatalf("Inc failed for name=%s inc=%d wait=%d", name, inc, wait)
		}
	}

	tDec := func(name string, inc uint64, wait uint64) {
		if c.Dec(name, inc) != wait {
			t.Fatalf("Dec failed for name=%s inc=%d wait=%d", name, inc, wait)
		}
	}

	tSet("a", 3)
	tSet("b", 12)
	tSet("c", 9)
	tSet("d", 1)

	tSlice(0, false, []string{"d", "a", "c", "b"})
	tSlice(0, true, []string{"b", "c", "a", "d"})
	tSlice(4, false, []string{"c", "b"})

	tMax(12)
	tSum(25)

	tInc("d", 9, 10)
	tInc("f", 0, 0)

	tSlice(0, false, []string{"a", "c", "d", "b"})
	tSlice(0, true, []string{"b", "d", "c", "a"})

	tInc("f", 1, 1)

	tSlice(0, false, []string{"f", "a", "c", "d", "b"})
	tSlice(0, true, []string{"b", "d", "c", "a", "f"})

	tDec("a", 1, 2)

	tSlice(0, false, []string{"f", "a", "c", "d", "b"})
	tSlice(0, true, []string{"b", "d", "c", "a", "f"})

	tDec("a", 2, 0)

	tSlice(0, false, []string{"f", "c", "d", "b"})
	tSlice(0, true, []string{"b", "d", "c", "f"})

	tSet("c", 0)

	tSlice(0, false, []string{"f", "d", "b"})
	tSlice(0, true, []string{"b", "d", "f"})
}
