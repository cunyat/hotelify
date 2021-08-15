module github.com/cunyat/hotelify/internal/rooms

go 1.16

replace github.com/cunyat/hotelify/internal/common => ../common

require (
	github.com/cunyat/hotelify/internal/common v0.0.0-20210716152902-8c5d04fc12a4
	github.com/gin-gonic/gin v1.7.2
	github.com/google/uuid v1.3.0
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jackc/pgx/v4 v4.13.0
	github.com/stretchr/testify v1.7.0
)
