Architecture Document
=====================
![API_Architecture](Architecture.png)

The architecture of the API consists of a server that receives requests. It processes them and returns them to the clients.
- It has 4 basic functions:
   - `/ Login` - Receive the username and password and returns an access token.
   - `/ Logout` - Receives the generated token from the user and erases it.
   - `/ Upload` - Receive an image sent by the client and return the filename and the size.
   - `/ Status` - Returns the users logged into the server.