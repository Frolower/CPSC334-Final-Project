import React from 'react';
import "../css/TeamCard.css"

function TeamCard({ team }) {
    return (
        <div className="team-card">
            <h3>{team.team_name}</h3>
        </div>
    );
}

export default TeamCard;
