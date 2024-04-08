CREATE TABLE passwords (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID,
    modified_at TIMESTAMP,
    modified_by UUID,
    deleted_at TIMESTAMP,
    password VARCHAR(255) NOT NULL,
    user_id UUID,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
