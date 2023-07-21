package somagenerica_test

import (
	"testing"

	"github.com/flan6/estudos/somagenerica"
)

func TestSum(t *testing.T) {
	tests := map[string]struct {
		a    any
		b    any
		want any
	}{
		"success": {1, 1.2, 2.2},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			switch test.want.(type) {
			case int:
				got, err := somagenerica.Sum[int](test.a, test.b)
				if err != nil {
					t.Fatal(err)
				}
				if got != test.want {
					t.Errorf("Sum() = %v, want %v", got, test.want)
				}

			case float64:
				got, err := somagenerica.Sum[float64](test.a, test.b)
				if err != nil {
					t.Fatal(err)
				}
				if got != test.want {
					t.Errorf("Sum() = %v, want %v", got, test.want)
				}
			}
		})
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		somagenerica.Sum[int](float64(1000.1), 1000)
	}
}
