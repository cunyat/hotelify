module github.com/cunyat/hotelify/internal/rooms

go 1.16

replace github.com/cunyat/hotelify/internal/common => ../common

require (
	github.com/cunyat/hotelify/internal/common v0.0.0-20210716152902-8c5d04fc12a4
	github.com/gin-gonic/gin v1.7.2
	github.com/google/uuid v1.3.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/stretchr/testify v1.7.0
)
