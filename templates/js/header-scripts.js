// Websocket code adapted from github.com/gorilla/websocket/examples/chat
var conn;
var msg;
var doc;
var userInputTimer;

window.onload = function () {
    doc = document.getElementById("document");

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://{{$}}/ws");
        conn.onclose = function (evt) {
            window.alert("Connection closed.");
        };
        conn.onmessage = function (evt) {
          /*
          var doScroll = doc.scrollTop === doc.scrollHeight - doc.clientHeight;
          if (doScroll) {
              doc.scrollTop = doc.scrollHeight - doc.clientHeight;
          }*/
          
          // Sets the HTML contents of the activeEditor editor
          tinyMCE.activeEditor.setContent(evt.data);
        }
    } else {
        window.alert("Your browser does not support WebSockets.");
    }
  };

function timerUpdate(){

// Timer adapted from: http://stackoverflow.com/questions/4220126/run-javascript-function-when-user-finishes-typing-instead-of-on-key-up

clearTimeout(userInputTimer);

userInputTimer = setTimeout(function() {
  if (!conn) {
    console.log("No ws connection");
    return false;
  }
  if (!(tinyMCE.get('tiny_content').getContent())) {
    console.log("Message is null");
    return false;
  }
  conn.send(tinyMCE.get('tiny_content').getContent());
  return false;
  console.log("keydown in doc");
}, 1000);

}

//Apapted from: https://www.tinymce.com/docs/
tinymce.init({
  selector: 'textarea',
    setup: function(ed) {
      ed.on('keyup', function(e) {
          //console.log('Editor contents was modified. Contents: ' + ed.getContent());
          timerUpdate();
      });
  },
  statusbar: false,
  height: 850,
  width: 810,
  max_width: 720,
  browser_spellcheck: true,
  toolbar: 'save undo redo styleselect bold italic alignleft aligncenter alignright bullist numlist outdent indent code forecolor fontsizeselect pagebreak',
  fontsize_formats: '8pt 10pt 12pt 14pt 18pt 24pt 36pt 46pt 56pt',
  themes: "inlite",
  plugins: "save code fullpage textcolor pagebreak ",
  save_enablewhendirty: true,

  save_onsavecallback: function () {
  	save_content(); 
  }


});

//Adapted from: http://archive.tinymce.com/wiki.php/API3:method.tinymce.Editor.getContent
// & https://github.com/eligrey/FileSaver.js/
function save_content(){
	// Get content of a specific editor:
	var editor_text = tinyMCE.get('tiny_content').getContent();

	var blob = new Blob([editor_text], {type: "text/plain;charset=utf-8"});
	saveAs(blob, "PasteHub.txt");
}
