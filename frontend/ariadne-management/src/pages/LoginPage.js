import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

const LoginPage = () => {
    const { login, error, loading } = useAuth(); // Get login function and states from context
    const navigate = useNavigate();

    const [formData, setFormData] = useState({
        username: '',
        password: '',
    });

    // Handle input changes
    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({
            ...prevData,
            [name]: value,
        }));
    };

    // Handle form submission
    const handleSubmit = async (e) => {
        e.preventDefault();

        // Use the login function from AuthContext
        await login(formData);

        // If login is successful, redirect to dashboard
        if (!error) {
            navigate('/dashboard');
        }
    };

    return (
        <div>
            <h2>Login</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Username</label>
                    <input
                        type="text"
                        name="username"
                        value={formData.username}
                        onChange={handleChange}
                    />
                </div>
                <div>
                    <label>Password</label>
                    <input
                        type="password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                    />
                </div>
                {error && <div>{error}</div>} {/* Show error message if login failed */}
                {loading && <div>Loading...</div>} {/* Show loading indicator */}
                <button type="submit" disabled={loading}>Login</button>
            </form>
        </div>
    );
};

export default LoginPage;
