package object

import "testing"

func TestHashKey(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		testHashKey(t,
			&String{Value: "Hello World"},
			&String{Value: "Hello World"},
			&String{Value: "My name is johnny"},
			&String{Value: "My name is johnny"},
		)
	})

	t.Run("Integer", func(t *testing.T) {
		testHashKey(t,
			&Integer{Value: 1},
			&Integer{Value: 1},
			&Integer{Value: 2},
			&Integer{Value: 2},
		)
	})

	t.Run("Boolean", func(t *testing.T) {
		testHashKey(t,
			&Boolean{Value: true},
			&Boolean{Value: true},
			&Boolean{Value: false},
			&Boolean{Value: false},
		)
	})
}

// testHashKey asserts that a1 and a2 hash equally, b1 and b2 hash equally,
// and that the a and b groups hash differently.
func testHashKey(t *testing.T, a1, a2, b1, b2 Hashable) {
	t.Helper()
	if a1.HashKey() != a2.HashKey() {
		t.Errorf("same-content objects have different hash keys")
	}
	if b1.HashKey() != b2.HashKey() {
		t.Errorf("same-content objects have different hash keys")
	}
	if a1.HashKey() == b1.HashKey() {
		t.Errorf("different-content objects have same hash keys")
	}
}
