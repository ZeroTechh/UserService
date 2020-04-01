package meta

// AccountStatuses refers to status of account
type AccountStatuses int

const (
	// VERIFIED is when user's email is verified
	VERIFIED AccountStatuses = iota
	// UNVERIFIED is when user's email is not verified
	UNVERIFIED
	// DELETED is when user's account is deleted
	DELETED
)

func (a AccountStatuses) String() string {
	statuses := []string{"verified", "unverified", "deleted"}
	return accountStatuses.Str(statuses[a])
}
