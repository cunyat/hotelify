package room

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBedCapacity(t *testing.T) {
	tests := []struct {
		bed  bedType
		want int
	}{
		{
			bed:  SingleBed,
			want: 1,
		},
		{
			bed:  SofaBed,
			want: 1,
		},
		{
			bed:  RollawayBed,
			want: 1,
		},
		{
			bed:  DoubleBed,
			want: 2,
		},
		{
			bed:  QueenBed,
			want: 2,
		},
	}

	for _, tt := range tests {
		roomBed := RoomBed{bedType: tt.bed}
		t.Run(roomBed.String(), func(t *testing.T) {
			capacity := roomBed.Capacity()
			assert.Equal(t, tt.want, capacity)
		})
	}
}
