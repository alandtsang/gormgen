package biz

import (
	"context"
	"fmt"
	"github.com/alandtsang/gormgen/dal"
	"github.com/alandtsang/gormgen/dal/model"
	"github.com/alandtsang/gormgen/dal/query"
	"gorm.io/gen"
	"log"
	"time"
)

type Contact struct {
	Name            string
	Mobile          string
	MobileConfirmed int
	Email           string
	EmailConfirmed  int
}

func ContactsCrud() {
	//createContactsDemo()
	listContactsDemo()
	//getContactsDemo()
	//updateContactsDemo()
	//deleteContactsDemo()
}

func createContactsDemo() {
	ctx := context.Background()

	contacts := []*Contact{
		{
			Name: "Alan",
		},
		{
			Name: "Tom",
		},
	}

	if err := createContacts(ctx, contacts); err != nil {
		log.Fatal(err)
	}
}

func listContactsDemo() {
	ctx := context.Background()

	contactIDs := []int{1, 2, 4}

	modelContacts, err := listContacts(ctx, contactIDs)
	if err != nil {
		log.Fatal(err)
	}
	for _, con := range modelContacts {
		fmt.Printf("contact: %+v\n", con)
	}
}

func getContactsDemo() {
	ctx := context.Background()
	contactID := 4

	contact, err := fetchContacts(ctx, contactID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("contact: %+v\n", contact)
}

func updateContactsDemo() {
	ctx := context.Background()

	contactID := 1

	if err := updateContacts(ctx, contactID, nil); err != nil {
		log.Fatal(err)
	}
}

func deleteContactsDemo() {
	ctx := context.Background()

	contactIDs := []int{5}

	if err := deleteContacts(ctx, contactIDs); err != nil {
		log.Fatal(err)
	}
}

func createContacts(ctx context.Context, contacts []*Contact) (err error) {
	var contactModels []*model.Contacts

	now := time.Now()

	for _, con := range contacts {
		contactModels = append(contactModels, &model.Contacts{
			Name:            con.Name,
			Mobile:          con.Mobile,
			MobileConfirmed: con.MobileConfirmed,
			Email:           con.Email,
			EmailConfirmed:  con.EmailConfirmed,
			CreatedAt:       now,
			UpdatedAt:       now,
		})
	}

	c := query.Use(dal.DB).Contacts
	if err = c.WithContext(ctx).Create(contactModels...); err != nil {
		fmt.Printf("Create contact failed, %v", err)
		return err
	}
	return err
}

func listContacts(ctx context.Context, contactIDs []int) (contacts []*model.Contacts, err error) {
	c := query.Use(dal.DB).Contacts
	if len(contactIDs) > 0 {
		contacts, err = c.WithContext(ctx).WithContext(ctx).Where(c.ID.In(contactIDs...)).Find()
	} else {
		contacts, err = c.WithContext(ctx).WithContext(ctx).Find()
	}
	if err != nil {
		fmt.Printf("List contacts failed, %v", err)
		return nil, err
	}
	return contacts, nil
}

func fetchContacts(ctx context.Context, contactID int) (*model.Contacts, error) {
	c := query.Use(dal.DB).Contacts
	contact, err := c.WithContext(ctx).Where(c.ID.Eq(contactID)).First()
	if err != nil {
		fmt.Printf("Fetch contact failed, %v\n", err)
		return nil, err
	}
	return contact, nil
}

func updateContacts(ctx context.Context, contactID int, updates map[string]interface{}) (err error) {
	c := query.Use(dal.DB).Contacts

	if len(updates) > 0 {
		var result gen.ResultInfo
		result, err = c.WithContext(ctx).Where(c.ID.Eq(contactID)).Updates(updates)
		if err != nil {
			fmt.Printf("Update contact failed, %v", err)
			return err
		}
		fmt.Printf("update result: %+v\n", result)
	}
	return nil
}

func deleteContacts(ctx context.Context, contactIDs []int) (err error) {
	c := query.Use(dal.DB).Contacts
	if _, err = c.WithContext(ctx).Where(c.ID.In(contactIDs...)).Delete(); err != nil {
		fmt.Printf("Delete contact failed, %v", err)
		return err
	}
	return nil
}
