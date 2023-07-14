package py

import "testing"

func Test_hi1(t *testing.T) {
	t.Run("hi1", func(t *testing.T) {
		err := hi()
		if err != nil {
			t.Error(err)
		}
	})
}
