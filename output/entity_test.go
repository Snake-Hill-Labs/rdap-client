package output

import (
	"testing"
	"time"

	"github.com/registrobr/rdap-client/Godeps/_workspace/src/github.com/registrobr/rdap/protocol"
)

var expectedEntityOutput = `handle:   XXXX
person:   Joe User
e-mail:   joe.user@example.com
address:  Av Naçoes Unidas, 11541, 7 andar, Sao Paulo, SP, 04578-000, BR
phone:    tel:+55-11-5509-3506;ext=3506
created:  20150301
changed:  20150310

`

func TestEntityPrint(t *testing.T) {
	entityResponse := protocol.Entity{
		ObjectClassName: "entity",
		Handle:          "XXXX",
		VCardArray: []interface{}{
			"vcard",
			[]interface{}{
				[]interface{}{"version", struct{}{}, "text", "4.0"},
				[]interface{}{"fn", struct{}{}, "text", "Joe User"},
				[]interface{}{"kind", struct{}{}, "text", "individual"},
				[]interface{}{"email", struct{ Type string }{Type: "work"}, "text", "joe.user@example.com"},
				[]interface{}{"lang", struct{ Pref string }{Pref: "1"}, "language-tag", "pt"},
				[]interface{}{"adr", struct{ Type string }{Type: "work"}, "text",
					[]interface{}{
						"Av Naçoes Unidas", "11541", "7 andar", "Sao Paulo", "SP", "04578-000", "BR",
					},
				},
				[]interface{}{"tel", struct{ Type string }{Type: "work"}, "uri", "tel:+55-11-5509-3506;ext=3506"},
			},
		},
		LegalRepresentative: "Joe User",
		Entities: []protocol.Entity{
			{
				ObjectClassName: "entity",
				Handle:          "XXXX",
				VCardArray: []interface{}{
					"vcard",
					[]interface{}{
						[]interface{}{"version", struct{}{}, "text", "4.0"},
						[]interface{}{"fn", struct{}{}, "text", "Joe User"},
						[]interface{}{"kind", struct{}{}, "text", "individual"},
						[]interface{}{"email", struct{ Type string }{Type: "work"}, "text", "joe.user@example.com"},
						[]interface{}{"lang", struct{ Pref string }{Pref: "1"}, "language-tag", "pt"},
						[]interface{}{"adr", struct{ Type string }{Type: "work"}, "text",
							[]interface{}{
								"Av Naçoes Unidas", "11541", "7 andar", "Sao Paulo", "SP", "04578-000", "BR",
							},
						},
						[]interface{}{"tel", struct{ Type string }{Type: "work"}, "uri", "tel:+55-11-5509-3506;ext=3506"},
					},
				},
				Events: []protocol.Event{
					{
						Action: protocol.EventActionRegistration,
						Date:   protocol.Date(2015, 03, 01, 12, 00, 00, 00, time.UTC),
					},
					{
						Action: protocol.EventActionLastChanged,
						Date:   protocol.Date(2015, 03, 10, 14, 00, 00, 00, time.UTC),
					},
				},
			},
		},
		Events: []protocol.Event{
			{
				Action: protocol.EventActionRegistration,
				Date:   protocol.Date(2015, 03, 01, 12, 00, 00, 00, time.UTC),
			},
			{
				Action: protocol.EventActionLastChanged,
				Date:   protocol.Date(2015, 03, 10, 14, 00, 00, 00, time.UTC),
			},
		},
	}

	entityOutput := Entity{Entity: &entityResponse}

	var w WriterMock
	if err := entityOutput.Print(&w); err != nil {
		t.Fatal(err)
	}

	if string(w.Content) != expectedEntityOutput {
		for _, l := range diff(expectedEntityOutput, string(w.Content)) {
			t.Log(l)
		}
		t.Fatal("error")
	}
}
