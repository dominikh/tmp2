package pkg

import "time"

func Fn() {
	time.Parse("123", "456")
}

func whatever() {}