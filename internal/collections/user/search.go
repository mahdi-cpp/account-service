package user

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mahdi-cpp/iris-tools/search"
)

type SearchOptions struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Bio         string    `json:"bio"`
	AvatarURL   string    `json:"avatarURL"`
	IsOnline    *bool     `json:"isOnline"`

	UsernameQuery string `json:"usernameQuery"`

	// Date filters
	LastSeen      *time.Time `json:"lastSeen"`
	CreatedAfter  *time.Time `form:"createdAfter,omitempty"`
	CreatedBefore *time.Time `form:"createdBefore,omitempty"`
	ActiveAfter   *time.Time `form:"activeAfter,omitempty"`
	SearchA
}

type SearchA struct {
	// Sorting
	Sort      string `json:"sort,omitempty"`      // "title", "created", "members", "lastActivity"
	SortOrder string `json:"sortOrder,omitempty"` // "asc" or "desc"

	// Pagination
	Page int `json:"page,omitempty"`
	Size int `json:"size,omitempty"`
}

var PHAssetLessFunks = map[string]search.LessFunction[*User]{
	"id":               func(a, b *User) bool { return a.ID.String() < b.ID.String() },
	"creationDate":     func(a, b *User) bool { return a.CreatedAt.Before(b.CreatedAt) },
	"modificationDate": func(a, b *User) bool { return a.UpdatedAt.Before(b.UpdatedAt) },
	"title":            func(a, b *User) bool { return a.Username < b.Username },
}

func GetLessFunc(sortBy, sortOrder string) search.LessFunction[*User] {

	fn, exists := PHAssetLessFunks[sortBy]
	if !exists {
		return nil
	}

	if sortOrder == "end" {
		return func(a, b *User) bool { return !fn(a, b) }
	}
	return fn
}

func BuildUserSearchCriteria(with *SearchOptions) search.Criteria[*User] {

	return func(c *User) bool {

		// ID filter
		if with.ID != uuid.Nil && c.ID != with.ID {
			return false
		}

		// Title search_manager (case-insensitive)
		if with.UsernameQuery != "" {
			query := strings.ToLower(with.UsernameQuery)
			username := strings.ToLower(c.Username)
			if !strings.Contains(username, query) {
				return false
			}
		}

		// Username exact match
		if with.Username != "" && c.Username != with.Username {
			return false
		}

		// Boolean flags
		if with.IsOnline != nil && c.IsOnline != *with.IsOnline {
			return false
		}

		// Date filters
		if with.CreatedAfter != nil && c.CreatedAt.Before(*with.CreatedAfter) {
			return false
		}
		if with.CreatedBefore != nil && c.CreatedAt.After(*with.CreatedBefore) {
			return false
		}
		//if with.ActiveAfter != nil && c.LastMessage != nil &&
		//	c.LastMessage.CreationDate.Before(*with.ActiveAfter) {
		//	return false
		//}

		return true
	}
}

func Search(chats []*User, with *SearchOptions) []*User {

	// Build criteria
	criteria := BuildUserSearchCriteria(with)

	// Execute search_manager
	results := search.Find(chats, criteria)

	// Sort results if needed
	if with.Sort != "" {
		lessFn := GetLessFunc(with.Sort, with.SortOrder)
		if lessFn != nil {
			search.SortIndexedItems(results, lessFn)
		}
	}

	// Extract final assets
	final := make([]*User, len(results))
	for i, item := range results {
		final[i] = item.Value
	}

	// Apply pagination
	start := with.Page
	// Check if the start index is out of bounds. If so, return an empty slice.
	if start >= len(final) {
		return []*User{}
	}

	end := start + with.Size
	if end > len(final) {
		end = len(final)
	}
	return final[start:end]
}
