// initialize the folder -- creates go.mod file
go mod init folder-name

// install gin library
go get github.com/gin-gonic/gin

// install gorm library
go get gorm.io/gorm

// gorm postgresql library
go get gorm.io/driver/postgres

// for logging
go get github.com/rs/zerolog/log

// env var
go get github.com/joho/godotenv

// run the main file
go run .\main.go

//golang migrations up and down
go get github.com/golang-migrate/migrate/v4/cmd/migrate

// migration files
migrate create -ext sql -dir migrations/ -format unix create_user 
// migrate up
migrate -path migrations/ -database "postgresql://postgres:postgres@localhost:5432/sample?sslmode=disable" up
