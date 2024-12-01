import React from 'react';
import { Link } from 'react-router-dom';

function TeamMenu({ teamId }) {
    return (
        <nav>
            <ul>
                <li><Link to={`/team/${teamId}/cars`}>Cars</Link></li>
                <li><Link to={`/team/${teamId}/pilots`}>Pilots</Link></li>
                <li><Link to={`/team/${teamId}/staff`}>Staff</Link></li>
                <li><Link to={`/team/${teamId}/championships`}>Championships</Link></li>
            </ul>
        </nav>
    );
}

export default TeamMenu;
