package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alandtsang/gormgen/biz"
)

func main() {
	contact, err := biz.FetchContactByID(context.Background(), 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("contact: %+v\n", contact)
}
