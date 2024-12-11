import React, { createContext, useState, useContext, useEffect } from 'react';

const AuthContext = createContext();

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);

    // Load user from local storage if available
    useEffect(() => {
        const token = localStorage.getItem('authToken');
        if (token) {
            // Use the token to fetch user data or just mark as logged in
            setUser({ token });
        }
    }, []);

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
                body: JSON.stringify({
                    username: credentials.username,
                    password: credentials.password,
                }),
            });

            const data = await response.json();

            if (response.ok) {
                // If login is successful, store the token in localStorage
                localStorage.setItem('authToken', data.token);
                setUser({ username: credentials.username, token: data.token });
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
        localStorage.removeItem('authToken');
        setUser(null);
    };

    // Function to create a team
    const createTeam = async (teamData) => {
        const token = localStorage.getItem('authToken'); // Get the JWT token from localStorage

        if (!token) {
            setError('No token found, please log in.');
            return;
        }

        setLoading(true);
        setError(null);

        try {
            const response = await fetch('http://localhost:8080/createTeam', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`, // Include the JWT token here
                },
                body: JSON.stringify(teamData),
            });

            const data = await response.json();

            if (response.ok) {
                // Successfully created the team, handle success
                console.log('Team created successfully:', data);
            } else {
                // Handle errors
                setError(data.error || 'Failed to create team');
            }
        } catch (err) {
            setError('An error occurred. Please try again.');
        } finally {
            setLoading(false);
        }
    };

    return (
        <AuthContext.Provider value={{ user, login, logout, createTeam, loading, error }}>
            {children}
        </AuthContext.Provider>
    );
};
