-- Drop triggers first
DROP TRIGGER IF EXISTS update_inferences_updated_at ON inferences;
DROP TRIGGER IF EXISTS update_analyzed_documents_updated_at ON analyzed_documents;
DROP TRIGGER IF EXISTS update_unprocessed_documents_updated_at ON unprocessed_documents;
DROP TRIGGER IF EXISTS update_patients_updated_at ON patients;
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- Drop the trigger function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop tables in reverse dependency order
DROP INDEX IF EXISTS idx_inferences_inferable;
DROP TABLE IF EXISTS inferences;

DROP INDEX IF EXISTS idx_analyzed_docs_created_at;
DROP INDEX IF EXISTS idx_analyzed_docs_unprocessed_id;
DROP INDEX IF EXISTS idx_analyzed_docs_patient_id;
DROP TABLE IF EXISTS analyzed_documents;

DROP TABLE IF EXISTS unprocessed_documents;

DROP INDEX IF EXISTS idx_patients_name;
DROP TABLE IF EXISTS patients;

DROP INDEX IF EXISTS idx_users_email;
DROP TABLE IF EXISTS users;