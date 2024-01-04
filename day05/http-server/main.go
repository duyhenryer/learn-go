// You can edit this code!
// Click here and start typing.
package main

import (
	"ct-backend-course-baonguyen/server"
)

func main() {
	server.StartServer()
}

func variadicSum(nums ...int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return
}
