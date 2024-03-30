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

INSERT INTO authors (id, name, bio)
VALUES
(1, 'William Shakespeare', 'William Shakespeare was an English playwright, poet and actor. He is widely regarded as the greatest writer in the English language and the world''s pre-eminent dramatist. He is often called England''s national poet and the "Bard of Avon".' ),
(2, 'Charles Dickens', 'Charles John Huffam Dickens was an English novelist and social critic who created some of the world''s best-known fictional characters, and is regarded by many as the greatest novelist of the Victorian era.' ),
(3, 'Jane Austen', 'Jane Austen was an English novelist known primarily for her six novels, which implicitly interpret, critique, and comment upon the British landed gentry at the end of the 18th century. Austen''s plots often explore the dependence of women on marriage for the pursuit of favourable social standing and economic security.' ),
(4, 'Fyodor Dostoevsky', 'Fyodor Mikhailovich Dostoevsky, sometimes transliterated as Dostoyevsky, was a Russian novelist, short story writer, essayist and journalist. Numerous literary critics regard him as one of the greatest novelists in all of world literature, as many of his works are considered highly influential masterpieces.' ),
(5, 'Leo Tolstoy', 'Count Lev Nikolayevich Tolstoy, usually referred to in English as Leo Tolstoy, was a Russian writer. He is regarded as one of the greatest and most influential authors of all time.' ),
(6, 'Author with no bio', null );

INSERT INTO items (id, name, description)
VALUES
(1, 'Hamlet', 'The Tragedy of Hamlet, Prince of Denmark, often shortened to Hamlet, is a tragedy written by William Shakespeare sometime between 1599 and 1601. It is Shakespeare''s longest play, with 29,551 words.' ),
(2, 'Othello', 'Othello is a tragedy written by William Shakespeare, around 1603. The story revolves around two characters, Othello and Iago. Othello is a Moorish military commander who was serving as a general of the Venetian army in defence of Cyprus against invasion by Ottoman Turks.' ),
(3, 'Romeo and Juliet', 'Romeo and Juliet is a tragedy written by William Shakespeare early in his career about the romance between two Italian youths from feuding families. It was among Shakespeare''s most popular plays during his lifetime and, along with Hamlet, is one of his most frequently performed.' ),
(4, 'Item without description', null ),
(5, 'Item without author', null ),
(6, 'Book with multiple authors', null );

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