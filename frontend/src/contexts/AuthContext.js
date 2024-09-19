
import React, { createContext, useState, useEffect } from 'react';
import authService from '../services/authService';

const AuthContext = createContext();

const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    const storedUser = localStorage.getItem('user');
  
    if (token && storedUser) {
      try {
        const parsedUser = JSON.parse(storedUser);
        setUser(parsedUser);
        setIsAuthenticated(true);
      } catch (error) {
        console.error("Failed to parse stored user data:", error);
        // Handle the error gracefully, e.g., clear invalid data from localStorage
        localStorage.removeItem('user');
      }
    }
  }, []);  

  const login = async (username, password) => {
    const data = await authService.login(username, password);
    setUser(data.user);
    setIsAuthenticated(true);
  };

  const register = async (fullname, username, password, role) => {
    return await authService.register(fullname, username, password, role);
  };

  const logout = () => {
    authService.logout();
    setUser(null);
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider value={{ user, isAuthenticated, login, register, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export { AuthContext, AuthProvider };
