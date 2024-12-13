CREATE TABLE laps (
    lap_number INT NOT NULL,
    session_id INT NOT NULL REFERENCES sessions (session_id),
    lap_time TIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);