CREATE TABLE parameters (
    attribute VARCHAR(255) NOT NULL,
    setup_id INT NOT NULL REFERENCES setups (setup_id),
    value VARCHAR(255) NOT NULL,
    PRIMARY KEY (attribute, setup_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);