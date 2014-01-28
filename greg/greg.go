// play with regular expressions in Go a la rubular.com
//
// The worknings are simple:
//
//  - Anything but '/compile' brings you the front page
//  - /compile expects a regular expressions (mandatory) and optionally a list of parameters to try out.
package main

import (
	"io"
	"net/http"
	"strings"
	"encoding/json"
	"regexp"
)

// The header
const header = `<!DOCTYPE html>
<html lang="en">
<head>
	<title>greg: a Go regular expression editor and tester</title>

	<!-- Latest compiled and minified CSS -->
	<!-- link rel="stylesheet"
	href="//netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css"
	-->

	<!-- Optional theme -->
	<!-- http://bootswatch.com/slate/bootstrap.min.css -->
	<link rel="stylesheet"	href="//bootswatch.com/cyborg/bootstrap.min.css">

	<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    	<script src="https://code.jquery.com/jquery.js"></script>

	<!-- Latest compiled and minified JavaScript -->
	<script src="//netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>

	<style type="text/css">
	#regex, #regex-input, #match-result, #match-groups {
		font-family: monospace;
	}
	span.match {
		background-color: #f80;
	}
	</style>
</head>
<body>
<div class="container">
`

const footer = `
<script type="text/javascript">
$(function() {

	var Greg = {
		showMatches: function(json) {
			if (json.error) {
				$("#regex-match").addClass("hidden");
				$("#regex-fail").html(json.error);
				$("#regex-fail").removeClass("hidden");
				return;
			}
			console.log(json);
			var result = $("#match-result");
			var groups = $("#match-groups");

			var list = '<ul>';
			for (var key in json.matches) {
				list +=  '<li>';
				var input = json.matches[key].input;
				var im = json.matches[key].im; // im is [][]int
				if (im == null) {
					list += input;
				} else {
					var pos = 0;
					im.forEach(function(m) {
						var i = m[0], j = m[1];
						console.log("before:  (" + pos + ", " + i + ") #=> " + input.substring(pos,i));
						console.log("matched: (" + i + ", " + j + ") #=> " + input.substring(i,j));
						list += input.substring(pos, i); // before match
						list += '<span class="match">' + input.substring(i, j) + '</span>'; // match
						pos = j; // move pos up	
					});
					list += input.substring(pos, input.length);

				}
				list += '</li>';
			}
			list += '</ul>';
			result.html(list);
			$("#regex-fail").addClass("hidden");
			//$("#regex-match").html(json.regex);
			$("#regex-match").removeClass("hidden");
		},
		showSpinner: function() {
			$("#spin").removeClass("hidden");
		},
		hideSpinner: function() {
			$("#spin").addClass("hidden");
		},
		update: function() {
			$("#greeting").addClass("hidden");
			$.ajax({
				type:    "POST",
				url:     "/compile",
				data:    $("#greg").serialize(),
				success: function(data, status, jqXHR) {
					$("#regex-match").removeClass("hidden");
					Greg.showMatches(data);
				},
				error:   function(jqXHR, status, err) {
					$("#regex-fail").removeClass("hidden");
					$("#regex-fail").html(status);
				},
				dataType: "json"
			});
		}
	};

	$("#regex").change(Greg.update);
	$("#regex-input").change(Greg.update);
	$("#regex-mod").change(Greg.update);
});
</script>

<!--
<div class="row well">
	<h4>Oh, well</h4>
	<ul>
		<li>Make regex <em>actually</em> match the test input
			<ul>
				<li>backend compiles regex</li>
				<li>matches against input</li>
				<li>return list of matches (as JSON)</li>
				<li>matches are lined up</li>
			</ul>
		Instead of returning just a list of matching subindices we might be smarter to return a little more information. Return: a list of pairs, where each pair is (input, submatches) in which submatches is a list of indices.
		</li>
		<li>Is there any use for flag?</em>
		<li>Split up files, so js files go in the right directory</li>
		<li>Provide simplified regex for convenience</li>
	</ul>
</div>
-->
</div>
</body>
</html>`

const gregHeader = `<div class="row">
	<h1 class="text-center">Greg<br /><small>A Go regular expression editor and tester</small></h1>
</div>`

const regForm = `
<div class="row">
<form action="/compile" role="form" method="POST" id="greg">
	<div class="controls">
		<div class="col-xs-10">
			<div class="input-group">
				<span class="input-group-addon">/</span>
				<input type="text" class="form-control" name="regex" id="regex" placeholder="Enter regex here" />
			</div>
			<span class="help-block">Your regular expression</span>
		</div>
		<div class="col-xs-2">
			<input type="text" class="form-control" name="regex-mod" id="regex-mod" placeholder="" />
			<span class="help-block">Modifiers</span>
		</div>
		<div class="col-xs-6">
			<textarea class="form-control" rows="10" name="regex-input" id="regex-input"></textarea>
			<span class="help-block">Your test input</span>
		</div>
		<div class="col-xs-6">
			<div class="text-center alert alert-success" id="greeting">
				Greg is a <a href="http://golang.org">Go</a>-based regular expression editor.
				It's a handy way to test regular expressions as you write them.

				To start, enter a regular expression and a test string.
			</div>
			<div class="hidden" id="regex-match">
				<div class="row panel panel-default">
					<div class="panel-heading">Match Result</div>
					<div class="panel-body" id="match-result"></div>
				</div>
				<div class="row panel panel-default">
					<div class="panel-heading">Match Groups</div>
					<div class="panel-body" id="match-groups"></div>
				</div>
			</div>
			<div class="hidden alert alert-danger" id="regex-fail">
			</div>
		</div>
	</div>
</form>
</div>
<div class="row">
	<div class="col-xs-12">
		<span id="spin" class="glyphicon glyphicon-record pull-right hidden"></span>
	</div>
</div>`

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
func regCompile(w http.ResponseWriter, req *http.Request) {
	j := json.NewEncoder(w) // Must, but could fail

	err := req.ParseForm()
	if err != nil {
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

	ret := make(map[string]interface{}, 4)
	ret["regex"] = re.String()
	ret["echo"] = input
	// TODO(tj): Provide simplified expression (using regexp/syntax)
	//ret["simple"] = re.Simplify().String()
	var matches []map[string]interface{}

	// Do matching
	for _, in := range inputs {
		im := re.FindAllStringSubmatchIndex(in, -1)
		matches = append(matches, map[string]interface{}{ "input": in, "im": im })
	}

	ret["matches"] = matches

	err = j.Encode(ret)
	if err != nil {
		j.Encode(map[string]string{"error": err.Error()})
	}
}

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/compile", regCompile)
	//fmt.Println("Go to http://localhost:9000")
	/*
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println(err)
	}
	*/
}
