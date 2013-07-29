$(document).ready(function() {
  $("pre").each(function() {
    var pre = $(this);
    var unhighlighted = pre.text();
    var regex = /\:\:\:(\w+)/g;
    var lang = regex.exec(unhighlighted)[1];
    var code = unhighlighted.replace(regex, "")

    $.post("http://glacial-thicket-4470.herokuapp.com", { "lang": lang, "code": code }, function(data) {
      pre.replaceWith($(data))
      pre.replaceWith(document.createTextNode('text < and > not & html!'))
    });
  });
});
