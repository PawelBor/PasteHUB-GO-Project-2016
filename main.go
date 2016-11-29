// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"net/http"
	"math/rand"
	"text/template"
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var addr = flag.String("addr", ":80", "http service address")
var homeTemplate = template.Must(template.ParseFiles("templates/document.html"))

type Document struct{
	Uri string
	Text string
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.Execute(w, r.Host)
}

func generateUri() string{
	chars := "abcdefghijklmnopqrstuvwxyz1234567890"
	uri := ""
	for i := 0; i < 6; i++{
		random := rand.Intn(36)
		uri += string(chars[random])
	}

	return uri
}

func checkMongoForUri(uri string) bool{
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		panic(err)
		
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("doceditor").C("documents")

	// Index
	index := mgo.Index{
		Key:        []string{"Uri"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		panic(err)
		
	}
	
	//var results []Document
	result := Document{}

	//err = c.Find(bson.M{}).All(&results)

	//FAIL HERE
	err = c.Find(bson.M{"uri": uri}).One(&result)
	
	if result.Uri != ""{
		return true
	}
	
	return false
}

func getDocumentData(uri string) string{

	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		panic(err)
		
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("doceditor").C("documents")

	// Index
	index := mgo.Index{
		Key:        []string{"Uri"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		panic(err)
		
	}
	
	//var results []Document
	result := Document{}

	//err = c.Find(bson.M{}).All(&results)

	//FAIL HERE
	err = c.Find(bson.M{"uri": uri}).One(&result)
	
	return result.Text
}

func createDocument(uri string){
	session, err := mgo.Dial("127.0.0.1:27017")

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("doceditor").C("documents")

	err = c.Insert(&Document{Uri: uri, Text: ""})

	if err != nil {
		panic(err)
	}
}

func main() {
	styleHandler := http.FileServer(http.Dir("public/css"))
    http.Handle("/css/", http.StripPrefix("/css/", styleHandler))

	js := http.FileServer(http.Dir("public/js"))
    http.Handle("/js/", http.StripPrefix("/js/", js))

	img := http.FileServer(http.Dir("public/images"))
    http.Handle("/images/", http.StripPrefix("/images/", img))

	m := macaron.Classic()
	m.Use(macaron.Renderer())

	flag.Parse()

	// Read existing document
	m.Get("/:uri", func (ctx *macaron.Context){
		// Check first if the front-end generated uri exists
		exists := checkMongoForUri(ctx.Params(":uri"))

		// If exists, keep generating a new id for document
		if(exists){
			// Load page from DB
			ctx.Data["Text"] = getDocumentData(ctx.Params(":uri"))
			ctx.HTML(200, "document")
		}else {
			ctx.HTML(404, "notFound")
		}
	})

	m.Post("/:uri", func (ctx *macaron.Context){
		// Check first if the front-end generated uri exists
		exists := checkMongoForUri(ctx.Params(":uri"))

		if exists{
			newUri := generateUri()
			for exists{
				// If exists, keep generating a new id for document
				exists = checkMongoForUri(newUri)
				newUri = generateUri()
			}
			// Save it toDB
			createDocument(newUri)
			ctx.Data["uri"] = newUri		
			ctx.HTML(200, "uri")
		} else {
			// Save it toDB
			createDocument(ctx.Params(":uri"))
			ctx.Data["uri"] = ctx.Params(":uri")		
			ctx.HTML(200, "uri")
		}
	})

	m.Get("/:uri/ws", func(w http.ResponseWriter, r *http.Request){
		// start hub in new thread
		// DOESN'T WORK :(
		hub := newHub()
		go hub.run()

		serveWs(hub, w, r)
	})

	m.Run(80)
}