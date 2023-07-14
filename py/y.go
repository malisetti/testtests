package py

import (
	"fmt"
	"math/rand"
	"time"
)

func hi() error {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		return fmt.Errorf("error hi")
	}
	return nil
}
