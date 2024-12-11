import React, { useState } from 'react';
import {Link, useNavigate} from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext'; // Import the AuthContext hook
import "../css/Auth.scss"

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
                // If signup is successful, log the user in with the returned token
                await login({
                    username: formData.username,
                    password: formData.password,
                    token: data.token,  // Pass the token received from backend
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
            <nav>
                <Link className={"link"} to={"/"}>
                    <h2 className={"logo"}>
                        Ariadne Management
                    </h2>
                </Link>
            </nav>
            <div className={"content"}>
                <div className={"form-container"}>
                    <h2>Create Account</h2>
                    <form onSubmit={handleSubmit}>
                        <div className={"input-content"}>
                            <label>Username</label>
                            <input
                                className={"form-input"}
                                type="text"
                                name="username"
                                value={formData.username}
                                onChange={handleChange}
                            />
                        </div>
                        <div className={"input-content"}>
                            <label>Email</label>
                            <input
                                className={"form-input"}
                                type="email"
                                name="email"
                                value={formData.email}
                                onChange={handleChange}
                            />
                        </div>
                        <div className={"input-content"}>
                            <label>First name</label>
                            <input
                                className={"form-input"}
                                type="text"
                                name="first_name"
                                value={formData.first_name}
                                onChange={handleChange}
                            />
                        </div>
                        <div className={"input-content"}>
                            <label>Last name</label>
                            <input
                                className={"form-input"}
                                type="text"
                                name="last_name"
                                value={formData.last_name}
                                onChange={handleChange}
                            />
                        </div>
                        <div className={"input-content"}>
                            <label>Password</label>
                            <input
                                className={"form-input"}
                                type="password"
                                name="password"
                                value={formData.password}
                                onChange={handleChange}
                            />
                        </div>
                        <div className={"input-content"}>
                            <label>Confirm Password</label>
                            <input
                                className={"form-input"}
                                type="password"
                                name="confirmPassword"
                                value={formData.confirmPassword}
                                onChange={handleChange}
                            />
                        </div>
                        {error && <div>{error}</div>} {/* Show error message */}
                        {loading && <div>Loading...</div>} {/* Show loading indicator */}
                        <button className={"button-style"} type="submit" disabled={loading}>Create Account</button>
                    </form>
                </div>
            </div>
        </div>
    );
};

export default SignUpPage;
