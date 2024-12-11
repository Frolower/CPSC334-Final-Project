import React from 'react';
import "../css/Landing.scss"
import {Link} from "react-router-dom";

function LandingPage() {
    return (
        <div className={"content-container"}>
            <h1>Ariadne Management</h1>
            <h2>A simple way to do your job.</h2>
            <div className={"link-container"}>
                <Link to="/login" className={"link-style"}>Log In</Link>
                <Link to="/signup" className={"link-style"}>Sign Up</Link>
            </div>
        </div>
    );
}

export default LandingPage;
