package partition

import (
	"reflect"
	"testing"
)

type testData struct {
	slice       any
	predicate   PredicateFunc
	expectedIdx int
	description string
}

func copyAnySlice(slice any) any {
	s := reflect.ValueOf(slice)
	copy := reflect.MakeSlice(s.Type(), s.Len(), s.Len())

	n := reflect.Copy(copy, s)
	if n != s.Len() {
		panic("copyAnySlice failed")
	}

	return copy.Interface()
}

func TestBasic(t *testing.T) {
	sizes := []int{13, 255, 512, 4096, 57361}

	for _, sz := range sizes {
		slice := make([]int, sz)
		n := 0
		i := Slice(slice, func(_ int) bool {
			n++
			return false
		})

		if n < len(slice) {
			t.Error("Slice has not looked at every element")
		}
		if i != 0 {
			t.Error("Slice should return 0 on contradiction predicate")
		}

		i = Slice(slice, func(_ int) bool { return true })
		if i != len(slice) {
			t.Error("Slice should return len(slice) on tautology predicate")
		}

		// b has to start at true to work with both odd and even sizes.
		b := true
		i = Slice(slice, func(_ int) bool {
			b = !b
			return b
		})
		if i != len(slice)/2 {
			t.Error("Slice should return len(slice)/2 when predicate alternates between true and false")
		}
	}

}

func TestPartitionIntSlice(t *testing.T) {

	slice1 := []int{1278, 27, 481, 437, 18, 49, -28, 127, 589, 294, 111, 586, 3781, 238, 345, 1, 4584}
	slice2 := []int{342, 107, -130, 256, 266, 106, 264, -48, -384, 193, 353, 96, 321, 245, 169, 370, 481, -58, 328, 193}
	slice3 := []int{-83, -354, 41, 445, -231, 269, 439, 205, 148, 392, -148, -283, 278, 182, 427, -443, -270, 209, -5, 277}

	tests := []testData{
		{
			slice:       slice1,
			predicate:   func(i int) bool { return slice1[i]%2 != 0 }, // isOdd
			expectedIdx: 10,
			description: "Partition slice1 with is-odd predicate.",
		},
		{
			slice:       slice2,
			predicate:   func(i int) bool { return slice2[i]%2 != 0 }, // isOdd
			expectedIdx: 8,
			description: "Partition slice2 with is-odd predicate.",
		},
		{
			slice:       slice3,
			predicate:   func(i int) bool { return slice3[i]%2 != 0 }, // isOdd
			expectedIdx: 13,
			description: "Partition slice3 with is-odd predicate.",
		},
		{
			slice:       slice1,
			predicate:   func(i int) bool { return slice1[i] < 0 }, // isNegative
			expectedIdx: 1,
			description: "Partition slice1 with is-negative predicate.",
		},
		{
			slice:       slice2,
			predicate:   func(i int) bool { return slice2[i] < 0 }, // isNegative
			expectedIdx: 4,
			description: "Partition slice2 with is-negative predicate.",
		},
		{
			slice:       slice3,
			predicate:   func(i int) bool { return slice3[i] < 0 }, // isNegative
			expectedIdx: 8,
			description: "Partition slice3 with is-negative predicate.",
		},
	}

	for _, test := range tests {
		copy := copyAnySlice(test.slice)
		i := Slice(test.slice, test.predicate)
		if i != test.expectedIdx {
			t.Errorf("Test failed: %s (Slice returned %d, expected %d)\n", test.description, i, test.expectedIdx)
			s := reflect.ValueOf(test.slice)
			if s.Len() < 30 {
				t.Errorf("result: %+v\n", s.Interface())
			}
			t.FailNow()
		}

		reflect.Copy(reflect.ValueOf(test.slice), reflect.ValueOf(copy))
	}

}
