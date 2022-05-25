# gormgen
go-gorm gen demo

## Documentation

[go-gorm/gen README](https://github.com/go-gorm/gen/blob/master/README.md)

## Demo

Step 1: Define the model


```go
type Contacts struct {
	ID              int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name            string    `gorm:"column:name"`
	Mobile          string    `gorm:"column:mobile"`
	MobileConfirmed int       `gorm:"column:mobile_confirmed"`
	Email           string    `gorm:"column:email"`
	EmailConfirmed  int       `gorm:"column:email_confirmed"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (Contacts) TableName() string {
	return "contacts"
}
```

Step 2: Generate the code by gorm 

```bash
make model && ./model
```
```
GO111MODULE=on go build -o model cmd/modelgen/modelgen.go
2022/05/25 14:38:06 Start generating code.
2022/05/25 14:38:06 generate query file: /Users/aland/gowork/src/github.com/alandtsang/gormgen/dal/query/contacts.gen.go
2022/05/25 14:38:06 generate query file: /Users/aland/gowork/src/github.com/alandtsang/gormgen/dal/query/gen.go
2022/05/25 14:38:06 Generate code done.
```

Step 3: Compile the demo code

```bash
make build && ./main
```

```
GO111MODULE=on go build -o main cmd/gormgen/gormgen.go

2022/05/25 14:38:40 /Users/aland/gowork/src/github.com/alandtsang/gormgen/vendor/gorm.io/gen/do.go:603
[5.839ms] [rows:2] SELECT * FROM `contacts` WHERE `contacts`.`id` IN (1,2,4)
contact: &{ID:1 Name:Alan Mobile: MobileConfirmed:0 Email: EmailConfirmed:0 CreatedAt:2022-05-12 02:59:39 +0000 UTC UpdatedAt:2022-05-12 08:38:15 +0000 UTC}
contact: &{ID:2 Name:Alan Mobile: MobileConfirmed:0 Email: EmailConfirmed:0 CreatedAt:2022-05-12 08:14:30 +0000 UTC UpdatedAt:2022-05-12 08:14:30 +0000 UTC}
```

## License

Please refer to [LICENSE](https://github.com/alandtsang/gormgen/blob/master/LICENSE) file.
