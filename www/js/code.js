// Generated by CoffeeScript 1.7.1
(function() {
  var appendError, lclock, listDocs, main, me, seenClock, showing;

  me = "";

  showing = "";

  lclock = 0;

  seenClock = function(c) {
    if (c > lclock) {
      lclock = c;
      console.log("lclock=" + lclock);
    }
  };

  listDocs = function(data) {
    var doc, docs, li, ret, ul, _i, _len, _ref; 
   ret = JSON.parse(data);
    if (ret.Err !== "") {
      appendError(ret.Err);
      return;
    }
    docs = $("div#docs");
    docs.empty();
    if (ret.Docs === null || ret.Docs.length === 0) {
      docs.append("No Documents");
      return;
    }
    ul = $('<div class="list-group-justified"/>');
    ret.Docs.reverse();
    _ref = ret.Docs;
    for (_i = 0, _len = _ref.length; _i < _len; _i++) {
      doc = _ref[_i];
      li = $('<a href="#" class="list-group-item"/>');
      li.append('<span class="glyphicon glyphicon-file"></span>' + doc + '<span class="glyphicon glyphicon-chevron-right"></span> <span class="badge badge-primary">14</span>');
      li.find("a.author").click(function(ev) {
        var name;
        ev.preventDefault();
        name = $(this).text();
        if (name.length > 0 && name.indexOf('@') === 0) {
          name = name.substring(1);
        }
        return _showUser(name);
      });
      li.hover((function(ev) {
        if (me !== "") {
          $(this).find("a.retrib").show();
        }
      }), (function(ev) {
        $(this).find("a.retrib").hide();
      }));
      ul.append(li);
    }
    docs.append(ul);
  };

  appendError = function(e) {
    $("div#errors").show();
    return $("div#errors").append('<div class="error">Error: ' + e + '</div>');
  };

  showHome = function(ev) {
    ev.preventDefault();
    _showHome();
  };

  _showHome = function() {
    $.ajax({
      url: "api/list-docs",
      type: "POST",
      success: listDocs,
      cache: false
    });
  };

  main = function() {
    _showHome()
    $("div#errors").hide();
    listDocs();
  };

  $(document).ready(main);

}).call(this);
