create table Person (
                        id uuid primary key default gen_random_uuid(),
                        name text not null,
                        surname text not null,
                        patronymic text,
                        age integer,
                        gender text,
                        nationality text
)

INSERT INTO Person (name,surname) VALUES ('Den','Ivanov')

select name,surname from Person

Drop table Person