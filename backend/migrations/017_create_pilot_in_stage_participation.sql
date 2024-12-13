CREATE TABLE pilotInStage (
    document_number INT NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    stage_id INT NOT NULL REFERENCES stages (stage_id) ON DELETE CASCADE,
    chassis_number VARCHAR(255) NOT NULL REFERENCES cars (chassis_number) ON DELETE CASCADE,
    result INT NOT NULL,
    FOREIGN KEY (document_number, first_name, last_name) REFERENCES pilots (document_number, first_name, last_name) ON DELETE CASCADE,
    PRIMARY KEY (document_number, first_name, last_name, chassis_number, stage_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);