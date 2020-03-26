package configfile

import (
    "encoding/json"
    "log"
    "strings"
    "text/template"
)

// JSON stringify a config file
func (cf ConfigFile) String() string {
    stringified, err := json.MarshalIndent(cf, "", "    ")
    if err != nil {
        return ""
    }

    return string(stringified)
}

// JSON stringify a host entry
func (he HostEntry) String() string {
    stringified, err := json.MarshalIndent(he, "", "    ")
    if err != nil {
        return ""
    }

    return string(stringified)
}

// Iterates through the config file's host entries,
// calling their FormattedText() method
// @NOTE the use of json.MarshalIndent will not properly render
// some characters.
func (cf ConfigFile) FormattedText() string {
    outstr := ""
    for _, host := range cf.Hosts {
        outstr += host.FormattedText() + "\n"
    }

    return outstr
}

// Template for output formatted Host Entries
const TEMPLATE = `
Host {{.Host}}
    HostName {{.HostName}}{{if .User}} 
    User {{.User}}{{ end }}{{ if .IdentityFilePath}}
    IdentityFile {{.IdentityFilePath}}{{ end }}
`

// Populates the template with the HostEntry's data
// @NOTE the use of json.MarshalIndent will not properly render
// some characters.
func (he HostEntry) FormattedText() string {

    buf := &strings.Builder{}
    t, err := template.New("host").Parse(TEMPLATE)
    if err != nil {
        log.Println("Could not hold back.")
        return ""
    }

    err = t.Execute(buf, he)
    if err != nil {
        log.Println("error: ", err)
    }
    return buf.String()
}
