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
    content_type varchar not null,
    title        varchar not null,
    payload      varchar not null,
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
    is_mentor  integer,
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
    status        integer not null,
    topic_code    varchar not null
        constraint topic_pk
            primary key
);


create table question
(
    id             bigserial not null,
    topic_code     varchar not null,
    content_text   varchar not null,
    content_image  varchar not null,
    correct_answer varchar not null,
    difficult      integer not null,
    subject_code   varchar not null,
    status         integer not null
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

create table conversation
(
    id          bigserial not null ,
    user_id     varchar                                                         not null,
    mentor_id varchar not null,
    valid_time  timestamp with time zone default (now() + '01:00:00'::interval) not null,
    status integer not null
);

create table subject
(
	subject_code varchar not null
		constraint subject_pk
			primary key,
	subject_name varchar not null,
	status int default 1 not null,
	order_number int not null,
	image varchar
);
