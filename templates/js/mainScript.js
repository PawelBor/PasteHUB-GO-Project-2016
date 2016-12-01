//Get the password from Go
		var passwordDB = {{.Password}};
		var conn;
		var uri = window.location.pathname;
		var userInputTimer;
		var saveToDbTimer;

		$(document).ready(function() {
			checkIfPrivate();
		});

		function checkIfPrivate(){
			if(passwordDB === ""){// Document is private

				// Generate the authentication form
				document.getElementById("visibility_wrapper").innerHTML
				='<div class="mce-doc"><textarea id="tiny_content"></textarea></div>'

				//tinyEMC initializer
				initializeMCE();

				//Populate data from mongoDB
				tinyMCE.get('tiny_content').setContent("{{.Text}}", {format : 'raw'});

				//Keep cursor at the bottom of the page after file saves to mongoDB
				tinyMCE.activeEditor.selection.select(tinyMCE.activeEditor.getBody(), true);
				tinyMCE.activeEditor.selection.collapse(false);

				//Get current uri location
				var loc = window.location, new_uri;

				if (loc.protocol === "https:") {
				    new_uri = "wss:";
				} else {
				    new_uri = "ws:";
				}
				new_uri += "//" + loc.host;
				new_uri += loc.pathname + "/ws";

			  	// Checks if browser supports websockets
			    if (window["WebSocket"]) {
			    	var uri = window.location.pathname;
			    	// Upgrade connection for current uri to a websocket
			        conn = new WebSocket(new_uri);
			        console.log(conn);
			        conn.onclose = function (evt) {
			            window.alert("Connection closed.");
			        };
			        conn.onmessage = function (evt) {
			        	// Update content of the editor on message retrieve
			        	console.log(evt);
			        	tinyMCE.get('tiny_content').setContent(evt.data, {format : 'raw'});	
			        	tinyMCE.activeEditor.selection.select(tinyMCE.activeEditor.getBody(), true);
						tinyMCE.activeEditor.selection.collapse(false); 
			        }
			    } else {
			        window.alert("Your browser does not support WebSockets.");
			    }

			}else{// Document is public - Generate the document
				document.getElementById("visibility_wrapper").innerHTML
				= '<div class="centered_box" style="width: 500px;height: 220px;margin:0 auto;margin-top:200px;background-color: white;border-radius: 5px;box-shadow: 5px 5px 15px 5px;text-align: center;padding-left:10px;padding-right:15px;padding-top:50px;"><p><b style="font-size:23px;">This document is password protected.</b><br/>Please enter a password to access this document.</p><br/><input type="text" class="form-control" style="width:300px;float:left;" id="password" placeholder="Password"><button class="btn btn-primary" id="password_btn" onClick="login()" style="width:150px;" >Log In</button></div>';
			}
		}


		function login(){
			if (document.getElementById("password").value === passwordDB){
				// Generate the document
				document.getElementById("visibility_wrapper").innerHTML
			='<div class="mce-doc"><textarea id="tiny_content"></textarea></div>'

			// Init the tinyEMC editor.
			initializeMCE();

			var x = "{{.Text}}";
			console.log(x);
			tinyMCE.get('tiny_content').setContent(x, {format : 'raw'});

			tinyMCE.activeEditor.selection.select(tinyMCE.activeEditor.getBody(), true);
			tinyMCE.activeEditor.selection.collapse(false);


			var loc = window.location, new_uri;
			if (loc.protocol === "https:") {
			    new_uri = "wss:";
			} else {
			    new_uri = "ws:";
			}
			new_uri += "//" + loc.host;
			new_uri += loc.pathname + "/ws";

			  	

		  	// Checks if browser supports websockets
		    if (window["WebSocket"]) {
		    	var uri = window.location.pathname;
		    	// Upgrade connection for current uri to a websocket
		        conn = new WebSocket(new_uri);
		        console.log(conn);
		        conn.onclose = function (evt) {
		            window.alert("Connection closed.");
		        };
		        conn.onmessage = function (evt) {
		        	console.log(evt);
		        	tinyMCE.get('tiny_content').setContent(evt.data, {format : 'raw'});	
		        	tinyMCE.activeEditor.selection.select(tinyMCE.activeEditor.getBody(), true);
					tinyMCE.activeEditor.selection.collapse(false); 
		        }
		    } else {
		        window.alert("Your browser does not support WebSockets.");
		    }
			}else{
				window.alert("Wrong Password");
				return;
			}
		}


		function initializeMCE(){

			//Apapted from: https://www.tinymce.com/docs/
			tinymce.init({
				selector: 'textarea',
				setup: function(ed) {
					ed.on('keyup', function(e) {

					// adapted from http://stackoverflow.com/questions/10406930/how-to-construct-a-websocket-uri-relative-to-the-page-uri

				    clearTimeout(userInputTimer);
					clearTimeout(saveToDbTimer);

					userInputTimer = setTimeout(function() {
						if (!conn) {
							console.log("No ws connection");
							return false;
						}
						if (!(tinyMCE.get('tiny_content').getContent({format : 'raw'}))) {
							console.log("Message is null");
							return false;
						}
						conn.send(tinyMCE.get('tiny_content').getContent({format : 'raw'}));
						return false;
						console.log("keydown in doc");
					}, 500);

					// Save text to database *timer*
			      	saveToDbTimer = setTimeout(function() {
			        	// Save text to database
			        	var text = tinyMCE.get('tiny_content').getContent({format : 'raw'});
			        	$.ajax({
					        type: "PUT",
					        url: uri,
					        async: true,
					        data: {
					          uri : uri,
					          text : text
					        },
					        success: function(response){
					          console.log("Saved to db");
					          tinyMCE.activeEditor.selection.select(tinyMCE.activeEditor.getBody(), true);
				              tinyMCE.activeEditor.selection.collapse(false);
					        },
					        error: function(response){
					          console.log("PUT failed");
					        }
					   	});
			      	}, 5000);
					// Timer adapted from: http://stackoverflow.com/questions/4220126/run-javascript-function-when-user-finishes-typing-instead-of-on-key-up

					});
					},
					statusbar: false,
					height: 850,
					width: 810,
					max_width: 720,
					browser_spellcheck: true,
					toolbar: 'save undo redo styleselect bold italic alignleft aligncenter alignright bullist numlist outdent indent code forecolor fontsizeselect ',
					fontsize_formats: '8pt 10pt 12pt 14pt 18pt 24pt 36pt 46pt 56pt 66pt 76pt 86pt',
					themes: "inlite",
					plugins: "save code fullpage textcolor  ",
					save_enablewhendirty: true,

					save_onsavecallback: function () {
						save_content(); 
					}
	      	});

			//Adapted from: https://github.com/MrRio/jsPDF
			function save_content(){
				var doc = new jsPDF();
		        doc.fromHTML(tinyMCE.get('tiny_content').getContent({format : 'raw'}));
		        doc.save('PasteHub.pdf');
			}

		}// End InitializeMCE