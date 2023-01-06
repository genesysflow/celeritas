CREATE TABLE sessions(
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX session_expiry_idx ON sessions (expiry);