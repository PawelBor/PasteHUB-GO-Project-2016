// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"net/http"
	"math/rand"
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var addr = flag.String("addr", ":8080", "http service address")
// Map of hubs (uri:* Hub)
var hubs = make(map[string]*Hub)

// Structure for the Mongo document
type Document struct{
	Uri string
	Text string
	Password string
}

// Function to generate a random 6 char long uri
func generateUri() string{
	chars := "abcdefghijklmnopqrstuvwxyz1234567890"
	uri := ""
	for i := 0; i < 6; i++{
		// Generate a number between 0 and 36
		random := rand.Intn(36)
		// Pick out a character at index "random" from chars
		uri += string(chars[random])
	}

	return uri
}

func checkMongoForUri(uri string) bool{
	session, err := mgo.Dial("127.0.0.1:27017")

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("doceditor").C("documents")

	// Index is an options structure for return from db
	index := mgo.Index{
		Key:        []string{"Uri"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
		
	}

	result := Document{}

	err = c.Find(bson.M{"uri": uri}).One(&result)
	
	// If there is any results for the uri, return true, else false
	if result.Uri != ""{
		return true
	}
	
	return false
}

func getDocumentData(uri string) (string,string){
	// Connect to mongo
	session, err := mgo.Dial("127.0.0.1:27017")

	// Keep connection from closing until end of method
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Get a cursor to the "documents" collection in db "doceditor"
	c := session.DB("doceditor").C("documents")

	// Index is an options structure for return from db
	index := mgo.Index{
		Key:        []string{"Uri"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	// Ensures the options are being used on the query
	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
		
	}
	
	// Make a variable to store our results
	result := Document{}

	// Launch query on MongoDB and put stuff in result
	err = c.Find(bson.M{"uri": uri}).One(&result)
	if err != nil {
		panic(err)
		
	}

	// Return the text saved in DB for the specific uri
	return result.Text, result.Password
}

func createDocument(uri string, password string){
	// Connect to mongo
	session, err := mgo.Dial("127.0.0.1:27017")

	// Keep connection from closing until end of method
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Get a cursor to the "documents" collection in db "doceditor"
	c := session.DB("doceditor").C("documents")

	// Insert a new Document struct into the DB with the desired uri
	err = c.Insert(&Document{Uri: uri, Text: "", Password: password})
	// Panic if there is an error in inserting
	if err != nil {
		panic(err)
	}
}

func updateText(uri string, text string){
	session, err := mgo.Dial("127.0.0.1:27017")

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("doceditor").C("documents")

	// Object will be the handle on our db object
	object := bson.M{"uri": uri}
	// Update 
	update := bson.M{"$set": bson.M{"text": text}}

	err = c.Update(object, update)

	if err != nil {
		panic(err)
	}
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	flag.Parse()

	// Put used to replace the document text in the DB
	m.Put("/:uri", func (ctx *macaron.Context, r *http.Request){
		// Get text value from HTTP header
		text := r.FormValue("text")
		updateText(ctx.Params(":uri"), text)
	})

	// Get used to read existing document and load template
	m.Get("/:uri", func (ctx *macaron.Context){
		// Check first if the front-end generated uri exists
		exists := checkMongoForUri(ctx.Params(":uri"))

		// If exists, keep generating a new id for document
		if(exists){
			text, password := getDocumentData(ctx.Params(":uri"))
			// Load page from DB
			ctx.Data["Text"] = text
			ctx.Data["Password"] = password
			ctx.HTML(200, "document")
		}else {
			ctx.HTML(404, "notFound")
		}
	})

	// Post used to create a new document and put it into the DB
	m.Post("/:uri", func (ctx *macaron.Context, r *http.Request){
		// Check first if the front-end generated uri exists
		exists := checkMongoForUri(ctx.Params(":uri"))
		password := r.FormValue("password")

		if exists{
			newUri := generateUri()
			// While the uri is already in the db
			for exists{
				// generate a new id for document and check if exists
				exists = checkMongoForUri(newUri)
				newUri = generateUri()
			}
			// Save it toDB
			createDocument(newUri, password)
			ctx.Data["uri"] = newUri		
			ctx.HTML(200, "uri")
		} else {
			// Save it toDB
			createDocument(ctx.Params(":uri"), password)
			ctx.Data["uri"] = ctx.Params(":uri")		
			ctx.HTML(200, "uri")
		}
	})

	// Get/ws used to serve the websocket for each client and either start or check a hub
	m.Get("/:uri/ws", func(ctx *macaron.Context, w http.ResponseWriter, r *http.Request){

		// Check if hub already in map
		i := hubs[ctx.Params(":uri")]
		if i != nil{
			// if yes, servews to that hub
			serveWs(i, w, r)
		}else{
			// else make a new hub, put it in the map and start the goroutine
			hub := newHub()
			hubs[ctx.Params(":uri")] = hub
			go hub.run()

			serveWs(hub, w, r)
		}
	})

	// Start running the server on port :80
	m.Run(8080)
}

