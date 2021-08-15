package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/cunyat/hotelify/internal/rooms/domain/room"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RoomRepository struct {
	db *pgxpool.Pool
}

var _ room.Repository = (*RoomRepository)(nil)

type sqlRoom struct {
	UUID     string `db:"uuid"`
	Num      string `db:"num"`
	Floor    int    `db:"floor"`
	Capacity int    `db:"capacity"`
	Services string `db:"services"`
}

func NewMysqlRoomRepository(db *pgxpool.Pool) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) Save(ctx context.Context, entity room.Room) error {
	rm := domainToSQL(entity)
	q := `insert into rooms (uuid, num, floor, capacity, services) values (?, ?, ?, ?, ?)`
	_, err := r.db.Exec(ctx, q, rm.UUID, rm.Num, rm.Floor, rm.Capacity, rm.Services)
	if err != nil {
		return fmt.Errorf("could not insert room: %w", err)
	}

	return nil
}

func (r *RoomRepository) Get(ctx context.Context, uuid string) (room.Room, error) {
	row := r.db.QueryRow(ctx, "select * from rooms where uuid = ? LIMIT 1", uuid)
	rm, err := scanRoom(row)

	if err != nil {
		return room.Room{}, err
	}

	return sqlToDomain(rm)
}

func (r *RoomRepository) List(ctx context.Context) ([]room.Room, error) {
	rooms := []room.Room{}
	roomRows, err := r.db.Query(ctx, "select uuid, num, floor, capacity, services from rooms")
	if err != nil {
		return nil, err
	}

	defer roomRows.Close()
	for roomRows.Next() {
		rm, err := scanRoom(roomRows)
		if err != nil {
			return nil, err
		}

		domain, err := sqlToDomain(rm)
		if err != nil {
			return nil, fmt.Errorf("error parsing entity: %w", err)
		}

		rooms = append(rooms, domain)
	}

	return rooms, nil
}

func (r *RoomRepository) Update(ctx context.Context, entity room.Room) error {
	rm := domainToSQL(entity)
	q := `update rooms set num = ?, floor = ?, services = ?, capacity = ? where uuid = ?`
	_, err := r.db.Exec(ctx, q, rm.Num, rm.Floor, rm.Services, rm.Capacity, rm.UUID)
	if err != nil {
		return fmt.Errorf("error updating room: %w", err)
	}

	return nil
}

func (r *RoomRepository) Delete(ctx context.Context, entity room.Room) error {
	q := `delete from rooms where uuid = ?`
	_, err := r.db.Exec(ctx, q, entity.UUID())
	if err != nil {
		return fmt.Errorf("error deleting room: %w", err)
	}

	return nil
}

func scanRoom(row pgx.Row) (sqlRoom, error) {
	var rm sqlRoom
	err := row.Scan(&rm.UUID, &rm.Num, &rm.Floor, &rm.Capacity, &rm.Services)
	if err != nil {
		return sqlRoom{}, err
	}

	return rm, nil
}

func sqlToDomain(r sqlRoom) (room.Room, error) {
	services := strings.Split(r.Services, ";")
	return room.NewRoom(r.UUID, r.Num, r.Floor, r.Capacity, services)
}

func domainToSQL(r room.Room) sqlRoom {
	services := strings.Join(r.Services(), ";")
	return sqlRoom{
		UUID:     r.UUID(),
		Num:      r.Num(),
		Floor:    r.Floor(),
		Capacity: r.Capacity(),
		Services: services,
	}
}
