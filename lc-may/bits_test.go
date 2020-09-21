package lcmay

import "testing"


func TestBIts1(t *testing.T) {
	var a []int
	var b []int
	a = []int{0,1,2,3,4,5,6,7,8}
	b = []int{0,1,1,2,1,2,2,3,1}

	bits := countBits(len(a))
	for i:=0;i<len(a);i++ {
		if bits[i] != b[i] {
			t.Error(i, ":", bits[i],"should be", b[i])
		}
	}
}

func BenchmarkBitsK(b *testing.B) {
	countBits(0x1000000)
}

func BenchmarkBitsC(b *testing.B) {
	countBitsC(0x1000000)
}


func BenchmarkBitsD(b *testing.B) {
	countBitsD(0x1000000)
}

func BenchmarkAll(b *testing.B) {
	b.Run("K",BenchmarkBitsK)
	b.Run("C",BenchmarkBitsC)
	b.Run("D",BenchmarkBitsD)
}