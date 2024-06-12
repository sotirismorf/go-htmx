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

CREATE TYPE filetype AS ENUM ('application/pdf', 'image/jpeg', 'image/png');

CREATE TABLE uploads (
  id          BIGSERIAL PRIMARY KEY,
  sum         char(32)  NOT NULL UNIQUE,
  name        text      NOT NULL,
  size        INTEGER   NOT NULL,
  type        filetype  NOT NULL
);

CREATE TABLE item_has_author(
    item_id BIGSERIAL NOT NULL,
    author_id BIGSERIAL NOT NULL,
    FOREIGN KEY (item_id) REFERENCES items(id) on delete cascade,
    FOREIGN KEY (author_id) REFERENCES authors(id) on delete cascade,
    UNIQUE (item_id, author_id)
);

INSERT INTO authors (name, bio)
VALUES
('William Shakespeare', 'William Shakespeare was an English playwright, poet and actor. He is widely regarded as the greatest writer in the English language and the world''s pre-eminent dramatist. He is often called England''s national poet and the "Bard of Avon".' ),
('Charles Dickens', 'Charles John Huffam Dickens was an English novelist and social critic who created some of the world''s best-known fictional characters, and is regarded by many as the greatest novelist of the Victorian era.' ),
('Jane Austen', 'Jane Austen was an English novelist known primarily for her six novels, which implicitly interpret, critique, and comment upon the British landed gentry at the end of the 18th century. Austen''s plots often explore the dependence of women on marriage for the pursuit of favourable social standing and economic security.' ),
('Fyodor Dostoevsky', 'Fyodor Mikhailovich Dostoevsky, sometimes transliterated as Dostoyevsky, was a Russian novelist, short story writer, essayist and journalist. Numerous literary critics regard him as one of the greatest novelists in all of world literature, as many of his works are considered highly influential masterpieces.' ),
('Leo Tolstoy', 'Count Lev Nikolayevich Tolstoy, usually referred to in English as Leo Tolstoy, was a Russian writer. He is regarded as one of the greatest and most influential authors of all time.' ),
('Author with no bio', null );

INSERT INTO items (name, description)
VALUES
('Hamlet', 'The Tragedy of Hamlet, Prince of Denmark, often shortened to Hamlet, is a tragedy written by William Shakespeare sometime between 1599 and 1601. It is Shakespeare''s longest play, with 29,551 words.' ),
('Othello', 'Othello is a tragedy written by William Shakespeare, around 1603. The story revolves around two characters, Othello and Iago. Othello is a Moorish military commander who was serving as a general of the Venetian army in defence of Cyprus against invasion by Ottoman Turks.' ),
('Romeo and Juliet', 'Romeo and Juliet is a tragedy written by William Shakespeare early in his career about the romance between two Italian youths from feuding families. It was among Shakespeare''s most popular plays during his lifetime and, along with Hamlet, is one of his most frequently performed.' ),
('Item without description', null ),
('Item without author', null ),
('Book with multiple authors', null );

INSERT INTO item_has_author (item_id, author_id)
VALUES
(1, 1),
(2, 1),
(3, 1),
(4, 2),
(6, 1),
(6, 2),
(6, 3),
(6, 4);
