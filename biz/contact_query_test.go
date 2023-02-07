package biz

import (
	"context"
	"testing"
	"time"

	"github.com/alandtsang/gormgen/dal/model"
	"github.com/stretchr/testify/assert"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

var now time.Time

func init() {
	timeStr := "2023-02-01 17:35:00"
	now, _ = time.Parse(timeFormat, timeStr)
}

func TestCreateContact(t *testing.T) {
	ctx := context.Background()
	Tom := &Contact{
		Name:            "Tom",
		Mobile:          "15011111111",
		MobileConfirmed: 1,
		Email:           "tom@gmail.com",
		EmailConfirmed:  0,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	err := CreateContact(ctx, Tom)
	assert.Nil(t, err)

	Jerry := &Contact{
		Name:            "Jerry",
		Mobile:          "15022222222",
		MobileConfirmed: 0,
		Email:           "jerry@gmail.com",
		EmailConfirmed:  1,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	err = CreateContact(ctx, Jerry)
	assert.Nil(t, err)
}

func TestListContactsByNames(t *testing.T) {
	ctx := context.Background()
	names := []string{"Tom", "Jerry"}

	expect := []*model.Contact{
		{
			ID:              1,
			Name:            "Tom",
			Mobile:          "15011111111",
			MobileConfirmed: 1,
			Email:           "tom@gmail.com",
			EmailConfirmed:  0,
			CreatedAt:       now,
			UpdatedAt:       now,
		},
		{
			ID:              2,
			Name:            "Jerry",
			Mobile:          "15022222222",
			MobileConfirmed: 0,
			Email:           "jerry@gmail.com",
			EmailConfirmed:  1,
			CreatedAt:       now,
			UpdatedAt:       now,
		},
	}

	got, err := ListContactsByNames(ctx, names...)
	assert.Nil(t, err)
	if !assert.ObjectsAreEqual(expect, got) {
		t.Errorf("TestListContactsByNames failed, expect: %v, got: %v", expect, got)
	}
}

func TestFetchContactByID(t *testing.T) {
	ctx := context.Background()
	var contactID uint64 = 1

	expect := &model.Contact{
		ID:              1,
		Name:            "Tom",
		Mobile:          "15011111111",
		MobileConfirmed: 1,
		Email:           "tom@gmail.com",
		EmailConfirmed:  0,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	got, err := FetchContactByID(ctx, contactID)
	assert.Nil(t, err)
	assert.Equal(t, expect, got)
}

func TestUpdateContact(t *testing.T) {
	ctx := context.Background()

	var contactID uint64 = 1
	updates := map[string]interface{}{
		"email_confirmed": 1,
	}

	err := UpdateContact(ctx, contactID, updates)
	assert.Nil(t, err)
}

func TestDeleteContacts(t *testing.T) {
	ctx := context.Background()
	contactIDs := []uint64{1, 2}

	err := DeleteContacts(ctx, contactIDs)
	assert.Nil(t, err)
}
