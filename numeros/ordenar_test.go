package numeros

import "testing"

func TestSumPares_nil_array(t *testing.T) {
	var numeros []int = nil
	result := 0
	r := SumPares(numeros)
	if r != result {
		t.Errorf("Se esperaba: %d, Se obtuvo: %d", result, r)
	}
}
func TestSumPares_empty_array(t *testing.T) {
	var numbers []int
	result := 0
	r := SumPares(numbers)
	if r != result {
		t.Errorf("Se esperaba: %v, y se obtuvo: %v", result, r)
	}
}
func TestSumPares_inpar_numbers(t *testing.T) {
	numbers := []int{1, 3, 5}
	result := 0
	r := SumPares(numbers)
	if r != result {
		t.Errorf("Se esperaba %v y se obtuvo %v", result, r)
	}
}
func TestSumPares_mix_array(t *testing.T) {
	numbers := []int{2, 4, 5}
	result := 6
	r := SumPares(numbers)
	if r != result {
		t.Errorf("Se esperaba: %v, se obtuvo: %v", result, r)
	}
}
func TestSumPares_one_number(t *testing.T) {
	numbers := []int{2}
	result := 2
	r := SumPares(numbers)
	if r != result {
		t.Errorf("Se esperaba %v y se obtuvo %v", result, r)
	}
}

func TestSumParesGeneral(t *testing.T) {
	datosTest := map[string]struct {
		numbers []int
		result  int
	}{
		"NilArray":       {nil, 0},
		"EmptyArray":     {[]int{}, 0},
		"InparNumbers":   {[]int{1, 3, 5}, 0},
		"MixArray":       {[]int{2, 4, 5}, 6},
		"OneNumber":      {[]int{2}, 2},
		"OneNumberInpar": {[]int{3}, 0},
	}
	for name, dt := range datosTest {
		t.Run(name, func(t *testing.T) {
			r := SumPares(dt.numbers)
			if r != dt.result {
				t.Errorf("Se esperaba: %v, y se obtuvo: %v", dt.result, r)
			}
		})
	}
}
