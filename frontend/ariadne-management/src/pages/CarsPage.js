import React, { useState } from 'react';
import Modal from '../components/Modal';  // Reuse the Modal component
import axios from 'axios';

function CarsPage() {
    const [showModal, setShowModal] = useState(false);
    const [carDetails, setCarDetails] = useState({ make: '', model: '', chassisNumber: '' });

    const handleCreateCar = (e) => {
        e.preventDefault();

        axios
            .post('http://localhost:8080/createCar', carDetails)  // Replace with actual API endpoint
            .then(response => {
                console.log('Car added:', response.data);
                setShowModal(false);  // Close modal after success
            })
            .catch(error => {
                console.error('Error adding car:', error);
            });
    };

    return (
        <div>
            <h1>Cars Page</h1>
            <button onClick={() => setShowModal(true)}>Add Car</button>

            {/* Modal for adding a car */}
            <Modal
                showModal={showModal}
                setShowModal={setShowModal}
                handleSubmit={handleCreateCar}
            >
                <h2>Add Car</h2>
                <div>
                    <label htmlFor="make">Make:</label>
                    <input
                        type="text"
                        id="make"
                        value={carDetails.make}
                        onChange={(e) => setCarDetails({ ...carDetails, make: e.target.value })}
                        required
                    />
                </div>
                <div>
                    <label htmlFor="model">Model:</label>
                    <input
                        type="text"
                        id="model"
                        value={carDetails.model}
                        onChange={(e) => setCarDetails({ ...carDetails, model: e.target.value })}
                        required
                    />
                </div>
                <div>
                    <label htmlFor="chassisNumber">Chassis Number:</label>
                    <input
                        type="text"
                        id="chassisNumber"
                        value={carDetails.chassisNumber}
                        onChange={(e) => setCarDetails({ ...carDetails, chassisNumber: e.target.value })}
                        required
                    />
                </div>
            </Modal>
        </div>
    );
}

export default CarsPage;
