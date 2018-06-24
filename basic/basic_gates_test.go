package logicgates

import (
	"testing"
)

func Test_Nand(t *testing.T) {
	var testData = []struct {
		first    bool
		second   bool
		expected bool
	}{
		{false, false, true},
		{true, false, true},
		{true, true, false},
		{false, true, true},
	}
	for _, tt := range testData {
		if tt.expected != Nand(tt.first, tt.second) {
			t.Errorf("Not expected output. Nand(%v, %v) suppose to be %v.", tt.first, tt.second, tt.expected)
		}
	}
}

func Test_And(t *testing.T) {
	var testData = []struct {
		first    bool
		second   bool
		expected bool
	}{
		{false, false, false},
		{true, false, false},
		{true, true, true},
		{false, true, false},
	}
	for _, tt := range testData {
		if tt.expected != And(tt.first, tt.second) {
			t.Errorf("Not expected output. And(%v, %v) suppose to be %v.", tt.first, tt.second, tt.expected)
		}
	}
}

func Test_Not(t *testing.T) {
	var testData = []struct {
		first    bool
		expected bool
	}{
		{false, true},
		{true, false},
	}
	for _, tt := range testData {
		if tt.expected != Not(tt.first) {
			t.Errorf("Not expected output. Not(%v) suppose to be %v.", tt.first, tt.expected)
		}
	}
}

func Test_Or(t *testing.T) {
	var testData = []struct {
		first    bool
		second   bool
		expected bool
	}{
		{false, false, false},
		{true, false, true},
		{true, true, true},
		{false, true, true},
	}
	for _, tt := range testData {
		if tt.expected != Or(tt.first, tt.second) {
			t.Errorf("Not expected output. Or(%v, %v) suppose to be %v.", tt.first, tt.second, tt.expected)
		}
	}
}

func Test_Nor(t *testing.T) {
	var testData = []struct {
		first    bool
		second   bool
		expected bool
	}{
		{false, false, true},
		{true, false, false},
		{true, true, false},
		{false, true, false},
	}
	for _, tt := range testData {
		if tt.expected != Nor(tt.first, tt.second) {
			t.Errorf("Not expected output. Nor(%v, %v) suppose to be %v.", tt.first, tt.second, tt.expected)
		}
	}
}

func Test_And3(t *testing.T) {
	var testData = []struct {
		first    bool
		second   bool
		third    bool
		expected bool
	}{
		{false, false, true, false},
		{false, false, false, false},
		{false, true, true, false},
		{false, true, false, false},
		{true, false, true, false},
		{true, false, false, false},
		{true, true, true, true},
		{true, true, false, false},
	}
	for _, tt := range testData {
		if tt.expected != And3(tt.first, tt.second, tt.third) {
			t.Errorf("Not expected output. And3(%v, %v, %v) suppose to be %v.", tt.first, tt.second, tt.third, tt.expected)
		}
	}
}

func Test_Xor(t *testing.T) {
	var testData = []struct {
		first    bool
		second   bool
		expected bool
	}{
		{false, false, false},
		{true, false, true},
		{true, true, false},
		{false, true, true},
	}
	for _, tt := range testData {
		if tt.expected != Xor(tt.first, tt.second) {
			t.Errorf("Not expected output. Xor(%v, %v) suppose to be %v.", tt.first, tt.second, tt.expected)
		}
	}
}

func Test_Xnor(t *testing.T) {
	var testData = []struct {
		first    bool
		second   bool
		expected bool
	}{
		{false, false, true},
		{true, false, false},
		{true, true, true},
		{false, true, false},
	}
	for _, tt := range testData {
		if tt.expected != Xnor(tt.first, tt.second) {
			t.Errorf("Not expected output. Xnor(%v, %v) suppose to be %v.", tt.first, tt.second, tt.expected)
		}
	}
}

func Benchmark_And(b *testing.B) {

	for n := 0; n < b.N; n++ {
		And(true, true)
	}
}
func Benchmark_Nand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Nand(true, true)
	}
}
func Benchmark_Not(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Not(true)
	}
}
func Benchmark_Nor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Nor(true, true)
	}
}
func Benchmark_Or(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Or(true, true)
	}
}
func Benchmark_And3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		And3(true, true, false)
	}
}
func Benchmark_Xor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Xor(true, false)
	}
}
func Benchmark_Xnor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Xnor(true, false)
	}
}
