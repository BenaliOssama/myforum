package models

import "time"

type Category struct {
	Id          int    `json:"Id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Post struct {
	PostId       int
	UserId       int
	UserName     string
	Title        string
	Categories   []string
	Content      string
	LikeCount    int
	DislikeCount int
	Created_At   time.Time
	Clicked      bool
	DisClicked   bool
}

var PostsTable = `
    CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        like_count INTEGER DEFAULT 0,
        dislike_count INTEGER DEFAULT 0,
        created_at DATETIME DEFAULT (strftime('%Y-%m-%d %H:%M:%S', 'now')),
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );
`

var CategoriesTable = `
    CREATE TABLE IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
        description TEXT NOT NULL
    );
    INSERT OR IGNORE INTO categories (name, description) VALUES
        ('Liked Posts', 'This category contains all posts you liked'), 
        ('Created Posts', 'This category contains all posts you created'), 
        ('Music', 'Discuss everything related to music, including genres, artists, and concerts'), 
        ('Sports', 'Talk about all types of sports, games, and tournaments'), 
        ('Movies & TV Shows', 'Share recommendations and discuss your favorite films and series'), 
        ('Technology', 'Discuss the latest trends in tech, gadgets, and software'), 
        ('Gaming', 'A place for gamers to discuss games, consoles, and tips'), 
        ('Books & Literature', 'Share and discover books, authors, and literary genres'), 
        ('Travel', 'Exchange travel tips, favorite destinations, and experiences'), 
        ('Food & Cooking', 'Discuss recipes, restaurants, and all things culinary');
`

var PostCategoriesTable = `
    CREATE TABLE IF NOT EXISTS post_categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        post_id INTEGER NOT NULL,
        category_id INTEGER NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
        FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
    );
`
var UsersTable = `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`

var SessionsTable = `CREATE TABLE IF NOT EXISTS sessions (
    session_id TEXT PRIMARY KEY,      
    user_id INTEGER NOT NULL,         
    expires_at DATETIME NOT NULL,    
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP, 
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`
var LikesTable = `CREATE TABLE IF NOT EXISTS likes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    post_id INTEGER,
    comment_id INTEGER,
    target_type TEXT CHECK(target_type IN ('post', 'comment')) NOT NULL,
    type TEXT CHECK(type IN ('like', 'dislike')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
    CHECK (
        (target_type = 'post' AND post_id IS NOT NULL AND comment_id IS NULL) OR 
        (target_type = 'comment' AND comment_id IS NOT NULL AND post_id IS NULL)
    )
);
`
var CommentsTable = `CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        post_id INTEGER NOT NULL,
        content TEXT NOT NULL,
        like_count INTEGER DEFAULT 0,
        dislike_count INTEGER DEFAULT 0,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE
    );
`
