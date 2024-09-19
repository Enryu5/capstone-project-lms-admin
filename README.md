## THIS IS STILL WORK IN PROGRESS AND INCOMPLETE ##

To setup this project, please follow this instruction (make sure you already installed postgresql as this project use postgresql for the database):
1. Clone the repository
2. Open the backend folder and open the .env.example
3. Change the value to match your environment (for JWT secret key you can use online generator)
4. Change the .env.example file name to .env
5. Open the repository in your code editor and create two terminal
6. Change  1 terminal to backend directory (cd backend) and the other to frontend (cd frontend)
7. Run go mod tidy in the backend terminal and npm install in the frontend terminal
8. Run go run main.go in the backend terminal and npm start in the frontend terminal
