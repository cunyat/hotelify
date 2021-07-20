package create

import (
	"context"
	"errors"
	"testing"

	"github.com/cunyat/hotelify/internal/rooms/adapters/storagemocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateRoomCommand_Saved(t *testing.T) {
	repo := new(storagemocks.Repository)

	repo.On("Save", mock.Anything, mock.Anything).Once().Return(nil)

	cmd := RoomCommand{
		UUID:     uuid.NewString(),
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

func TestCreateRoomCommand_RepositoryError(t *testing.T) {
	repo := new(storagemocks.Repository)

	repo.On("Save", mock.Anything, mock.Anything).Once().Return(errors.New("error in repository"))

	cmd := RoomCommand{
		UUID:     uuid.NewString(),
		Num:      "123",
		Floor:    1,
		Beds:     map[string]int{"double-bed": 2},
		Services: []string{"tv"},
	}

	handler := RoomCommandHandler(repo)
	err := handler(context.Background(), cmd)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "repository")
	assert.Contains(t, err.Error(), "saving")
	repo.AssertExpectations(t)
}

func TestCreateRoomCommand_BadBedType(t *testing.T) {
	repo := new(storagemocks.Repository)

	cmd := RoomCommand{
		UUID:     uuid.NewString(),
		Num:      "123",
		Floor:    1,
		Beds:     map[string]int{"bad-bed": 2},
		Services: []string{"tv"},
	}

	handler := RoomCommandHandler(repo)
	err := handler(context.Background(), cmd)

	require.Error(t, err)
	repo.AssertExpectations(t)
}
