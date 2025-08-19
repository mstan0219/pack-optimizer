package pack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCorrect(t *testing.T) {
	tests := []struct {
		ordered int
		want    map[int]int
	}{
		{1, map[int]int{250: 1}},
		{250, map[int]int{250: 1}},
		{500, map[int]int{500: 1}},
		{750, map[int]int{500: 1, 250: 1}},
		{1000, map[int]int{1000: 1}},
		{12001, map[int]int{5000: 2, 2000: 1, 250: 1}},
		{4999, map[int]int{5000: 1}},
		{9999, map[int]int{5000: 2}},
		{10000, map[int]int{5000: 2}},
		{500000, map[int]int{5000: 100}},
		{12500, map[int]int{5000: 2, 2000: 1, 500: 1}},
		{15000, map[int]int{5000: 3}},
		{2250, map[int]int{2000: 1, 250: 1}},
		{3750, map[int]int{2000: 1, 1000: 1, 500: 1, 250: 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Size %d", tt.ordered), func(t *testing.T) {
			got := Correct(tt.ordered, GetSizes())

			if len(got) != len(tt.want) {
				t.Errorf("cuurect(%d) = %v, want %v", tt.ordered, got, tt.want)
				return
			}
			for pack, count := range tt.want {
				if gotCount, exists := got[pack]; !exists || gotCount != count {
					t.Errorf("cuurect(%d)[%d] = %d, want %d", tt.ordered, pack, gotCount, count)
				}
			}
		})
	}
}

func TestCorrectCustomSize(t *testing.T) {
	tests := []struct {
		ordered int
		want    map[int]int
	}{
		{500000, map[int]int{53: 9429, 31: 7, 23: 2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Size %d", tt.ordered), func(t *testing.T) {
			got := Correct(tt.ordered, []int{
				23, 31, 53,
			})

			if len(got) != len(tt.want) {
				t.Errorf("cuurect(%d) = %v, want %v", tt.ordered, got, tt.want)
				return
			}
			for pack, count := range tt.want {
				if gotCount, exists := got[pack]; !exists || gotCount != count {
					t.Errorf("cuurect(%d)[%d] = %d, want %d", tt.ordered, pack, gotCount, count)
				}
			}
		})
	}
}

func TestInCorrect(t *testing.T) {
	SetSizes([]int{
		5000, 2000, 1000, 500, 250,
	})
	tests := []struct {
		ordered int
		want    []map[int]int
	}{
		{
			ordered: 1,
			want: []map[int]int{
				{5000: 1},
				{2000: 1},
				{1000: 1},
			},
		},
		{
			ordered: 251,
			want: []map[int]int{
				{5000: 1},
				{2000: 1},
				{1000: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Size %d", tt.ordered), func(t *testing.T) {
			got := InCorrect(tt.ordered, GetSizes())
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InCorrect(%d) = %v, want %v", tt.ordered, got, tt.want)
			}
		})
	}
}
