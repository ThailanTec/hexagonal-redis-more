package pkg

import "math/rand"

func GerenrateID() int {
	n := rand.Intn(10000000000) + 10000000000

	return n
}
