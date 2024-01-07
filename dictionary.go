package main

type Dictionary map[string]string
type DictionaryErr string

const ErrUnknown = DictionaryErr("unknown word. cannot define")
const ErrWordExists = DictionaryErr("word already defined. cannot add")
const ErrWordNotFound = DictionaryErr("word is not defined. cannot update")

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrUnknown
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrUnknown:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrUnknown:
		return ErrWordNotFound
	case nil:
		d[word] = newDefinition
		return nil
	default:
		return err
	}
}
