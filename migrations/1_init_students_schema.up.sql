-- nft table
CREATE TABLE IF NOT EXISTS students (
    id BIGSERIAL PRIMARY KEY,
    username CHARACTER VARYING NOT NULL,
    score CHARACTER VARYING,
    time_created TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS unique_idx_students ON students (username);

COMMENT ON COLUMN students.username IS 'username';
COMMENT ON COLUMN students.score IS 'score';
COMMENT ON COLUMN students.time_created IS 'the time when this record of row is inserted into pg';

