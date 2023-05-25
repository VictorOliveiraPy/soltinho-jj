
-- Creation of the "users" table
CREATE TABLE users (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    academy_name VARCHAR(100) NOT NULL,
    instructor_belt VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Creation of the "students" table
CREATE TABLE students (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    user_id VARCHAR REFERENCES users(id),
    name VARCHAR(100) NOT NULL,
    age INTEGER NOT NULL,
    graduation VARCHAR(50) NOT NULL,
    attendance INTEGER DEFAULT 0,
    absences INTEGER DEFAULT 0,
    payment BOOLEAN NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- Grant permission to the user to update the "payment" field in the "students" table
GRANT UPDATE(payment) ON students TO users;

-- Grant permission to the user to update the "attendance" field in the "students" table
GRANT UPDATE(attendance) ON students TO users;

-- Grant permission to the user to update the "absences" field in the "students" table
GRANT UPDATE(absences) ON students TO users;

-- Grant permission to the user to update the "graduation" field in the "students" table
GRANT UPDATE(graduation) ON students TO users;