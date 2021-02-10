package service

import "testing"

func TestStub(t *testing.T) {
	dat := []ArithmeticDS{
		{A: 20, B: 12, OP: '+'},
		{A: 20, B: 13, OP: '-'},
		{A: 12, B: 15, OP: '*'},
		{A: 72, B: 12, OP: '/'},
	}

	expectedResult := []int64{
		20 + 12,
		20 - 13,
		12 * 15,
		72 / 12,
	}

	for i, v := range dat {
		result := GetResult(&v)
		if result != expectedResult[i] {
			t.Errorf("Case %d: unexpected value: %v expected %v", i, result, expectedResult[i])
		}
	}
}
