package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/update"
)

type UpdateOptions struct {

	// Core Identity & Basic Information
	ID uuid.UUID `json:"id"` // Unique identifier for the user

	Username     *string `json:"username"`    // User's login username (must be unique)
	DisplayName  *string `json:"displayName"` // Name displayed publicly (can differ from Username)
	PhoneNumber  *string `json:"phoneNumber"` // User's phone number
	Email        *string `json:"email"`       // User's email address
	FirstName    *string `json:"firstName"`   // User's first name
	LastName     *string `json:"lastName"`    // User's last name
	Bio          *string `json:"bio"`         // Short biography or "About Me" section
	OriginalURL  *string `json:"originalUrl"`
	ThumbnailURL *string `json:"thumbnailUrl"`
	IsVerified   *bool   `json:"isVerified"` // Indicates if the application is verified (e.g., for official accounts)

	// Presence & Connectivity (for chat, social apps)
	IsOnline      *bool      `json:"isOnline"`      // Current online status of the user
	LastSeen      *time.Time `json:"lastSeen"`      // Last timestamp the user was seen online
	StatusMessage *string    `json:"statusMessage"` // User's custom status message (e.g., "Busy", "Available")

	// Privacy & Social Features (for social, photo, music apps)
	ProfileVisibility *string `json:"profileVisibility"` // Profile visibility setting: "public", "private", "friendsOnly"
	FollowerCount     *int    `json:"followerCount"`     // Number of followers this user has
	FollowingCount    *int    `json:"followingCount"`    // Number of Users this user is following

	BlockedUsers       *[]string `json:"blockedUsers,omitempty"`
	AddBlockedUsers    []string  `json:"addBlockedUsers,omitempty"`
	RemoveBlockedUsers []string  `json:"removeBlockedUsers,omitempty"`

	// Preferences & Customization
	PreferredLanguage *string  `json:"preferredLanguage"` // User's preferred language (e.g., "en-US", "fa-IR")
	Timezone          *string  `json:"timezone"`          // User's timezone (e.g., "Asia/Tehran")
	ThemePreference   *string  `json:"themePreference"`   // User's UI theme preference: "light", "dark", "system"
	Interests         []string `json:"interests"`         // List of user's interests (for content recommendations)

	// Account Status & Usage
	SubscriptionTier   *string `json:"subscriptionTier"`   // User's subscription level (e.g., "free", "premium", "pro")
	AccountStatus      *string `json:"accountStatus"`      // Current status of the user's application: "active", "suspended", "deactivated"
	IsTwoFactorEnabled *bool   `json:"isTwoFactorEnabled"` // Indicates if two-factor authentication is enabled

	// Timestamps & Tracking
	LastActivity *time.Time `json:"lastActivity"` // Timestamp of the user's last public activity in the app

	// Generic/Extensible Metadata (for highly specific or future data)
	Metadata map[string]string `json:"metadata"` // Flexible field for storing additional, application-specific data
}

// Initialize updater
var chatUpdater = update.NewUpdater[User, UpdateOptions]()

func init() {

	// Basic Info Updates
	chatUpdater.AddScalarUpdater(func(c *User, u UpdateOptions) {
		if u.Username != nil {
			c.Username = *u.Username
		}
		if u.PhoneNumber != nil {
			c.PhoneNumber = *u.PhoneNumber
		}
		if u.Username != nil {
			c.Username = *u.Username
		}
		if u.Email != nil {
			c.Email = *u.Email
		}
		if u.FirstName != nil {
			c.FirstName = *u.FirstName
		}
		if u.LastName != nil {
			c.LastName = *u.LastName
		}
		if u.Bio != nil {
			c.Bio = *u.Bio
		}
		if u.ThumbnailURL != nil {
			c.ThumbnailURL = *u.ThumbnailURL
		}
		if u.IsVerified != nil {
			c.IsVerified = *u.IsVerified
		}
	})

	// Presence & Connectivity (for chat, social apps)
	chatUpdater.AddScalarUpdater(func(c *User, u UpdateOptions) {
		if u.IsOnline != nil {
			c.IsOnline = *u.IsOnline
		}
		if u.LastSeen != nil {
			c.LastSeen = *u.LastSeen
		}
		if u.StatusMessage != nil {
			c.StatusMessage = *u.StatusMessage
		}
	})

	/// Privacy & Social Features (for social, photo, music apps)
	chatUpdater.AddScalarUpdater(func(c *User, u UpdateOptions) {
		if u.ProfileVisibility != nil {
			c.ProfileVisibility = *u.ProfileVisibility
		}
		if u.IsVerified != nil {
			c.IsVerified = *u.IsVerified
		}
		if u.FollowerCount != nil {
			c.FollowerCount = *u.FollowerCount
		}
		if u.FollowingCount != nil {
			c.FollowingCount = *u.FollowingCount
		}
	})

	// Banned Users Collection Updates
	chatUpdater.AddCollectionUpdater(func(c *User, u UpdateOptions) {
		op := update.CollectionUpdateOp[string]{
			FullReplace: u.BlockedUsers,
			Add:         u.AddBlockedUsers,
			Remove:      u.RemoveBlockedUsers,
		}
		c.BlockedUserIDs = update.ApplyCollectionUpdate(c.BlockedUserIDs, op)
	})

}

func Update(user *User, update UpdateOptions) *User {
	chatUpdater.Apply(user, update)
	return user
}
