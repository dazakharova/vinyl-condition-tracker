CREATE TABLE record_sides (
                              id INTEGER PRIMARY KEY AUTOINCREMENT,
                              record_id INTEGER NOT NULL,
                              name TEXT NOT NULL,

                              FOREIGN KEY (record_id) REFERENCES records(id) ON DELETE CASCADE,
                              CONSTRAINT record_sides_name_check CHECK (name IN ('A','B','C','D','E','F')),
                              CONSTRAINT record_sides_unique UNIQUE (record_id, name)
);