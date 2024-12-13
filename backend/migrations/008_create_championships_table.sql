CREATE TABLE championships (
    championship_id SERIAL PRIMARY KEY,
    team_id INT NOT NULL REFERENCES teams (team_id) ON DELETE CASCADE,
    championship_name VARCHAR(255) NOT NULL,
    team_standings INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);