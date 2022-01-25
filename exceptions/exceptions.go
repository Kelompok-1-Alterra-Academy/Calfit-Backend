package exceptions

import "errors"

var (
	// ErrInvalidCredentials is thrown when the user credentials are invalid
	ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrInternalServerError is thrown when the server encounters an error
	ErrInternalServerError = errors.New("internal server error")
	// ErrNotFound is thrown when the object is not found
	ErrNotFound = errors.New("not found")
	// ErrUserAlreadyExists is thrown when the user already exists
	ErrUserAlreadyExists = errors.New("user already exists")
	// ErrBookingNotFound is thrown when the booking is not found
	ErrBookingNotFound = errors.New("booking not found")
	// ErrGymNotFound is thrown when the gym is not found
	ErrGymNotFound = errors.New("gym not found")
	// ErrClassNotFound is thrown when the class is not found
	ErrClassNotFound = errors.New("class not found")
	// ErrMembershipNotFound is thrown when the membership is not found
	ErrMembershipNotFound = errors.New("membership not found")
	// ErrUserNotFound is thrown when the user is not found
	ErrUserNotFound = errors.New("user not found")
	// ErrEmptyInput is thrown when the input is empty
	ErrEmptyInput = errors.New("empty input")
	// ErrValidationFailed is thrown when the input validation is failed
	ErrValidationFailed = errors.New("validation failed")
	// ErrSessionNotFound is thrown when the session is not found
	ErrSessionNotFound = errors.New("session not found")
	// ErrScheduleNotFound is thrown when the schedule is not found
	ErrScheduleNotFound = errors.New("schedule not found")
	// ErrMissingId is thrown when id is missing
	ErrMissingId = errors.New("missing require id")
	// ErrBadRequest is thrown when request message is invalid
	ErrBadRequest = errors.New("invalid request message")
	// ErrOperationalAdminExists is thrown when the superadmin already exists
	ErrOperationalAdminExists = errors.New("operational admin already exists")
	// ErrOperationalAdminNotFound is thrown when the superadmin is not found
	ErrOperationalAdminNotFound = errors.New("operational admin not found")
	// ErrSuperadminExists is thrown when the superadmin already exists
	ErrSuperadminExists = errors.New("superadmin already exists")
	// ErrSuperadminNotFound is thrown when the superadmin is not found
	ErrSuperadminNotFound = errors.New("superadmin not found")
)
