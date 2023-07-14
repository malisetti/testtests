package px

import "testing"

func Test_hello1(t *testing.T) {
	t.Run("hello1", func(t *testing.T) {
		err := hello()
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_hello2(t *testing.T) {
	t.Run("hello2", func(t *testing.T) {
		err := hello()
		if err != nil {
			t.Error(err)
		}
	})
}

func Test_hello3(t *testing.T) {
	t.Run("hello3", func(t *testing.T) {
		err := hello()
		if err != nil {
			t.Error(err)
		}
	})
}
