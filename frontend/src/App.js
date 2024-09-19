import './assets/css/App.css';
import { Routes, Route, Navigate } from 'react-router-dom';
import AuthLayout from './layouts/auth';
import AdminLayout from './layouts/admin';
import RTLLayout from './layouts/rtl';
import { ChakraProvider } from '@chakra-ui/react';
import initialTheme from './theme/theme';
import { useState } from 'react';
import ProtectedRoute from './components/protectedRoute/ProtectedRoute';
import { AuthProvider } from './contexts/AuthContext'; // Import AuthProvider

// Importing the components that will be used in the routes
import MainDashboard from 'views/admin/default';
import NFTMarketplace from 'views/admin/marketplace';
import Profile from 'views/admin/profile';
import DataTables from 'views/admin/dataTables';
import RTL from 'views/admin/rtl';
import SignInCentered from 'views/auth/signIn';
import Register from 'views/auth/register';

export default function Main() {
  const [currentTheme, setCurrentTheme] = useState(initialTheme);

  return (
    <ChakraProvider theme={currentTheme}>
      <AuthProvider> {/* Wrap your app with AuthProvider */}
        <Routes>
          {/* Public Auth Routes */}
          <Route path="auth/*" element={<AuthLayout />}>
            <Route path="sign-in" element={<SignInCentered />} />
            <Route path="register" element={<Register />} />
          </Route>

          {/* Protected Admin Routes */}
          <Route
            path="admin/*"
            element={
              <ProtectedRoute>
                <AdminLayout theme={currentTheme} setTheme={setCurrentTheme} />
              </ProtectedRoute>
            }
          >
            <Route path="default" element={<MainDashboard />} />
            <Route path="nft-marketplace" element={<NFTMarketplace />} />
            <Route path="profile" element={<Profile />} />
            <Route path="data-tables" element={<DataTables />} />
          </Route>

          {/* Protected RTL Routes */}
          <Route
            path="rtl/*"
            element={
              <ProtectedRoute>
                <RTLLayout theme={currentTheme} setTheme={setCurrentTheme} />
              </ProtectedRoute>
            }
          >
            <Route path="rtl-default" element={<RTL />} />
          </Route>

          {/* Default Redirection */}
          <Route path="/" element={<Navigate to="/auth/sign-in" replace />} />
        </Routes>
      </AuthProvider>
    </ChakraProvider>
  );
}
