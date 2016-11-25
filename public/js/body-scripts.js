// Websocket code adapted from github.com/gorilla/websocket/examples/chat
      var conn;
      var doc;
      window.onload = function () {
        doc = document.getElementById("tiny_content");

        if (window["WebSocket"]) {
            conn = new WebSocket("ws://{{$}}/ws");
            conn.onclose = function (evt) {
                window.alert("Connection closed.");
            };
            conn.onmessage = function (evt) {
              var doScroll = doc.scrollTop === doc.scrollHeight - doc.clientHeight;
              if (doScroll) {
                  doc.scrollTop = doc.scrollHeight - doc.clientHeight;
              }
              doc.value = evt.data;
            }
        } else {
            window.alert("Your browser does not support WebSockets.");
        }
      };

      // Timer adapted from: http://stackoverflow.com/questions/4220126/run-javascript-function-when-user-finishes-typing-instead-of-on-key-up
      var userInputTimer;

      $("#document")
        .on('input',function(e){
          // From: https://medium.freecodecamp.com/how-to-reverse-a-string-in-javascript-in-3-different-ways-75e4763c68cb#.vdkzw58fq
          // $("#output").val($("#userinput").val().split("").reverse().join(""))

          clearTimeout(userInputTimer);
          userInputTimer = setTimeout(function() {
            if (!conn) {
              console.log("No ws connection");
              return false;
            }
            if (!doc.value) {
              console.log("Message is null");
              return false;
            }
            conn.send(doc.value);
            return false;
            console.log("keydown in doc");
          }, 3000);
        });