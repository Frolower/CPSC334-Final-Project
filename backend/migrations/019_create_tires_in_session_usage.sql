CREATE TABLE tiresInSession (
    tire_id VARCHAR(255) NOT NULL REFERENCES tires(tire_id) ON DELETE CASCADE,
    session_id INT NOT NULL REFERENCES sessions(session_id) ON DELETE CASCADE,
    position VARCHAR(255) NOT NULL,
    PRIMARY KEY (tire_id, session_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);