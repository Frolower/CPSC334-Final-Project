CREATE TABLE staffToCar (
    document_number INT NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    chassis_number VARCHAR(255) NOT NULL REFERENCES cars (chassis_number) ON DELETE CASCADE,
    FOREIGN KEY (document_number, first_name, last_name) REFERENCES staff (document_number, first_name, last_name) ON DELETE CASCADE,
    PRIMARY KEY (document_number, first_name, last_name, chassis_number),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);