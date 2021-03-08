package parser

// Match contains all of the
// fields required to check
// and print matched patterns
type Match struct{}

// func (m *Match) isMatch(reg regexp.Regexp, s string) bool {
// 	return reg.MatchString(s)
// }

// // print prints out the results of a match
// // to an io.Writer
// func (m *Match) print(w io.Writer) {
// 	_, _ = fmt.Fprintf(w, "[MATCH] \nPattern: %q\nLink: %q\n---\n", m.Pattern, m.Feed.Link)
// }

// // render prints out the results of a match
// // to an io.Writer using a go text/template
// // for formatting
// func (m *Match) render(w io.Writer, tmpl string) {
// 	t, err := template.New("match").Parse(tmpl)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	if err := t.ExecuteTemplate(w, "match", &m.Feed); err != nil {
// 		panic(err)
// 	}
// }
