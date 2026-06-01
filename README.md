# Documentation for course_scheduler_backend
`course_scheduler_backend` is a RESTful API designed to work with a database of college course and faculty data. Using Go's HTTP server, it handles all the functionality specified by the front end, which includes CRUD operations and specialized data reports. In this implementation, we chose to use MongoDB as it would be easier than a relational database to implement our choice of data structure for the courses. The backend uses the MongoDB Go driver to query the database and other common Go packages to format and return the information to the frontend.

### Endpoints
#### GET /api/courses ==> genericSearch()
* The most basic search feature that queries by a single value, selected on the front end.
* Can filter by `instructor`, `building`, `meeting days`, or course `faculty ratio`.

#### GET /api/schedule ==> getProfessorSchedule()
* Retrieves a professor's schedule (all classes they teach) by their `email`.

#### GET /api/report ==> handlerGetReport()
* Retrieves a report of erroneous or incomplete data in the database.
* Searches for unavailable meeting days and instructor information.

#### GET /api/rooms ==> handlerGetRooms()
* Queries for all available rooms (rooms not in use at the given time and day)

#### GET /api/all ==> handlerGetAllCourses()
* Simply queries for all courses in the database.

#### POST /api/descriptions ==> handlerUpdateCourseDescription()
* Update queries for a course by unique `CRN` value and updates the `description`.

#### POST /api/archive ==> handlerDeleteCourse()
* Query course by unique `CRN` and delete.

#### POST /api/courses ==> handlerCreateCourse()
* Creates a course in database with information from provided `Course` struct.

### Notes
* All `GET` requests send data through HTTP parameters and `POST` requests send data through HTTP Body
