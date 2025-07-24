CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    date_of_birth DATE NOT NULL,
    avatar_path TEXT,
    user_name TEXT UNIQUE NOT NULL,
    about_me TEXT,
    privacy_status TEXT NOT NULL DEFAULT 'public' CHECK(privacy_status IN ('public', 'private', 'almost_private')),
    gender TEXT NOT NULL CHECK(gender IN ('male', 'female')),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);