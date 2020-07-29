CREATE TABLE IF NOT EXISTS users (
        id BIGSERIAL PRIMARY KEY,
        user_name VARCHAR(32) UNIQUE NOT NULL,
        auth_provider VARCHAR(100) NOT NULL,
        -- pass_key VARCHAR(64) NOT NULL,
        --hash value here?
        first_name CHAR(64) NOT NULL,
        last_name CHAR(64) NOT NULL,
        email VARCHAR(100) NOT NULL,
        created_at TIMESTAMP WITH TIMEZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL
);
CREATE TABLE IF NOT EXISTS notebooks (
        id BIGSERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        public BOOLEAN,
        color VARCHAR(30),
        user_id BIGINT NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
);
--color -hex?
CREATE TABLE IF NOT EXISTS sections (
        id BIGSERIAL PRIMARY KEY,
        created_at TIMESTAMP,
        updated_at TIMESTAMP,
        label_type VARCHAR,
        signifier VARCHAR,
        notebook_id NOT NULL,
        FOREIGN KEY (notebook_id) REFERENCES notebooks(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS notes (
        id BIGSERIAL PRIMARY KEY,
        folder_id NOT NULL REFERENCES sections(id) ON DELETE CASCADE,
        body TEXT,
);
/*
 - a note can have many tags; and a tag can represent many notes
 - a note can have many signifiers and a signifier can belong to many notes 
 
 */
CREATE TABLE IF NOT EXISTS tags (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(80),
);
CREATE TABLE IF NOT EXISTS signifiers (
        id BIGSERIAL PRIMARY KEY,
        name VARCHAR(80),
);
CREATE TABLE IF NOT EXISTS note_tags_junction (
        id BIGSERIAL NOT NULL,
        note_id BIGINT NOT NULL,
        tag_id BIGINT NOT NULL,
        FOREIGN KEY (note_id) REFERENCES notes(id) ON DELETE CASCADE,
        FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
        UNIQUE (note_id, tag_id)
);
CREATE TABLE IF NOT EXISTS note_signifiers_junction (
        id BIGSERIAL PRIMARY KEY,
        note_id BIGINT NOT NULL,
        signifier_id BIGINT NOT NULL,
        FOREIGN KEY (note_id) REFERENCES notes(id) ON DELETE CASCADE,
        FOREIGN KEY (signifier_id) REFERENCES signifiers(id) ON DELETE CASCADE,
        UNIQUE (note_id, signifier_id)
);
-- CREATE TYPE role AS ENUM ('guest', 'member', 'admin');
-- label_type in ["to-do", "in-progress", "complete", "canceled", "delayed", "event/appointment", "deadline"]
-- signifier in ["priority", "inspiration", "explore"]
-- CREATE TABLE technologies (
--         name VARCHAR(255),
--         details VARCHAR(255)
-- -- ); 
-- CREATE TABLE Music (
--         Artist VARCHAR(20) NOT NULL,
--         SongTitle VARCHAR(30) NOT NULL,
--         AlbumTitle VARCHAR(25),
--         Year INT,
--         Price FLOAT,
--         Genre VARCHAR(10),
--         -- CriticRating FLOAT Tags TEXT, PRIMARY KEY(Artist, SongTitle));