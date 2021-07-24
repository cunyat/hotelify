CREATE DATABASE rooms
DEFAULT CHARACTER SET utf8mb4
  DEFAULT COLLATE utf8mb4_bin;

use rooms;

create table if not exists rooms (
    uuid varchar(38) not null,
    num varchar(61) not null,
    floor int not null,
    services varchar(255) not null,
    beds json not null,

    primary key (uuid)
) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;
