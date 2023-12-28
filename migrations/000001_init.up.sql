CREATE TABLE IF NOT EXISTS ApodData (
    id SERIAL PRIMARY KEY,
    copyright VARCHAR(255),
    date DATE,
    explanation TEXT,
    media_type VARCHAR(50),
    service_version VARCHAR(10),
    title VARCHAR(255),
    url VARCHAR(255),
    hdurl VARCHAR(255),
    data bytea
);

CREATE INDEX IF NOT EXISTS idx_apod_date ON ApodData(date);
