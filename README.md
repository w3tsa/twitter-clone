### Go curd API

- Docker API
- [ ] Get to the PSQL database from within the Docker image
  - [ ] run: `docker ps -a` -> should give the name/Id of the container that has postgress option
  - [ ] copy the container id and run: `docker exec -it <container id> bash`
  - [ ] From their run: `psql -U <dbusername> <dbname>`
  - [ ] should open the db and we can run our db query here
