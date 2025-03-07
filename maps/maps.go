package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("key not found in dictionary")
	ErrWordExists       = DictionaryErr("word already exists in Map")
	ErrWordDoesNotExist = DictionaryErr("Could not perform action on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		d[key] = value
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) error {
	delete(d, key)
}
