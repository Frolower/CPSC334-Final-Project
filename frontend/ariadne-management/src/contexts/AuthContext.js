import React, { createContext, useState, useContext } from 'react';

const AuthContext = createContext();

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);

    // Verifying user login
    const login = async (credentials) => {
        setLoading(true);
        setError(null); // Clear any previous errors

        try {
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(credentials), // Send user credentials (username, password)
            });

            const data = await response.json();

            if (response.ok) {
                // If login is successful
                setUser({ username: data.username });  // Save user details in context
            } else {
                // If login fails
                setError(data.message || 'Login failed');
            }
        } catch (err) {
            // If network or other error occurs
            setError('An error occurred. Please try again.');
        } finally {
            setLoading(false);
        }
    };

    const logout = () => {
        setUser(null);
    };

    return (
        <AuthContext.Provider value={{ user, login, logout, loading, error }}>
            {children}
        </AuthContext.Provider>
    );
};
