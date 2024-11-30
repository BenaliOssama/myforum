package modules

import (
	"time"
)

// Define a Snippet type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets
// table?
type User struct {
	UserId   int
	UserName int
	JoinedAt time.Time
}

// This will insert a new snippet into the database.
func (m *ForumModel) InsertPost(title string, content string, expires int) (int, error) {
	return 0, nil
}

// This will return a specific snippet based on its id.
func (m *User) GetPost(id int) (*User, error) {
	return nil, nil
}

// This will return the 10 most recently created snippets.
func (m *ForumModel) LatestPost() ([]*User, error) {
	return nil, nil
}
