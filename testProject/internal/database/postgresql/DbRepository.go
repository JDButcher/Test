package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	operation "testProject/pkg"
)

type (
	Config struct {
		PostgreSQLDBServerURL string `envconfig:"PostgreSQLDB_URL"`
		PostgreSQLDBPort      string `envconfig:"PostgreSQLDB_PORT"`
		PostgreSQLDBName      string `envconfig:"PostgreSQLDB_NAME"`
		PostgreSQLDBUser      string `envconfig:"PostgreSQLDB_USER"`
		PostgreSQLDBPassword  string `envconfig:"PostgreSQLDB_PASSWORD"`
	}
)

func NewDBRepository(postgreSqlConfig *Config) *sql.DB {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", postgreSqlConfig.PostgreSQLDBServerURL, postgreSqlConfig.PostgreSQLDBPort, postgreSqlConfig.PostgreSQLDBUser, postgreSqlConfig.PostgreSQLDBPassword, postgreSqlConfig.PostgreSQLDBName)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	return db
}

func PostHandler(ginCtx gin.Context, operation operation.Computer) {

}
func ReadHandler(ginCtx gin.Context, operation operation.Computer) {

}
func DeleteHandler(ginCtx gin.Context, operation operation.Computer) {

}
