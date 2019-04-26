# catcher
https://fortytwosearch.herokuapp.com/
Creation of an instant search engine built with Go

This is the beggining of my personnal "instant search engine".
The goal is not to have the most powerfull search engine but to acquire some knowledges

# Building :
 I have choose to use a trie, https://en.wikipedia.org/wiki/Trie, then i index all of my data contain in /backend/to_index.json and i fill my trie based on login and firstname.
 To do the research i go search in specific node and i compare the login in the node to the request and if the distance between this two words is inferior to specific value i return the node.
 To calculate the distance i use : https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance
 
 # Install
 `git clone`
 
 `cd frontend/`
 
 `npm install`
 
 `npm start`
 
 `cd ../backend`
 
 `go run main.go`
 
 
 You need to have minimum the version 1.11 of Go because i use gomodules.
