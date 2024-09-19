import axios from 'axios';

const API_URL = "http://localhost:8080/api";  // Adjust to match your backend API URL

// Login function
const login = async (username, password) => {
  const response = await axios.post(`${API_URL}/login`, { username, password });
  
  if (response.data.token) {
    localStorage.setItem('token', response.data.token);  // Store JWT token in localStorage
    localStorage.setItem('user', JSON.stringify(response.data.user));  // Store user info
  }
  
  return response.data;
};

// Register function
const register = async (fullname, username, password, role) => {
  const response = await axios.post(`${API_URL}/register`, { fullname, username, password, role });
  
  return response.data;
};

// Logout function
const logout = () => {
  localStorage.removeItem('token');
  localStorage.removeItem('user');
};

export default { login, register, logout };
