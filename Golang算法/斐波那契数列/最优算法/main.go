package main

import (
	"fmt"
	"math/big"
	"time"
)

//求矩阵的n次幂
func MatPow(a Matrix, b int) Matrix {
	arr0 := [4]*big.Int{big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(1)}
	s := New(2, 2, arr0[0:4])
	for b > 0 {
		if b&1 == 1 {
			s = *Multiply(s, a)
			b = b >> 1
		} else {
			b = b >> 1
		}
		a = *Multiply(a, a)
	}
	return s
}

//矩阵的N次幂与fib(1)和Fib(0)组成的2行1列的矩阵相乘求fib(n+1)和Fib(n)组成的2行1列的矩阵
//从fib(n+1)和Fib(n)的2行1列的矩阵中取出fib(n)
func Fib(n int) *big.Int {
	arr0 := [6]*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(0), big.NewInt(2), big.NewInt(1)}
	k := New(2, 2, arr0[0:4])
	s := MatPow(k, n)
	d := New(2, 1, arr0[0:2])
	s = *Multiply(s, d)
	return s.Get(2, 1)

}

func main() {

	start := time.Now()
	n := 1000000
	m := Fib(n)
	fmt.Println("f(n)的结果是", m)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}
