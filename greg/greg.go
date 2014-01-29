// play with regular expressions in Go a la rubular.com
//
// The workings are simple:
//
//  - Anything but '/compile' brings you the front page
//  - /compile expects a regular expressions (mandatory) and optionally a list of parameters to try out.
package greg

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, header)
	io.WriteString(w, gregHeader)
	io.WriteString(w, regForm)
	io.WriteString(w, footer)
}

// Compile a given regular expression and write out the result in JSON
//
// There are quite a few switches to take care of
//
//  - Posix?
//  -
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
	if len(inputs) == 1 && inputs[0] == "" { // TODO: Better input validation
		j.Encode(map[string]string{"error": "greg: Input cannot be empty"})
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

	// TODO(tj): Provide simplified expression (using regexp/syntax)
	//ret["simple"] = re.Simplify().String()

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
}
