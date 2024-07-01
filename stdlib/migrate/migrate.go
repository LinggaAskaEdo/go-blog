package migrate

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"

	preference "github.com/linggaaskaedo/go-blog/stdlib/preference"
)

type Options struct {
	Path            string
	MySQLEnabled    bool
	PostgresEnabled bool
}

func Init(log zerolog.Logger, opt Options, sql *sqlx.DB) {
	if sql == nil {
		return
	}

	filePath := opt.Path
	driverName := sql.DriverName()

	log.Debug().Str("db_path", filePath).Send()
	log.Debug().Str("db_type", driverName).Send()

	switch driverName {
	case preference.POSTGRES:
		driver, err := postgres.WithInstance(sql.DB, &postgres.Config{})
		if err != nil {
			log.Panic().Err(err).Str("migrate_"+sql.DriverName(), "FAILED").Send()
		}

		m, err := migrate.NewWithDatabaseInstance("file://"+filePath, driverName, driver)
		if err != nil {
			log.Panic().Err(err).Msg("migrate failed !!!")
		}

		m.Up()

	case preference.MYSQL:
		driver, err := mysql.WithInstance(sql.DB, &mysql.Config{})
		if err != nil {
			log.Panic().Err(err).Str("migrate_"+sql.DriverName(), "FAILED").Send()
		}

		m, err := migrate.NewWithDatabaseInstance("file://"+filePath, driverName, driver)
		if err != nil {
			log.Panic().Err(err).Msg("migrate failed !!!")
		}

		m.Up()

	default:
		log.Panic().Msg("Migration failed, unknown driver name. Please check you configuration")
	}

	log.Debug().Str("migrate_status", "OK").Send()
}
