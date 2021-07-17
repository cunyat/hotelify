package create

import (
	"context"
	"testing"

	"github.com/cunyat/hotelify/internal/rooms/adapters/storagemocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateRoomCommand_Saved(t *testing.T) {
	repo := new(storagemocks.Repository)

	repo.On("Save", mock.Anything, mock.Anything).Once().Return(nil)

	cmd := RoomCommand{
		Num:      "123",
		Floor:    1,
		Beds:     map[string]int{"double-bed": 2},
		Services: []string{"tv"},
	}

	handler := RoomCommandHandler(repo)
	err := handler(context.Background(), cmd)

	require.NoError(t, err)
	repo.AssertExpectations(t)
}
