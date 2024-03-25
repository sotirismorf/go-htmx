CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text    NOT NULL,
  bio  text
);

CREATE TABLE groups (
  id          BIGSERIAL PRIMARY KEY,
  name        text    NOT NULL CHECK ( name != '' )
);

CREATE TABLE items (
  id          BIGSERIAL PRIMARY KEY,
  name        text    NOT NULL CHECK ( name != '' ),
  description text CHECK ( description != '' )
);

CREATE TABLE publishers (
  id          BIGSERIAL PRIMARY KEY,
  name        text    NOT NULL,
  description text
);

CREATE TABLE item_has_author(
    item_id BIGSERIAL NOT NULL,
    author_id BIGSERIAL NOT NULL,
    FOREIGN KEY (item_id) REFERENCES items(id),
    FOREIGN KEY (author_id) REFERENCES authors(id),
    UNIQUE (item_id, author_id)
);

-- name: InitAuthors :exec
INSERT INTO authors ( name, bio)
VALUES
( 'William Shakespeare', 'William Shakespeare was an English playwright, poet and actor. He is widely regarded as the greatest writer in the English language and the world''s pre-eminent dramatist. He is often called England''s national poet and the "Bard of Avon".' ),
( 'Charles Dickens', 'Charles John Huffam Dickens was an English novelist and social critic who created some of the world''s best-known fictional characters, and is regarded by many as the greatest novelist of the Victorian era.' ),
( 'Jane Austen', 'Jane Austen was an English novelist known primarily for her six novels, which implicitly interpret, critique, and comment upon the British landed gentry at the end of the 18th century. Austen''s plots often explore the dependence of women on marriage for the pursuit of favourable social standing and economic security.' ),
( 'Fyodor Dostoevsky', 'Fyodor Mikhailovich Dostoevsky, sometimes transliterated as Dostoyevsky, was a Russian novelist, short story writer, essayist and journalist. Numerous literary critics regard him as one of the greatest novelists in all of world literature, as many of his works are considered highly influential masterpieces.' ),
( 'Leo Tolstoy', 'Count Lev Nikolayevich Tolstoy, usually referred to in English as Leo Tolstoy, was a Russian writer. He is regarded as one of the greatest and most influential authors of all time.' ),
( 'Author with no bio', null );

-- name: InitItems :exec
INSERT INTO items ( name, description)
VALUES
( 'Book 1', 'Description' ),
( 'Item without description', null );
