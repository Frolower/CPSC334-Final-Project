CREATE TABLE setupInSession (
    setup_id INT NOT NULL REFERENCES setups(setup_id) ON DELETE CASCADE,
    session_id INT NOT NULL REFERENCES sessions(session_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    PRIMARY KEY (setup_id, session_id)
)