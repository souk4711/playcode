DO $$
BEGIN
  IF NOT EXISTS (SELECT * FROM languages) THEN
    INSERT INTO languages
      (name, langcode, cr_langcode, cr_main_filename)
    VALUES
      ('C (GCC 10.2.1)', 'c', 'c', 'main.c'),
      ('C++ (GCC 10.2.1)', 'cpp', 'cpp', 'main.cpp'),
      ('D (2.100.0)', 'd', 'd', 'main.d'),
      ('Erlang (25.0.4)', 'erlang', 'erlang', 'main.erl'),
      ('Golang (1.19)', 'go', 'go', 'main.go'),
      ('Haskell (GHC 9.4.2)', 'haskell', 'haskell', 'main.hs'),
      ('Java (OpenJDK 18)', 'java', 'java', 'main.java'),
      ('Javascript (Node.js 18.8.0)', 'javascript', 'javascript', 'main.js'),
      ('Python (3.10.6)', 'python', 'python', 'main.py'),
      ('Ruby (3.1.2)', 'ruby', 'ruby', 'main.rb'),
      ('Rust (1.63.0)', 'rust', 'rust', 'main.rs'),
      ('Scala (3.1.3)', 'scala', 'scala', 'main.scala'),
      ('Typescript (4.8.2)', 'typescript', 'typescript', 'main.js');
  END IF;
END
$$
