create
extension if not exists "uuid-ossp";

create
extension if not exists "pgcrypto";

drop table if exists "user";
drop table if exists "team_task";
drop table if exists "team_secret";
drop table if exists "team";
drop table if exists "role";
drop table if exists "answer";
drop table if exists "task";
drop table if exists "task_type";
drop table if exists "task_difficulty";
drop table if exists "answer_secret";
drop table if exists "answer_type";
drop table if exists "secret";

create table "team"
(
    id    uuid not null default uuid_generate_v4(),
    title text not null,

    primary key (id)
);

create table "role"
(
    id    int generated always as identity,
    title text not null,

    primary key (id)
);

create table "user"
(
    id              uuid    not null default uuid_generate_v4(),
    password        text    not null,
    phone_number    text    not null,
    email           text    not null, --либо искать дублирующие, либо убираем unique
    email_confirmed boolean not null default false,
    telegram        text    not null,
    vk              text    not null,
    "group"         text    not null,
    name            text    not null,
    team_id         uuid             default null,
    role_id         int              default null,
    is_admin        boolean          default false,

    primary key (id),

    foreign key (team_id)
        references "team" (id)
        on delete set default,
    foreign key (role_id)
        references "role" (id)
        on delete set default
);

create table "task_type"
(
    id    int generated always as identity,
    title text not null,

    primary key (id)
);

create table "task"
(
    id             uuid not null default uuid_generate_v4(),
    title          text not null,
    description    text,
    type_id        int  not null,
    max_points     int  not null,
    min_points     int  not null,
    answer_type_id int  not null,
    primary key (id),

    foreign key (type_id)
        references "task_type" (id)
);

create table "team_task"
(
    id                uuid        not null default uuid_generate_v4(),
    task_id           uuid        not null,
    task_type_id        int        not null,
    team_id           uuid        not null,
    start_time        timestamptz not null default now(),
    additional_points int         not null default 0,
    answer_text       text                 default null,
    answerImageBase64  text                 default null,
    result            bool                 default null,
    primary key (id),

    foreign key (task_id)
        references "task" (id),
    foreign key (team_id)
        references "team" (id)
);


insert into "task_type" (title)
values ('Данные удалены'), ('НОЦ'), ('Онлайн задачи'), ('Видео задачи');

delete from task_type where id = 1;

insert into "role" (title)
values ('Участник'), ('Заместитель'), ('Капитан');

insert into "task" ("title", "description",  "type_id", "max_points", "min_points", "answer_type_id") values ('**Великие умы МГТУ**', 'Важной частью истории Университета являются известные выпускники и прославленные преподаватели, на которых вы можете равняться, с кого можете брать пример. Многие из них внесли весомый вклад в науку ещё в период обучения в МГТУ. Одним из таких людей является **Шухов Владимир Григорьевич**.

Деятельность этого великого инженера не ограничилась одной отраслью. Владимир Григорьевич развивал нефтяную индустрию и строительство, теплотехнику и судостроение, военное и реставрационное дело. В память о нем и об одном его великом изобретении в мастерской главного здания установлена *табличка* . Большинство из вас, дорогие участники, столкнется с ней во время Учебно-Технологической Практики.

***

*Ваша задача - отсыкать ее, узнать о каком изобретении она повествует и сделать командное фото на фоне таблички с вашей реакцией на создание чего-то нового одним из членов вашей команды.*', 3, 35, 0, 1);
insert into "task" ("title", "description",  "type_id", "max_points", "min_points", "answer_type_id") values ('**Дебют на Бауманской сцене**', 'Главные концерты в Бауманке проходят в ***Большом Зале Центра Креативных Индустрий (БЗЦКИ)***, но помимо концертного зала в Здании также есть коворкинг "Парус" и малая сцена, где часто проводятся "Квартирники".

***

*Пока вы еще не успели выступить в Большом Зале, можете попробовать записать дебют своей команды в одном из юмористических жанров, которые у нас есть в МГТУ: Стендап(RUSH Stand Up Club), КВН (Баманская лига КВН) или Импровизация (Bauman Improv) на этой небольшой сцене*.', 4, 35, 0, 4);
insert into "task" ("title", "description",  "type_id", "max_points", "min_points", "answer_type_id") values ('**Чистая комната**', 'В МГТУ есть множество Научно-Образовательных Центров, которые занимаются исследованиями в различных сферах. Одним из самых необычных и продвинутых является **НОЦ "Функциональные микро/наносистемы"**. Этот НОЦ уникален и тем, что прямо на территории УЛК рядом с кабинетами, в которых проходят пары, они создали, так называемую **"Чистую комнату"**.

Это комната в которой соблюдаются определенные условия окружающей среды, например, температура, влажность, давление и многое другое. Люди внутри находятся только в стерильных костюмах, а в случае "разгерметизации" эта комната по новой очищается и работы не проводятся там еще в течение недели, пока комната не вернется к исходному состоянию. Понаблюдать за работой в ней можно через стекло на первом этаже. Также на этом стекле есть комиксы выпускаемые Бауманкой.

***

*Но и вы сможете создать свой комикс, придумав название и обложку для него на фоне этого НОЦа на тему того, что бы произошло в случае "разгерметизации" этой комнаты.*', 3, 35, 0, 1);
insert into "task" ("title", "description",  "type_id", "max_points", "min_points", "answer_type_id") values ('**Музей МГТУ**', '**Музей МГТУ** является хранилищем Бауманской истории. Внутри него собрано множество экспонатов, рассказывающих о прошлом студентов, их жизни, учебе и достижениях. Самым старым экспонатом музея является... А впрочем...

***

*Лучше узнать это прямо внутри музея и сделать фотографию с тем, что в этом году отмечает свой 2-вековой юбилей. А также написать еще 3 понравившихся экспоната.*

*Если окажется так, что музей будет закрыт, вам стоит обратиться к портрету ректора, который 55 лет назад официально основал музей, сфотографироваться с ним и написать его фамилию. Найти его можно не на улице и не на бульваре, а в компании других известных выпускников.*', 3, 35, 0, 1);
insert into "task" ("title", "description",  "type_id", "max_points", "min_points", "answer_type_id") values ('**Дом физики**', 'Для многих МГТУ становится вторым домом. Но и в самой Бауманке есть свой дом - **Дом физики**.
Главным местом в Доме физики является маятник Фуко, который доказывает то, что Земля вращается вокруг своей оси.

***

*Теперь вам нужно продемонстрировать на фото рядом с маятником, что бы произошло, если бы Земля резко остановилась.*', 3, 35, 0, 2);
insert into "task" ("title", "description",  "type_id", "max_points", "min_points", "answer_type_id") values ('**Будущее уже близко**', 'Писатели фантасты и режиссёры современных блокбастеров про спасение планеты много размышляли о будущем человечества в космическом пространстве. Одним из главных страхов космонавтов может стать **встреча с инопланетной жизнью**. Пока будущее не наступило, предлагаем вам пофантазировать на эту тему.

***

*Найдите космонавта, "искажающего реальность" свои скафандром, рядом с корпусом самого "космического" факультета и напугайте его, выступив в роли монстра с другой планеты. Такой памятный момент обязательно необходимо запечатлить.*', 3, 35, 0, 2);
insert into "task" ("title", "description",  "type_id", "max_points", "min_points", "answer_type_id") values ('**Наши дети**', 'МГТУ им. Н.Э. Баумана является одним из старейших университетов нашей страны. На протяжении многих лет наши ученые развивают различные области науки. Но некоторые направления их деятельности требовали настолько углубленного изучения, что **часть факультетов и лабораторий МГТУ стали впоследствии самостоятельными высшими учебными заведениями**.

***

*Все эти факты об истории МГТУ можно найти рядом с центральной лестницей Главного Учебного Корпуса. Для того, чтобы не забыть свою историю - найдите 5 вузов (кроме МГТУ) на этом месте и сделайте рядом с ним фото в ретро стиле.*', 3, 35, 0, 1);
insert into "task" ("title", "description",  "type_id", "max_points", "min_points", "answer_type_id") values ('**Фатальная ошибка**', 'Многим из вас перед началом обучения удалось побывать на **самом масштабном выездном мероприятии Студенческого Совета - "Школе Молодого Бауманца"**. Одним из самых популярных мемов этого года стала шутка про "Фатальную ошибку".

Человеком, повторившим этот мем для бауманцев, стал **известный выпускник МГТУ**, который в настоящее время играет важнейшую роль в жизни нашего Университета.

***

*Ваша задача - записать на видео короткую сценку, которая закончится фразой "фатальная ошибка" на фоне любой аудитории, находящейся на том же этаже, что и кабинет человека, о котором идет речь в этой загадке. Просмотр ШМБлога поможет вам освежить в памяти имя этого человека.*', 4, 35, 0, 4);
insert into "task" ("title", "description",  "type_id", "max_points", "min_points", "answer_type_id") values ('**Что-то новенькое**', '**Сегодня МГТУ им. Н.Э. Баумана расширяет свою площадь за счет строительства новых корпусов**.

Этим процессом можно полюбоваться на макете, расположенном на "Красной площади". Однако с недавнего времени в главном здании нашего Университета есть еще одно место в южном крыле, в котором вы можете узнать о новых кампусах.

***

*Найдите этот стенд и сделайте командное фото указав какой корпус вы ждете больше всего.*', 3, 35, 0, 2);
--
-- COPY public.team (id, title) FROM stdin;
-- 7b1a54ca-dbe6-40f6-bd58-955ed42987f4	команда разработки сайта
-- \.
--
--
-- --
-- -- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
-- --
--
-- COPY public."user" (id, password, phone_number, email, email_confirmed, telegram, vk, "group", name, team_id, role_id, is_admin) FROM stdin;
-- 23b542a0-e275-4665-bc6e-4345f3011d3b	$2a$08$0vikwKywPXQQu6dKwe.lcumaAbNDU96DTOLUWNheg0bjMG3jNVwyO	89678244577	pavelpawlenko@ya.ru	f	niles_penrose	niles_penrose	РК6-11	павленко антон александрович	7b1a54ca-dbe6-40f6-bd58-955ed42987f4	3	f
-- 5219672a-bd7a-4588-8100-995c6ec17b36	$2a$08$NM9NUkyJl18uoMVOHTYhrO/FoXODRFm0K.vBhzSSYdmt7tggsrHzi	88005553535	cringe@ya.ru	f	cringe	cringe	РК5-11	пипяо ты кринд	\N	\N	f
-- \.

