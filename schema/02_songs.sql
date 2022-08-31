CREATE TABLE IF NOT EXISTS songs
(
    song_id          SERIAL PRIMARY KEY, 
    title            VARCHAR(255) NOT NULL,
    performer        VARCHAR(255) NOT NULL,
    genre            VARCHAR(255) NOT NULL,
    duration         REAL NOT NULL,
    created_at       TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_songs_pk ON songs (song_id);