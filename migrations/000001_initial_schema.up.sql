-- Create users table
CREATE TABLE users (
                       id SERIAL UNIQUE not null PRIMARY KEY,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       first_name VARCHAR(100),
                       last_name VARCHAR(100),
                       last_sign_in_at TIMESTAMP WITH TIME ZONE,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_users_firebase_uid ON users(firebase_uid);
CREATE INDEX idx_users_email ON users(email);

-- Create patients table
CREATE TABLE patients (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          possible_species TEXT[],
                          possible_breed TEXT[],
                          sex VARCHAR(20),
                          date_of_birth DATE,
                          weight DECIMAL(10, 2),
                          height DECIMAL(10, 2),
                          color VARCHAR(100),
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
                          updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_patients_name ON patients(name);

-- Create unprocessed_documents table
CREATE TABLE unprocessed_documents (
                                       id SERIAL PRIMARY KEY,
                                       content TEXT NOT NULL,
                                       num_lines BIGINT NOT NULL DEFAULT 0,
                                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
                                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Create analyzed_documents table
CREATE TABLE analyzed_documents (
                                    id SERIAL PRIMARY KEY,
                                    title VARCHAR(500) NOT NULL,
                                    content TEXT NOT NULL,
                                    num_lines BIGINT NOT NULL DEFAULT 0,
                                    patient_id INTEGER REFERENCES patients(id) ON DELETE SET NULL,
                                    start_line BIGINT NOT NULL,
                                    end_line BIGINT NOT NULL,
                                    unprocessed_document_id INTEGER NOT NULL REFERENCES unprocessed_documents(id) ON DELETE CASCADE,
                                    window_lines TEXT[],
                                    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
                                    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_analyzed_docs_patient_id ON analyzed_documents(patient_id);
CREATE INDEX idx_analyzed_docs_unprocessed_id ON analyzed_documents(unprocessed_document_id);
CREATE INDEX idx_analyzed_docs_created_at ON analyzed_documents(created_at DESC);

-- Create inferences table
CREATE TABLE inferences (
                            id SERIAL PRIMARY KEY,
                            request TEXT NOT NULL,
                            response TEXT NOT NULL,
                            config JSONB,
                            inferable_type VARCHAR(100),
                            inferable_id INTEGER,
                            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
                            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_inferences_inferable ON inferences(inferable_type, inferable_id);

-- Create trigger function for auto-updating updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Apply triggers to all tables
CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_patients_updated_at
    BEFORE UPDATE ON patients
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_unprocessed_documents_updated_at
    BEFORE UPDATE ON unprocessed_documents
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_analyzed_documents_updated_at
    BEFORE UPDATE ON analyzed_documents
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_inferences_updated_at
    BEFORE UPDATE ON inferences
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();