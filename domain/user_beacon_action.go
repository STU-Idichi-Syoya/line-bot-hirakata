package domain

import (
	"time"
)

type UserBeaconAction struct {
	UserID      UserID
	ActionPlace Place
	CreatedAt   time.Time
}
type Place string

// もし現在時刻を追加するならtimestampをnilにする
func NewBeaconAction(userId UserID, actionPlace Place, timestamp_utc time.Time) (*UserBeaconAction, error) {
	return &UserBeaconAction{
		UserID:      userId,
		ActionPlace: actionPlace,
		CreatedAt:   time.Now(),
	}, nil
}
