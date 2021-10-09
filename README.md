# APPOINTY-GOLANG

### TASK uploaded in master branch

Use POSTMAN to to run the APIs

To run the project:-

run the following command in the terminal of the root directory

go run main.go

The server runs on:-

http://localhost:8000


### To post user POST method on Postman

http://localhost:8000/users


(JSON data)


{
 
 "name" : "username",
  "email" : "abc@gmail.com",
  "password: : "user1234"
  
}

### To get user data GET method on Postman

http://localhost:8000/users/{id}

replace {id} with the user id from mongodb


### To post userposts POST method on Postman

http://localhost:8000/posts

(JSON data)

{

  "id" : "userid"
  "caption" : "your captions"
  "url" : "image url"
  "time" : "time of post"
  
}


### To get post GET method on Postman

http://localhost:8000/posts/{id}

replace the {id} with the post id from mongodb


### To get all posts of a user GET method on Postman

http://localhost:8000/posts/users/{id}

replace the {id} with the user id for which you want the posts




