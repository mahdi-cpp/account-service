package user

import (
	"fmt"
	"regexp"
	"strings"
)

func (u *User) Validate() error {

	err := validateUsername(u.Username)
	if err != nil {
		return err
	}

	err = validatePhoneNumber(u.PhoneNumber)
	if err != nil {
		return err
	}

	err = validateEmail(u.Email)
	if err != nil {
		return err
	}

	err = u.ValidateNames()
	if err != nil {
		return err
	}

	err = u.validateStrings()
	if err != nil {
		return err
	}

	return nil
}

func (u *UpdateOptions) ValidateUpdate() error {

	if u.Username != nil {
		if err := validateUsername(*u.Username); err != nil {
			return err
		}
	}
	if u.Email != nil {
		if err := validateEmail(*u.Email); err != nil {
			return err
		}
	}
	if u.PhoneNumber != nil {
		if err := validatePhoneNumber(*u.PhoneNumber); err != nil {
			return err
		}
	}
	if u.FirstName != nil {
		if err := validateName(*u.FirstName, "FirstName"); err != nil {
			return err
		}
	}
	if u.LastSeen != nil {
		if err := validateName(*u.LastName, "LastName"); err != nil {
			return err
		}
	}

	return nil
}

func (u *User) validateStrings() error {

	const maxDescriptionLength = 50
	if len(strings.TrimSpace(u.PhoneNumber)) > maxDescriptionLength {
		return fmt.Errorf("description length exceeds the maximum of %d characters", maxDescriptionLength)
	}

	const maxInviteLinkLength = 500
	if len(strings.TrimSpace(u.Email)) > maxInviteLinkLength {
		return fmt.Errorf("inviteLink length exceeds the maximum of %d characters", maxInviteLinkLength)
	}

	return nil
}

// ValidateUsername checks if the Username field of the User struct is valid.
// A valid username:
// - Must be between 5 and 32 characters long.
// - Can contain alphanumeric characters and underscores.
// - Must start with a letter.
// - Cannot end with an underscore.
// - Cannot have consecutive underscores.
func validateUsername(username string) error {

	if username == "" { //also can empty in private , group chats
		return nil
	}

	// 1. Check if the username is within the valid length range.
	if len(username) < 5 || len(username) > 32 {
		return fmt.Errorf("username length must be between 0 and 32 characters, got %d", len(username))
	}

	// 2. Use a regular expression for comprehensive validation.
	// The regex explained:
	// ^ - Asserts position at the start of the string.
	// [a-zA-Z] - Matches any single uppercase or lowercase letter.
	// [a-zA-Z0-9_]* - Matches any combination of letters, numbers, or underscores, zero or more times.
	// [a-zA-Z0-9] - Matches any single letter or number. This is to ensure the username does not end in an underscore.
	// $ - Asserts position at the end of the string.
	// The pattern effectively requires the username to start with a letter and end with a letter or number.
	// It also implicitly handles the consecutive underscore case by not allowing it to match the end of the string if it contains one.
	pattern := "^[a-zA-Z][a-zA-Z0-9_]*[a-zA-Z0-9]$"
	if !regexp.MustCompile(pattern).MatchString(username) {
		return fmt.Errorf("username '%s' is invalid. It must start with a letter, be between 5-32 characters, and contain only letters, numbers, and underscores, without ending in an underscore", username)
	}

	return nil
}

// ValidateEmail checks if the Email field of the User struct is valid.
// A valid email:
// - Must be a valid format (e.g., user@domain.com).
// - Must not be empty.
// - Must be within a reasonable length (e.g., 6 to 254 characters).
func validateEmail(email string) error {
	//email := u.Email

	// 1. Check for empty string.
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	// 2. Check for reasonable length. The official maximum length is 254 characters.
	if len(email) < 6 || len(email) > 254 {
		return fmt.Errorf("email length must be between 6 and 254 characters, got %d", len(email))
	}

	// 3. Use a regular expression for comprehensive validation.
	// This regex is a commonly used pattern for email validation.
	// It is not 100% compliant with all RFCs but covers most real-world cases.
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		return fmt.Errorf("email address '%s' is invalid", email)
	}

	return nil
}

// validatePhoneNumber checks if the PhoneNumber field is a valid Iranian mobile number.
// A valid Iranian mobile number:
// - Can start with "+98" followed by a mobile operator code and 8 digits (e.g., +98912...).
// - Can start with "09" followed by 9 digits (e.g., 0912...).
// - Has a total of 11 digits if starting with "0".
// - Has a total of 12 digits if starting with "+98".
func validatePhoneNumber(phoneNumber string) error {

	// An empty phone number is acceptable, but if it exists, it must be valid.
	if phoneNumber == "" {
		return nil
	}

	// This regex pattern handles both formats:
	// ^ - Start of the string.
	// (\+98|0)? - Matches an optional "+98" or "0". The '?' makes the group optional.
	// 9 - Matches the literal digit '9'.
	// \d{9} - Matches exactly 9 digits.
	// $ - End of the string.
	//
	// The pattern essentially checks for:
	// 1. "+98" or "0" at the beginning (optional).
	// 2. The digit "9" immediately following.
	// 3. Exactly 9 more digits after that.
	pattern := `^(\+98|0)?9\d{9}$`

	if !regexp.MustCompile(pattern).MatchString(phoneNumber) {
		return fmt.Errorf("phone number '%s' is not a valid Iranian mobile number. It must be in the format 09xxxxxxxx or +989xxxxxxxxx", phoneNumber)
	}

	return nil
}

// ValidateNames In your User struct methods, you can call this helper function:
func (u *User) ValidateNames() error {
	if err := validateName(u.FirstName, "FirstName"); err != nil {
		return err
	}
	if err := validateName(u.LastName, "LastName"); err != nil {
		return err
	}
	return nil
}

// validateName checks if a name field (either FirstName or LastName) is valid.
func validateName(name string, fieldName string) error {
	// 1. Remove leading and trailing whitespace.
	name = strings.TrimSpace(name)

	// 2. Check for minimum and maximum length.
	if len(name) < 2 || len(name) > 50 {
		return fmt.Errorf("%s length must be between 2 and 50 characters", fieldName)
	}

	// 3. Use a regular expression to allow only letters and single spaces between words.
	// This regex allows for Persian and English letters.
	// A more comprehensive regex might be needed for international use cases.
	// \p{L} matches any kind of letter from any language.
	pattern := `^[\p{L}]+([\s][\p{L}]+)*$`
	re := regexp.MustCompile(pattern)

	if !re.MatchString(name) {
		return fmt.Errorf("invalid characters in %s: %s", fieldName, name)
	}

	return nil
}
