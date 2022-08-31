CREATE TABLE IF NOT EXISTS playlists_collaborations
(
    playlist_id      INTEGER, 
    user_id          INTEGER,

    CONSTRAINT playlist_collaborations_fkey_playlist_id FOREIGN KEY (playlist_id)
    REFERENCES playlists (playlist_id) ON DELETE CASCADE,

    CONSTRAINT playlist_collaborations_fkey_user_id FOREIGN KEY (user_id)
    REFERENCES users (user_id) ON DELETE CASCADE,
    
    PRIMARY KEY(playlist_id, user_id)
);

-- CREATE UNIQUE INDEX IF NOT EXISTS unique_playlists_collaborations ON playlists_collaborations(playlist_id, user_id);