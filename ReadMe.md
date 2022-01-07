# Shopify Backend/Production Engineer Challenge

## Design:
I kept my design really simple to showcase best what I know about the web and my command of tools like docker. For this reason, I am going to use docker-compose as the orchestrator between containers of which are there are just three big portions:
1. Actual CRUD - processes requests from the frontend, written in my favourite language, Go.
2. Frontend - provides the html page as the ui (ideally would use some framework, but since it's small enough, there's no need, simple go container again).
3. Postgres container for storing the data.

## Reasoning behind my choices
In terms of using Docker + microservices, I used this combination because it allows for more flexibility of changing the backend in the future. What I discovered with microservices is that if you want more features, but in different languages because of a certain library, or one team is better at that framework, etc., instead of rewriting your old code and updating it, you can split up your code into smaller chunks that communicate with each other. One fun thing I did before was do a hackathon where the backend was in 4 or so different frameworks.

While this expansion of different frameworks is not ideal in an actual backend environment, it does favour these smaller projects where it can go through a lot of changes before it gets too big to do more changes. Also, Docker makes sure of a common runtime amongst computers.
