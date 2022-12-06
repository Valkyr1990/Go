package main

import "fmt"

func main() {
	var num, num1 int
	var q, w, e, r, t, y, u, i, o, p, a, s, d, f, g, h, j, k, l, z int
	fmt.Scan(&num)
	fmt.Scan(&num1)
	a = num1 / 100000
	s = num1 / 10000
	d = s % 10
	f = num1 / 1000
	g = f % 10
	h = num1 / 100
	j = h % 10
	k = num1 / 10
	l = k % 10
	z = num1 % 10
	q = num / 100000
	w = num / 10000
	e = w % 10
	r = num / 1000
	t = r % 10
	y = num / 100
	u = y % 10
	i = num / 10
	o = i % 10
	p = num % 10
	if (q == a || q == d || q == g || q == j || q == l || q == z) && q != 0 {
		fmt.Print(q, " ")
	}
	if (e == a || e == d || e == g || e == j || e == l || e == z) && e != 0 {
		fmt.Print(e, " ")
	}
	if (t == a || t == d || t == g || t == j || t == l || t == z) && t != 0 {
		fmt.Print(t, " ")
	}
	if (u == a || u == d || u == g || u == j || u == l || u == z) && u != 0 {
		fmt.Print(u, " ")
	}
	if (o == a || o == d || o == g || o == j || o == l || o == z) && o != 0 {
		fmt.Print(o, " ")
	}
	if (p == a || p == d || p == g || p == j || p == l || p == z) && p != 0 {
		fmt.Print(p, " ")
	}

}
