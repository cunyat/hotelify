module github.com/cunyat/hotelify/internal/rooms

go 1.16

replace github.com/cunyat/hotelify/internal/common => ../common

require (
	github.com/cunyat/hotelify/internal/common v0.0.0-20210716152902-8c5d04fc12a4
	github.com/gin-gonic/gin v1.7.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.3.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/mattn/go-sqlite3 v1.14.8
	github.com/stretchr/testify v1.7.0
)
