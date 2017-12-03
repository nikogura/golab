package template

import (
	"text/template"
	"bytes"
)

var card1 = map[int]string{1: "S", 2: "H", 3: "C", 4: "D"}

func SimpleTemplate() string {
	t := template.Must(template.New("t1").Parse("Dot:{{.}}"))

	buf := new(bytes.Buffer)

	t.Execute(buf, card1)

	return buf.String()
}
