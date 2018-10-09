package main

import (
	"fmt"
	"math"
)

//func phi(x float64) float64 {
//	return (1./3)*math.Pow(x, 4) - (4./3)*math.Pow(x, 3) + (11./6)*math.Pow(x, 2) + (1. / 6)
//}
//
//func diffPhi(x float64) float64 {
//	return (4./3)*math.Pow(x, 3) - 4*math.Pow(x, 2) + (11./3)*x
//}
//
//func diff2Phi(x float64) float64 {
//	return 4*math.Pow(x, 2) - 8*x + 11./3
//}
//func phi(x float64) float64 {
//	return 1/3000.*(math.Pow(x, 4) - 4.* math.Pow(x, 3) + (11/2.)*math.Pow(x, 2) - 3*x + (1. / 2)) + x
//}
//
//func diffPhi(x float64) float64 {
//	return 1/3000.*(4.*math.Pow(x, 3) - 12*math.Pow(x, 2) + 11.*x - 3) +1
//}
//
//func diff2Phi(x float64) float64 {
//	return 1/3000.*(8.*math.Pow(x, 2) - 24*x + 11.)
//}
func phi(x float64) float64 {
	return math.Pow(4.*math.Pow(x, 3)-(11/2.)*math.Pow(x, 2)+3*x-(1./2), 1./4)
}

func diffPhi(x float64) float64 {
	return (phi(x+0.00001) - phi(x-0.00001)) / 0.00002
	//return 1/4.*math.Pow(4.* math.Pow(x, 3) - (11/2.)*math.Pow(x, 2) + 3*x - (1. / 2), -3./4)*(12.* math.Pow(x, 2) - 11*x + 3)
}

func diff2Phi(x float64) float64 {
	return (diffPhi(x+0.00001) - diffPhi(x-0.00001)) / 0.00002
	//return 0
	//return 1/3000.*(8.*math.Pow(x, 2) - 24*x + 11.)
}

func f(x float64) float64 {
	return math.Pow(x, 4) - 4.*math.Pow(x, 3) + (11/2.)*math.Pow(x, 2) - 3*x + (1. / 2)
}

func diffF(x float64) float64 {
	return (phi(x+0.00001) - phi(x-0.00001)) / 0.00002
	//return 1/4.*math.Pow(4.* math.Pow(x, 3) - (11/2.)*math.Pow(x, 2) + 3*x - (1. / 2), -3./4)*(12.* math.Pow(x, 2) - 11*x + 3)
}

func diff2F(x float64) float64 {
	return (diffPhi(x+0.00001) - diffPhi(x-0.00001)) / 0.00002
	//return 0
	//return 1/3000.*(8.*math.Pow(x, 2) - 24*x + 11.)
}

func q(a, b float64) float64 {
	if math.Abs(diffPhi(a)) > math.Abs(diffPhi(b)) {
		return math.Abs(diffPhi(a))
	}
	return math.Abs(diffPhi(b))
}

func x0(a, b float64) float64 {
	return (a + b) / 2
}

func check(a, b float64) bool {

	currX0 := x0(a, b)

	m := func(x0 float64) float64 {
		var x1 float64
		x1 = phi(x0)
		return math.Abs(x1 - x0)
	}

	currM := m(currX0)

	rho := func(a, x0 float64) float64 {
		return math.Abs(x0 - a)
	}

	currRho := rho(a, currX0)

	if currM/(1-q(a, b)) > currRho {
		return false
	}

	return true
}

func iter(a, b, q, epsilon float64) int {
	return int(math.Abs(math.Log((b-a)/(epsilon*(1-q)))/math.Log(1/q)) + 1)
}

func FixedPointMethod(a, b float64, maxIter int, epsilon float64) float64 {
	//var currEpsilon, currX0 float64
	currX0 := x0(a, b)
	curr := currX0
	next := currX0
	for i := 0; i < maxIter; i++ {
		next = phi(curr)
		//fmt.Printf("| x%v | : | %.10f |\n", i+1, next)
		//currEpsilon = math.Abs(curr - next)
		curr = next
		//if currEpsilon > 0.644*epsilon/0.366 {
		//	break
		//}
	}
	//if currEpsilon > epsilon {
	//fmt.Printf("Solution not found!")
	//panic(fmt.Sprintf("%f.4 > %f.4", currEpsilon, epsilon))
	//}
	return curr
}

func methodStefen(a, b, epsilon float64) float64 {
	var X0, x1, x2, x, currEpsilon, i float64
	X0 = x0(a, b)
	currEpsilon = 1
	for i = 0; currEpsilon > epsilon && X0-2*x1+x2 != 0; i++ {
		x1 = phi(X0)
		x2 = phi(x1)
		x = (X0*x2 - math.Pow(x1, 2)) / (X0 - 2*x1 + x2)
		currEpsilon = math.Abs(x - X0)
		//fmt.Printf("| x%v | : | %.10f |\n", i+1, x)
		X0 = x1
	}
	return x
}

func chord(a, b, epsilon float64) float64 {
	var currEpsilon, x1, x2, i float64
	x1 = a
	currEpsilon = 1000
	for i = 0; currEpsilon > epsilon; i++ {
		x2 = x1 - f(x1)*(x1-b)/(f(x1)-f(b))
		currEpsilon = math.Abs(x2 - x1)
		//fmt.Printf("| x%v | : | %.10f |\n", i+1, x2)
		x1 = x2
	}
	return x1
}

func main() {
	fmt.Printf("Checking : %t\n", check(2, 1))
	fmt.Sprintf("Testing Fixed Point iteration")
	epsilon := 0.000001
	b := 2.
	a := 1.5
	qq := q(a, b)
	n := iter(a, b, qq, epsilon)
	x := FixedPointMethod(a, b, n, epsilon)
	fmt.Printf("Our q: %.3f\n", qq)
	fmt.Printf("Our n: %d\n", n)
	fmt.Printf("Solution from fixed point method: %.7f\n", x)

	fmt.Printf("Solution from stefen method: %.7f\n", methodStefen(a, b, epsilon))

	fmt.Printf("Chord sign: %f.3\n", diffF((a+b)/2)*diff2F((a+b)/2))
	fmt.Printf("Chord sign: %f.3\n", diffF(a)*diff2F(a))
	fmt.Printf("Chord sign: %f.3\n", diffF(b)*diff2F(b))

	fmt.Printf("Solution from chord method: %.7f\n", chord(a, b, epsilon))
	fmt.Printf("Our phi(a): %.3f\n", phi(a))
	fmt.Printf("Our phi(b): %.3f\n", phi(b))
	fmt.Printf("Our diffphi(x): %.3f\n", diff2Phi(1.8))
	fmt.Printf("Our diff()x-diff2(x): %.3f\n", diffPhi(1.8)-diff2Phi(1.8))
}
