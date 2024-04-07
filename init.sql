-- Subscription table
CREATE TABLE IF NOT EXISTS subscription (
    id VARCHAR(16) PRIMARY KEY DEFAULT (LOWER(HEX(RANDOMBLOB(16)))),
    name VARCHAR(255) NOT NULL UNIQUE,
    frequency VARCHAR(255) NOT NULL,
    rateLimit INT NOT NULL,
    deprecated BOOLEAN NOT NULL DEFAULT FALSE
    );

-- Token table
CREATE TABLE IF NOT EXISTS token (
    id VARCHAR(16) PRIMARY KEY DEFAULT (LOWER(HEX(RANDOMBLOB(16)))),
    passphrase VARCHAR(255),
    subscription_id VARCHAR(16) NOT NULL,
    currentUsage INT NOT NULL DEFAULT 0,
    FOREIGN KEY (subscription_id) REFERENCES Subscription(id)
    );
