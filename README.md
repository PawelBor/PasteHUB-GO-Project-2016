# *Pastehub - Emerging Technologies Project 2016 - Year 4*
##The Team


**Name:** Pawel Borzym, Edvardas Lasauskas, Niks Gurins, Gediminas Saparauskas </br>
**College:** Galway-Mayo Institute of Technology </br>
**Course:** Software Development - Y4 </br>
**Module:** Emerging Technologies </br>
**Lecturer:** Dr.Ian Mcloughlin </br>

##Project Motivation
The reason we decided to take on this project is because we were looking for a challenge and something new.</br>
Imagine a website, where you can, without downloading any software, write fully fledged documents and retrieve them at any time.</br>
Now imagine on that same website you can have an entire team working on a single document simultaneously.</br>
</br>
Sounds good right? (Try Google Docs) Well, we decided to make an attempt at making that.


##Project Outline
Upon entering the site the user will see a "Generate document" button, which when clicked will generate a random link to a document.</br>
Once inside the document, the connection to the page is upgraded to a websocket.</br>
With the help of the websocket, the user can now share the link to their document, with friends or colleagues, and edit the document together in real-time.</br>
This will, however, mean that anyone joining the link will be able to edit the document. To get around this, we have implemented an option to make a document private. </br>
To make a document private, simply tick the "Private" option before clicking "Generate document" and enter a password for the document accordingly.</br>
Now, when another user joins your private document, they will prompted to enter the password before they get to see it.


##Technologies
*Architecture* | *Technology*
---------|----------
Languages| HTML5, CSS3, JavaScript, GO
Libraries| Bootstrap, jQuery, jsPDF
Frameworks| Macaron, Gorilla
Database| MongoDB


##Design decisions
In terms of the server-side language, we used GO, both because we hadn't had any prior experience with it, and because it was required for our module.</br>
We started the project by taking the websockets/chat exmaple from Gorilla and adapting it for our needs.</br>
The hub.go and client.go have been barely changed at all, but our main.go logic is far from the example.</br>
We use the Macaron framework to handle our http requests across the entire website.</br>
Bootstrap helped us make our application look nice without reinventing the wheel, and jQuery is there to simplify our javascript.</br>
For the database, mongodb seemed like a good decision as it was something new to learn, as well as the fact that our data didn't need any relationships.


##Using the live application
We've focused on making the UI as straight forward as possible to make the web-application easy to use.</br>
On entering the website, the user can either create a public(non-private) document or a private one, for which a password is required.</br>
The creator of the document is then forwarded to their the page where the document resides.</br>
Above the text area, are all the options avaliable to the user for editing the document.</br>
When another user tries to connect to an already created document, the code will check if it is private.</br>
If yes, the user will be prompted to enter the password before showing the document.</br>
A document without a password will skip the above step and take the user straight to the editor.</br>
Users attempting to connect to a non-existant document will get a 404, Not Found error.


##Running the application for yourself
Prerequisites to running the app
* GO installed and a GO workspace made
* Macaron and Gorilla packages installed
```
go get gopkg.in/macaron.v1     and     go get github.com/gorilla/websocket
```
* MongoDB installed and a database called "doceditor" with a collection "documents" inside
* A browser that supports HTML5, CSS3 and Javascript along with Websockets
* Make sure port 8080 is open on your system

To run the actual code, download the entire app, extract and use the 
```
go build
```
command on the extraction directory. Then simply launch the *.exe file created and open localhost:8080 in your browser.


##Features of the application
*  The application will use the database to store links to documents as well as the text in them and their password
*  Whenever a change to document has been made, there are 2 timers that get started. If another change is made, the time waited is reset and wait again
*  One timer is responsible for when to send messages across the websocket, and the other for when to save the document to the database(minimize HTTP request)
*  Websockets allowing users to edit document simultaneously 
*  Creation of private and public documents
*  Document editing (Fonts, font sizes, font colours, alignment of text etc.)
*  Save a written document in PDF format


##References
Reference | Link
---------|----------
Bootstrap | https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-alpha.4/css/bootstrap.min.css 
JQuery | https://ajax.googleapis.com/ajax/libs/jquery/3.0.0/jquery.min.js
Tether | https://github.com/HubSpot/tether
TinyMCE | https://cdn.tinymce.com/4/tinymce.min.js
jsPDF | https://cdnjs.cloudflare.com/ajax/libs/jspdf/1.3.2/jspdf.debug.js
Gorilla | https://github.com/gorilla/
MGO | https://github.com/go-mgo/mgo
Mongo | https://www.mongodb.com/
Macaron | https://github.com/go-macaron/macaron
