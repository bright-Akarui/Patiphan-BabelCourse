package migrations

import (
	"go-gin/config"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
)

func Migrate() {
	m := gormigrate.New(
		config.GetDB(),
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			m1747737106CreateArticlesTable(),
			m1747841916CreateArticlesTable(),
			m1747844730CreateArticlesTable(),
			m1747910241CreateUserTable(),
		},
	)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
