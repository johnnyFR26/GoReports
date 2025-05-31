package reporter

import (
	"bytes"
	"html/template"
)

const tpl = `
<h1>Relatório de Tarefas - {{ .Date }}</h1>
<table border="1" cellpadding="8" cellspacing="0">
    <tr>
        <th>Funcionário</th>
        <th>Tarefa</th>
        <th>Quantidade</th>
    </tr>
    {{ range .Entries }}
    <tr>
        <td>{{ .EmployeeName }}</td>
        <td>{{ .TaskName }}</td>
        <td>{{ .Count }}</td>
    </tr>
    {{ end }}
</table>
`

type ReportEntry struct {
	EmployeeName string
	TaskName     string
	Count        int
}

type ReportData struct {
	Date    string
	Entries []ReportEntry
}

func GenerateHTML(data ReportData) (string, error) {
	t, err := template.New("report").Parse(tpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
