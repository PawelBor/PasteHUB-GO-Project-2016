# PasteHub
Year 4 Project - Emerging Technologies


# *Emerging Technologies Project 2016 - Year 4*
## PasteHub


**Name:** Pawel Borzym, Edvardas Lasauskas, Niks Gurins, Gediminas Saparauskas </br>
**College:** Galway-Mayo Institute of Technology </br>
**Course:** Software Development - Y4 </br>
**Module:** Emerging Technologies </br>
**Lecturer:** Dr.Ian Mcloughlin </br>

#Project Overview

For our Go project we have decided to do a live document typing application.</br> 
One user will need to login and will then be able to type into a blank document. This user will have a random generated link for sharing.</br>
Other users with that link will be able to join that room/session and watch as the main user types into the document.</br>
A key will be used to identify which user is currently able to type into the document. The user holding the key will be able to pass it to another user in the room/session. The key is used to control **concurrency**.


##TECHNOLOGIES

*Architecture* | *Technology*
---------|----------
FRONT-END| HTML, CSS, JavaScript, Bootstrap
SERVER| GO-Lang + Macaron Framework
DATABASE| CouchDB


####**Design of the application**
*  The application will use the DB for storing user login details.
*  Ability to save to DB and store the data there and then read off it (Too many requests to keep it being updated constantly). 
*  User can type and if stopped typing, the data will be sent to DB after short interval and saved + read. If the user will type for too long, other members in room/session will see blank screen until the editor stops typing.
*  Text stored in String format on the server-side and display it back to the front-end(user). 
*  Possiblility of storing the document on the database or saving it locally.



####**Functionality of the application**
*  Register/Login functionality
*  Be able to export documents after finished typing to some format.
*  Pass around some type of key between users to determine who is the editor. 
*  Have guest accounts be able to only watch the room/session.
