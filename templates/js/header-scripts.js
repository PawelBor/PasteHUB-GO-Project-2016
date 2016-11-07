//Apapted from: https://www.tinymce.com/docs/
tinymce.init({

  selector: 'textarea',
  statusbar: false,
  height: 850,
  width: 810,
  browser_spellcheck: true,
  toolbar: 'save undo redo styleselect bold italic alignleft aligncenter pagebreak alignright bullist numlist outdent indent code forecolor backcolor fontsizeselect',
  fontsize_formats: '8pt 10pt 12pt 14pt 18pt 24pt 36pt 46pt 56pt',
  themes: "inlite",
  plugins: "save code fullpage textcolor pagebreak",
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
