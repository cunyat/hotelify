package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cunyat/hotelify/internal/common/adapters/query"
	"github.com/cunyat/hotelify/internal/rooms/adapters/storage"
	"github.com/cunyat/hotelify/internal/rooms/app/get"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetRoom(t *testing.T) {
	repo := storage.NewInMemoryRoomRepository()
	qBus := query.NewInMemoryQueryBus()
	qBus.Register(get.RoomQueryType, get.RoomQueryHandler(repo))

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/rooms/:uuid", GetRoomHandler(qBus))

	entity, _ := room.CreateRoom(
		uuid.NewString(),
		"32",
		1,
		map[string]int{"double-bed": 2},
		[]string{"tv", "drinks"},
	)
	_ = repo.Save(context.Background(), entity)

	t.Run("it return 200 and the room object", func(t *testing.T) {
		t.Parallel()

		url := fmt.Sprintf("/rooms/%s", entity.UUID())
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response get.RoomQueryResponse
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(&response)
		require.NoError(t, err)
	})

	t.Run("it return 404 when room not found", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequest(http.MethodGet, "/rooms/random-id", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}
