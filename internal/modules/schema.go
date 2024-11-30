package modules

var PostsTable = `
    CREATE TABLE  IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        created DATETIME NOT NULL
    );
`
