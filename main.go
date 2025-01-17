package main

import (
	"fmt"

	"github.com/docker/distribution/uuid"
	"github.com/grangerp/test-pkg-repo/tenant"
)

func main() {
	t := tenant.Tenant{}
	_ = uuid.Size
	fmt.Println(t)
}
