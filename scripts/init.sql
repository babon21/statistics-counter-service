--
-- Table structure for table `statistics`
--

DROP TABLE IF EXISTS statistics;

CREATE TABLE statistics
(
    id     SERIAL PRIMARY KEY,
    date   DATE NOT NULL,
    views  INTEGER,
    clicks INTEGER,
    cost   NUMERIC,
    cpc    NUMERIC,
    cpm    NUMERIC
);

CREATE
INDEX statistics_date_index ON statistics (date);