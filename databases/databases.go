package databases

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/tora0091/webapps/environments"
)

type Databases struct {
	Env *environments.Env
}

func NewDatabases(env *environments.Env) *Databases {
	return &Databases{
		Env: env,
	}
}

func (d *Databases) GetConnection() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(d.makeDns()), &gorm.Config{})
}

func (d *Databases) makeDns() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		d.Env.DbUser(), d.Env.DbPass(), d.Env.DbHost(),
		d.Env.DbPort(), d.Env.DbSchema(), d.Env.DbCharset())
}
