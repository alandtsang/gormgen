package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/alandtsang/gormgen/dal"
	"github.com/alandtsang/gormgen/dal/model"
	"github.com/alandtsang/gormgen/dal/query"
)

type Contact struct {
	Name            string
	Mobile          string
	MobileConfirmed int
	Email           string
	EmailConfirmed  int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func CreateContact(ctx context.Context, contact *Contact) (err error) {
	contactModel := &model.Contact{
		Name:            contact.Name,
		Mobile:          contact.Mobile,
		MobileConfirmed: contact.MobileConfirmed,
		Email:           contact.Email,
		EmailConfirmed:  contact.EmailConfirmed,
		CreatedAt:       contact.CreatedAt,
		UpdatedAt:       contact.UpdatedAt,
	}

	c := query.Use(dal.DB).Contact
	if err = c.WithContext(ctx).Create(contactModel); err != nil {
		fmt.Printf("Create contact failed, %v", err)
		return err
	}
	return err
}

func ListContactsByNames(ctx context.Context, names ...string) (contacts []*model.Contact, err error) {
	c := query.Use(dal.DB).Contact

	contacts, err = c.WithContext(ctx).ListByNames(names)
	if err != nil {
		fmt.Printf("List contacts failed, %v", err)
		return nil, err
	}
	return contacts, nil
}

func FetchContactByID(ctx context.Context, contactID uint64) (*model.Contact, error) {
	c := query.Use(dal.DB).Contact

	contact, err := c.WithContext(ctx).GetByID(contactID)
	if err != nil {
		fmt.Printf("Fetch contact failed, %v\n", err)
		return nil, err
	}
	return contact, nil
}

func UpdateContact(ctx context.Context, contactID uint64, updates map[string]interface{}) (err error) {
	c := query.Use(dal.DB).Contact

	if len(updates) > 0 {
		if _, err = c.WithContext(ctx).Where(c.ID.Eq(contactID)).Updates(updates); err != nil {
			fmt.Printf("Update contact failed, %v", err)
			return err
		}
	}
	return nil
}

func DeleteContacts(ctx context.Context, contactIDs []uint64) (err error) {
	c := query.Use(dal.DB).Contact

	if err = c.WithContext(ctx).DeleteByIDs(contactIDs); err != nil {
		fmt.Printf("Delete contact failed, %v", err)
		return err
	}
	return nil
}
