--
-- PostgreSQL database dump
--
INSERT INTO
    posts (title, content, author, created_at, slug)
VALUES
    (
        'Understanding Algorithms',
        'This is a blog post about algorithms...',
        'John Doe',
        CURRENT_TIMESTAMP,
        'understanding-algorithms'
    );

--
