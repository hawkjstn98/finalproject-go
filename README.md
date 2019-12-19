# finalproject-go

This is a university project for the Mobile Cross Platform course in Multimedia Nusantara University.


## Pre-requisite
+ Installation of **latest** [golang](https://golang.org/)
+ Installation of **latest** [MongoDB](https://www.mongodb.com/) (if the project runs locally)
+ IDEs, either [Visual Studio Code](https://code.visualstudio.com/) or [GoLand](https://www.jetbrains.com/go/)
+ [Heroku](https://devcenter.heroku.com/articles/heroku-cli) for deployment (make sure to login to heroku using `heroku login` in a terminal)

## Steps to use this repository
1. Clone this repository from the master branch
2. Open the folder with command line or your favorite IDE or the ones mentioned above.
3. Enter `go run main/main.go` to run the program
4. Hit `localhost:1323` or other port according to your ***environtment*** with [Postman](https://www.getpostman.com/)

## Steps to make a rest-api, like this repository, and the deployment process
***ATTENTION, THIS STEPS MEANT TO BE USED TO MAKE A NEW PROJECT AND NOT HOW TO USE AND DEPLOY THIS REPOSITORY TO YOUR OWN HEROKU AND MONGODB SERVER***
+ If Visual Studio is the IDE choosen for this project then you might want to read [this](https://rominirani.com/setup-go-development-environment-with-visual-studio-code-7ea5d643a51a)
1. Make a new repository in a new folder using `git init PROJECT-NAME` and push the repository to your github
2. Run `go mod init github.com/USERNAME/YOUR-REPO` to initiate go mod that used by golang to track depedency (necessary for deployment)
3. Run `go mod tidy` to tidy up the repository, it will pull all necessary depedency automatically (in theory)
4. Then make a file named `main.go` with a package named `main`
5. Get echo using `go get -u github.com/labstack/echo/...` to get the repo that we need to make a RESTFUL-API (docs can be found [here](https://echo.labstack.com/guide))
6. Get [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) to connect and interact with MongoDB database locally or hosted on the internet (docs on this can be found in [here](https://godoc.org/go.mongodb.org/mongo-driver/mongo) or [here](https://docs.mongodb.com/ecosystem/drivers/go/))
7. To deploy the repository to heroku you need to first run `heroku create` on the repository that you want to deploy
8. Check if your code pull the `PORT` needed to host the server from heroku's enviroment (how to [here](https://gobyexample.com/environment-variables), or check the code in this repository on how to do it)
8. Make sure all the commits have been pushed to master branch in github, then run `git push heroku master` (this will push the master branch from github to heroku and automatically deployed)
9. Run `heroku logs --tail` (preferably in a new terminal or terminal tab) to monitor the ongoing deployment process by heroku and to monitor other logs that have been made by you
10. Log in to [MongoDB Cloud](https://cloud.mongodb.com/) and create a new cluster
11. Go to Collections tab and ***create*** a new database on the newly created cluster with database named according to your code or reversed (referenced to step 6, best practice on this is to put the database name in an ENV or a constant variable)
12. Create new colections with names that you want to have as if it was a mysql table by clicking on the database and pressed the green create button (Optional mongo can create the collection itself by adding the data via code)
13. Setup the `Database Access` for database user and `Network Access` to whitelist certain IP for connections
...The links should be something like `cloud.mongodb.com/v2/...#security/database/users` and `cloud.mongodb.com/v2/...#security/network/whitelist`
14. Connect the newly hosted MongoDB by choosing to "Connect to you app" option in MongoDB Cloud then copy the provided string to your codes (docs on how to connect using go can be found [here](https://docs.atlas.mongodb.com/driver-connection/), and once again the best practice is to put the string on an ENV or a constant variable)
