create table scripts (
    code varchar(255) not null constraint scripts_pkey primary key,
    action_processor varchar(255),
    action_retriever varchar(255),
    message_failed varchar(255),
    message_image varchar(255),
    message_succeed varchar(255),
    message_text varchar(255),
    order_number integer,
    parent_code varchar(255)
);
create table quick_reply (
    id bigserial constraint quick_reply_pkey primary key,
    script_code varchar,
    content_type varchar,
    title varchar,
    payload varchar,
    image_url varchar,
    status integer not null,
    order_number integer
);
-- auto-generated definition
create table users (
    id varchar not null constraint users_pkey primary key,
    phone varchar,
    email varchar,
    created_at timestamp with time zone default now() not null,
    status integer not null,
    full_name varchar(255),
    is_mentor varchar(255),
    user_name varchar(255)
);