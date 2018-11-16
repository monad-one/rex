create database rex;

use rex;

create table movies(
    movie_id integer primary key,
    title varchar(255),
    genres varchar(255)
);

create table ml_youtube(
    youtube_id varchar(30) primary key,
    movie_id integer,
    title varchar(255),
    foreign key (movie_id) references movies(movie_id)
);

create table ratings(
    user_id integer,
    movie_id integer,
    rating float,
    primary key (user_id, movie_id),
    foreign key (movie_id) references movies(movie_id)
);

create table recommendations(
    user_id integer,
    movie_id integer,
    prediction float,
    primary key (user_id, movie_id),
    foreign key (movie_id) references movies(movie_id)
);
