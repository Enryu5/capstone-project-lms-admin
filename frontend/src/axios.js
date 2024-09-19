import axios from 'axios';

// Create an Axios instance
const api = axios.create({
  baseURL: 'http://localhost:8080/api', // Your backend URL
  withCredentials: true, // Enable credentials (cookies, tokens)
  headers: {
    'Content-Type': 'application/json', // Define the content type
  },
});

export default api;
