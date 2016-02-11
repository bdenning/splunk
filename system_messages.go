package splunk

import (
	"fmt"
	"net/url"
)

const (
	MessageSeverityInfo  = "info"
	MessageSeverityWarn  = "warn"
	MessageSeverityError = "error"
)

type Message struct {
	Name     string
	Value    string
	Severity string
}

func DisplayMessage(s *Session, m Message) (err error) {
	v := url.Values{}
	v.Set("name", m.Name)
	v.Set("value", m.Value)
	v.Set("severity", m.Severity)

	_, err = s.Request("POST", "/services/messages", v)
	if err != nil {
		return err
	}

	return nil
}

func RemoveMessage(s *Session, m Message) (err error) {
	_, err = s.Request("DELETE", fmt.Sprintf("/services/messages/%s", m.Name), url.Values{})
	if err != nil {
		return err
	}

	return nil
}
