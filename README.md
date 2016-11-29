# *Pastehub - Emerging Technologies Project 2016 - Year 4*
##The team


**Name:** Pawel Borzym, Edvardas Lasauskas, Niks Gurins, Gediminas Saparauskas </br>
**College:** Galway-Mayo Institute of Technology </br>
**Course:** Software Development - Y4 </br>
**Module:** Emerging Technologies </br>
**Lecturer:** Dr.Ian Mcloughlin </br>

#Project Motivation

The reason we decided to take on this project is because there is nothing like it out there.</br>
Imagine a website, where you can, without downloading any software, write fully fledged document and retrieve it at any time.</br>
Now imagine on that same website you can have an entire team working on a single document simultaneously.</br>
</br>
Sounds good right? Well, we decided to make an attempt at making it a reality.


#Project Outline

Upon entering the site the user will see a "Generate document" button, which when clicked will generate a random link to a document.</br>
Once inside the document, the connection to the page is upgraded to a websocket.</br>
With the help of the websocket, the user can now share the link to their document, with friends or colleagues, and edit the document together in real-time.</br>
This will, however, mean that anyone joining the link will be able to edit the document. To get around this, we have implemented an option to make a document private. </br>
To make a document private, simply tick the "Private" option before clicking "Generate document" and enter a password for the document accordingly.


#Technologies

*Architecture* | *Technology*
---------|----------
Languages| HTML5, CSS3, JavaScript, GO
Libraries| Bootstrap, jQuery
Frameworks| Macaron, Gorilla
Database| MongoDB


####**Features of the application**
*  The application will use the database to store links to documents as well as the text in them
*  Whenever a change to document has been made, there are 2 timers that get started. If another change is made, the time waited is reset and wait again
*  One timer is responsible for when to send messages across the websocket, and the other for when to save the document to the database(minimize HTTP request)
*  Ability to save to DB and store the data there and then read off it (Too many requests to keep it being updated constantly). 
*  User can type and if stopped typing, the data will be sent to DB after short interval and saved + read. If the user will type for too long, other members in room/session will see blank screen until the editor stops typing.
*  Text stored in String format on the server-side and display it back to the front-end(user). 
*  Possiblility of storing the document on the database or saving it locally.



####**Functionality of the application**
*  Register/Login functionality
*  Be able to export documents after finished typing to some format.
*  Pass around some type of key between users to determine who is the editor. 
*  Have guest accounts be able to only watch the room/session.
*  Import .txt files for editing online.

####**Design decisions**