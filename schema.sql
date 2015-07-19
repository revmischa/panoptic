BEGIN;

CREATE EXTENSION IF NOT EXISTS "postgis";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DROP TABLE IF EXISTS "stream_source";
DROP TABLE IF EXISTS "account";
DROP TABLE IF EXISTS "login";
DROP TYPE IF EXISTS "stream_state";


/** TYPES **/

CREATE TYPE stream_state AS ENUM('unknown', 'starting', 'playing', 'recording', 'paused', 'stopped');
COMMENT ON TYPE stream_state IS 'state that a stream is in. Used to synchronize with clients.';


/** TABLES **/

-- everything lives under account
CREATE TABLE "account" (
       id SERIAL PRIMARY KEY,
       created TIMESTAMP NOT NULL DEFAULT NOW(),
       name TEXT NOT NULL
);

-- user login
CREATE TABLE "login" (
       id SERIAL PRIMARY KEY,
       created TIMESTAMP NOT NULL DEFAULT NOW(),
       account INTEGER NOT NULL REFERENCES account(id),
       name TEXT,
       email TEXT NOT NULL,
       password TEXT,
       is_active BOOLEAN NOT NULL DEFAULT 't',
       is_verified BOOLEAN NOT NULL DEFAULT 'f',
       last_login_ip INET
);

CREATE TABLE "stream_source" (
       id SERIAL PRIMARY KEY,
       created TIMESTAMP NOT NULL DEFAULT NOW(),
       account INTEGER NOT NULL REFERENCES account(id),
       url TEXT NOT NULL,
       mac_address MACADDR,
       src_width INTEGER,
       src_height INTEGER,
       rec_width INTEGER,
       rec_height INTEGER,
       rec_frame_rate NUMERIC,
       last_frame_time TIMESTAMP,
       last_frame_s3_key TEXT,
       location geography,
       state stream_state NOT NULL DEFAULT 'unknown'
);
COMMENT ON TABLE stream_source IS 'Represents a streaming source (e.g. camera)';


COMMIT;
