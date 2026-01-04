DROP INDEX IF EXISTS idx_patients_doctor_id;

ALTER TABLE patients
    DROP COLUMN doctor_id;
