/**
 * Greg JS component
 */
var Greg = {
	showMatches: function(json) {
		if (json.error) {
			$("#regex-match").addClass("hidden");
			$("#error-msg").html(json.error);
			$("#regex-fail").removeClass("hidden");
			return;
		}
		console.log(json);
		var result = $("#match-result");
		var groups = $("#match-groups");

		var list = '<ul class="list-unstyled">';
		var mcnt = 0;
		var subm = '';
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
					list += input.substring(pos, i); // before match
					list += '<span class="match">' + input.substring(i, j) + '</span>'; // match
					pos = j; // move pos up
					if (m.length > 2) {
						for (var k = 2; k < m.length; k += 2) {
							var i = m[k], j = m[k+1];
							subm += '<tr><td>'
							if (json.names[k/2]) {
								subm += '<span class="label label-primary">';
								subm += json.names[k/2];
								subm += '</span>';
							} else {
								subm += (++mcnt);
							}
							subm += '</td><td>'
							subm += input.substring(i, j);
							subm += '</td></tr>\n';
						}
					}
				});
				list += input.substring(pos, input.length);
			}
			list += '</li>';
		}
		list += '</ul>';
		result.html(list);
		if (subm.length > 0) {
			groups.html('<table>' + subm + '</table>');
		} else {
			groups.html("");
		}
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
		if ( $("#regex").val() == "" || $("#regex-input").val() == "")
			return;
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
				$("#regex-match").addClass("hidden");
			},
			dataType: "json"
		});
	},
	example: function() {
		var d = new Date;
		$("#regex").val('(?P<day>\\d{1,2})/(?P<month>\\d{1,2})/(?P<year>\\d{4})');
		$("#regex-input").val("Today's date is: " + d.getDate() + "/" + (d.getMonth() + 1) + "/" + d.getFullYear() + ".");
		Greg.update();
	},
	clear: function() {
		$("#regex").val("");
		$("#regex-input").val("");
		$("#regex-match").addClass("hidden");
		$("#regex-fail").addClass("hidden");
		$("#greeting").removeClass("hidden");
	},
};
