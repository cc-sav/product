package storage

import (
	"encoding/gob"
	"os"
	"savvie/users"
	"time"
)

// Comment stores everything about an option's comment. User contains a username
// (email address). Option contains an Option.ID. Body will be escaped, so HTML
// inside will not be rendered.
type Comment struct {
	User   string
	Option string
	Date   time.Time
	Body   string
	Type   string // "Comment", "Pro", or "Con".
}

// NiceDate returns an output-ready, human-recognizable date string.
// The format of the date is not guaranteed to remain constant. It may change without notice.
func (c *Comment) NiceDate() string {
	return c.Date.Local().Format("Jan 2, 2006 at 15:04 MST")
}

// NiceName returns the human name of the comment's creator.
func (c *Comment) NiceName() string {
	return users.GetUser(c.User).Name
}

const commentsFolder = "/vagrant/data/comments/"

// LoadComments returns all of the comments on the given option, but they're not guaranteed to be sorted.
// They will be sorted in a future version.
func LoadComments(optID string) []Comment {
	filename := commentsFolder + optID + "-comments.txt"
	file, err := os.Open(filename)
	if err != nil {
		return []Comment{}
	}
	defer file.Close()

	result := []Comment{}
	dec := gob.NewDecoder(file)
	dec.Decode(&result)
	return result
}

// SaveNewComment persists a new comment with the given information.
func SaveNewComment(user, optID, body string) error {
	if user == "" || optID == "" || body == "" {
		return Error{"Could not create comment with given information."}
	}

	all := LoadComments(optID)
	filename := commentsFolder + optID + "-comments.txt"
	file, err := os.Create(filename)
	if err != nil {
		return Error{"Could not open file: " + filename}
	}
	defer file.Close()

	c := Comment{
		User:   user,
		Option: optID,
		// TODO: set location when we set the time
		Date: time.Now(),
		Body: body,
	}
	all = append(all, c)

	enc := gob.NewEncoder(file)
	enc.Encode(all)
	return nil
}

// Error provides information about what went wrong.
type Error struct {
	e string
}

func (e Error) Error() string {
	return "comments: " + e.e
}
