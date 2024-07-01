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
  description text CHECK ( description != '' ),
  year        SMALLINT NOT NULL
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

CREATE TABLE item_has_upload(
    item_id BIGSERIAL NOT NULL,
    upload_id BIGSERIAL NOT NULL,
    FOREIGN KEY (item_id) REFERENCES items(id) on delete cascade,
    FOREIGN KEY (upload_id) REFERENCES uploads(id) on delete cascade,
    UNIQUE (item_id, upload_id)
);

CREATE EXTENSION unaccent;

INSERT INTO authors (name, bio)
VALUES
('William Shakespeare', 'William Shakespeare was an English playwright, poet and actor. He is widely regarded as the greatest writer in the English language and the world''s pre-eminent dramatist. He is often called England''s national poet and the "Bard of Avon".' ),
('Charles Dickens', 'Charles John Huffam Dickens was an English novelist and social critic who created some of the world''s best-known fictional characters, and is regarded by many as the greatest novelist of the Victorian era.' ),
('Jane Austen', 'Jane Austen was an English novelist known primarily for her six novels, which implicitly interpret, critique, and comment upon the British landed gentry at the end of the 18th century. Austen''s plots often explore the dependence of women on marriage for the pursuit of favourable social standing and economic security.' ),
('Fyodor Dostoevsky', 'Fyodor Mikhailovich Dostoevsky, sometimes transliterated as Dostoyevsky, was a Russian novelist, short story writer, essayist and journalist. Numerous literary critics regard him as one of the greatest novelists in all of world literature, as many of his works are considered highly influential masterpieces.' ),
('Leo Tolstoy', 'Count Lev Nikolayevich Tolstoy, usually referred to in English as Leo Tolstoy, was a Russian writer. He is regarded as one of the greatest and most influential authors of all time.' ),
('Author with no bio', null );

INSERT INTO items (name, year, description)
VALUES
('Hamlet', 1623, 'The Tragedy of Hamlet, Prince of Denmark, often shortened to Hamlet, is a tragedy written by William Shakespeare sometime between 1599 and 1601. It is Shakespeare''s longest play, with 29,551 words.' ),
('Othello', 1622, 'Othello is a tragedy written by William Shakespeare, around 1603. The story revolves around two characters, Othello and Iago. Othello is a Moorish military commander who was serving as a general of the Venetian army in defence of Cyprus against invasion by Ottoman Turks.' ),
('Romeo and Juliet', 1597, 'Romeo and Juliet is a tragedy written by William Shakespeare early in his career about the romance between two Italian youths from feuding families. It was among Shakespeare''s most popular plays during his lifetime and, along with Hamlet, is one of his most frequently performed.' ),
('Item without description', 2024, null ),
('Item without author', 2023, null ),
('Test item no1',  2000, 'Description of test item no1' ),
('Test item no2',  2000, 'Description of test item no2' ),
('Test item no3',  2000, 'Description of test item no3' ),
('Test item no4',  2000, 'Description of test item no4' ),
('Test item no5',  2000, 'Description of test item no5' ),
('Test item no6',  2000, 'Description of test item no6' ),
('Test item no7',  2000, 'Description of test item no7' ),
('Test item no8',  2000, 'Description of test item no8' ),
('Test item no9',  2000, 'Description of test item no9' ),
('Test item no10', 2000, 'Description of test item no10' ),
('Test item no11', 2000, 'Description of test item no11' ),
('Test item no12', 2000, 'Description of test item no12' ),
('Test item no13', 2000, 'Description of test item no13' ),
('Test item no14', 2000, 'Description of test item no14' ),
('Test item no15', 2000, 'Description of test item no15' ),
('Test item no16', 2000, 'Description of test item no16' ),
('Test item no17', 2000, 'Description of test item no17' ),
('Test item no18', 2000, 'Description of test item no18' ),
('Test item no19', 2000, 'Description of test item no19' ),
('Test item no20', 2000, 'Description of test item no20' ),
('Book with multiple authors', 2015, null );

INSERT INTO uploads (sum, name, size, type)
VALUES
('1aab5ab4e85254f70c8a8c44a41dacae', 'booklet_03.pdf', 15686086, 'application/pdf'),
('17373e9fa0f0794450959f774cf18bba', 'booklet_04.pdf',  7460195, 'application/pdf'),
('b6e8c10e6543b0fa824b4838c6138946', 'booklet_07.pdf',  7248127, 'application/pdf'),
('9d3d1b7ecb0376e62d4cdb81545ad3be', 'booklet_08.pdf',  6299017, 'application/pdf'),
('2963d75b69ab5b558b89395eb5a6dfa3', 'booklet_09.pdf', 24757689, 'application/pdf'),
('cbf5488c1b1df533c70365cf0251a55c', 'booklet_10.pdf', 31101464, 'application/pdf'),
('19cc8b35683c99e322fd7ff07c3e931e', 'booklet_11.pdf', 11732923, 'application/pdf'),
('becd5eaa8c36bc5c4a6765a953bdbc93', 'booklet_12.pdf',  5749750, 'application/pdf');

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

INSERT INTO item_has_upload (item_id, upload_id)
VALUES
(1,  1),
(2,  1),
(3,  1),
(4,  2),
(5,  2),
(6,  3),
(7,  3),
(8,  4),
(9,  4),
(10, 5),
(11, 5),
(12, 6),
(13, 6),
(14, 7),
(15, 7),
(16, 8),
(17, 8),
(18, 8),
(19, 8),
(20, 8);
