CREATE TABLE user
(
    userID INT NOT NULL,
    userName varchar (255) UNIQUE NOT NULL,
    password varchar (255) NOT NULL,
    deviceName varchar (255) NOT NULL,
    createdTime timestamp without time zone default now(),
    latestLogin timestamp without time zone default now(),
    constraint user_pkey primary key(userID),
)

CREATE TABLE deviceInfo
(
    deviceID INT NOT NULL,
    deviceName varchar (255) NOT NULL,
    platform varchar (255) NOT NULL,
    userID INT NOT NULL,
    constraint deviceinfo_pkey primary key(deviceID),
)
