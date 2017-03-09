# DBS-test


How to build and run the application,

1. downlaod all the files
2. install go
3. go to command prompt, and download imports if not available
4. go get -u gopkg.in/mgo.v2
5. go build main.go
6. main

Restaurent search page
---------------
http://localhost:8080/search/

created web page using go 
	1. Templates
	2. Regx - valid path
	3. Common handler to call http services

MongoDB - database setup
-----------------
I do not have aws or vpc accounts to use cloud, please find local database attached.
https://mlab.com/environments/e-45b07757

download data(2).zip and setuo local mongo database

	1. change direcory to models
	2. go get -u gopkg.in/mgo.v2
	3. go build main
	4. main
	5. should able to see data from mongo

mongodb.png
------------
Shows available restaurants in mongodb

SSL connection
-----------

db.CreateUser({
    user: "mongoUser",
    pwd: "userPassword",
    roles: [
        { role: "readWrite", db: "myDatabase" }
    ]})

security:
    authorization: enabled  


Enhancement TODO list
---------------
	1. MCV implementation
	2. Templating with bootstrap
	3. mlab.com setup for cloud databse

