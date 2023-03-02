package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"os"
)

func part1() { // Network intensive
	for i := 0; i < 10; i++ {
		resp, err := http.Get("https://google.com")
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
	}

}

func part2() { // CPU intensive
	Pi(100000)
}

func part3() { //Disk intensive
	f, err := os.CreateTemp("", "*")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10000; i++ {
		data := make([]byte, 100)
		_, err := rand.Read(data)
		if err != nil {
			panic(err)
		}
		_, err = f.Write(data)
		if err != nil {
			panic(err)
		}
		err = f.Sync()
		if err != nil {
			panic(err)
		}
	}

	f.Close()
	err = os.Remove(f.Name())
	if err != nil {
		panic(err)
	}

}

func arccot(x int64, unity *big.Int) *big.Int {
	bigx := big.NewInt(x)
	xsquared := big.NewInt(x * x)
	sum := big.NewInt(0)
	sum.Div(unity, bigx)
	xpower := big.NewInt(0)
	xpower.Set(sum)
	n := int64(3)
	zero := big.NewInt(0)
	sign := false

	term := big.NewInt(0)
	for {
		xpower.Div(xpower, xsquared)
		term.Div(xpower, big.NewInt(n))
		if term.Cmp(zero) == 0 {
			break
		}
		if sign {
			sum.Add(sum, term)
		} else {
			sum.Sub(sum, term)
		}
		sign = !sign
		n += 2
	}
	return sum
}

func Pi(ndigits int64) string {
	if ndigits <= 7 {
		return "3.141595"
	} else {
		digits := big.NewInt(ndigits + 10)
		unity := big.NewInt(0)
		unity.Exp(big.NewInt(10), digits, nil)
		pi := big.NewInt(0)
		four := big.NewInt(4)

		pi.Mul(four, pi.Sub(pi.Mul(four, arccot(5, unity)), arccot(239, unity)))
		output := fmt.Sprintf("%s.%s", pi.String()[0:1], pi.String()[1:ndigits])
		return output
	}
}
