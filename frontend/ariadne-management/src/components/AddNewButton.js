import React from 'react';
import { Link } from 'react-router-dom';

function AddNewButton({ label, url }) {
    return (
        <Link to={url}>
            <button>{label}</button>
        </Link>
    );
}

export default AddNewButton;
