package main

import (
	"e-comerce.com/m/internal/routers"
)

func main() {
  r := routers.NewRouter();

  r.Run(":8002")
}
