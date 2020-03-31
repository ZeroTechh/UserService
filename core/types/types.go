package types

// Main is used to store user main data
type Main struct {
	UserID   string `bson:"_id,omitempty,-"`
	Username string `bson:"Username,omitempty,-"`
	Email    string `bson:"Email,omitempty,-"`
	Password string `bson:"Password,omitempty,-"`
}

// Extra is used to store user extra data
type Extra struct {
	UserID      string `bson:"_id,omitempty,-"`
	FirstName   string `bson:"FirstName,omitempty,-"`
	LastName    string `bson:"LastName,omitempty,-"`
	Gender      string `bson:"Gender,omitempty,-"`
	BirthdayUTC int64  `bson:"BirthdayUTC,omitempty,-"`
}

// Meta is used to store user meta data
type Meta struct {
	UserID             string `bson:"_id,omitempty,-"`
	AccountStatus      string `bson:"AccountStatus,omitempty,-"`
	AccountCreationUTC int64  `bson:"AccountCreationUTC,omitempty,-"`
}
