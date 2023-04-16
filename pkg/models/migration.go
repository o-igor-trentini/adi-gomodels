package models

import "time"

// Migration tabela de registro de migrações.
type Migration struct {
	ID        uint64    `gorm:"column:mi_id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:mi_name;not null;type:varchar(255);comment:Nome da função"`
	CreatedAt time.Time `gorm:"column:mi_created_at;not null;comment:Data de criação do registro"`
}

func (Migration) TableName() string {
	return "mi_migrations"
}
