import React, { useState, useEffect } from 'react';
import axios from 'axios';
import TeamCard from '../components/TeamCard'; // TeamCard component
import Modal from '../components/Modal';

function Dashboard() {
    const [teams, setTeams] = useState([]);
    const [showModal, setShowModal] = useState(false);
    const [teamName, setTeamName] = useState('');

    // Fetch teams on component mount
    useEffect(() => {
        const fetchTeams = async () => {
            const token = localStorage.getItem('authToken'); // Get the token from localStorage
            try {
                const response = await axios.get('http://localhost:8080/getTeams', {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                setTeams(response.data.teams); // Set teams from the API response
            } catch (error) {
                console.error('Error fetching teams:', error.response ? error.response.data : error);
            }
        };

        fetchTeams();
    }, []);

    const handleCreateTeam = async (e) => {
        e.preventDefault();
        const token = localStorage.getItem('authToken');

        try {
            const response = await axios.post(
                'http://localhost:8080/createTeam',
                { team_name: teamName },
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                }
            );
            setTeams([...teams, response.data.team]); // Add the newly created team to the list
            setShowModal(false); // Close modal
        } catch (error) {
            console.error('Error creating team:', error.response ? error.response.data : error);
        }
    };

    return (
        <div>
            <h1>Dashboard</h1>
            <button onClick={() => setShowModal(true)}>Create Team</button>
            <Modal showModal={showModal} setShowModal={setShowModal} handleSubmit={handleCreateTeam}>
                <h2>Create Team</h2>
                <input
                    type="text"
                    value={teamName}
                    onChange={(e) => setTeamName(e.target.value)}
                    placeholder="Enter team name"
                    required
                />
            </Modal>
            <div className="team-grid">
                {teams.map((team) => (
                    <TeamCard key={team.team_id} team={team} />
                ))}
            </div>
        </div>
    );
}

export default Dashboard;
