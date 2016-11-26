// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"fmt"
	"net/http"
	"text/template"
	"math/rand"
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var addr = flag.String("addr", ":8080", "http service address")
var homeTemplate = template.Must(template.ParseFiles("public/index.html"))

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
		uri += chars[random:len(chars)]
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
	results := Document{}

	//err = c.Find(bson.M{}).All(&results)

	//FAIL HERE
	err = c.Find(bson.M{"Uri": uri}).One(&results)
	
	fmt.Println(uri)
	fmt.Println(results)
	if err != nil {
		fmt.Println("YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY")
		panic(err)
		
	}
	
	return false
}

func getDocumentData(uri string) string{

	session, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("doceditor").C("documents")

	index := mgo.Index{
		Key:        []string{"Uri"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)

	results := Document{}
	//var results []Document

	//err = c.Find(bson.M{}).All(&results)
	err = c.Find(bson.M{"Uri": uri}).One(&results)
	
	
	fmt.Println(results)

	if err != nil {
		panic(err)
	}

	var text string

	text = results.Text	

	return text
}

func createHub(uri string){
	flag.Parse()
	hub := newHub()
	// start hub in new thread
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
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

	m.Run(8080)
}