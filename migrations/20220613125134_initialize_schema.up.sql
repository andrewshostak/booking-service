create table if not exists bookings (
    id bigserial primary key,
    first_name varchar(64) not null,
    last_name varchar(64) not null,
    gender varchar(64) not null,
    birthday date not null,
    launchpad_id varchar(24) not null,
    destination_id varchar(24) not null,
    launch_date timestamp with time zone not null
);