CREATE TABLE tires (
    tire_id VARCHAR(255) PRIMARY KEY,
    tread_remaining INT NOT NULL,
    compound VARCHAR(255) NOT NULL,
    chassis_number VARCHAR(255) NOT NULL REFERENCES cars (chassis_number) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);