create table scripts
(
    code             varchar(255) not null
        constraint scripts_pkey
            primary key,
    action_processor varchar(255) not null,
    message_failed   varchar(255),
    message_image    varchar(255),
    message_succeed  varchar(255),
    message_text     varchar(255),
    order_number     integer,
    parent_code      varchar(255)
);


create table quick_reply
(
    id           bigserial
        constraint quick_reply_pkey
            primary key,
    script_code  varchar not null,
    content_type varchar,
    title        varchar,
    payload      varchar,
    image_url    varchar,
    status       integer not null,
    order_number integer
);


create table users
(
    id         varchar                                not null
        constraint users_pkey
            primary key,
    phone      varchar,
    email      varchar,
    created_at timestamp with time zone default now() not null,
    status     integer                                not null,
    full_name  varchar(255),
    is_mentor  varchar(255),
    user_name  varchar(255)
);


create table topic
(
    topic_name    varchar,
    description   varchar,
    document_link varchar,
    subject_code  varchar not null,
    image         varchar,
    order_number  integer,
    topic_code    varchar not null
        constraint topic_pk
            primary key
);


create table question
(
    id             bigserial,
    topic_code     bigint,
    content_text   varchar,
    content_image  varchar,
    correct_answer varchar,
    difficult      integer,
    subject_code   varchar
);

comment on column question.difficult is '1 - de
2 - trung binh
3 - kho
4 - rat kho';


create table session_context
(
    id          bigserial,
    user_id     varchar not null,
    script_code varchar not null,
    valid_time  timestamp with time zone not null default (now() + '00:05:00'::interval)
);