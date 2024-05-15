CREATE TABLE sessions (
    id uuid PRIMARY KEY NOT NULL,
    user_id uuid NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    refresh_token varchar NOT NULL,
    user_agent varchar NOT NULL,
    client_ip varchar NOT NULL,
    is_blocked boolean NOT NULL DEFAULT FALSE,
    expires_at timestamptz NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);
