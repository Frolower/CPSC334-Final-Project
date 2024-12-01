import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import LoginPage from './pages/LoginPage';
import Dashboard from './pages/Dashboard';
import { AuthProvider } from './contexts/AuthContext';
import ProtectedRoute from './components/ProtectedRoute';
import SignUpPage from "./pages/SignUpPage";
import TeamPage from './pages/TeamPage';
import CarsPage from './pages/CarsPage';

const App = () => {
    return (
        <AuthProvider>
            <Router>
                <Routes>
                    {/* Public Routes */}
                    <Route path="/" element={<LandingPage />} />
                    <Route path="/login" element={<LoginPage />} />
                    <Route path="/signup" element={<SignUpPage />} />

                    {/* Protected Routes */}
                    <Route
                        path="/dashboard"
                        element={
                            <ProtectedRoute>
                                <Dashboard />
                            </ProtectedRoute>
                        }
                    />
                    <Route
                        path="/team/:teamId"
                        element={
                            <ProtectedRoute>
                                <TeamPage />
                            </ProtectedRoute>
                        }
                    />
                    <Route
                        path="/team/:teamId/cars"
                        element={
                            <ProtectedRoute>
                                <CarsPage />
                            </ProtectedRoute>
                        }
                    />
                </Routes>
            </Router>
        </AuthProvider>
    );
};

export default App;
