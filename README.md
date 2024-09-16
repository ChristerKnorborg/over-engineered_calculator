# Over-engineered Calculator


This project implements a RESTful calculator API in Go. It uses Firestore as the backend storage, deployed on Google Cloud Run. 
The calculator supports the operations: addition, subtraction, multiplication, division, modulo, and exponentiation. Furthermore, it has a history feature that stores calculations on the database.
The project also contains a small webpage that communicates with the backend via JavaScript to perform the calculations. The frontend is deployed on Firebase Hosting. 

All of the endpoints are `\GET` methods, and are as follows:

| #Endpoint   | #Parameters             |  
| `\add`      | `\operand1`, `\operand2`|  
| `\subtract` | `\operand1`, `\operand2`|  
| `\multiply` | `\operand1`, `\operand2`|  
| `\divide`   | `\operand1`, `\operand2`|  
| `\modulo`   | `\operand1`, `\operand2`|  
| `\power`    | `\operand1`, `\operand2`|  
| `\history`  |                         |  


| Endpoint  | Parameters        |
| ------------- | ------------- |
| `\add`  | `\operand1`, `\operand2`  |
| `\add`  | `\operand1`, `\operand2`  |
| `\add`  | `\operand1`, `\operand2`  |
| `\divide`  | `\operand1`, `\operand2`  |

My solution to the problem contains the following (implemented) files:


calculator/  
│  
├── operations/  
│ &nbsp ├── calculator.go              # Core calculator operations (add, subtract, etc.)  
│ &nbsp ├── storage.go                 # Storage logic for saving and retrieving history  
│   └── calculator_test.go           # Unit tests for calculator operations and history  
│  
├── api/  
│   ├── handlers.go                # HTTP handlers for API requests   
│   ├── routes.go                  # API routing  
│   └── api_test.go                # Unit tests for API handlers and routes  
│  
├── main.go                        # Entry point of the application  
│  
├── setup/  
│   ├── setup.go                   # Firestore initialization (also includes an emulator for local testing)  
│  
├── web/                             
│   ├── styles.css                 # Styling for frontend calculator  
│   ├── index.html                 # Simple webpage that interacts with the calculator API  
│   └── script.js                  # JavaScript for making HTTP requests to the API  
│  
├── PostmanRequests.json           # Postman requests for API testing  
