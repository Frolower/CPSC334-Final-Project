import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';  // To extract team_id from URL
import TeamMenu from '../components/TeamMenu';
import AddNewButton from '../components/AddNewButton';

function TeamPage() {
    const { teamId } = useParams();  // Get the team ID from the URL
    const [team, setTeam] = useState(null);

    useEffect(() => {
        // Fetch team details by ID
        const fetchTeam = async () => {
            const response = await fetch(`/api/team/${teamId}`);
            const data = await response.json();
            setTeam(data.team);
        };
        fetchTeam();
    }, [teamId]);

    return (
        <div>
            {team && (
                <>
                    <h1>{team.team_name}</h1>
                    <TeamMenu teamId={teamId} />  {/* Render TeamMenu component */}
                    <div className="team-details">
                        <AddNewButton label="Add New Car" url={`/team/${teamId}/add-car`} />
                        {/* More AddNewButton components for pilots, staff, etc. */}
                    </div>
                </>
            )}
        </div>
    );
}

export default TeamPage;
