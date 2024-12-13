CREATE TABLE sessions (
    session_id SERIAL PRIMARY KEY,
    stage_id INT NOT NULL REFERENCES stages (stage_id),
    type VARCHAR(255) NOT NULL,
    session_date DATE NOT NULL,
    start_time TIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);