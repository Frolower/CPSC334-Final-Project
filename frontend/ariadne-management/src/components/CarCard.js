import React from 'react';
import { Link } from 'react-router-dom';

function CarCard({ car }) {
    return (
        <div className="car-card">
            <h3>{car.make} {car.model}</h3>
            <p>Chassis Number: {car.chassis_number}</p>
            <Link to={`/team/${car.team_id}/car/${car.chassis_number}`}>View Details</Link>
        </div>
    );
}

export default CarCard;
