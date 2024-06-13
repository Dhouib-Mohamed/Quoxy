-- Subscription table
CREATE TABLE subscription (
                              id VARCHAR(16) PRIMARY KEY,
                              name VARCHAR(255) NOT NULL UNIQUE,
                              frequency VARCHAR(255) NOT NULL,
                              rate_limit INT NOT NULL,
                              deprecated TINYINT NOT NULL DEFAULT 0
);

-- Token table
CREATE TABLE token (
                       id VARCHAR(16) PRIMARY KEY,
                       passphrase VARCHAR(255),
                       subscription_id VARCHAR(16) NOT NULL,
                       current_usage INT NOT NULL DEFAULT 0,
                       FOREIGN KEY (subscription_id) REFERENCES subscription(id)
);
