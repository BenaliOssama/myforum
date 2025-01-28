-- Create a `snippets` table in SQLite
CREATE TABLE IF NOT EXISTS snippets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

-- -- Insert some dummy records
-- INSERT INTO snippets (title, content, created, expires) VALUES (
--     'An old silent pond',
--     'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
--     CURRENT_TIMESTAMP,
--     DATETIME(CURRENT_TIMESTAMP, '+365 days')
-- );

-- INSERT INTO snippets (title, content, created, expires) VALUES (
--     'Over the wintry forest',
--     'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
--     CURRENT_TIMESTAMP,
--     DATETIME(CURRENT_TIMESTAMP, '+365 days')
-- );

-- INSERT INTO snippets (title, content, created, expires) VALUES (
--     'First autumn morning',
--     'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
--     CURRENT_TIMESTAMP,
--     DATETIME(CURRENT_TIMESTAMP, '+7 days')
-- );

-- Add an index on the created column
CREATE INDEX IF NOT EXISTS idx_snippets_created ON snippets(created);


-- No "USE" statement in SQLite; just directly create the table in the active database
CREATE TABLE IF NOT EXISTS sessions (
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP NOT NULL
);

-- Create an index on the expiry column for faster queries
CREATE INDEX IF NOT EXISTS sessions_expiry_idx ON sessions (expiry);



-- SQLite doesn't use the USE statement. The database is specified when opening the connection.

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    hashed_password TEXT NOT NULL,
    created DATETIME NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS users_uc_email ON users (email);
