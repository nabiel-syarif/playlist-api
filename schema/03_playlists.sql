CREATE TABLE IF NOT EXISTS playlists
(
    playlist_id      SERIAL PRIMARY KEY, 
    name             TEXT NOT NULL,
    owner_id         INTEGER NOT NULL,
    created_at       TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP NULL DEFAULT NOW(),

    CONSTRAINT playlist_fkey_user_id FOREIGN KEY (owner_id)
    REFERENCES users (user_id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_playlists_pk ON songs (song_id);
