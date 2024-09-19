import React, { useContext } from 'react';
import { Navigate, Outlet } from 'react-router-dom';
import { AuthContext } from '../../contexts/AuthContext';

// ProtectedRoute Component for React Router v6
const ProtectedRoute = ({ redirectPath = '/login' }) => {
  const { isAuthenticated } = useContext(AuthContext);

  // Check if the user is authenticated; otherwise, redirect to the login page
  if (!isAuthenticated) {
    return <Navigate to={redirectPath} replace />;
  }

  return <Outlet />; // Renders the nested route component if authenticated
};

export default ProtectedRoute;
