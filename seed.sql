
INSERT INTO users (username, password) VALUES ('dudung', 'maman');
SELECT setval('users_id_seq'::regclass, (SELECT MAX(id) FROM users));

INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (1, 'Lorem Ipsum', 1, 1, '978-3-16-148410-0', '2023-08-12', 10);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (2, 'Dolor Sit', 2, 2, '978-1-40-289462-3', '2022-07-14', 15);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (3, 'Amet Consectetur', 3, 3, '978-0-14-200067-0', '2021-06-15', 20);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (4, 'Adipiscing Elit', 1, 4, '978-1-56619-909-4', '2020-05-16', 25);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (5, 'Sed Do', 2, 5, '978-1-4028-9462-6', '2019-04-17', 30);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (6, 'Eiusmod Tempor', 3, 1, '978-3-16-148410-1', '2018-03-18', 35);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (7, 'Incididunt Ut', 1, 2, '978-1-40-289462-4', '2017-02-19', 40);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (8, 'Labore Et', 2, 3, '978-0-14-200067-1', '2016-01-20', 45);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (9, 'Dolore Magna', 3, 4, '978-1-56619-909-5', '2015-12-21', 50);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (10, 'Aliqua Ut', 1, 5, '978-1-4028-9462-7', '2014-11-22', 55);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (11, 'Enim Ad', 2, 1, '978-3-16-148410-2', '2013-10-23', 60);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (12, 'Minim Veniam', 3, 2, '978-1-40-289462-5', '2012-09-24', 65);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (13, 'Quis Nostrud', 1, 3, '978-0-14-200067-2', '2011-08-25', 70);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (14, 'Exercitation Ullamco', 2, 4, '978-1-56619-909-6', '2010-07-26', 75);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (15, 'Laboris Nisi', 3, 5, '978-1-4028-9462-8', '2009-06-27', 80);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (16, 'Ut Aliquip', 1, 1, '978-3-16-148410-3', '2008-05-28', 85);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (17, 'Ex Ea', 2, 2, '978-1-40-289462-6', '2007-04-29', 90);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (18, 'Commodo Consequat', 3, 3, '978-0-14-200067-3', '2006-03-30', 95);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (19, 'Duis Aute', 1, 4, '978-1-56619-909-7', '2005-02-28', 100);
INSERT INTO books (id, title, author_id, category_id, isbn, published_at, stock) VALUES (20, 'Irure Dolor', 2, 5, '978-1-4028-9462-9', '2004-01-01', 105);
SELECT setval('books_id_seq'::regclass, (SELECT MAX(id) FROM books));

INSERT INTO authors (created_at, updated_at, deleted_at, first_name, last_name, biography, birth_date) VALUES
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'John', 'Doe', 'An accomplished author with many bestsellers.', '1980-01-01'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Jane', 'Smith', 'A renowned poet and writer.', '1975-05-15'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Michael', 'Johnson', 'Expert in historical fiction.', '1988-09-23'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Emily', 'Davis', 'Author of several science fiction novels.', '1990-02-28'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Robert', 'Brown', 'Critically acclaimed mystery writer.', '1982-12-05'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Olivia', 'Wilson', 'Young adult fiction writer with a large following.', '1995-11-20'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'James', 'Taylor', 'Famous for his inspirational books.', '1970-07-16'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Sophia', 'Anderson', 'Notable for her historical biographies.', '1984-03-10'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'William', 'Thomas', 'An award-winning travel writer.', '1968-06-30'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Ava', 'Jackson', 'Best known for her contemporary romance novels.', '1992-10-25');
SELECT setval('authors_id_seq'::regclass, (SELECT MAX(id) FROM authors));

INSERT INTO categories (created_at, updated_at, deleted_at, name) VALUES
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Fiction'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Non-Fiction'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Science Fiction'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Fantasy'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Mystery'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Biography'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Historical Fiction'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Romance'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Thriller'),
('2024-01-01 12:00:00', '2024-01-01 12:00:00', NULL, 'Self-Help');
SELECT setval('categories_id_seq'::regclass, (SELECT MAX(id) FROM categories));

INSERT INTO borrowing_records (user_id, book_id, borrowed_at, returned_at, created_at, updated_at) VALUES
(1, 1, '2024-01-01 10:00:00', '2024-01-10 10:00:00', NOW(), NOW()),
(2, 1, '2024-01-02 11:00:00', '2024-01-12 11:00:00', NOW(), NOW()),
(3, 1, '2024-01-03 12:00:00', '2024-01-13 12:00:00', NOW(), NOW()),
(4, 2, '2024-01-04 13:00:00', '2024-01-14 13:00:00', NOW(), NOW()),
(5, 2, '2024-01-05 14:00:00', '2024-01-15 14:00:00', NOW(), NOW()),
(6, 3, '2024-01-06 15:00:00', '2024-01-16 15:00:00', NOW(), NOW()),
(7, 3, '2024-01-07 16:00:00', '2024-01-17 16:00:00', NOW(), NOW()),
(8, 3, '2024-01-08 17:00:00', '2024-01-18 17:00:00', NOW(), NOW()),
(9, 4, '2024-01-09 18:00:00', '2024-01-19 18:00:00', NOW(), NOW()),
(10, 4, '2024-01-10 19:00:00', '2024-01-20 19:00:00', NOW(), NOW()),
(1, 5, '2024-01-11 20:00:00', '2024-01-21 20:00:00', NOW(), NOW()),
(2, 5, '2024-01-12 21:00:00', '2024-01-22 21:00:00', NOW(), NOW()),
(3, 5, '2024-01-13 22:00:00', '2024-01-23 22:00:00', NOW(), NOW()),
(4, 6, '2024-01-14 23:00:00', '2024-01-24 23:00:00', NOW(), NOW()),
(5, 6, '2024-01-15 10:00:00', '2024-01-25 10:00:00', NOW(), NOW()),
(6, 7, '2024-01-16 11:00:00', '2024-01-26 11:00:00', NOW(), NOW()),
(7, 7, '2024-01-17 12:00:00', '2024-01-27 12:00:00', NOW(), NOW()),
(8, 7, '2024-01-18 13:00:00', '2024-01-28 13:00:00', NOW(), NOW()),
(9, 8, '2024-01-19 14:00:00', '2024-01-29 14:00:00', NOW(), NOW()),
(10, 8, '2024-01-20 15:00:00', '2024-01-30 15:00:00', NOW(), NOW()),
(1, 9, '2024-01-21 16:00:00', '2024-01-31 16:00:00', NOW(), NOW()),
(2, 9, '2024-01-22 17:00:00', '2024-02-01 17:00:00', NOW(), NOW()),
(3, 9, '2024-01-23 18:00:00', '2024-02-02 18:00:00', NOW(), NOW()),
(4, 10, '2024-01-24 19:00:00', '2024-02-03 19:00:00', NOW(), NOW()),
(5, 10, '2024-01-25 20:00:00', '2024-02-04 20:00:00', NOW(), NOW()),
(6, 10, '2024-01-26 21:00:00', '2024-02-05 21:00:00', NOW(), NOW()),
(7, 1, '2024-01-27 22:00:00', '2024-02-06 22:00:00', NOW(), NOW()),
(8, 1, '2024-01-28 23:00:00', '2024-02-07 23:00:00', NOW(), NOW()),
(9, 2, '2024-01-29 10:00:00', '2024-02-08 10:00:00', NOW(), NOW()),
(10, 2, '2024-01-30 11:00:00', '2024-02-09 11:00:00', NOW(), NOW()),
(1, 3, '2024-01-31 12:00:00', '2024-02-10 12:00:00', NOW(), NOW()),
(2, 3, '2024-02-01 13:00:00', '2024-02-11 13:00:00', NOW(), NOW()),
(3, 4, '2024-02-02 14:00:00', '2024-02-12 14:00:00', NOW(), NOW()),
(4, 4, '2024-02-03 15:00:00', '2024-02-13 15:00:00', NOW(), NOW()),
(5, 5, '2024-02-04 16:00:00', '2024-02-14 16:00:00', NOW(), NOW()),
(6, 5, '2024-02-05 17:00:00', '2024-02-15 17:00:00', NOW(), NOW()),
(7, 6, '2024-02-06 18:00:00', '2024-02-16 18:00:00', NOW(), NOW()),
(8, 6, '2024-02-07 19:00:00', '2024-02-17 19:00:00', NOW(), NOW()),
(9, 7, '2024-02-08 20:00:00', '2024-02-18 20:00:00', NOW(), NOW()),
(10, 7, '2024-02-09 21:00:00', '2024-02-19 21:00:00', NOW(), NOW()),
(1, 8, '2024-02-10 22:00:00', '2024-02-20 22:00:00', NOW(), NOW()),
(2, 8, '2024-02-11 23:00:00', '2024-02-21 23:00:00', NOW(), NOW()),
(3, 9, '2024-02-12 10:00:00', '2024-02-22 10:00:00', NOW(), NOW()),
(4, 9, '2024-02-13 11:00:00', '2024-02-23 11:00:00', NOW(), NOW()),
(5, 10, '2024-02-14 12:00:00', '2024-02-24 12:00:00', NOW(), NOW()),
(6, 10, '2024-02-15 13:00:00', '2024-02-25 13:00:00', NOW(), NOW());
SELECT setval('borrowing_records_id_seq'::regclass, (SELECT MAX(id) FROM borrowing_records));