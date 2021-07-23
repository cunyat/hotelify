package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cunyat/hotelify/internal/common/adapters/querymocks"
	"github.com/cunyat/hotelify/internal/common/domain/query"
	"github.com/cunyat/hotelify/internal/rooms/app/get"
	"github.com/cunyat/hotelify/internal/rooms/app/response"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetRoom(t *testing.T) {
	entity, _ := room.CreateRoom(
		uuid.NewString(),
		"32",
		1,
		map[string]int{"double-bed": 2},
		[]string{"tv", "drinks"},
	)

	qBus := new(querymocks.Bus)
	qBus.
		On(
			"Ask",
			mock.Anything,
			mock.MatchedBy(func(q query.Query) bool {
				getRoom, _ := q.(get.RoomQuery)
				return getRoom.UUID == entity.UUID()
			}),
		).Return(response.FromDomain(entity), nil)

	qBus.
		On("Ask", mock.Anything, mock.AnythingOfType("get.RoomQuery")).
		Return(nil, room.ErrRoomNotFound)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/rooms/:uuid", GetRoomHandler(qBus))

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

		var resp response.Room
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(&resp)
		require.NoError(t, err)
		assert.Equal(t, entity.UUID(), resp.UUID)
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
