package main

import (
	"fmt"
	"time"

	us "./user"
)

type matias struct {
	us.User
}

func main() {
	u := new(matias)
	u.CompleteUser(1, "leon matias", time.Now(), true)

	fmt.Println(u.User)
}
