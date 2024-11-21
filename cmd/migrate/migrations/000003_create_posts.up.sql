CREATE TABLE
    IF NOT EXISTS posts (
        id bigserial PRIMARY KEY,
        content TEXT,
        title TEXT,
        user_id bigint,
        created_at timestamp(0)
        with
            time zone NOT NULL DEFAULT NOW ()
    );