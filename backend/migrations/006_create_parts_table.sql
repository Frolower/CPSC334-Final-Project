CREATE TABLE parts (
    part_id VARCHAR(255) PRIMARY KEY,
    quantity INT NOT NULL,
    chassis_number VARCHAR(255) NOT NULL REFERENCES cars (chassis_number) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);