package greg

// The header
const header =
`<!DOCTYPE html>
<html lang="en">
<head>
	<title>greg: a Go regular expression editor and tester</title>

	<!-- Latest compiled and minified CSS
	-->
	<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap.min.css">

	<!-- Optional theme -->
	<!-- http://bootswatch.com/slate/bootstrap.min.css
	<link rel="stylesheet"	href="http://bootswatch.com/slate/bootstrap.min.css">
	-->

	<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    	<script src="https://code.jquery.com/jquery.js"></script>
	<script src="/js/greg.js"></script>

	<!-- Latest compiled and minified JavaScript -->
	<script src="//netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>

	<style type="text/css">
	#regex, #regex-input, #match-result, #match-groups {
		font-family: monospace;
	}
	#match-result, #match-groups {
		/*color: #fff;*/
	}
	span.match {
		background-color: #f0ad4e;
	}
	span.rx-syntax {
		font-weight: bold;
		font-family: monospace;
		white-space: nowrap;
		padding-right: 2px;
		min-width: 40px;
		display: inline-block;
	}
	.table td {
		border-top: none;
	}
	table.qref {
		/*background-color: #272b30;*/
		margin: 2px;
		width: 100%;
	}
	table.qref td {
		vertical-align: top;
	}
	span.match-count {
		font-family: sans-serif;
	}
	</style>
</head>
<body>
<div class="container">
`

const footer =
`<script type="text/javascript">
$(function() {
	$("#regex").on('keyup change', Greg.update);
	$("#regex-input").on('keyup change', Greg.update);
	$("#regex-mod").on('keyup change', Greg.update);
});
</script>
</div>
</body>
</html>`

const gregHeader = `<div class="row">
	<h1 class="text-center">Greg<br /><small>A Go regular expression editor and tester</small></h1>
</div>`

const regForm = `
<div class="well">
<form role="form" id="greg">
<div class="row">
	<div class="controls">
		<div class="col-xs-12">
			<div class="input-group">
			<span class="input-group-addon">/</span>
				<input type="text" class="form-control" name="regex" id="regex" placeholder="Enter regex here" />
				<span class="input-group-addon">/</span>
			</div>
			<span class="help-block">Your regular expression</span>
		</div>
		<div class="col-xs-6">
			<textarea class="form-control" rows="10" name="regex-input" id="regex-input"></textarea>
			<!-- span class="help-block">Your test input</span -->
		</div>
		<div class="col-xs-6">
			<div class="text-center alert alert-success" id="greeting">
				Greg is a <a href="http://golang.org">Go</a>-based regular expression editor.
				It's a handy way to test regular expressions as you write them.

				To start, enter a regular expression and a test
				string, or try an <a href="#" onclick="Greg.example(); return false;">example</a>.
			</div>
			<div class="hidden" id="regex-match">
				<div class="panel panel-primary">
					<div class="panel-heading">Match Result</div>
					<div class="panel-body" id="match-result"></div>
				</div>
				<div class="panel panel-primary">
					<div class="panel-heading">Match Groups</div>
					<div class="panel-body" id="match-groups"></div>
				</div>
			</div>
			<div class="hidden" id="regex-fail">
				<div class="panel panel-danger">
					<div class="panel-heading"><strong>DANGER ZONE</strong></div>
					<div class="panel-body" id="error-msg"></div>
				</div>
			</div>
		</div>	
	</div>
</div>
<div class="row">
	<div class="col-xs-12">
	<div class="pull-left"><img src="img/ajax-loader.gif" id="spin" class="hidden" /></div>
	<div class="pull-right">
		<div class="btn-group text-left">
			<!--
			<button type="button" class="btn btn-sm btn-default">Simplify</button>
			<div class="btn-group">
			-->
			<button type="button" class="btn btn-default btn-sm dropdown-toggle" data-toggle="dropdown">
				Options <span class="caret"></span>
			</button>
			<ul class="dropdown-menu" role="menu">
				<li><a href="#" onclick="Greg.clear(); return false;">Clear fields</a></li>
			</ul>
			<!-- /div -->
		</div>
	</div>
	</div>
</div>
</form>
</div>
<div class="row">
<div class="col-xs-12">
		<ul class="nav nav-tabs">
			<li class="active"><a href="#qref1" data-toggle="tab">Syntax</a></li>
			<li><a href="#qref2" data-toggle="tab">Escape Sequences</a></li>
			<li><a href="#about" data-toggle="tab">About</a></li>
		</ul>
		<div class="tab-content">
			<div class="tab-pane fade active in" id="qref1">
			<div class="row">
				<div class="col-xs-4">
			<table class="qref">
					<tr>
						<td colspan=2><span class="label label-primary">Single characters</span></td>
					</tr>
					<tr>
						<td><span class="rx-syntax">.</span></td>
						<td>any character, possibly including newline</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">[xyz]</span></td>
						<td>character class</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">[^xyz]</span></td>
						<td>negated character class</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\d</span></td>
						<td>Perl character class</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\D</span></td>
						<td>negated Perl character class</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">[:alpha:]</span></td>
						<td>ASCII character class</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">[:^alpha:]</span></td>
						<td>negated ASCII character class</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\pN</span></td>
						<td>Unicode character class (one-letter name)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\p{Greek}</span></td>
						<td>Unicode character class</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\PN</span></td>
						<td>negated Unicode character class (one-letter name)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">P{Greek}</span></td>
						<td>negated Unicode character class</td>
					</tr>
				</table>
				</div>
				<div class="col-xs-4">
				<table class="qref">
					<tr>
						<td colspan=2><span class="label label-primary">Composites</span></td>
					</tr>
					<tr>
						<td><span class="rx-syntax">xy</span></td>
						<td>x followed by y</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x|y</span></td>
						<td>x or y (prefer x)</td>
					</tr>
					<tr>
						<td colspan=2><span class="label label-primary">Repetitions</span></td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x*</span></td>
						<td>zero or more x, prefer more</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x+</span></td>
						<td>one or more x, prefer more</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x?</span></td>
						<td>zero or one x, prefer one</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x{n,m}</span></td>
						<td>n or n+1 or ... or m x, prefer more</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x{n,}</span></td>
						<td>n or more x, prefer more</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x{n}</span></td>
						<td>exactly n x</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x*?</span></td>
						<td>zero or more x, prefer fewer</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x+?</span></td>
						<td>one or more x, prefer fewer</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x??</span></td>
						<td>zero or one x, prefer zero</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x{n,m}?</span></td>
						<td>n or n+1 or ... or m x, prefer fewer</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x{n,}?</span></td>
						<td>n or more x, prefer fewer</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x{n}?</span></td>
						<td>exactly n x</td>
					</tr>
				</table>

				</div>
				<div class="col-xs-4">
				<table class="qref">
					<tr>
						<td colspan=2><span class="label label-primary">Grouping</span></td>
					</tr>
					<tr>
						<td><span class="rx-syntax">(re)</span></td>
						<td>numbered capturing group</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">(?P&lt;name&gt;re)</span></td>
						<td>named & numbered capturing group (submatch)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">(?:re)</span></td>
						<td>non-capturing group (submatch)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">(?flags)</span></td>
						<td>set flags within current group</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">(?flags:re)</span></td>
						<td>set flags during re; non-capturing</td>
					</tr>
					<tr>
						<td><span class="label label-default">flag</span> <span class="rx-syntax">i</span> </td>
						<td>case-insensitive (default false)</td>
					</tr>
					<tr>
						<td><span class="label label-default">flag</span> <span class="rx-syntax">m</span> </td>
						<td>multi-line mode: ^ and $ match begin/end line in addition to begin/end text (default false)</td>
					</tr>
					<tr>
						<td><span class="label label-default">flag</span> <span class="rx-syntax">s</span> </td>
						<td>let . match \n (default false)</td>
					</tr>
					<tr>
						<td><span class="label label-default">flag</span> <span class="rx-syntax">U</span> </td>
						<td>ungreedy: swap meaning of x* and x*?, x+ and x+?, etc (default false)</td>
					</tr>
				</table>
				</div>
			</div>

			</div>
			<div class="tab-pane fade" id="qref2">
			<div class="row">
				<div class="col-xs-4">
			<table class="qref">
					<tr>
						<td colspan=2><span class="label label-primary">Empty strings</span></td>
					</tr>
					<tr>
						<td><span class="rx-syntax">^</span></td>
						<td>at beginning of text or line (flag m=true)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">$</span></td>
						<td>at end of text (like \z not \Z) or line (flag m=true)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\A</span></td>
						<td>at beginning of text</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\b</span></td>
						<td>at ASCII word boundary (\w on one side and \W, \A, or \z on the other)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\B</span></td>
						<td>not an ASCII word boundary</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\z</span></td>
						<td>at end of text</td>
					</tr>
				</table>
				</div>
				<div class="col-xs-4">
				<table class="qref">
					<tr>
						<td colspan=2><span class="label label-primary">Escape sequences</span></td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\a</span></td>
						<td>bell (== \007)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\f</span></td>
						<td>form feed (== \014)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\t</span></td>
						<td>horizontal tab (== \011)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\n</span></td>
						<td>newline (== \012)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\r</span></td>
						<td>carriage return (== \015)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\v</span></td>
						<td>vertical tab character (== \013)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\*</span></td>
						<td>literal *, for any punctuation character *</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\123</span></td>
						<td>octal character code (up to three digits)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\x7F</span></td>
						<td>hex character code (exactly two digits)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\x{10FFFF}</span></td>
						<td>hex character code</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\Q...\E</span></td>
						<td>literal text ... even if ... has punctuation</td>
					</tr>
				</table>

				</div>
				<div class="col-xs-4">
				<table class="qref">
					<tr>
						<td colspan=2><span class="label label-primary">Character class elements</span></td>
					</tr>
					<tr>
						<td><span class="rx-syntax">x</span></td>
						<td>single character</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">A-Z</span></td>
						<td>character range (inclusive)</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\d</span></td>
						<td>Perl character class</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">[:foo:]</span></td>
						<td>ASCII character class foo</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\p{Foo}</span></td>
						<td>Unicode character class Foo</td>
					</tr>
					<tr>
						<td><span class="rx-syntax">\pF</span></td>
						<td>Unicode character class F (one-letter name)</td>
					</tr>
				</table>
				</div>
			</div>
			</div>
			<div class="tab-pane fade" id="about">
			<p>Created by <a href="http://github.com/laumann">Thomas
			Jespersen</a>. Heavily inspired
				by <a href="http://rubular.com">Rubular</a>. The
				source code can be found on <a
				href="http://github.com/laumann/greg">Github</a>
				</p>
				<p><span class="label label-primary">Features to come</span></p>
				<ul>
					<li><span class="label label-danger">important!</span> Display named submatches with their names</li>
					<li>Use package <code>regexp/syntax</code> to provide regex simplification</li>
					<li>Permalink generation under options</li>
					<li>Further regex options should be tick-off-able (aligned to the left of options dropdown). Options: POSIX</li>
					<li>Loading spinner, positioned where? &mdash; Maybe put it to the right of the regex field, and remove the modifiers input (unused)</li>
				</ul>
			</div>
		</div>
</div>
</div>`
