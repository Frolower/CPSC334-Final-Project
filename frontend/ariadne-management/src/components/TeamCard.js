import React from 'react';

function TeamCard({ team }) {
    return (
        <div className="team-card">
            <h3>{team.team_name}</h3>
            <p>{team.created_at}</p>
        </div>
    );
}

export default TeamCard;
