CREATE TABLE side_listenings (
                                 id INTEGER PRIMARY KEY AUTOINCREMENT,
                                 record_side_id INTEGER NOT NULL,
                                 cartridge TEXT NOT NULL,
                                 min_grade INTEGER NOT NULL,
                                 max_grade INTEGER NOT NULL,
                                 comments TEXT NOT NULL DEFAULT '',
                                 listened_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

                                 FOREIGN KEY (record_side_id) REFERENCES record_sides(id) ON DELETE CASCADE,
                                 CHECK (min_grade <= max_grade)
);