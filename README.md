# How to build your first rest api with Go

### Why Choose Go For Building APIs?

* Go is better option for raw performance and computation. 
* It is a fast, lightweight platform. 
* Go offers great performance out of the box, it is easy to deploy, and has many of the necessary tools you need to build and deploy a scalable web service in its standard library. 
* Go can handle multiple tasks concurrently using goroutines and channels.
* This makes it well-suited for building high-performance, scalable APIs that can manage many requests together.

With this description I believe every programmer likes to read this article 🐎

### Scope of Story
This tutorial will walk you through a practical example of building REST APIs in Go. As an HTTP framework, 
I personally like to use [Gin Gonic](https://gin-gonic.com/), since it is straightforward and powerful.

### Folder Structure

I preferred to hold folder structure simple to understand. For big projects we need more divation. Versions of apis can be in different folders and maybe we will need more directories for example separating logic from the controller as service files.
```
- api
  - controller
    - main_test.go
    - main.go
  - view
    - main_test.go
    - main.go
- .gitignore
- go.mod
- go.sum
- main.go
```
<code>api/</code> contains one or more versions of the REST API  
<code>controller/</code> contains the service logic for the models <code>Note</code> For big project we can move logic to the service files  
<code>view/</code> expose services via REST API  
<code>main.go/</code> where everything is going to start

### Design API Endpoints
We’ll build an API that provides access to a file on  a specific path. 

When developing an API, you typically begin by designing the endpoints. Your API’s users will have more success if the endpoints are easy to understand.

Here are the endpoints we’ll create in this tutorial.

GET – Get a file which are placed into a specific path by giving file name.  
POST – Upload a specific file into the determined directory.

POST /api/v1/files/  
GET /api/v1/files/:name/

### REST Endpoints using Gin Gonic
After designing core logic we need to expose it via REST API. Therefore, I mostly create an api folder that contains versioned APIs, e.g. in this simple case v1 inside it, in more complex APIs it might also be clever to create dedicated folders for each version.

To specify the http method on your final endpoint use <code>http</code> method name in upper case letters <code>(eg. POST, GET)</code>. The first argument of these functions are the endpoint address and the second one is a function to handle API calls.

<code>Gin</code> puts requests in <code>*gin.Context</code> type, so you can access request headers, files etc from a *gin.Context variable. You should pass this type as the argument to the second argument of <code>http</code> methods.

To run your server on a custom port use the Run method with <code>“:PORT_NUM”</code>. For example to run on port <code>8095</code> use <code>router.Run(“:8095”)</code>

Now <code>main.go</code> file in view’s directory now should look like this:

```
package view

import (
   "github.com/gin-gonic/gin"
)

func StartServer() {
   // Create a router using Gin framework
   router := gin.Default()
   
   api := router.Group("/api")
   // Grouping a bunch of same endpoints
   
   v1 := api.Group("/v1")
   files := v1.Group("/files")
   files.POST("/", func(c *gin.Context) {
     // Controller code's goes here
})
   router.Run(":8080")
}
```

We can group a bunch of same endpoints using Gin’s Group function on our router. In this project we’re going to have two endpoints, one for Post and another for Get.

#### How Grouping Works?
We group our router on <code>/api</code> endpoint. Because all of endpoints starts with <code>/api</code>, then group <code>/api</code> with <code>/v1</code> because (of course) this is the first version of this API and finally group <code>/v1</code> on <code>/files</code>. Regarding above description final result would be like below:

```
POST /api/v1/files/ 
GET /api/v1/files/:name/
```

### Gins Functions

* <code>FormFile</code> Gin’s FormFile function retrieves a single file from the request
* <code>H</code>  If any errors occurs returns a JSON response using Gin’s JSON function alongside with Gin’s H type  

Now the body of <code>POST</code> function should look:

``
file, err := c.FormFile("file")
if err != nil {
   c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
``
* <code>ShouldBindUri</code> bind url variables
We’re going to use this type with Gin’s <code>ShouldBindUri</code> to read a variable from <code>uri</code>.

```
type File struct {
   Name string `uri:"name" binding:"required"`
}
```

Add the endpoint to download a file using it’s name.

```
files.GET("/:name/", func(c *gin.Context) {
   var file File
   // Bind url variables
   if err := c.ShouldBindUri(&file); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"msg": err})
      return
   }
})
```
After binding we have to pass it to the controller’s Download function to handle downloading that file. Finally we can serve the file using Gin’s Data function.

```
var f File
// Bind url variables
if err := c.ShouldBindUri(&f); err != nil {
   c.JSON(http.StatusBadRequest, gin.H{"error": err})
   return
}
m, cn, err := controller.Download(f.Name)
if err != nil {
   c.JSON(http.StatusNotFound, gin.H{"error": err})
   return
}
c.Header("Content-Disposition", "attachment; filename="+n)
// Serve the file
c.Data(http.StatusOK, m, cn)
```

### Build project
In the root folder of your module run <code>go build -o ./rest-api main.go</code> command to build your module into one file. Now you can run it using <code>./rest-api</code> . Alternatively you can run your project without building it using <code>go run main.go</code> .

Gin’s default environment is debug mode. You can change it to release mode before calling <code>view.StartServer()</code> function using <code>gin.SetMode(gin.ReleaseMode)</code>. if you’re using debug mode you should see an output similar to this that means your server is running and waiting for you call it’s APIs.


#### Please find quick reference for GO commands from [Here](https://github.com/hakimehmordadi/GOlang-Cheat-Sheet)










