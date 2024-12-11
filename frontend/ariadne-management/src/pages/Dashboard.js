import React, { useState } from 'react';
import Modal from '../components/Modal';
import axios from 'axios';

function Dashboard() {
    const [showModal, setShowModal] = useState(false);
    const [teamName, setTeamName] = useState('');

    // Function to handle form submission
    const handleCreateTeam = (e) => {
        e.preventDefault();

        const token = localStorage.getItem('authToken'); // Get the token from localStorage

        // API call to create a team
        axios
            .post(
                'http://localhost:8080/createTeam',
                { team_name: teamName },
                {
                    headers: {
                        'Authorization': `Bearer ${token}` // Add token to the Authorization header
                    }
                }
            )
            .then(response => {
                console.log('Team created successfully:', response.data);
                setShowModal(false);  // Close the modal after successful creation
            })
            .catch(error => {
                console.error('Error creating team:', error.response ? error.response.data : error);
            });
    };



    return (
        <div>
            <h1>Dashboard</h1>

            {/* Button to open modal */}
            <button onClick={() => setShowModal(true)}>Create Team</button>

            {/* Modal for creating team */}
            <Modal
                showModal={showModal}
                setShowModal={setShowModal}
                handleSubmit={handleCreateTeam}
            >
                <h2>Create Team</h2>
                <div>
                    <label htmlFor="teamName">Team Name:</label>
                    <input
                        type="text"
                        id="teamName"
                        value={teamName}
                        onChange={(e) => setTeamName(e.target.value)}
                        required
                    />
                </div>
            </Modal>
        </div>
    );
}

export default Dashboard;
