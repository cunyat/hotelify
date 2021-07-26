package storage

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cunyat/hotelify/internal/rooms/domain/room"
	"github.com/jmoiron/sqlx"
)

type RoomRepository struct {
	db *sqlx.DB
}

var _ room.Repository = (*RoomRepository)(nil)

type sqlRoom struct {
	UUID     string `db:"uuid"`
	Num      string `db:"num"`
	Floor    int    `db:"floor"`
	Services string `db:"services"`
}

type sqlBed struct {
	ID       int    `db:"id"`
	RoomUUID string `db:"room_uuid"`
	BedType  string `db:"bed_type"`
	Count    int    `db:"count"`
}

func NewMysqlRoomRepository(db *sqlx.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}
func (r *RoomRepository) Save(ctx context.Context, entity room.Room) error {
	rm, beds := domainToSQL(entity)

	_, err := r.db.ExecContext(ctx, "insert into rooms (uuid, num, floor, services) values (?, ?, ?, ?)", rm.UUID, rm.Num, rm.Floor, rm.Services)
	if err != nil {
		return fmt.Errorf("could not insert room: %w", err)
	}

	err = r.insertBeds(ctx, beds)
	if err != nil {
		return err
	}

	return nil
}

func (r *RoomRepository) Get(ctx context.Context, uuid string) (room.Room, error) {
	beds := []sqlBed{}
	rm := sqlRoom{}

	err := r.db.Select(&beds, "select * from beds where room_uuid = ?", uuid)
	if err != nil {
		return room.Room{}, err
	}

	err = r.db.Get(&rm, "select * from rooms where uuid = ? LIMIT 1", uuid)
	if err != nil {
		return room.Room{}, err
	}

	return sqlToDomain(rm, beds)
}

func (r *RoomRepository) List(ctx context.Context) ([]room.Room, error) {
	rooms := []sqlRoom{}

	err := r.db.Select(&rooms, "select * from rooms")
	if err != nil {
		return nil, err
	}

	entities := make([]room.Room, len(rooms))
	for i, rm := range rooms {
		beds := []sqlBed{}
		err = r.db.Select(&beds, "select * from beds where room_uuid = ?", rm.UUID)
		if err != nil {
			return nil, err
		}

		entities[i], err = sqlToDomain(rm, beds)
		if err != nil {
			return nil, fmt.Errorf("error parsing entity: %w", err)
		}
	}

	return entities, nil
}

func (r *RoomRepository) Update(ctx context.Context, entity room.Room) error {
	rm, beds := domainToSQL(entity)

	_, err := r.db.ExecContext(ctx, "delete from beds where room_uuid = ?", rm.UUID)
	if err != nil {
		return fmt.Errorf("error removing previous beds %w", err)
	}

	err = r.insertBeds(ctx, beds)
	if err != nil {
		return fmt.Errorf("error updating beds %w", err)
	}

	_, err = r.db.ExecContext(ctx, "update rooms set num = ?, floor = ?, services = ? where uuid = ?", rm.Num, rm.Floor, rm.Services, rm.UUID)
	if err != nil {
		return fmt.Errorf("error updating room: %w", err)
	}

	return nil
}

func (r *RoomRepository) insertBeds(ctx context.Context, beds []sqlBed) error {
	for _, bed := range beds {
		_, err := r.db.ExecContext(ctx, "insert into beds (room_uuid, bed_type, count) values (?, ?, ?)", bed.RoomUUID, bed.BedType, bed.Count)
		if err != nil {
			return fmt.Errorf("could not insert bed: %w", err)
		}
	}

	return nil
}

func sqlToDomain(r sqlRoom, bb []sqlBed) (room.Room, error) {
	beds := make([]room.RoomBed, len(bb))
	for i, bed := range bb {
		btype, err := room.NewBedTypeFromString(bed.BedType)
		if err != nil {
			return room.Room{}, err
		}
		beds[i] = room.NewRoomBed(btype, bed.Count)
	}

	services := strings.Split(r.Services, ";")
	return room.NewRoom(r.UUID, r.Num, r.Floor, beds, services)
}

func domainToSQL(r room.Room) (sqlRoom, []sqlBed) {
	beds := make([]sqlBed, len(r.Beds()))
	for i, bed := range r.Beds() {
		beds[i] = sqlBed{RoomUUID: r.UUID(), BedType: bed.String(), Count: bed.Count()}
	}

	services := strings.Join(r.Services(), ";")
	return sqlRoom{
		UUID:     r.UUID(),
		Num:      r.Num(),
		Floor:    r.Floor(),
		Services: services,
	}, beds
}

// InMemoryRoomRepository is for test mocking, so we can simulate real behavior
type InMemoryRoomRepository struct {
	rooms map[string]room.Room
}

var _ room.Repository = (*InMemoryRoomRepository)(nil)

func NewInMemoryRoomRepository() *InMemoryRoomRepository {
	return &InMemoryRoomRepository{
		rooms: make(map[string]room.Room),
	}
}

func (r *InMemoryRoomRepository) Save(ctx context.Context, entity room.Room) error {
	_, ok := r.rooms[entity.UUID()]
	if ok {
		return errors.New("duplicated room uuid")
	}

	r.rooms[entity.UUID()] = entity

	fmt.Printf("New room: %v", entity)
	return nil
}

func (r *InMemoryRoomRepository) Get(ctx context.Context, uuid string) (room.Room, error) {
	entity, ok := r.rooms[uuid]
	if !ok {
		return room.Room{}, room.ErrRoomNotFound
	}

	return entity, nil
}

func (r *InMemoryRoomRepository) List(ctx context.Context) ([]room.Room, error) {
	var rooms []room.Room
	for _, rm := range r.rooms {
		rooms = append(rooms, rm)
	}

	return rooms, nil
}

func (r *InMemoryRoomRepository) Update(ctx context.Context, entity room.Room) error {
	r.rooms[entity.UUID()] = entity
	return nil
}
