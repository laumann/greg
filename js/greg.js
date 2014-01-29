/**
 * Greg JS component
 */
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

		var list = '<ul class="list-unstyled">';
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
					//console.log("before:  (" + pos + ", " + i + ") #=> " + input.substring(pos,i));
					//console.log("matched: (" + i + ", " + j + ") #=> " + input.substring(i,j));
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
	},
	example: function() {
		$("#regex").html("a(a*a)b");
		$("#regex-input").html("aaaab\nxabx");
		Greg.update();
	}
};
