package output

import (
	"io"
	"text/template"

	"github.com/registrobr/rdap-client/Godeps/_workspace/src/github.com/registrobr/rdap/protocol"
)

type Entity struct {
	Entity *protocol.Entity

	CreatedAt string
	UpdatedAt string

	ContactsInfos []contactInfo
}

func (e *Entity) AddContact(c contactInfo) {
	e.ContactsInfos = append(e.ContactsInfos, c)
}

func (e *Entity) setDates() {
	for _, event := range e.Entity.Events {
		date := event.Date.Format("20060102")

		switch event.Action {
		case protocol.EventActionRegistration:
			e.CreatedAt = date
		case protocol.EventActionLastChanged:
			e.UpdatedAt = date
		}
	}
}

func (e *Entity) Print(wr io.Writer) error {
	e.setDates()
	var contactInfo contactInfo
	contactInfo.setContact(*e.Entity)
	e.ContactsInfos = append(e.ContactsInfos, contactInfo)

	t, err := template.New("entity template").
		Funcs(contactInfoFuncMap).
		Parse(contactTmpl)

	if err != nil {
		return err
	}

	return t.Execute(wr, e)
}
