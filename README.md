#React & Go Simple To-Do List App

### Features

- Todo item can be added
- Todo list can be listed

###Used Technologies
In summary, **react** for the frontend, **go** for backend, and **docker** for deployment, and PostgreSQL for the database was used.

### How to Run?
>Docker must be installed and running already

Run on command line `docker-compose up` in the root project folder
Backend side will work on 5000 port you can access on your browser
> sample: localhost: 5000/api/todo

Frontend side will work on 3000 port you can access it too
> sample: localhost: 3000

### How to Run Tests?
For backend side go to `./backend/test` run `go test` on the command line
Frontend side tests will be added

####Backend Packages
You can see all mods that I used in the backend.

	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/gofiber/fiber/v2 v2.23.0
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.4.0
	github.com/lib/pq v1.10.4
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.17.0

####Frontend
Actually there is no additional package. But it should add some test in the future.
