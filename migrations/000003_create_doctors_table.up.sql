CREATE TABLE doctors (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    specialization VARCHAR(100) NOT NULL,
    available_days VARCHAR(100) NOT NULL,
    department_id UUID NOT NULL,

    CONSTRAINT fk_doctors_department
        FOREIGN KEY (department_id)
        REFERENCES departments(id)
);