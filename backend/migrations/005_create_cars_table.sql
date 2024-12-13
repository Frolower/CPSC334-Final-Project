CREATE TABLE cars (
    chassis_number VARCHAR(255) PRIMARY KEY,
    make VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    team_id INT NOT NULL REFERENCES teams (team_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);