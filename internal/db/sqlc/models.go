// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Bus struct {
	ID               int32       `json:"id"`
	RouteID          int32       `json:"route_id"`
	DepartureTime    time.Time   `json:"departure_time"`
	ArrivalTime      time.Time   `json:"arrival_time"`
	Capacity         int32       `json:"capacity"`
	Price            int32       `json:"price"`
	BusType          string      `json:"bus_type"`
	Corporation      pgtype.Text `json:"corporation"`
	SuperCorporation pgtype.Text `json:"super_corporation"`
	ServiceNumber    pgtype.Text `json:"service_number"`
	IsVip            pgtype.Bool `json:"is_vip"`
}

type BusSeat struct {
	ID         int32  `json:"id"`
	BusID      int32  `json:"bus_id"`
	SeatNumber int32  `json:"seat_number"`
	Status     string `json:"status"`
}

type City struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Penalty struct {
	ID                int32         `json:"id"`
	BusID             int32         `json:"bus_id"`
	ActualHoursBefore pgtype.Float8 `json:"actual_hours_before"`
	HoursBefore       pgtype.Float8 `json:"hours_before"`
	Percent           int32         `json:"percent"`
	CustomText        pgtype.Text   `json:"custom_text"`
}

type Route struct {
	ID                    int32           `json:"id"`
	OriginTerminalID      int32           `json:"origin_terminal_id"`
	DestinationTerminalID int32           `json:"destination_terminal_id"`
	Duration              pgtype.Interval `json:"duration"`
	Distance              int32           `json:"distance"`
}

type SeatReservation struct {
	ID          int32              `json:"id"`
	BusID       int32              `json:"bus_id"`
	BusSeatID   int32              `json:"bus_seat_id"`
	UserID      int32              `json:"user_id"`
	Status      string             `json:"status"`
	ReservedAt  pgtype.Timestamptz `json:"reserved_at"`
	PurchasedAt pgtype.Timestamptz `json:"purchased_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Terminal struct {
	ID     int32  `json:"id"`
	CityID int32  `json:"city_id"`
	Name   string `json:"name"`
}

type Ticket struct {
	ID                int32              `json:"id"`
	UserID            int32              `json:"user_id"`
	BusID             int32              `json:"bus_id"`
	SeatReservationID int32              `json:"seat_reservation_id"`
	Status            string             `json:"status"`
	ReservedAt        pgtype.Timestamptz `json:"reserved_at"`
	PurchasedAt       pgtype.Timestamptz `json:"purchased_at"`
}

type User struct {
	ID                int32     `json:"id"`
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
