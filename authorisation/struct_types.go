package authorisation

// Role Define User Role entity
type Role struct{
	id int
	description string
	parentID int
	childRoles []int
}

// User Define user entity
type User struct {
	id int64
	name string
	roleID int
}

// ProcessingResult define processing result entity
type ProcessingResult struct {
	itemIndex int
	message *string
}