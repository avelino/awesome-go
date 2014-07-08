var getFileData = function(url) {
  $.ajax({
    url: url,
    headers: {
      "Accept": "application/vnd.github.v3.raw"
    }
  }).done(function(data) {
    var converter = new Showdown.converter();
    var html = converter.makeHtml(data);
    $("#content").html(html)
  })
};

$(document).ready(function(){
  $.getJSON("https://api.github.com/repos/avelino/awesome-go/git/trees/HEAD").
    done(function(data){
      for (var i = data.tree.length - 1; i >= 0; i--) {
        if(data.tree[i].path == "README.md") {
          getFileData(data.tree[i].url);
          break;
        }
      };
    });
});
