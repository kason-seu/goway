package slicesop

import "fmt"

func slicesop_fn() {

	u := []int{1, 2, 3, 4, 5}

	fmt.Println("u", u, "len :", len(u), "cap : ", cap(u))

	s := u[1:3]
	fmt.Println("s", s, "len :", len(s), "cap : ", cap(s))

	s = append(s, 6)
	fmt.Println("s", s, "len :", len(s), "cap : ", cap(s))

	s = append(s, 7)
	fmt.Println("s", s, "len :", len(s), "cap : ", cap(s))

	s = append(s, 8)
	fmt.Println("s", s, "len :", len(s), "cap : ", cap(s))

}
