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
	<!-- http://bootswatch.com/slate/bootstrap.min.css -->
	<link rel="stylesheet"	href="http://bootswatch.com/slate/bootstrap.min.css">

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
		color: #fff;
	}
	span.match {
		background-color: #f89406;
	}
	span.rx-syntax {
		font-weight: bold;
		font-family: monospace;
	}
	.table td {
		border-top: none;
	}
	</style>
</head>
<body>
<div class="container">
`

const footer =
`<script type="text/javascript">
$(function() {
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
<div class="well">
<form role="form" id="greg">
<div class="row">
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
	<div class="col-xs-12 text-right">
		<a href="#" onclick="Greg.clear(); return false;" class="btn btn-default btn-sm" role="button">Clear fields</a>
	</div>

</div>
</form>
</div>
<div class="row">
	<div class="col-xs-12">
		<div class="panel panel-info">
			<div class="panel-heading">Quick Reference</div>
			<div class="panel-body">
				<div class="row">
					<div class="col-xs-4">
						<table class="table table-condensed">
							<tr>
								<td><span class="rx-syntax">.</span></td>
								<td>any character, possibly including newline (flag s=true)</td>
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
								<td><span class="rx-syntax">[:^alpha]</span></td>
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
						<table class="table table-condensed">
							<tr>
								<td><span class="rx-syntax">xy</span></td>
								<td>x followed by y</td>
							</tr>
							<tr>
								<td><span class="rx-syntax">x|y</span></td>
								<td>x or y (prefer x)</td>
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
					<div class="col-xs-4">Column 3</div>
				</div>
			</div>
		</div>
	</div>
</div>
<div class="row">
	<div class="col-xs-12">
		<span id="spin" class="glyphicon glyphicon-record pull-right hidden"></span>
	</div>
</div>`
