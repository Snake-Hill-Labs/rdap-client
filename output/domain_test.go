package output

import (
	"testing"
	"time"

	"github.com/registrobr/rdap-client/Godeps/_workspace/src/github.com/registrobr/rdap/protocol"
)

func TestDomainPrint(t *testing.T) {
	domain := Domain{
		Domain: &protocol.Domain{
			ObjectClassName: "domain",
			LDHName:         "example.br",
			Status: []protocol.Status{
				"active",
			},
			Links: []protocol.Link{
				{
					Value: "https://rdap.registro.br/domain/example.br",
					Rel:   "self",
					Href:  "https://rdap.registro.br/domain/example.br",
					Type:  "application/rdap+json",
				},
			},
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
				{
					ObjectClassName: "entity",
					Handle:          "YYYY",
					VCardArray: []interface{}{
						"vcard",
						[]interface{}{
							[]interface{}{"version", struct{}{}, "text", "4.0"},
							[]interface{}{"fn", struct{}{}, "text", "Joe User 2"},
							[]interface{}{"kind", struct{}{}, "text", "individual"},
							[]interface{}{"email", struct{ Type string }{Type: "work"}, "text", "joe.user2@example.com"},
							[]interface{}{"lang", struct{ Pref string }{Pref: "1"}, "language-tag", "pt"},
							[]interface{}{"adr", struct{ Type string }{Type: "work"}, "text",
								[]interface{}{
									"Av Naçoes Unidas", "11541", "7 andar", "Sao Paulo", "SP", "04578-000", "BR",
								},
							},
							[]interface{}{"tel", struct{ Type string }{Type: "work"}, "uri", "tel:+55-11-5509-3506;ext=3507"},
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
			Nameservers: []protocol.Nameserver{
				{
					ObjectClassName: "nameserver",
					LDHName:         "a.dns.br",
					Events: []protocol.Event{
						{
							Action: protocol.EventDelegationCheck,
							Status: []protocol.Status{protocol.StatusNSAA},
						},
					},
				},
			},
			SecureDNS: &protocol.SecureDNS{
				DSData: []protocol.DS{
					{
						KeyTag:    12345,
						Digest:    "0123456789ABCDEF0123456789ABCDEF01234567",
						Algorithm: 5,
						Events: []protocol.Event{
							{
								Action: protocol.EventActionRegistration,
								Date:   protocol.Date(2015, 03, 01, 12, 00, 00, 00, time.UTC),
							},
							{
								Action: protocol.EventDelegationSignCheck,
								Status: []protocol.Status{protocol.StatusDSOK},
								Date:   protocol.Date(2015, 03, 01, 12, 00, 00, 00, time.UTC),
							},
						},
					},
				},
			},
		},
	}

	expected := `
domain:   example.br
nserver:  a.dns.br
dsrecord: 12345 RSASHA1 0123456789ABCDEF0123456789ABCDEF01234567
dsstatus: 20150301 ds ok
dslastok: 00010101
created:  20150301
changed:  20150310
status:   active

handle:   XXXX
person:   Joe User
e-mail:   joe.user@example.com
address:  Av Naçoes Unidas, 11541, 7 andar, Sao Paulo, SP, 04578-000, BR
phone:    tel:+55-11-5509-3506;ext=3506
created:  20150301
changed:  20150310

handle:   YYYY
person:   Joe User 2
e-mail:   joe.user2@example.com
address:  Av Naçoes Unidas, 11541, 7 andar, Sao Paulo, SP, 04578-000, BR
phone:    tel:+55-11-5509-3506;ext=3507
created:  20150301
changed:  20150310

`

	var w WriterMock
	if err := domain.Print(&w); err != nil {
		t.Fatal(err)
	}

	if string(w.Content) != expected {
		for _, l := range diff(expected, string(w.Content)) {
			t.Log(l)
		}
		t.Fatal("error")
	}
}
