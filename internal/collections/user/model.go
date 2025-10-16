package user

import (
	"time"

	"github.com/google/uuid"
)

func (u *User) GetID() uuid.UUID         { return u.ID }
func (u *User) SetID(id uuid.UUID)       { u.ID = id }
func (u *User) SetCreatedAt(t time.Time) { u.CreatedAt = t }
func (u *User) SetUpdatedAt(t time.Time) { u.UpdatedAt = t }
func (u *User) GetRecordSize() int       { return 4000 }

type User struct {
	// Core Identity & Basic Information
	ID           uuid.UUID `json:"id"`          // unique: true
	Username     string    `json:"username"`    // unique: true
	Email        string    `json:"email"`       // unique: true
	PhoneNumber  string    `json:"phoneNumber"` // unique: true
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	DisplayName  string    `json:"displayName"`
	Bio          string    `json:"bio"`
	OriginalURL  string    `json:"originalUrl"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	IsVerified   bool      `json:"isVerified"`

	// Presence & Connectivity
	IsOnline      bool      `json:"isOnline"` // Note: omitempty for bool is complex, depends on default value
	LastSeen      time.Time `json:"lastSeen"`
	StatusMessage string    `json:"statusMessage"`

	// Privacy & Social Features
	ProfileVisibility string   `json:"profileVisibility"`
	FollowerCount     int      `json:"followerCount"`
	FollowingCount    int      `json:"followingCount"`
	BlockedUserIDs    []string `json:"blockedUserIds"` // changed json tag to follow camelCase

	// Preferences & Customization
	PreferredLanguage string   `json:"preferredLanguage"`
	Timezone          string   `json:"timezone"`
	ThemePreference   string   `json:"themePreference"`
	Interests         []string `json:"interests"`

	// Account Status & Usage
	SubscriptionTier   string `json:"subscriptionTier"`
	AccountStatus      string `json:"accountStatus"`
	IsTwoFactorEnabled bool   `json:"isTwoFactorEnabled"`

	// Timestamps & Tracking
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Version   string    `json:"version"`
}
