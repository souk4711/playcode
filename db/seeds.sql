DO $$
BEGIN
  IF NOT EXISTS (SELECT * FROM languages) THEN
    INSERT INTO languages
      (name, langcode, cr_langcode)
    VALUES
      ('C (GCC 10.2.1)', 'c', 'c'),
      ('C++ (GCC 10.2.1)', 'cpp', 'cpp'),
      ('D (2.100.0)', 'd', 'd'),
      ('Erlang (25.0.4)', 'erlang', 'erlang'),
      ('Golang (1.19)', 'go', 'go'),
      ('Haskell (GHC 9.4.2)', 'haskell', 'haskell'),
      ('Java (OpenJDK 18)', 'java', 'java'),
      ('Javascript (Node.js 18.8.0)', 'javascript', 'javascript'),
      ('Python (3.10.6)', 'python', 'python'),
      ('Ruby (3.1.2)', 'ruby', 'ruby'),
      ('Rust (1.63.0)', 'rust', 'rust'),
      ('Scala (3.1.3)', 'scala', 'scala'),
      ('Typescript (4.8.2)', 'typescript', 'typescript');
  END IF;
END
$$
