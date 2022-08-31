CREATE TABLE IF NOT EXISTS playlists_songs
(
    playlist_id      INTEGER, 
    song_id          INTEGER,

    CONSTRAINT playlist_song_fkey_playlist_id FOREIGN KEY (playlist_id)
    REFERENCES playlists (playlist_id) ON DELETE CASCADE,

    CONSTRAINT playlist_song_fkey_song_id FOREIGN KEY (song_id)
    REFERENCES songs (song_id) ON DELETE CASCADE,

    PRIMARY KEY (playlist_id, song_id)
);

-- CREATE UNIQUE INDEX IF NOT EXISTS unique_playlist_song ON playlists_songs(playlist_id, song_id);