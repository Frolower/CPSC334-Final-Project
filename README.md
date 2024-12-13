<h1>Ariadne Management</h1>

<h2>Prerequisites</h2>
<ol>
<li>
<a href="https://nodejs.org/en/download/package-manager">Install Node.js</a>
</li>
<li>
<a href="https://go.dev/dl/">Install Golang</a>
</li>
<li>
Setup a <a href="https://www.postgresql.org/download/">PSQL</a> database
</li>
</ol>


<h2>To install the project</h2>
<ol>
<li>
Clone github repo <code>git clone https://github.com/Gonzaga-CPSC-321-Fall-2024/final-project-Frolower.git</code>
</li>
<li>
Navigate into the folder <code> cd final-project-Frolower</code>
</li>
</ol>

<h2>To run frontend</h2>
<ol>
<li>
Navigate into the frontend folder <code>cd frontend</code>
</li>
<li>
Navigate into the React folder <code>cd ariadne management</code>
</li>
<li>
Download dependencies <code>npm install</code>
</li>
<li>
Run <code>npm start</code> to start the frontend on <a href="http://localhost:3000">http://localhost:3000</a>
</li>
</ol>

<h2>To run backend</h2>
<ol>
<li>
Navigate into the backend folder <code>cd backend</code>
</li>
<li>
Download dependencies <code>go mod tidy</code>
</li>
<li>
Create a <code>.env</code> file with a following structure <br>
<code>
DB_USER=your_user <br>
DB_PASSWORD=your_password <br>
DB_HOST=your_host <br>
DB_PORT=your_port <br>
DB_NAME=your_db_name <br>

JWT_SECRET= A string of symbold e.g., sruhguosfeifosjifg
</code>
</li>
<li>
Navigate into cmd folder <code>cd cmd</code>
</li>
<li>
Run <code>go run main.go</code> to start the backend app on http://localhost:8080
</li>
</ol>

<h2>To setup databse</h2>
<p>All required tables are located in <code>backend/migration</code> folder</p>

<h2>To use the app</h2>
<p>Currently, the frontend is broken, so <a href="https://www.postman.com/">postman</a> is highly recommended</p>
<ol>
<li>
In postman type <code>http://localhost:8080/signup</code> to create an account
</li>
<li>
In body choose RAW JSON option and create a file with such configuration <code> <br>
{ <br>
    "username": "username", <br>
    "email": "test@gmail.com", <br>
    "first_name": "Name", <br>
    "last_name": "Last Name", <br>
    "password": "password" <br>
} <br>
</code>
and submit it, you will receive JWT you must use to access all other routes
</li>
<li>
When testing other routes in Header add <code>Authorization</code> key and <code>JWT</code> as value
</li>
<li>
All working routes are located in <code>main.go</code>
</li>
</ol>