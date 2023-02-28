# Budget WebApp
This is a simple budgetting webapp built in Python Flask. It is using a micro service architecture compriced of three services.
1. Web page frontend served by Flask.
2. HTTP backend served by Gin-Gonic (GO)
3. MariaDB database.

The application works by showing a HTML form on a webpage. The user enters the amounts for the different categories. When pressing Save ("Spara") the /add-expense url will be called with the form data. This data is then converted to a new dictionary with correct types and the JSONified and sent in another POST request to the backend/expenses endpoint. I the backend the data is bound to a struct to be accessible and then inserted into the database.

On GET requests to the backend (such as building the table shown on the main page) the last entered data is retrieved and sent to the frontend from the backend together with a total field summarizing the total amounts and the ID and timestamp fields from the DB (but they aren't currently used).

Screenshots are available in the Screenshots folder in the repo. The Frontend image shows the UI with the form and the table. The backend shows the data returned from the backend on a GET request.

All three services are very loosely coupled and will run independently.

The services are containerized using Docker and Docker Compose. They communicate with each other using the internal Docker DNS. All sensitive data are passed to the containers using Docker secrets combined with environment variables.
