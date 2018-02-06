-- Игры, команды, игроки, турниры, матчи
-- Команда
DROP TABLE IF EXISTS Commands;
CREATE TABLE Commands (
    ID INT NOT NULL PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Icon VARCHAR(100) NOT NULL
);


DROP TABLE IF EXISTS CommandsPlayer;
CREATE TABLE CommandsPlayer (
    CommandsID INT,
    PlayerID INT,
    IsCommander BOOLEAN
);

DROP TABLE IF EXISTS Players;
CREATE TABLE Players (
    ID INT,
    FirstName VARCHAR(100),
    SecName VARCHAR(100)
    Avatar VARCHAR(100),
    IsCommander BOOLEAN
);

-- Игры 

DROP TABLE IF EXISTS Games;
CREATE TABLE Games (
    ID INT,
    Player1ID INT,
    Player2ID INT,
    VideoTranslationUrl VARCHAR(400)
)

DROP TABLE IF EXISTS GameResults;
CREATE TABLE Games (
    GamesID INT,
    Player1ID INT,
    Player2ID INT
)

