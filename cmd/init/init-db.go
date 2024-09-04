package init

import "github.com/jmoiron/sqlx"

type InitDb struct {
	Db *sqlx.DB
}

func NewInitDb(db *sqlx.DB) *InitDb {
	return &InitDb{
		Db: db,
	}
}

func (i *InitDb) Init() {

	var person = `CREATE TABLE IF NOT EXISTS user (
    		id INT AUTO_INCREMENT PRIMARY KEY,
    		email VARCHAR(50) NOT NULL,
    		phone VARCHAR(15) NOT NULL,
    		password VARCHAR(100) NOT NULL);`

	i.Db.MustExec(person)
}
