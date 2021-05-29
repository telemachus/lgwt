package dictionary

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("could not find the world you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

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

func (d Dictionary) Add(k, v string) error {
	_, err := d.Search(k)

	switch err {
	case ErrNotFound:
		d[k] = v
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(k, v string) error {
	_, err := d.Search(k)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[k] = v
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(k string) {
	delete(d, k)
}

func Search(m map[string]string, k string) string {
	return m[k]
}
