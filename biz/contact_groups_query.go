package biz

import (
	"context"
	"fmt"
	"log"

	"github.com/alandtsang/gormgen/dal"
	"github.com/alandtsang/gormgen/dal/model"
	"github.com/alandtsang/gormgen/dal/query"
)

func ContactGroupsCrud() {
	listContactGroupsDemo()
}

func listContactGroupsDemo() {
	ctx := context.Background()

	contactGroupIDs := []int{1}
	modelContactGroups, err := listContactGroups(ctx, contactGroupIDs)
	if err != nil {
		log.Fatal(err)
	}
	for _, con := range modelContactGroups {
		fmt.Printf("contact group: %+v\n", con)
	}
}

func listContactGroups(ctx context.Context, contactGroupsIDs []int) (contactGroups []*model.ContactGroups, err error) {
	c := query.Use(dal.DB).ContactGroups
	if len(contactGroupsIDs) > 0 {
		contactGroups, err = c.WithContext(ctx).WithContext(ctx).Where(c.ID.In(contactGroupsIDs...)).Find()
	} else {
		contactGroups, err = c.WithContext(ctx).WithContext(ctx).Find()
	}
	if err != nil {
		fmt.Printf("List contact groups failed, %v", err)
		return nil, err
	}
	return contactGroups, nil
}
