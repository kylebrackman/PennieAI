ALTER TABLE patients
    ADD COLUMN doctor_id INTEGER REFERENCES users(id) ON DELETE SET NULL;

CREATE INDEX idx_patients_doctor_id ON patients(doctor_id);