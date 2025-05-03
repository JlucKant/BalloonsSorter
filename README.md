### Welcome to Balloons Sorter project.

This is a repository with a test task from YADRO.

 #### How to launch project:
 ##### Standard launch
1. Clone repository to your local directory:
```
git clone https://github.com/JlucKant/BalloonsSorter.git
```
2. Check if Golang `>= 1.23.4` is installed
3. Go to the directory with the copied repository and run the `main.go` file:
```
go run app/main.go
```
##### Launch in a Docker container
1. Clone repository to your local directory:
```
git clone https://github.com/JlucKant/BalloonsSorter.git
```
2. Start Docker daemon
3. Go to the directory with the copied repository
4. Run the `docker-compose` file:
```
docker-compose run --rm balloons-sorter
```
#### To run autotests:
1. Go to the directory with the copied repository and run:
```
go test ./...
```