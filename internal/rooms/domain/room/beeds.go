package room

import "errors"

type bedType int

const (
	// SingleBed represents an individual bed
	SingleBed bedType = iota + 1
	// SofaBed reprensets a sofa that can be converted into a bed
	SofaBed
	// RollawayBed represents a bed that can be renomed or placed
	RollawayBed
	// DoubleBed reprensets a double bed for two persons
	DoubleBed
	// QueenBed represents a bigger size bed, usually for two persons and a child
	QueenBed
)

var ErrInvalidBedType = errors.New("invalid bed type")

func NewBedTypeFromString(value string) (bedType, error) {
	switch value {
	case "single-bed":
		return SingleBed, nil
	case "sofa-bed":
		return SofaBed, nil
	case "rollaway-bed":
		return RollawayBed, nil
	case "double-bed":
		return DoubleBed, nil
	case "queen-bed":
		return QueenBed, nil
	default:
		return 0, ErrInvalidBedType
	}
}

type RoomBed struct {
	bedType bedType
	count   int
}

func NewRoomBed(bedType bedType, count int) RoomBed {
	return RoomBed{
		bedType: bedType,
		count:   count,
	}
}

func parseRoomBeds(beds map[string]int) ([]RoomBed, error) {
	var roomBeds []RoomBed

	for key, count := range beds {
		bedType, err := NewBedTypeFromString(key)
		if err != nil {
			return nil, err
		}

		bed := NewRoomBed(bedType, count)
		roomBeds = append(roomBeds, bed)
	}

	return roomBeds, nil
}

func (b RoomBed) Capacity() int {
	if b.bedType == DoubleBed {
		return 2
	}

	if b.bedType == QueenBed {
		return 2
	}
	return 1
}

func (b RoomBed) String() string {
	switch b.bedType {
	case SingleBed:
		return "single-bed"
	case SofaBed:
		return "sofa-bed"
	case RollawayBed:
		return "rollaway-bed"
	case DoubleBed:
		return "double-bed"
	case QueenBed:
		return "queen-bed"
	default:
		return ""
	}
}
