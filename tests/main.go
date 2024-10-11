package main

import (
	"fmt"

	"github.com/phongnd2802/go-ecommerce-backend-api/pkg/utils/crypto"
)

func main() {
	ranStr, _ := crypto.RandomSalt(16)
	fmt.Println(ranStr)
}