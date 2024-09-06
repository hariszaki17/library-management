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