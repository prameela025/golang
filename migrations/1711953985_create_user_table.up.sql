CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID,
    modified_by UUID,
    modified_at TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR(255) NOT NULL
);
