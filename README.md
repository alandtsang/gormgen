# gormgen

go-gorm gen demo

## Documentation

[go-gorm/gen README](https://github.com/go-gorm/gen/blob/master/README.md)

## Demo

Step 1: Create Table

```sql
CREATE TABLE `contact`
(
    `id`               bigint(20) unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
    `name`             varchar(255)        NOT NULL COMMENT 'name',
    `mobile`           varchar(255)        NOT NULL COMMENT 'mobile',
    `mobile_confirmed` tinyint(1)                   DEFAULT false COMMENT 'mobile_confirmed',
    `email`            varchar(255)                 DEFAULT '' COMMENT 'email',
    `email_confirmed`  tinyint(1)                   DEFAULT false COMMENT 'email_confirmed',
    `created_at`       timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created_at',
    `updated_at`       timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'updated_at'
) ENGINE = InnoDB
  COLLATE = utf8mb4_bin COMMENT = 'Contact Information';
```

Step 2: Define the model

```go
type Contact struct {
    ID              uint64    `gorm:"column:id;primary_key;AUTO_INCREMENT"`                 // id
    Name            string    `gorm:"column:name;NOT NULL"`                                 // name
    Mobile          string    `gorm:"column:mobile;NOT NULL"`                               // mobile
    MobileConfirmed int       `gorm:"column:mobile_confirmed;default:0"`                    // mobile_confirmed
    Email           string    `gorm:"column:email"`                                         // email
    EmailConfirmed  int       `gorm:"column:email_confirmed;default:0"`                     // email_confirmed
    CreatedAt       time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"` // created_at
    UpdatedAt       time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"` // updated_at
}

func (Contact) TableName() string {
    return "contact"
}
```

Step 3: Generate the code by gorm

```bash
make model && ./model
```

```
GO111MODULE=on go build -o model cmd/modelgen/modelgen.go
2022/05/25 15:53:00 Start generating code.
2022/05/25 15:53:00 generate query file: /Users/aland/gowork/src/github.com/alandtsang/gormgen/dal/query/contact.gen.go
2022/05/25 15:53:00 generate query file: /Users/aland/gowork/src/github.com/alandtsang/gormgen/dal/query/gen.go
2022/05/25 15:53:00 Generate code done.
```

Step 4: Run the demo code

```bash
make build && ./main
```

```
[1.448ms] [rows:1] SELECT * FROM `contact` WHERE `contact`.`id` IN (1,2,3)
contact: &{ID:1 Name:Tom Mobile:15911111111 MobileConfirmed:1 Email:tom@gmail.com EmailConfirmed:0 CreatedAt:2023-02-07 03:09:58 +0000 UTC UpdatedAt:2023-02-07 03:09:58 +0000 UTC}
```

## License

Please refer to [LICENSE](https://github.com/alandtsang/gormgen/blob/master/LICENSE) file.
