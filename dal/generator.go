package dal

import (
	"github.com/alandtsang/gormgen/dal/model"
	"gorm.io/gen"
)

const mysqlDSN = "root:11223344@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True"

func init() {
	DB = ConnectDB(mysqlDSN).Debug()
}

func GenerateModel() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		Mode:    gen.WithDefaultQuery,
	})

	// generate from struct in project
	g.ApplyBasic(model.Contacts{})

	g.Execute()
}
