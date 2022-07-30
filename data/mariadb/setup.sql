CREATE DATABASE IF NOT EXISTS `user`;
USE `user`;

CREATE TABLE `users` (
    `userId` int NOT NULL AUTO_INCREMENT,
    `mail` varchar(200) NOT NULL,
    `password` varchar(500) NOT NULL,
    PRIMARY KEY (`userId`)
);

CREATE TABLE `financialInformation` (
    `id` int NOT NULL AUTO_INCREMENT,
    `liquidAssets` double,
    `illiquidAssets` double,
    `income` double,
    `expenses` double,
    `liabilities` double,
    `userId` int NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`userId`) REFERENCES users(`userId`)
);

CREATE TABLE `goals` (
     `id` int NOT NULL AUTO_INCREMENT,
     `startDate` date,
     `endDate` date,
     `goal` varchar(200),
     `userId` int NOT NULL,
     PRIMARY KEY (`id`),
     FOREIGN KEY (`userId`) REFERENCES users(`userId`)
);

CREATE TABLE `periodicalDeposits` (
      `id` int NOT NULL AUTO_INCREMENT,
      `amount` double,
      `currency` varchar(20),
      `period` varchar(20),
      `userId` int NOT NULL,
      PRIMARY KEY (`id`),
      FOREIGN KEY (`userId`) REFERENCES users(`userId`)
);

CREATE TABLE `riskAssessments` (
       `id` int NOT NULL AUTO_INCREMENT,
       `riskAppetite` int,
       `maxVolatility` double,
       `userId` int NOT NULL,
       PRIMARY KEY (`id`),
       FOREIGN KEY (`userId`) REFERENCES users(`userId`)
);

CREATE TABLE `experiences` (
       `id` int NOT NULL AUTO_INCREMENT,
       `type` varchar(20) NOT NULL,
       `transactions` int NOT NULL,
       `timePeriod` varchar(20) NOT NULL,
       `userId` int NOT NULL,
       PRIMARY KEY (`id`),
       FOREIGN KEY (`userId`) REFERENCES users(`userId`)
);

CREATE TABLE `deposits` (
        `id` int NOT NULL AUTO_INCREMENT,
        `datetime` datetime NOT NULL,
        `amount` double NOT NULL,
        `currency` varchar(20) NOT NULL,
        `target` varchar(100) NOT NULL,
        `userId` int NOT NULL,
        PRIMARY KEY (`id`),
        FOREIGN KEY (`userId`) REFERENCES users(`userId`)
);

CREATE TABLE `securities` (
        `id` int NOT NULL AUTO_INCREMENT,
        `symbol` varchar(20) NOT NULL,
        `shares` double NOT NULL,
        `averageCostPerShare` double NOT NULL,
        `userId` int NOT NULL,
        PRIMARY KEY (`id`),
        FOREIGN KEY (`userId`) REFERENCES users(`userId`)
);