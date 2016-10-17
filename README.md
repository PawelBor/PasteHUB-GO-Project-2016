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
