package migrate

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	migrate "github.com/rubenv/sql-migrate"
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

	migrations := &migrate.FileMigrationSource{
		Dir: filePath,
	}

	n, err := migrate.Exec(sql.DB, driverName, migrations, migrate.Up)
	if err != nil {
		log.Fatal().Err(err).Msg("Migrations failed !!!")
	}

	log.Debug().Msg(fmt.Sprintf("Applied %d migrations !!!", n))

	// switch driverName {
	// case preference.POSTGRES:
	// 	driver, err := postgres.WithInstance(sql.DB, &postgres.Config{})
	// 	if err != nil {
	// 		log.Panic().Err(err).Str("migrate_"+sql.DriverName(), "FAILED").Send()
	// 	}

	// 	m, err := migrate.NewWithDatabaseInstance("file://"+filePath, driverName, driver)
	// 	if err != nil {
	// 		log.Panic().Err(err).Msg("migrate failed !!!")
	// 	}

	// 	if err := m.Up(); err != nil {
	// 		log.Panic().Err(err).Msg("steps migrate failed !!!")
	// 	}

	// case preference.MYSQL:
	// 	driver, err := mysql.WithInstance(sql.DB, &mysql.Config{})
	// 	if err != nil {
	// 		log.Panic().Err(err).Str("migrate_"+sql.DriverName(), "FAILED").Send()
	// 	}

	// 	m, err := migrate.NewWithDatabaseInstance("file://"+filePath, driverName, driver)
	// 	if err != nil {
	// 		log.Panic().Err(err).Msg("migrate failed !!!")
	// 	}

	// 	if err := m.Up(); err != nil {
	// 		log.Panic().Err(err).Msg("steps migrate failed !!!")
	// 	}

	// default:
	// 	log.Panic().Msg("Migration failed, unknown driver name. Please check you configuration")
	// }

	// log.Debug().Str("migrate_status", "OK").Send()
}
