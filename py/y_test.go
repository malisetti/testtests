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

func Test_hi2(t *testing.T) {
	t.Run("hi2", func(t *testing.T) {
		err := hi()
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_hi3(t *testing.T) {
	t.Run("hi3", func(t *testing.T) {
		err := hi()
		if err != nil {
			t.Error(err)
		}
	})
}
