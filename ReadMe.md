# Shopify Backend/Production Engineer Challenge

## Set-up:
I set up in WSL2, Ubuntu on Windows, though since this is essentially running linux, hopefully any unix distribution would work with all three options down below. Otherwise, use option #2 would probably work best for Windows. 

Option #1:
1. Ensure you have Make installed - gnu.org/software/make/. If not, I can list the instructions to type out.
2. Install docker

Option #2:

Option #3:
1. Install docker
2. Pull the images down from docker hub - 
3. Run `docker-compose up` in the command line.

## Extra Feature:


## Design:
I kept my design really simple to showcase best what I know about the web and my command of tools like docker. For this reason, I am going to use docker-compose as the orchestrator between containers of which are there are just three big portions:
1. Actual CRUD - processes requests from the frontend, written in my favourite language, Go.
2. Frontend - provides the html page as the ui (ideally would use some framework, but since it's small enough, there's no need, simple go container again).
3. Redis container for storing the data.

Another design decision I made that is somewhat unnecessary, but is good for long-term is decoupling the database operations from the http handlers.

## Reasoning behind my choices
In terms of using Docker + microservices, I used this combination because it allows for more flexibility of changing the backend in the future. What I discovered with microservices is that if you want more features, but in different languages because of a certain library, or one team is better at that framework, etc., instead of rewriting your old code and updating it, you can split up your code into smaller chunks that communicate with each other. One fun thing I did before was do a hackathon where the backend was in 4 or so different frameworks.

While this expansion of different frameworks is not ideal in an actual backend environment, it does favour these smaller projects where it can go through a lot of changes before it gets too big to do more changes. Also, Docker makes sure of a common runtime amongst computers.

So I also decided to use Redis as the backend for a couple reasons. First, if we were going to expand this to include user login, etc., we'd probably use a cache like redis anyways. Second, we're not at a scale where we need something bulky like a postgres container or a sql container. For now, a simple redis container works.

I used Golang for a couple reasons. First, I'm just really comfortable with it, at least in comparison to other languages since I used it at my last co-op. Second, Golang is quite fast and allows to scale up in comparison to lots of other perhaps slower languages.

## Tools used:
Docker/Docker-compose
Golang
Redis
