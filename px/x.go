package px

import (
	"fmt"
	"math/rand"
	"time"
)

func hello() error {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		return fmt.Errorf("error hello")
	}
	return nil
}
