package main

import (
	"errors"
	"strings"

	"github.com/wailsapp/wails"
)

//ContactsService Comment
type ContactsService interface {
	SaveContact(contact Contact) error
	ReadContact(fname string) (Contact, error)
	RemoveContact(fname string) error
	AllContacts() ([]Contact, error)
}

// Contact Comment
type Contact struct {
	Firstname, Lastname  string
	Mobile, Home, Office int
}

type contactService struct{}
type loggingService struct {
	log *wails.CustomLogger
}

func (m *loggingService) WailsInit(runtime *wails.Runtime) error {
	m.log = runtime.Log.New("LoggingService")
	return nil
}

var contacts map[string]Contact

// NewContactService Comment
func NewContactService() ContactsService {
	contacts = make(map[string]Contact)
	return &contactService{}
}

func (svc *contactService) SaveContact(contact Contact) error {
	if validateContact(contact) {
		contacts[buildKey(contact)] = contact
		return nil
	}
	return errors.New("Input Invalid")
}

func (svc *contactService) ReadContact(fname string) (Contact, error) {
	if len(fname) == 0 {
		return Contact{}, errors.New("Input Invalid")
	}

	for key, val := range contacts {
		if strings.Contains(key, fname) {
			return val, nil
		}
	}
	return Contact{}, errors.New("Not Found")
}

func (svc *contactService) RemoveContact(fname string) error {
	if len(fname) == 0 {
		return errors.New("Input Invalid")
	}

	rkey := ""
	for key := range contacts {
		if strings.Contains(key, fname) {
			rkey = key
			break
		}
	}

	if len(rkey) == 0 {
		return errors.New("Not Found")
	}
	delete(contacts, rkey)
	return nil
}

func (svc *contactService) AllContacts() ([]Contact, error) {
	contactlist := make([]Contact, 0)
	for _, val := range contacts {
		contactlist = append(contactlist, val)
	}
	return contactlist, nil
}

func buildKey(contact Contact) string {
	return contact.Firstname + "-" + contact.Lastname
}

func validateContact(contact Contact) bool {
	if len(contact.Firstname) == 0 {
		return false
	}
	if len(contact.Lastname) == 0 {
		return false
	}
	return true
}
