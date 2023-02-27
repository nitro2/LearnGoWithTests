package dictionary

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(s string) (string, error) {
	definition, ok := d[s]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	if d[key] != "" {
		return ErrWordExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}
	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	if err != nil {
		return err
	}

	delete(d, key)
	return nil

}
