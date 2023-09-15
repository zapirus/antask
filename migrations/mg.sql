CREATE TABLE IF NOT EXISTS table_data (
                                          id SERIAL PRIMARY KEY,
                                          headers TEXT,
                                          body TEXT
);

CREATE TABLE IF NOT EXISTS user_data (
                                         id SERIAL PRIMARY KEY,
                                         time TIMESTAMP,
                                         user_id VARCHAR(355),
                                         data INTEGER,
                                         FOREIGN KEY (data) REFERENCES table_data (id)
    );