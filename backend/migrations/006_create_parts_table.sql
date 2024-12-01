CREATE TABLE parts (
    part_id VARCHAR(255) NOT NULL PRIMARY KEY,
    quantity INT NOT NULL,
    chassis_number VARCHAR(255) REFERENCES cars (chassis_number),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);