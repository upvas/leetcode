package NextPermutation_31

import (
	"testing"
	"reflect"
)

func TestSort(t *testing.T)  {
	tests := []struct {
		arr []int
		want []int
	}{
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		got := make([]int, len(test.arr))
		copy(got, test.arr)
		sort(got)
		if want := test.want; !reflect.DeepEqual(want, got) {
			t.Fatalf("sort(%v) == %v, want %v", test.arr, got, want)
		}
	}
}