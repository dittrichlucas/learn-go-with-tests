package dictionary

// Dictionary is
type Dictionary map[string]string

// ErrDictionary is
type ErrDictionary string

// ErrNotFound is
const (
	ErrWordNotFound    = ErrDictionary("word not found")
	ErrExistingWord    = ErrDictionary("word already exists")
	ErrNonExistingWord = ErrDictionary("word does not exist")
)

func (e ErrDictionary) Error() string {
	return string(e)
}

// Add is
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Find(word)
	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrExistingWord
	default:
		return err
	}

	return nil
}

// Update is
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Find(word)
	switch err {
	case ErrWordNotFound:
		return ErrNonExistingWord
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete is
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

// Find is
func (d Dictionary) Find(word string) (string, error) {
	definition, exist := d[word]

	if !exist {
		return "", ErrWordNotFound
	}

	return definition, nil
}
