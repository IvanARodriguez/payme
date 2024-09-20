package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

// ULID type that wraps ulid.ULID
type ULID ulid.ULID

// Implementing the driver.Valuer interface for the ULID type
func (u ULID) Value() (driver.Value, error) {
	// Convert the ULID to its string representation
	return ulid.ULID(u).String(), nil
}

// Implementing the sql.Scanner interface for the ULID type
func (u *ULID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	// Convert the value to a string, then parse the ULID
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot scan type %T into ULID", value)
	}
	parsed, err := ulid.Parse(s)
	if err != nil {
		return err
	}
	*u = ULID(parsed)
	return nil
}

// Implementing the json.Marshaler interface for the ULID type
func (u ULID) MarshalJSON() ([]byte, error) {
	// Convert the ULID to its string representation and then to JSON
	return json.Marshal(ulid.ULID(u).String())
}

// Implementing the json.Unmarshaler interface for the ULID type
func (u *ULID) UnmarshalJSON(data []byte) error {
	// Convert the JSON data to a string
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	// Parse the string as a ULID
	parsed, err := ulid.Parse(s)
	if err != nil {
		return err
	}
	*u = ULID(parsed)
	return nil
}

type Business struct {
	ID        ULID      `json:"id" gorm:"type:char(26);primaryKey;"` // ULID as the primary key
	Name      string    `json:"name" gorm:"required;type:char(70)"`
	Email     string    `json:"email" gorm:"type:char(100);"`
	StripeId  string    `json:"stripeId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (b *Business) BeforeCreate(tx *gorm.DB) (err error) {
	newID := ulid.MustNew(ulid.Timestamp(time.Now()), entropy) // Use MustNew for safety
	b.ID = ULID(newID)                                         // Assign the generated ULID
	return
}

var entropy = ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
