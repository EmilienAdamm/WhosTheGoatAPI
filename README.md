
# WhosTheGoatAPI
## Content
This is an exemple of an API made in [GoLang](https://go.dev/). This API is part of a learning project: [Who's the GOAT ?](https://github.com/EmilienAdamm/WhosTheGoat) and is the API used by the web application to provide the players.
The goal of the application was to learn the GoLang API, and create a containerizable application to use in a Kubernetes cluster for instance.
## Link
The game is accessible here: https://goatest.bet
## Installation and running
>Note: Make sure you fill in the config.ini file with the information of your machines to run the API.
### RUNNING AS A SERVICE
 1. **Installing dependencies**
First of all, make sure git and GoLang are installed on your machine:
```
$ sudo apt update
$ sudo apt install git golang-go
```
2. **Fetching source code and building the app**
Now that the right packages are installed, you need to build the source code:
``` 
$ git clone https://github.com/EmilienAdamm/WhosTheGoatAPI.git
$ cd WhosTheGoatAPI
$ go build -o <app_name>
```
3. **Creating a sysmtemd service**
To create the service of the API, first create a file in the following directory like such:
``$ sudo nano /etc/systemd/system/app_name.service``
In this file, enter the following information:
```
[Unit] 
Description=WhosTheGoatAPI service
[Service] 
ExecStart=/path/to/your/app
Restart=always 
User=<User to run the service> 
Group=<Group to run the service>
Environment=PATH=/pathToGoBin:$PATH 
WorkingDirectory=/pathToApp
[Install] 
WantedBy=multi-user.target
```
4. **Activate and run the service**
```
$ sudo systemctl daemon-reload
$ sudo systemctl enable app_name
$ sudo systemctl start app_name
``` 
### RUNNING WITH DOCKER
*Will be implemented later*
