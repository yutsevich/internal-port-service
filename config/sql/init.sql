create
    database portdb;
use
    portdb;
create table ports
(
    id          int auto_increment primary key,
    name        varchar(30) not null,
    city        varchar(30) not null,
    country     varchar(30) not null,
    alias_id    int,
    FOREIGN KEY (alias_id) REFERENCES alias (id) ON DELETE CASCADE,
    region_id   int,
    FOREIGN KEY (region_id) REFERENCES regions (id) ON DELETE CASCADE,
    coordinates point,
    province    varchar(30) not null,
    timezone    varchar(30) not null,
    code        varchar(30) not null
);
create table alias
(
    id    int auto_increment primary key,
    alias varchar(30) not null
);
create table regions
(
    id     int auto_increment primary key,
    region varchar(30) not null
);
create table unlocs
(
    id    int auto_increment primary key,
    unloc varchar(30) not null
);