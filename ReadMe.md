# Shopify Backend/Production Engineer Challenge

## Set-up:
I set up in WSL2, Ubuntu on Windows, though since this is essentially running linux, hopefully any unix distributions would work with both options down below. If using Windows without WSL2, option #2 probably has a higher chance of working though not 100% sure that it would. There is a third option, though that involves changing the code to suit the windows environment and is very painful to set-up, in which case I'd suggest using a cloud vm with linux on it. In general however, I'd suggest using option #2 which uses pre-built images and saves a lot of setup and time.

Option #1 - manually build and run:
1. Ensure you have Make installed - gnu.org/software/make/. If not, I will elaborate on what's running.
2. Install docker here -  https://www.docker.com/products/docker-desktop and make sure docker-compose is also installed - https://docs.docker.com/compose/install/
3. Run `make` or `make build` or type out:
    `docker build \
		-t ehuan2/shopify-backend-challenge-backend:latest \
		-f ./src/backend/Dockerfile-golang \
	./src/backend/`
  followed by:
  `docker build \
		-t ehuan2/shopify-backend-challenge-frontend:latest \
		-f ./src/frontend/Dockerfile-frontend \
	./src/frontend/`
4. Run `make run` or `docker-compose up` 

Option #2 - pull previous pre-built :
1. Follow steps 2 to download docker (and 1 if you want simpler life with Make) from option #1.
2. Run `make run` or `docker-compose up` in the command line.

And that's it!

## Breakdown of features:
![image](https://user-images.githubusercontent.com/47696422/148782703-e3cde949-675e-43d7-879a-506f125178c2.png)
So this is a simple CRUD interface, where I said every item had some metadata associated with it, its type and its cost. Then the extra feature I used was the csv file generation. It should be self-explanatory here how to use it... no surprises.

## Design:
I kept my design really simple to showcase best what I know about the web and my command of tools like docker. For this reason, I am going to use docker-compose as the orchestrator between containers of which are there are just four big portions:
1. Actual CRUD - processes requests from the frontend, written in my favourite language, Go.
2. Frontend - provides the html page as the ui (ideally would use some framework, but since it's small enough, there's no need, simple js container).
3. Redis container for storing the data.
4. CSV Container that deals with generation of csv files.

## Reasoning behind my choices
I will explain the following choices:
1. Using Docker + microservice architecture
2. Use of basically no framework and Go instead of something like Apollo for a GraphQL API
3. Use of Redis

These are the reasons behind my choices:
1. In terms of using Docker + microservices, I used this combination because it allows for more flexibility of changing the backend in the future. What I discovered with microservices is that if you want more features, but in different languages because of a certain library, or one team is better at that framework, etc., instead of rewriting your old code and updating it, you can split up your code into smaller chunks that communicate with each other. One fun thing I did before was do a hackathon where the backend was in 4 or so different frameworks. Anyways, it's kinda perfect for this challenge because I was able to split up the csv generation and the crud into two different parts, where if I messed up the csv generation, then the crud would still be untouched by that mess up.

2. I used Golang for a couple reasons. First, I'm just really comfortable with it, at least in comparison to other languages since I used it at my last co-op. Second, Golang is quite fast and allows to scale up in comparison to lots of other perhaps slower languages. In terms of no framework, the truth is that I'm not too too familiar with a lot of the frameworks that Go can use, and using no framework is honestly just a lot clearer on what happens, instead of being obfuscated by lots and lots of set-up and library code.
 
3. So I also decided to use Redis as the backend for a couple reasons. First, if we were going to expand this to include user login, etc., we'd probably use a cache like redis anyways. Second, we're not at a scale where we need something bulky like a postgres container or a sql container. For now, a simple redis container works.

## Challenges + Stuff I can improve on:
One of the biggest challenges I faced was just getting use to web development again. It's been a hot minute since I lasted did anything web dev, so this was a pretty good refresher into it. Had to learn CORS all over again.

On that note, I feel as though there are a couple downsides to writing the http server from scratch, mainly being that the cors policy stuff + routing is a bit trickier and not as smooth as it should be. Usually we can just ignore this part and add it as an add-on, but for me it was something I found myself rewriting, which is not good design. In the future I would probably just still use frameworks to not start from scratch.
