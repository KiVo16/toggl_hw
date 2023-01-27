CREATE TABLE questions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    body TEXT
);

CREATE TABLE questions_options (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    question_id INTEGER,
    body TEXT,
    correct BOOLEAN,
    CONSTRAINT fk_options
        FOREIGN KEY (question_id)   
        REFERENCES questions (id)
        ON DELETE CASCADE 
);