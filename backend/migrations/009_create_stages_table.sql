CREATE TABLE stages (
    stage_id SERIAL PRIMARY KEY,
    stage_number INT NOT NULL,
    championship_id INT NOT NULL REFERENCES championships (championship_id) ON DELETE CASCADE,
    track VARCHAR(255) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);