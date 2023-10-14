package testing_example

import (
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	// askTilakForANumber() - Arrange
	ans := IntMin(2, -2, 5) // Act
	if ans != -2 {          // Assert
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
	// Cleanup
}

func ExampleIntMin() {
	fmt.Println(IntMin(2, -2, -2))
	fmt.Println(IntMin(4, 0, 0))
	fmt.Println(IntMin(100, 10000, 2))
	fmt.Println(IntMin(-5, -7, -2))
	// Output:
	// -2
	// 0
	// 2
	// -7
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b, c int
		want    int
	}{
		{0, 1, 5, 0},
		{1, 0, -3, -3},
		{2, -2, -2, -2},
		{0, -1, -2, -2},
		{-1, 0, -2, -2},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b, tt.c)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IntMin(1, 2, 4)
	}
}
