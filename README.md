# Instagram API using GO

Design and Develop an HTTP JSON API capable of the following operations,<br>
- Create an User
    - Should be a POST request
    - Use JSON request body
    - URL should be ‘/users'
- Get a user using id
    - Should be a GET request
    - Id should be in the url parameter
    - URL should be ‘/users/<id here>’
- Create a Post
    - Should be a POST request
    - Use JSON request body
    - URL should be ‘/posts'
- Get a post using id
    - Should be a GET request
    - Id should be in the url parameter
    - URL should be ‘/posts/<id here>’
- List all posts of a user
    - Should be a GET request
    - URL should be ‘/posts/users/<Id here>'


### JSON API created

#### controllers
- `getConnection` handles connecting to the DB
- `createUser` creats the user and pushing the user to DB
- `handleCreateUser` handles the creating of the user and showing the Id as res
- `createPost` creates the Post and push to the DB
- `handleCreatePost` handles the createPost and shows the Id as res
- `getUser` gets the user id and name from the DB and takes the `:userId` as param
- `getPost` gets the post complete information saved in the DB and takes `:postId` as param
- `listPost` lists all the posts created by the user and takes `:userId` as param

#### Two models

- Users
- Posts

#### MongoDB 

- instagram
    - posts
    - users

Testing the API endpoints and there result
![test report](/resource/test.png)

On Postman Testing results
1) ![getUser](/resource/getUser.png)
2) ![postPost](/resource/postPosts.png)
3) ![createUser](/resource/postUser.png)
4) ![listUserPosts](/resource/listuserposts.png)

Pagination has been used in the `listPost` and gives 2 posts for each page<br>
Passwords are hashed before it is saved in the DB (using `bcrypt`)<br>