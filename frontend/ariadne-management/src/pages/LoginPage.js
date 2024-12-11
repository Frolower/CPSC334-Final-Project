import React, { useState } from 'react';
import {Link, useNavigate} from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import "../css/Auth.scss"

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
            <nav>
                <Link className={"link"} to={"/"}>
                    <h2 className={"logo"}>
                        Ariadne Management
                    </h2>
                </Link>
            </nav>
            <div className={"content"}>
                <div className={"form-container"}>
                    <h2>Login</h2>
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
                            <label>Password</label>
                            <input
                                className={"form-input"}
                                type="password"
                                name="password"
                                value={formData.password}
                                onChange={handleChange}
                            />
                        </div>
                        {error && <div>{error}</div>} {/* Show error message if login failed */}
                        {loading && <div>Loading...</div>} {/* Show loading indicator */}
                        <button className={"button-style"} type="submit" disabled={loading}>Login</button>
                    </form>
                </div>
            </div>
        </div>
    );
};

export default LoginPage;
