CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE ,
    password TEXT NOT NULL ,
    token TEXT,
    token_ttl TEXT
);

CREATE TABLE user_plugins (
    user INTEGER NOT NULL REFERENCES users,
    plugin INTEGER NOT NULL REFERENCES plugins,
    role INTEGER,
    PRIMARY KEY (user, plugin)
);

CREATE TABLE plugins (
    id INTEGER PRIMARY KEY AUTOINCREMENT ,
    name TEXT NOT NULL UNIQUE ,
    shortDescription TEXT NOT NULL,
    version TEXT,
    repository TEXT,
    languages TEXT, --TODO: Extra table?
    allowMultipleInstallations BOOLEAN CHECK ( allowMultipleInstallations IN (0, 1, NULL) )
);

CREATE TABLE plugins_tags (
    plugin INTEGER NOT NULL REFERENCES plugins,
    tag TEXT NOT NULL --TODO: Extra table?
);