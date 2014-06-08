// play with regular expressions in Go a la rubular.com
//
// The workings are simple:
//
//  - Anything but '/compile' brings you the front page
//  - /compile expects a regular expressions (mandatory) and optionally a list of parameters to try out.
//
// TODO Use POSIX matcher if specified
// TODO Don't throw away blank lines
// TODO Provide simplify functionality
package greg

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"regexp/syntax"
	"strings"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, header)
	io.WriteString(w, gregHeader)
	io.WriteString(w, regForm)
	io.WriteString(w, footer)
}

// Encoding an error
func encErr(j *json.Encoder, err error) {
	j.Encode(map[string]string{"error": err.Error()})
}

// Simplifying a regular expression
// Just returns encoded as JSON: regex => simplified expression
func simplify(w http.ResponseWriter, req *http.Request) {
	j := json.NewEncoder(w)
	if err := req.ParseForm(); err != nil {
		encErr(j, err)
		return
	}
	rex := req.PostForm.Get("regex")

	// TODO(tj)
	re, err := syntax.Parse(rex, syntax.Perl)
	if err != nil {
		encErr(j, err)
		return
	}
	j.Encode(map[string]string{"regex": re.Simplify().String()})
}

// Compile a given regular expression and write out the result in JSON
// The result is something like
//
// 	{ "matches":
//		[ { "im":   [[15,24,15,16,17,19,20,24]] ,"input": "Today date is: 1/29/2014."}
//		]
//	, "names": ["","month","day","year"]
//	, "regex": "(?P\u003cmonth\u003e\\d{1,2})\\/(?P\u003cday\u003e\\d{1,2})\\/(?P\u003cyear\u003e\\d{4})"
//	}
//
// Where the matches are an array of objects. Each object then in turns contains
// a list of matches, im, which is [][]int as returned by
//
//	re.FindAllStringSubmatchIndex()
//
// for each input line. The array 'matches' is sorted according to the input
// lines.
func compile(w http.ResponseWriter, req *http.Request) {
	j := json.NewEncoder(w)

	if err := req.ParseForm(); err != nil {
		j.Encode(map[string]string{"error": err.Error()})
		return
	}
	exp := req.PostForm.Get("regex")
	if exp == "" {
		j.Encode(map[string]string{"error": "greg: Regex cannot be empty"})
		return
	}
	input := req.PostForm.Get("regex-input")
	if input == "" {
		j.Encode(map[string]string{"error": "greg: Input cannot be empty"})
		return
	}
	inputs := strings.Split(input, "\r\n")
	if len(inputs) == 1 && inputs[0] == "" {
		j.Encode(map[string]string{"error": "greg: Input cannot be empty"})
		return
	}

	re, err := regexp.Compile(exp)
	if err != nil {
		j.Encode(map[string]string{"error": err.Error()})
		return
	}

	// Return value
	ret := make(map[string]interface{}, 4)
	ret["regex"] = re.String()
	ret["names"] = re.SubexpNames()

	// Do matching
	var matches []map[string]interface{}
	for _, in := range inputs {
		im := re.FindAllStringSubmatchIndex(in, -1)
		matches = append(matches, map[string]interface{}{"input": in, "im": im})
	}

	ret["matches"] = matches

	err = j.Encode(ret)
	if err != nil {
		j.Encode(map[string]string{"error": err.Error()})
	}
}

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/compile", compile)
	http.HandleFunc("/simplify", simplify)
}
