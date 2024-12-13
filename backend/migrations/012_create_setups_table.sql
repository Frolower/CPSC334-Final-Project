CREATE TABLE setups (
    setup_id SERIAL PRIMARY KEY,
    setup_name VARCHAR(255) NOT NULL,
    track_name VARCHAR(255) NOT NULL,
    chassis_number VARCHAR(255) NOT NULL REFERENCES cars (chassis_number),
    championship_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);