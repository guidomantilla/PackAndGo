DROP TABLE IF EXISTS `pack-and-go`.`city`;
create table `pack-and-go`.`city`
(
    id   int auto_increment
        primary key,
    name varchar(400) not null
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;
begin;
insert into `pack-and-go`.`city` (id, name) values (1, 'Barcelona');
insert into `pack-and-go`.`city` (id, name) values (2, 'Seville');
insert into `pack-and-go`.`city` (id, name) values (3, 'Valencia');
insert into `pack-and-go`.`city` (id, name) values (4, 'Andorra la Vella');
insert into `pack-and-go`.`city` (id, name) values (5, 'Malaga');
commit;

DROP TABLE IF EXISTS `pack-and-go`.`trip`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
create table `pack-and-go`.`trip`
(
    id            int auto_increment
        primary key,
    originId      int     not null,
    destinationId int     not null,
    dates         text    not null,
    price         decimal not null,
    constraint trip_city_01_fk
        foreign key (originId) references city (id)
            on update cascade on delete cascade,
    constraint trip_city_02_fk
        foreign key (destinationId) references city (id)
            on update cascade on delete cascade
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;
