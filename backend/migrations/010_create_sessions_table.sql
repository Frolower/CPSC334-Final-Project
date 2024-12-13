CREATE TABLE sessions (
    session_id SERIAL PRIMARY KEY,
    stage_id INT NOT NULL REFERENCES stages (stage_id) ON DELETE CASCADE,
    type VARCHAR(255) NOT NULL,
    session_date DATE NOT NULL,
    start_time TIME NOT NULL,
    weather VARCHAR(255) NOT NULL,
    temperature INT NOT NULL,
    humidity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);