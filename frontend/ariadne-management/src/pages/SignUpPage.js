import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext'; // Import the AuthContext hook

const SignUpPage = () => {
    const navigate = useNavigate();
    const { login } = useAuth(); // Use the login method from AuthContext

    const [formData, setFormData] = useState({
        username: '',
        email: '',
        first_name: '',
        last_name: '',
        password: '',
        confirmPassword: '',
    });

    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(false);

    // Function to handle input changes
    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({
            ...prevData,
            [name]: value,
        }));
    };

    // Validate email format
    const validateEmail = (email) => {
        const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
        return emailRegex.test(email);
    };

    // Handle form submission and validate fields
    const handleSubmit = async (e) => {
        e.preventDefault();

        // Validate email format
        if (!validateEmail(formData.email)) {
            setError('Please enter a valid email address.');
            return;
        }

        // Validate that passwords match
        if (formData.password !== formData.confirmPassword) {
            setError('Passwords do not match.');
            return;
        }

        // Clear any previous errors
        setError(null);
        setLoading(true);

        // Proceed with the signup process
        try {
            const response = await fetch('http://localhost:8080/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    username: formData.username,
                    email: formData.email,
                    first_name: formData.first_name,
                    last_name: formData.last_name,
                    password: formData.password,
                }),
            });

            const data = await response.json();

            if (response.ok) {
                // If signup is successful, log the user in
                await login({
                    username: formData.username,
                    password: formData.password,
                });

                // Redirect to the dashboard
                navigate('/dashboard');
            } else {
                setError(data.message || 'Account creation failed');
            }
        } catch (err) {
            // If network or other error occurs
            setError('An error occurred. Please try again.');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div>
            <h2>Create Account</h2>
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
                    <label>Email</label>
                    <input
                        type="email"
                        name="email"
                        value={formData.email}
                        onChange={handleChange}
                    />
                </div>
                <div>
                    <label>First name</label>
                    <input
                        type="text"
                        name="first_name"
                        value={formData.first_name}
                        onChange={handleChange}
                    />
                </div>
                <div>
                    <label>Last name</label>
                    <input
                        type="text"
                        name="last_name"
                        value={formData.last_name}
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
                <div>
                    <label>Confirm Password</label>
                    <input
                        type="password"
                        name="confirmPassword"
                        value={formData.confirmPassword}
                        onChange={handleChange}
                    />
                </div>
                {error && <div>{error}</div>} {/* Show error message */}
                {loading && <div>Loading...</div>} {/* Show loading indicator */}
                <button type="submit" disabled={loading}>Create Account</button>
            </form>
        </div>
    );
};

export default SignUpPage;
