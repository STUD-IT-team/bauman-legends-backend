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

COPY public.team (id, title) FROM stdin;
73ce80d7-2798-442f-918e-bf40511dfb07	Захват ГЗ
7b1a54ca-dbe6-40f6-bd58-955ed42987f4	команда разработки сайта
40355e30-8138-4e9e-a09c-6d4b9ae9f4aa	Безудержные лягухи
deb66347-3dcb-4a55-8e42-b28de913c9b8	ФН2 moment
bf7be543-8627-4d3e-a526-e78e375769bc	 Безудержные лягухи
0916ff1b-8ad6-4faf-a78a-84390c65427e	ЗахватГЗ
3155c953-27df-4bf3-93a7-3c3cce34deb3	Орловские огурцы
0979f441-ea8c-4322-a670-3bb16cf27f6d	Легендарные абобы
71580588-3eca-4258-817b-a9d3ab09d703	Плазма до...
2d40c471-5ebc-4470-a419-97d182c2b096	ЖуКи
01dec7e5-69f0-4948-8adf-16def255ab41	Котики
6b2934e3-98fb-4a5a-9138-35a62df992b9	стоп это самое
9909f062-467c-44d2-a9a3-0e43553d6892	Тестовая команда
7b146ffd-3941-4851-ab0a-e557884042d4	Мама Квиз и сын Квизенок
4fb50154-1111-4f44-badb-e48f83e0d433	Жюль Ен
b9b02aae-0db5-4970-961c-9a5a198f7ee7	Летс гоу это самое
83c9e33a-e863-4ef5-9909-2ab2db18b7e4	Хортица
95fafbf5-1bbd-4f58-9ee3-1e1cf6eff939	216 стол
712edd72-aed6-4ec5-a521-cf4a8bbbf4f7	Winx Club
02e685fa-19e7-4f7e-8025-8dbfbee483a0	Коррекционный класс МГТУ им Баумана
5095faa5-51b3-4dcc-90cd-744b3566ccf2	Братство миндаля
543a4e0b-e35d-44f5-85df-43431142aab6	Авторский курс
f9edfee0-2a1c-46c9-8518-d5a92b0cff1b	Людочка
c8fafbd5-03e9-481f-bea4-93819d84a7bb	Кипятильники
28fd4455-c957-4065-97f6-847b632daf2b	ЯМыТищи
589132a7-7358-41cf-833c-7a5bd9f8e1b9	43 команда
300ef3ae-2e98-43cc-8151-a7f42980de56	алкожабы
fd8040f4-8e1e-40fe-8012-0398c80c5324	Торнадо
ab1abad2-ef51-48f7-a15c-797d70061adc	недопонятые гении
6fa56c83-f29f-4909-9d63-ed00152c144b	Энергогангстеры
2acc4cbe-c250-4830-a19e-8b8af134ce2b	Спец.Мед1-12
a845c25f-1727-40ef-90fa-8422bd94a20c	Легенды профсоюза
29103222-c68f-427b-ab8c-2cb674f1fb97	Легендарные котятки
ccbdb464-37b4-4fa5-88e0-8a97ea84dceb	Московская прописка
d68a442e-2c48-43d4-b184-15596fc420ec	Типо дизайнеры
d8d79fe2-6702-439e-bdda-ef7f2be4ebae	авоськи
a614efba-69c4-43fb-a9a4-d70d8be4c365	Мозги без инженеров
94abce80-a6ef-4ec6-8aae-cdfd55e4d47f	Масики
6148ed20-a1df-49f4-9e86-58163649a2d0	Adventure Time
05735fa8-6c86-4424-865c-79df5b9850c2	говнюки
1c6d8c8a-9bda-4ec8-bbf6-8dd0b83b98d6	Хеликоптер
307c360f-258f-45b6-b1cc-81c946df749e	ПивоПивы
32cc7a99-1f39-4dbb-a160-e293a14838a1	Шестое чувство
2b55f5a7-5872-4904-bdea-7d54243826e8	Литуны
6148d3c6-58fd-4ee4-830f-c8643ef8938f	МандАринОчки
786d1844-89ec-4d6a-bcc2-ed5a5b83ab95	жогры
0c68c5b1-5185-497b-9a60-b1c1e1763b6e	Троян
17beb218-e4fc-4644-935f-cfbcb462b0ec	 Чертилы
074448b3-2a26-4b58-b63a-cecbdad6590d	До копчёной infinitum
2b3b1e0a-1f6f-4f7d-a0da-7877dbf854f5	ФАНКЛУБ
8f8f8c5c-9df1-43ca-aa12-6eed5edfea17	Кафедра питания
a0ff82fd-80f4-4fbd-9e4a-9ee65ab57a3f	Кекики(тест 1)
9e61f1aa-978a-461a-90f6-99808580d083	Орден феникса
9f8e20ed-c2b2-4f77-9d95-ca22c7ad82d3	энергетики
966087bd-2504-4ec6-b863-422b8f9d24cf	ASAP
20cd8a22-243a-42f8-8253-5cadbeeae107	Джульетта
9b7aeb4b-946b-4491-957a-b10688ac71d5	Ну 10 типо
485a9afa-6f4c-457b-952a-203d110898b2	Экспрессо
1533d334-657f-455a-94b2-edb6a31ddc18	Великолепная четвёрка и шестёрка
bd9782f4-4e81-41b1-b6a6-be7bd5bfde4f	крутыши
572bfabc-f038-4482-97a6-cefd372cd700	утята
6e7bc851-c93b-4c94-bb2e-fdb0656c528a	Банка сгущёнки
899a1e33-3acd-4935-9330-ffa80a966412	Легенды
4b718a67-1bea-4e14-ada5-9b2eaa2731ed	Фанбаза Зуева
23845812-ffff-45e2-94e8-08b2581f1be2	Чертилы
811be257-7e9c-4847-91b3-3552791b9d7e	Пози👹
16956eba-5185-4ba2-bce2-a8239d6a89a9	Посиделки в ночи
0c21a026-7815-4005-ab4f-ac2b5192acb9	Хеллоу Китти
a5ea392d-d4f6-4c87-a04e-002e252cfb6c	Продолжение
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" (id, password, phone_number, email, email_confirmed, telegram, vk, "group", name, team_id, role_id, is_admin) FROM stdin;
75575339-4aa4-47aa-ac19-5aeb832ce636	$2a$08$gqAHaHVy3Z9kxyjO4d/tR.amvwpEpOj1dw.pavc8j8vV5zcpQCcCO	89167374550	roman_boyko.48@mail.ru	f	sourr_cream	weird418	ИУ9-12Б	Бойко Роман Александрович	\N	\N	f
13e04bc9-c1b7-489d-a3c4-180f86fc02dd	$2a$08$3gu3VAd5zL0d5BfMWnn8RewukuOp6LsrcCYyRxlRBTPhuekoqHFIm	89295650842	iamyuraleniviy@gmail.com	f	yuraleniviy	https://vk.com/yivinelaruy	ФН3-12Б	Юрков Даниил Денисович	\N	\N	f
a3fc09f4-6701-40c5-9411-ff21533c2e76	$2a$08$kztZAUOmZrio9StCE4qIg.FAHZ4EeMgkDfGJIty/GkI78UFuCoLda	8903 100-10-77	anyalad5@gmail.com	f	annlad5	id681348986	ИУ9-12	Ладонцева Анна Андрееана	\N	\N	f
b1e97967-71bd-4839-bd5e-53ef7e87cb87	$2a$08$0NO4G4mjGOro9MrkmQZUA.2fjaGpOYXgYcjD8djdy/XfKcwi8ABKK	8918-532-82-72	yukate.yudina@gmail.com	f	NeVstak	yukateyu	ИУ9-12Б	Юдина Екатерина Сергеевна	\N	\N	f
730db704-559c-4c03-b2ac-56b6e3715976	$2a$08$ZLXHrzAGSG3ItB9pP8SDOOSmFn5POlyvVvgrfDDH4XQrWeHT7j4PC	8 925 36130-32	arastun.misirov@mail.ru	f	arastun01	id615264221	Э1-13	Мисиров Арастун	\N	\N	f
db289e5d-5b0a-402b-a9b3-501cc4e7778c	$2a$08$9ncvwCSLAOpEop7GFeaHC.fuxpoQjM4hbpiX/aL.t7MGtF2RQ78R.	89774924467	chalenkotaya@yandex.ru	f	molodetsogyrets	molodetsogyrets	РК6-14	Чаленко Таисия Алексеевна	\N	\N	f
9c48bbca-d3f2-4b54-81f6-e99d4b23e005	$2a$08$xbNzKBNIn52HfSuKnGV8gOrvyTBICnq4HxugyYTuuxuK3MygjlTDC	8985-480-80-48	arbuzovtu@bmstu.ru	f	tima_arbuzov	tima_arov	ИУ9-12Б	Арбузов Тимофей Юрьевич	\N	\N	f
1fcfc807-b6ad-43d5-a080-9e429e8c3828	$2a$08$qLpGv3rVNlJsyE5/Ifluiu.v2n0NHV3Nrlr9P8WRRvbPaDnz8GOAq	89175902799	m-haustova2006@mail.ru	f	m_mksmvn	id672229995	ИУ9-12Б	Хаустова Мария Максимовна	\N	\N	f
21bc7856-de70-4601-af32-9fde1ab091db	$2a$08$UvH2mpbXVVZRt9wVFbzMLucMAPRfWZDVuzsJ4M3NUljpWVegmXB8i	89175288178	leraleksandrova005@gmail.com	f	chhikurova	bezzler	МТ4-11Б	Чикурова Лера Александровна	\N	\N	f
df897082-bd37-4347-a085-907818a02f61	$2a$08$oBUREO.LdURtAA69X7nPP.PTTxSoc5Ny8/QFx4OBDLJBMkri7CGUu	89534577744	ivan.plotnik.96@mail.ru	f	IvanPlotnikLuchshiy	ivanplotnik4	ИУ1-12	Плотник Иван Дмитриевич	\N	\N	f
91b60bee-0ad0-4a50-8e43-312ff6aa4c62	$2a$08$MVusvIZ/yRV4Y87NN9xw0ONJAl.J6j3p4j02GAGee9gWW7XX0r.mW	8 965 328 60 88	romanovls@ya.ru	f	kbeast	kbeast	БМТ5-11М	Романова Екатерина Анатольевна	\N	\N	f
22844855-23dc-45ca-bd8f-f048febd4096	$2a$08$u2xmSuUMA97b/zp1Mb86Y.nUUNA9a7vW.xe2BxHROu8xEwPhYKGau	8 920 54123-73	alexshedrina@yandex.ru	f	sashka170veu	sashka170veu	БМТ1-12Б	Щедрина Александра Владимировна	\N	\N	f
c5f861e7-cc14-47c1-bd42-341c57f88ef1	$2a$08$NfU3hDw6uZJI8oj.Z0qUeuLcxEvJmCpnVmWa22LcwZ4Wje8HghuSy	89161495651	orlov.kirill13@gmail.com	f	Tonatogead_56t	glyxapb	ИУ8-15	Орлов Кирилл Александрович	\N	\N	f
9bb53637-e57c-46ef-a753-cb9a5b91177b	$2a$08$jG/w2TXooT0GKpLVBcUp7.PR4AHdpAzQ5WCBMNMhS3cRTNwJGEY0i	89663285396	sharonov.2005@bk.ru	f	Garryyyyyyyyyyyyyyyyyyyyy	garryyyyyyy	СМ1-12	Игорь Александрович Шаронов	\N	\N	f
78153fbc-c8f4-42b9-befc-4d00eb3a3232	$2a$08$pXbckYmyY/irpFskHutinOJ09oWBeUZcqKwO6LTXeWrOiGUS6u4Pu	8916-298-82-72	meandarina@gmail.com	f	daechwitasoty	bicthboss	ИБМ6-13Б	Патрова Арина Александровна	\N	\N	f
6ee96519-8df8-4cf2-8e62-0538e2f2776a	$2a$08$Z3ZfK/XceZ/voGrO00AUue5I/wJ2Lkb2ekQlfEFCukEYBqY7a5E/C	89506808955	bregadze.zurab@yandex.ru	f	zurabbregadze	zura11	ФН2-12Б	Брегадзе Зураб Гелаевич	\N	\N	f
a7d51515-8435-467e-911c-365440e679b1	$2a$08$6mmWyaIOj7MW406zrY4BVeDVC3tvsny0qeBYjrcw4YMTygxAV8M7q	89852498044	dudchenko.kiri083@yandex.ru	f	TheGhostFox	ghostlyfox	МТ8-11Б	Дудченко Дудченко Николаевич	02e685fa-19e7-4f7e-8025-8dbfbee483a0	3	f
e4fde6e6-28b2-4c75-8a8b-14c35da34b4d	$2a$08$WlyhsfO7B8GodjtuCJ.1AuGAnBMPKRI9Kvc6nXiNABGvkXoIe9SGC	89771668165	morozow.4nrey@yandex.ru	f	LuciMorninStar	a_drew	РТ1-11	Морозов Андрей Игоревич	40355e30-8138-4e9e-a09c-6d4b9ae9f4aa	3	f
23b542a0-e275-4665-bc6e-4345f3011d3b	$2a$08$0vikwKywPXQQu6dKwe.lcumaAbNDU96DTOLUWNheg0bjMG3jNVwyO	89678244577	pavelpawlenko@ya.ru	f	niles_penrose	niles_penrose	РК6-11	павленко антон александрович	7b1a54ca-dbe6-40f6-bd58-955ed42987f4	3	f
cada0dec-025c-4e9e-8e53-66307bc7a249	$2a$08$YEEn4nV6vXfjL2pooYByY.wLOuFOxeHAxJxfzYl.AhFg7eM20isAy	89645617516	alya.mylnikova@yandex.ru	f	a_ma_ayla	aleksandra_mylnikova	РТ1-11	Мыльникова Александра Андреевна	40355e30-8138-4e9e-a09c-6d4b9ae9f4aa	2	f
5f89d7d3-499b-4b83-bad1-bd253fb22094	$2a$08$Es0tkXNB/XVk/l4zeYJrJ.c8HXB4hFaPmmmB9TXuzZmdEc9Iw3mQW	89779252075	kolbun.anna@yandex.ru	f	Bezumie333	bezumie666x	МТ8-11Б	Колбун Анна Андреевна	02e685fa-19e7-4f7e-8025-8dbfbee483a0	1	f
d4f046b7-b268-45ff-accc-cbf73ec54007	$2a$08$rH7rfUyuEZMlqtahyzIiyuyxkeF763jYze01Yu93Zc5oPs9FDm5C6	89160523694	terekhovva@ya.ru	f	trkhw	trkhw	Э3-13	Терехов Виталий Алексеевич	485a9afa-6f4c-457b-952a-203d110898b2	3	f
9c36a1e5-36ac-4d13-a139-fc83809b8d10	$2a$08$HBbH/HV.4rNGYOI.CMQifOnDdtPBUSpni/re73jnwMeX0ABtkJOku	89771793628	t13563378@mail.ru	f	bubalaka	bubalaka	ФН2-11Б	Татьяна Михайловна Панфилова	f9edfee0-2a1c-46c9-8518-d5a92b0cff1b	3	f
aaa4cd2d-44ae-4112-bbdd-6bfaeee917e5	$2a$08$0wFqFX392B0tNKdaEfOJzuYk34V.RYeun1e66o7q4s0xHRpBnbvpy	89153050206	lrsem@ya.ru	f	plhgx	leonid229	Э3-13	Семенов Леонид Романович	485a9afa-6f4c-457b-952a-203d110898b2	2	f
818fc880-452c-48c5-a96c-dd6bdb445e78	$2a$08$OP5NtyKe8lOf9NC5bjxBn.RHVyqoxOuXG16rp0weDP6s5u/09bRsG	89267917748	yar-fi@mail.ru	f	yarocabra	tranzefi	БМТ2-12	Чубаров Ярослав Валерьевич	300ef3ae-2e98-43cc-8151-a7f42980de56	1	f
f818ba60-a4a9-4553-9713-9c23ec3dcfe8	$2a$08$qyGgB8Ucymj64rxGxTcvsOYwjRo4F8oIsdiusaqfOSFYOU644Rwe2	8915-486-69-20	inkass.sv@gmail.com	f	laoffee	laoffee	СМ1-12	Соколова Светлана Валерьевна	2acc4cbe-c250-4830-a19e-8b8af134ce2b	3	f
72a26aad-95fe-4c28-bbd2-ca0664fb0d95	$2a$08$zs1GE6MfAINFAITtRtIcEekGqVhn2j8DnQao6CjNjEF8zhzcNTMeS	89324069735	4069735@gmail.com	f	Adrianovskaya_Yana	yana_adrian	МТ8-11Б	Адриановская Яна Олеговна	02e685fa-19e7-4f7e-8025-8dbfbee483a0	1	f
cacaaa87-9f1d-404f-88e5-d795e9f5e8a6	$2a$08$xQeeWstFA012Fo5/X3MicOma8.Rh4WdtT9.BgjrCE6xpx26OJwhP6	8 901 343 10 66	kolyapechkin88@gmail.com	f	pivnoybaronchik123	pivnoy____baron	СМ1-12	Серрано Сидоров Эбер	2acc4cbe-c250-4830-a19e-8b8af134ce2b	2	f
ab5d54e8-2e28-406e-89b4-c03dc7d95151	$2a$08$Fhen1shGMqi7b5U2Zy2ZL.9fii/7xVnYZ86zNZQkoFeRhSEZuCMpu	89152471758	karnazeeva05@gmail.com	f	Soniya_Sonia	ssoonniiaaa	ИБМ3-12Б	Карназеева Софья Олеговна	fd8040f4-8e1e-40fe-8012-0398c80c5324	1	f
30732d40-d384-4a47-bcea-99644092454e	$2a$08$ORrrZyM8fd9bQxB8uPIZ3e6H240Ybb96Gpz/maxBtcigzz6n5mHOe	89807137850	vlezko005@gmail.com	f	dmchk_jon	vlezko005	МТ11-13Б	Влезько Дмитрий Сергеевич	0c68c5b1-5185-497b-9a60-b1c1e1763b6e	1	f
c164acd8-a68b-4ea5-a7bf-e6d804d3199c	$2a$08$pgxS2JVSCWM84qOEgwi.E.PIID4u3LwFItQCZazhXwm0UBYrciyVa	89778770347	anko25082005@gmail.com	f	ehet0_0	ehet_0horat	РК6-14Б	Кошеленко Анна Андреевна	2b3b1e0a-1f6f-4f7d-a0da-7877dbf854f5	2	f
00ad9411-3f26-41e1-abf4-8ad605611fb6	$2a$08$HnkEKzgEM8zbNe.7.V0MDu74RWlRJg312h10keT0EjVMODyRoAG.q	89032385504	poleevea@gmail.com	f	elfalph	elfalph	МТ8-11Б	Козулина Полина Сергеевна	2b3b1e0a-1f6f-4f7d-a0da-7877dbf854f5	1	f
a155e7f4-664f-46d0-b3c3-b1c6ab121b82	$2a$08$YdOxBtgtdpKox3j9rkKjO..FrfMSzUnjVQIT2A3G981WqWyukviNG	8965-418-22-56	saa23u408@student.bmstu.ru	f	ADmi_OK	nowaynolove	ИУ5-12Б	Старкин Александр Андреевич	\N	\N	f
09c97de8-00a3-410c-b8aa-0ff87fd23b55	$2a$08$fekY1AUkivuJk5tvv99lcO3ems6D0A.d5CotKbVbqkqOAHSLKhs7y	89102013200	agarkova_ds@mail.ru	f	Dariaagg	dariaag	ИБМ3-11Б	Агаркова дарья сергеевна	\N	\N	f
337a32ad-a952-4b8f-a7bb-f567ad73565f	$2a$08$aVEzupI6SKafpiSxt3yO3.FL0MwwKh0u0zgWxxUX5E3NyP5dacwOq	89015696803	volkovnikolaj25@gmail.com	f	Sabitus	sab1tus	ИУ4-11	Городецкий Николай Андреевич	\N	\N	f
310198b3-c552-400e-b811-246e6992e815	$2a$08$on/QvU5vwDPGVPwAiLcs2.2aLeWPsTjYmLH1rVOAer0Ydm3bUgc2u	89052438924	maxim.zayka@outlook.com	f	v4mpus	v4mpus	СМ12-11	Зайка Максим Вадимович	\N	\N	f
4446d557-8461-4a2d-a6e0-de818c327693	$2a$08$dKDXs6K12xEeMiJHJrn2F.PdZmnD/WJ1zrf3YckDCUyp.Z11ajWJ2	89264810844	chizhovks@student.bmstu.ru	f	bu5y5leeper	bu5y_5leep	ИУ8-12	Чижов Константин Сергеевич	\N	\N	f
757f520a-e649-497e-878e-3f4781888dc4	$2a$08$1xws1.bnoPP1T5q4Gk6bd.M.o7oA9BvrMS8yWOGi/xkkSNkAcdnvi	89153785122	bobovav2504@gmail.com	f	Imurglitch	simpa_toru	МТ3-12	Бобова Виктория Юрьевна	\N	\N	f
e787920f-1c8b-431f-a9af-a438777dd84a	$2a$08$ExKWRCD2..j1Ctufta7cFeorYhbGGe3djCNqSqRLZ6Cp9xR/2EZBm	89393071363	scherbakov10000@gmail.com	f	ramade_inted	id283652621	СМ11-11Б	Щербаков Григорий Романович	\N	\N	f
e5203e1e-7762-4622-bc2d-d6af680c88fb	$2a$08$qlbY7KiLoL32iJePAUfX2.epLwT2MTDecGhywzpkZUOtV8YVOJSEq	89802724185	b.kompot2015@yandex.ru	f	bdinx	fanatviber	ФН11-12Б	Холматов Темур Муродбекович	\N	\N	f
a1f1f191-402c-4f02-b3c5-d64ffbf563b9	$2a$08$8pCNgfsfP7f6ngNL/eiOgenrh7i9eirGmuW51CjZCC2W1fqU5oC1O	89779567407	basovava@student.bmstu.ru	f	sugar_viky	vicki_tiki_tavi	ЮР-15	Басова Виктория Александровна	\N	\N	f
7f5e627b-5ca5-4256-9399-d66b119bf86b	$2a$08$YBkk3cIBE422.enV2BCfPOZrhEBPOyshAGSaPXuedwguQRFIza5mm	89833110359	knowalex7310@gmail.com	f	alexfir15	alesasha893	ИУ1-12	Фирсова Александра Сергеевна	\N	\N	f
88194d60-5200-4110-b6c9-649fe390824e	$2a$08$p9G2OQB5qjddOU86nAv9rO2PBk1nt6AT6XFFmh4d36RsB4v9pei3a	89160823542	bosarovakata@gmail.com	f	puuuummaa	puuuummaa	ИБМ3-12Б	Бошарова Екатерина Алексеевна	\N	\N	f
e3416369-69fa-4566-9b46-5ba9e8b74340	$2a$08$qDvGA8Sphe.RFY7aKS1G7.1rcmHW5qGNQ99u3iAepCIxChIrutse2	89050055755	nureevaangel@gmail.com	f	comoelsol	id745762483	АК2-11	Нуреева Ангелина Айратовна	\N	\N	f
8f00f249-287f-4118-a41c-2eb414d5fb1d	$2a$08$misDGkgVFjecToUm4iG5HeCDWYAuTWmzsjvS.pdQ6KOWIffXef2ri	89279751519	matvey.vershinin@gmail.com	f	reapman	reapman5885	ФН4-11Б	Вершинин Матвей Игоревич	\N	\N	f
80e1092c-91e1-44a5-b0cc-9a316b0d371b	$2a$08$C56raOUMhs3xTIPNA5g7Ae9.P35YIprcqYUdhSgk4f0nBbP8F0lka	8925-736 23 58	michael.lashkov@yandex.ru	f	okcid_vodoroda	okcid_vodoroda	СМ12-11	Лашков Михаил Дмитриевич	\N	\N	f
793c2bfe-e04a-45ff-aa9b-ad1e1d7dd42a	$2a$08$b2V1ZdKgikgE3Dhh2M/gnOGEtffge5sVhfjK37t2FDc/p0ArcY0oq	8 910 75311-01	cris200509@mail.ru	f	Fartovii232	fartovii232	ИУ8-12	Беляев Константин Станиславович	\N	\N	f
48204926-e31e-4117-87b8-78f575dab61a	$2a$08$/pU58K9HSunPNKjJT0JVLOc3WvrRnKR.AeHg9p9eE60B/4jJ47JIm	89100092903	alex.gaivoronskii1@gmail.com	f	miggers	miggers	ИУ4-12Б	Гайворонский Александр Иванович	\N	\N	f
a3ef6ac0-719b-466b-ac1e-e8dc99c092d2	$2a$08$QGQrQV04B6W/BtT4CNFWl.xUSF60ByIa4bWkp8bBb1xqp.q7apJ2m	89647856011	alina.begiyan@gmail.com	f	aalmora	aalmora	ИУ3-11Б	Бегиян Алина Норайровна	\N	\N	f
7558a03a-309f-4b53-9211-52a0e34eef45	$2a$08$uM1n4sbANzfyiFBwzJWx..AKZitZzS2NeTU2d4HnXPBot.I1WKZA.	89629777622	nikolaiveresoff@yandex.ru	f	Konnor_FLY	konnorxx	ФН11-12Б	Вересов Николай Алексеевич	\N	\N	f
6592dce3-3878-4f12-833e-ccbc5588cb51	$2a$08$TTGtWVL6GzkmgjFeXY7CTOgECtktTNtRpBIazWNV/w40uFjGWfaNe	89050005736	b.gayazov@list.ru	f	povykk	rrrrrrrrrrrrrrrpoprobuipovtori	ФН2-12Б	Гаязов Батыр Рависович	deb66347-3dcb-4a55-8e42-b28de913c9b8	2	f
9cc6a28d-4302-4a5e-a0a7-b4d99fd74204	$2a$08$X1D258cfrY5qD4FNQrkCZOPkhZ7W3XuIzcze1671zfi.vvNazIFDi	89160417774	elizavettapurtova@gmail.com	f	purtovaes	lizapurtova99	ФН2-12Б	Пуртова Елизавета Сергеевна	deb66347-3dcb-4a55-8e42-b28de913c9b8	1	f
484d0eca-b6b8-4e42-9f63-2d4f4b05e76a	$2a$08$aG4qTw/etwdEv1U337i0Lu2pAy38R.u/t6EZAH4PK8iPwLZYR2jc6	89639002343	sultan.gayazov13@mail.ru	f	Shtrudel7	shtrudel7	ФН2-12Б	Гаязов Султан Рависович	deb66347-3dcb-4a55-8e42-b28de913c9b8	2	f
09b0eafd-5f3f-4a22-b113-e3dde07184d5	$2a$08$lAoFTs2YNJdJRN4q/n3FHeL6MKLPNMssA7FIq/MIuPnFIwvzGOU3a	89652582851	kai23f278@student.bmstu.ru	f	Seny_ka	Se_nyka	ФН11-12Б	Кожухов Арсений Ильич	73ce80d7-2798-442f-918e-bf40511dfb07	1	f
3adb69a0-cdac-4e90-9318-2d1e1c6b0a53	$2a$08$8xDtfmlyWgfxTYrp94BqMuhsRpk8xfgMTvhXO5MGZlsyDxN9cIj4i	8 964 710 77 92	luschik2005@gmail.com	f	Voopoorazer	rock_n_roll_survivor	Э6-11	Лущик Андрей Григорьевич	4fb50154-1111-4f44-badb-e48f83e0d433	1	f
dd02a3fd-6f1c-47b4-9dea-20e732a4fdb9	$2a$08$JCxwDlOpueQw8fL1EUh0huUtKpsuJSIfFaqY1whKZEFtRDdjfwwG6	8912-201-84-85	julia.kuzminykh@mail.ru	f	ceacoacea	ceacoacea	ФН2-12Б	Кузьминых Юлия Владиславовна	73ce80d7-2798-442f-918e-bf40511dfb07	1	f
055f8322-be78-4cd6-839d-c401a7a0a21b	$2a$08$x7gefHwOCq9NClQ22EP37eqEvKqVThXOJ99wL4VXSqWJR5JbyfDzS	89277490099	ilinaea@student.bmstu.ru	f	sunny_katt	sunny__kat	РЛ6-11Б	Ильина Екатерина Александровна	589132a7-7358-41cf-833c-7a5bd9f8e1b9	1	f
065eb905-e775-4349-b615-baf43dc49928	$2a$08$5B40cyV4BQtUm8boN5sNm.oRYfW7L2il3vrvSiM9eJBUZbPSbniTu	89268518627	alexvtimof@gmail.com	f	strgzzr	id680976282	ИУ2-13	Тимофеев Алексей Валентинович	ccbdb464-37b4-4fa5-88e0-8a97ea84dceb	1	f
ead53b6f-4837-47cc-b08a-b92861e99e4e	$2a$08$B42xAOdl.A1gFYyFb9eCOuTSOnm.x9nxqAeakL6N4Z4NIVRtitPIK	89608226539	xaryok8@yandex.ru	f	zokha	madnoredboar	ЛТ6-15Б	Захар Игоревич Дударев	\N	\N	f
6628d608-17c2-4fa9-a66f-51a0cea12886	$2a$08$0Jwt4TjpalwnYj/NQ.TAwuZMDw0e46sxy.5vAIzEDsQN9ejMKybrC	89024440530	turdievihtiyar@gmail.com	f	turdievihtiyar	turdievihtiyar	ИУ5-12Б	Турдиев Ихтияр	\N	\N	f
b149c623-c24e-437a-ac4a-a6083dbd5a94	$2a$08$IqgHQPdrg/EG1iQ.OF9yqOY.tMorM5zhEEruabRL7n1PqP8ESIszG	8 910 48099-02	prokopenkoys@gmail.com	f	imleonvast	imleonvast	БМТ2-11	Прокопенко Яна Сергеевна	\N	\N	f
ede4a1a9-5601-4b6a-b946-55c52bf7ad1b	$2a$08$4yE4NFVEh7ZIbMwhW4dZvO5J4mYzpbuK7ds42PzwprJ0hzqygT3La	89435579867	eld475@gmail.com	f	telegram123	sdsudn	ИУ4-13Б	Ивано Иван Иванович	\N	\N	f
33e2d4d9-6e85-443a-bc80-58c719839034	$2a$08$rbZH0qkNQ.468GNwRo3oSOrNIIn1knVe/rtKFf6mLgLeqLdJJm5RW	89151071377	misha3052005@yandex.ru	f	chem1s3	chem1s3	МТ11-11Б	Чебыкин Михаил Сергеевич	\N	\N	f
9fefb0bc-9b5e-4a5b-9152-ea7d913ad545	$2a$08$H3YXxk.LYOm3N/RJeBHxsu6w0nCnyw/mS3fjlehMkz/ysgzpKXxrO	89250534011	fistovaas@bk.ru	f	firrrarii	firrrarii	СМ5-11Б	Фирстова Арина Сергеевна	572bfabc-f038-4482-97a6-cefd372cd700	1	f
e4c5a1ca-0127-48cb-bca3-a8f7f43078dc	$2a$08$lXyZJdRq0zidSP4CERqoAOCz9KFtetscum8N1AzY/L63yuyOftOz2	8916015-57-74	elizavetshcherbinina@gmail.com	f	eLiza_cuddlies	id242079298	ИУ9-12Б	Щербинина Елизавета Михайловна	\N	\N	f
065629fd-9441-4f7a-b010-ac590d78a5f4	$2a$08$1hHx68b/LeD/EV28ybpRd.hMmvilM6bycNAasjIbntZ25.Nz6zig6	8 950 67849-64	ihoniauks21@gmail.com	f	happy_yppah	ppunksnotdeadd	ИУ10-16	Суражевская Ингрид Натали	\N	\N	f
8cbd9dc8-01da-495d-82ef-d1735c65587e	$2a$08$Q.RLyEv21wvBTBRvdetssObvWAhP6y.cgJZFxFTj7IuxL2ptzsuQW	89772629930	sofiapleshkova@yandex.ru	f	gdeSofiia	id459419556	ФН11-12Б	Плешкова София Дмитриевна	\N	\N	f
829c5ca2-11d0-450d-92c4-f0a7e4cfadc3	$2a$08$2hDq6mrWxCbUHAfagmFDT.E9FSJ5ZYHDPeovoLdE5Qdy8m2k.5t/6	8916-865-77-49	sharova.mv@mail.ru	f	maren4ik	kvayou	РК6-12	Шарова Марина Владимировна	\N	\N	f
0a2656be-1cc8-4f60-ac25-9f28a0874363	$2a$08$dYuH3MiSFgpZbk.HQOKX4e8ycFPACeQ2xa0oKiS/iYpeowjxaYYma	89153270760	sony_orlova@mail.ru	f	Sofiarlv	sofia__rlv	БМТ1-12Б	Орлова Софья Андреевна	\N	\N	f
3273c6ec-7de0-4226-87c4-8c47681802b8	$2a$08$p1Udz3L8Dms2ynlcTguQIeXTStmB0JZoqL94Swa4mP2c1rlz7VzcW	89040669803	mishakr8987@gmail.com	f	dkevp8	mihakr	ИУ9-12Б	Краснобаев Михаил Сергеевич	\N	\N	f
ffa66dc3-89e6-4fde-a04b-404b40c404cd	$2a$08$gGxSBD6Ad2NfgcgGEFlF4uO8eyhQ9Vl9jecGHElpo0xh8CS77qkQi	89772780592	dorfyana@yandex.ru	f	dorfyanaa	dorfyana	ФН12-11Б	Дорфман Яна Эдуардовна	\N	\N	f
b2376188-822b-4a50-b469-77112176d26d	$2a$08$zFv3lhEL/CluowkbkQAmx.5qnVgaVAUqWzmLWRtlLJ8821DNE7oDS	89269227144	alexey.brysh@gmail.com	f	stepa_kakaha	muzik_tadzik	ИУ9-12Б	Брыш Алексей Олегович	\N	\N	f
7f13e8a0-ec2b-4620-9268-47b6d4908cb5	$2a$08$guuun42snA3oMqCnaNWMlu3UNMgybCRfvRNeVzesF3X.YLObuyMJC	8911-009-33-46	1artonaut1@gmail.com	f	art0naut	artonaut	ФН3-12Б	Ариткулов Артемий Владимирович	\N	\N	f
c7d4f223-db67-44df-844c-bb65939b236f	$2a$08$U2oDlPuNZ1MTDCkW/8k6YuIoFSUNcH8UVITanKRnSKIwWRsWToD02	8925 1106016	vovamatov@gmail.com	f	Matbl4	old_vobla	МТ8-11Б	Матов Владимир Витальевич	\N	\N	f
909c2721-6c10-4989-9640-62cfd99847dd	$2a$08$tBrW.2MO/zWSb3raCnVCruFnIk2gB/p3NReQB9RTaCLhytZpfIDCW	8 977 865 05 29	lpodlepenets@gmail.com	f	orange_lox	orange_lox	СМ1-11	Подлепенец Ларион Витальевич	7b146ffd-3941-4851-ab0a-e557884042d4	3	f
7d1b09aa-78ef-4fc5-8f5e-35803f50252d	$2a$08$cDho5mwSDLC9LSAPtVfWB.jc4DJhFZkUerg4yW3pcJWBjVP9JNhIG	89137151634	diana.barabash2005@yandex.com	f	Dianochka5901	id501387974	ИБМ3-14Б	Барабаш Диана Юрьевна	\N	\N	f
f6bf763f-3978-4b7f-bc29-0050c8fecd8c	$2a$08$imM11WMRosQydidt0iIUNuebLEl14X4l8jitAmbVL/zkxUmFwfPhW	89013660594	jmurj@vk.com	f	sattsrammu	jmurj	ИБМ4-13Б	Муртазина Юлия Вячеславовна	\N	\N	f
dce7b011-c2e0-4fc7-a736-b853fb4e1397	$2a$08$hVZkxFb01YUulJ.AUN2dZ./7Vmy9uiGegYCP.YUS5wTaKb1X0pihy	89202016818	iliya.pelikhov@gmail.com	f	ili_y_a	id416708602	ИУ3-11Б	Пелихов Илья Игоревич	\N	\N	f
92c22ba1-83d0-4310-88b7-d84d24d3d646	$2a$08$dKFP9P7ydGKvo7u3rW6HguS56yR0z9UxuZXRU1YsaxyLHb5Qx3vTW	8965-302-18-11	alexnoskov_44@mail.ru	f	lnoskov16	lnoskov16	ИУ5-12Б	Носков Алексей Александрович	\N	\N	f
627cb0dd-fbb5-4f10-9705-cfaba06fbce1	$2a$08$uEnSHz/UF0ttx2/lbO0.xuUHCVgsLP3mGTm7u51FQ9IR2ZLfg1lF6	896517739-01	miss.katenok3112@mail.ru	f	Death_0n_two_legs	deathontwolegs	МТ3-12	Зубкова Екатерина Дмитриевна	\N	\N	f
13d800b4-ed0b-49fd-9892-979867573555	$2a$08$uhloL9kLSMiNUjg1KdVhzOme9sYH7WQIIbaptGqD4Z0dpcXiGnjae	8926-377-90-33	alinamorein@gmail.com	f	aalmoree	aliiinxxss	Л4-12Б	Морейн Алина Анатольевна	28fd4455-c957-4065-97f6-847b632daf2b	3	f
48cd1f87-e9eb-4e0b-ab37-c23a131a8351	$2a$08$z1b9J7brQn7FYhXwISF/SuMxeME5Gm7Rj8ljQDSmrpjsFpdKAICPm	89253263691	r.karet@mail.ru	f	aruwoh	ramstik	ИУ3-12Б	Каретников Роман Вадимович	7b146ffd-3941-4851-ab0a-e557884042d4	1	f
4c33f85d-1f4c-400a-b0ae-787c54c20b54	$2a$08$EcHa8HB8wTpIdKSYU5whcO6tkWMqiYN5uF2vw6f04fKYK3l7U0TPW	8919-393-11-07	filatovaas@student.bmstu.ru	f	anyaffff	anyaaaaaf	ИБМ6-11Б	Филатова Анна Сергеевна	a614efba-69c4-43fb-a9a4-d70d8be4c365	1	f
c925f7ee-b3b4-48dd-b5a6-7db1cb48a6c7	$2a$08$rY7hl4xtPCTelbWDdeZxWOelXmpAhuBJPj6GvBfAH5l70lnrQK6R.	89154691181	terentii1029@mail.ru	f	yanterr	yanterr	СМ7-11Б	Попов Терентий Максимович	28fd4455-c957-4065-97f6-847b632daf2b	2	f
e98a3467-09ba-4d6b-825f-ce5efa4e7fe5	$2a$08$foEaBxoA/.4HaIX1Ku/zOulKkEtBeaxlPt2H/I8v.Uv878jJLmtOK	89263986108	pasha2005t10@gmail.com	f	Cum_monster3000	inarosos	ФН11-13Б	Увяткин Павел Михайлович	589132a7-7358-41cf-833c-7a5bd9f8e1b9	1	f
f6a3245f-46d1-4bf2-87e8-b77e37f7f211	$2a$08$BQsoPBTMcACQWuGb3pbkuex6F9TBBs/ruesua5jwWVLE0AZL4LfCW	8 964 71228-51	lizavoronkova20058@gmail.com	f	Leninzhivv	lizavetta_cccp	ЮР1-16	Воронкова Елизавета Максимовна	589132a7-7358-41cf-833c-7a5bd9f8e1b9	1	f
8552e35e-f417-4519-9112-5b79e2f40d3f	$2a$08$y.kmST1lDtO.wpQ81KC9TOOj1nVFr2MwuE5hF.dDFlbMICpp4DCiK	8 903 52477-75	oksiz2005@mail.ru	f	oksik_z_y	oksik_z	ИБМ6-14Б	Зайцева Оксана Юрьевна	4fb50154-1111-4f44-badb-e48f83e0d433	1	f
1290b992-004d-4458-a995-adff77fbb234	$2a$08$kdUAeRPlxS2QnEJZaloc5OranGijNCkVfj8.Gqc00Ehq5AmB9Kbby	89163407442	timekod6@gmail.com	f	KirZavy	kirushenechka	ИБМ6-14Б	Завьялов Кирилл Львович	4fb50154-1111-4f44-badb-e48f83e0d433	1	f
517ec3b8-436f-4b48-9e76-9923c49e2cde	$2a$08$JPgiTqEmQYatEeLjDdNO9u5zMyFw.jIkmo14/ifs0Jj7JxPOq1ICC	89232461519	jeeliz@mail.ru	f	Matthew2612	m.kalmykov2014	ИУ9-12Б	Калмыков Матвей Дмитриевич	\N	\N	f
48db253d-b410-47a4-9d62-82a7d44b0044	$2a$08$uEb9YO0dHidOKOJK7ENJm.NGKTTbWg7VHkeNhpAq5Wg503S.EY3yu	89146805502	king.76.04@mail.ru	f	changed_the_world	the_one_who_changed_this_world	Э2-11Б	Акимов Никита Андреевич	\N	\N	f
07dd7cd8-fc02-4e3a-a37f-8a0737ec9185	$2a$08$9Lva934QoUZGTLyBWN82dOlKIfcvoSkjwoppUnP5OBmQPTfgb6/.W	89777053043	timur.reychi@gmail.com	f	reychi_creatoruh	reymond_zero	СМ11-11Б	Желтяков Тимур Алекеевич	\N	\N	f
33e7a3a9-a959-47db-a4e2-635832569403	$2a$08$WfxE4Wm4YbnwRAA5FdJ5ReDBf8BDPs7Uq0KczGVWB18jkjpgwBl0O	8 985 95885-01	gosha_31@list.ru	f	Mafikur	id549492228	ИБМ6-11Б	Круглов Георгий Сергеевич	a614efba-69c4-43fb-a9a4-d70d8be4c365	1	f
ef5dd871-9477-46e0-9d4c-6e0e3974a4b4	$2a$08$y28RUWiHO1f372plewK79uyb2vPmn5x/CosOIbujGWxHKBhKVI57.	89852297433	smurov_arsenii@mail.ru	f	vr5bac	mrmarsi	БМТ2-12Б	Смуров Арсений Григорьевич	7b146ffd-3941-4851-ab0a-e557884042d4	2	f
fce5bf6c-7413-4cd8-b938-26e18ab0bd4a	$2a$08$kgZhzOzGFmJ9NNX3mHanguXCfFCE6zNDh7rv9y9pT4tDDGYlecMxS	89653147906	ewdokimovamaria@yandex.ru	f	kak_hosh	masha00342	ИУ10-13	Евдокимова Мария Константиновна	9b7aeb4b-946b-4491-957a-b10688ac71d5	3	f
bbf569bd-4fab-478f-a6bb-bbae284deafa	$2a$08$rGhC2bFMDWHaEtjKUpx1PeX0XtgXwAc3k6A.d6lVefFSjjyUWa7am	8 920 3665752	anastasiazueva37@gmail.com	f	l_Neoly_l	lnustyal	Э5-11Б	Зуева Анастасия Владимировна	9f8e20ed-c2b2-4f77-9d95-ca22c7ad82d3	1	f
9ce16d41-2455-4ed1-a8b2-9a2c3f163845	$2a$08$1uggyGj2sEQRdJfUqf3v6e1dH8ZXShJm4dgQcfTTXT83LrMq4J/cy	89874818675	milaguke@gmail.com	f	al_lien	milaguke	Э7-11	Ганеева Милана Артуровна	16956eba-5185-4ba2-bce2-a8239d6a89a9	3	f
dd2e3c4b-3a96-4e0b-8e99-0b69acdecf94	$2a$08$CbEt9nhY9zB7H.1UagBBZe7Z5CpnPuLOaMq4I4PBgLNrZn4NLUUyK	8 968 409 80-54	shaea23u670@student.bmstu.ru	f	anastas9999	n.shishkina0	ИУ3-11Б	Шишкина Анастасия Эдуардовна	\N	\N	f
50c791e2-8256-4aa0-96ec-c91cff931d8b	$2a$08$UuZhHXtcdy8SrFQ3p8Ay9ei5hEhRqgBiUJejWIkn2/f8GYrfx2GZO	89772871518	atv23yu104@student.bmstu.ru	f	ATTTTHL	id388262141	ЮР1-16	Абазов Тимур Васильевич	\N	\N	f
695d4a44-5ec4-4ae6-8ba2-b12d15b3008c	$2a$08$P8agINZdwBWnEA7qy3QZB.S2H/tBjOYJ2Wmchim.92AToxfcKq0fS	89851981430	lihorse_ch@rambler.ru	f	lihorse	lihorse	ИБМ7-13Б	Чернова Елизавета Германовна	\N	\N	f
c85b2934-0e28-4b4c-935f-1ec8d5f75b8d	$2a$08$1b4EPLKpocLBbUBdjrujF.IHHKaf8IPEK/vEVUZzUI.hyt./NzO3y	8 903 65069-43	ouw903@gmail.com	f	chetakslozhno	i_hlebushek	СМ13-13	Возиянова Владислава Владимировна	\N	\N	f
a22543b8-ceca-441f-b176-0935c9b85a5c	$2a$08$Lk/3I/p8wSFoU8PJiFmnBOmqDfTs3LzS0RIZpGu6dV.W47iqrVZnO	89857465680	a-lonka@yandex.ru	f	nnellaaa	id801310122	Э9-11Б	Алексеева Алена Леонидовна	\N	\N	f
3b338c79-7326-43ac-8a3b-48640cafd46e	$2a$08$MJCMDIEH1z2vEAtVidWduOZUeU0w.7KLOd3XwTaBME./teywB4sF2	89101422186	egelfanova@mail.ru	f	liiiiizzka	liiiizzzzka	СМ13-13	Елизавета Максимовна Гельфанова	\N	\N	f
d1939fba-9130-4695-96bb-aeb161241a93	$2a$08$Nshk0axiwt5zrhwm/TDR/u4sg2XP1a7CTS8AHybWeuhDzb8E7Jiyu	89112163826	xhklu@mail.ru	f	kilkil94	marus_37	СМ13-13	Волкова Мария Максимовна	\N	\N	f
34d9d2c4-1136-4bda-adf1-5c42b696c8db	$2a$08$oiYifv.ywwXJ9d.TNBhUI.c89sKYDUe1pZedqRlN9Tuq6ZO65ROj6	89995554148	mkizogyan@bk.ru	f	KizMary	mkizog	РК6-16Б	Кизогян Мери Артуровна	\N	\N	f
115eea3d-f94a-4210-a13d-5c85f650c52d	$2a$08$A9DyyewOAbOUygdl1g6QSu3zGge1gP5BRMPDfzmGr.NDScA5baaBG	8 960 417 07 15	elizabet.magomedova29@gmail.com	f	HolerovaYazvaGadukovna	elizabeth_ehh	ФН12-11Б	Магомедова Элизабет Камиловна	\N	\N	f
9ce1bb72-6252-440c-94e2-06f93e7a64fe	$2a$08$lVJhovj2YjWVJsmiaLC4PONqLg.97I8L0W4qdP23F/x4vp9g61GXW	89854338473	sandrakorneva73@gmail.com	f	kornevatoday	kornevatoday	РЛ6-11	Корнева Александра Васильевна	\N	\N	f
c0a2920c-a0c4-43c4-b816-30a4139694ba	$2a$08$XzWQaf3hA4DRiGo99DlaPuq5jPpbdT/VOHeuTyIUQ.he3fpb250cm	89779365288	gritsenkoyua@student.bmstu.ru	f	yuna_gri	yunona_gritsenko	ЮР-14	Гриценко Юнона	\N	\N	f
ad024941-c03a-4d5b-826f-b4b8597a5e69	$2a$08$GbLtHYYto1IuGg.x0WwozezfMzPqH.XicgxsKHSTHknQZNyywTfzG	89859578458	lbagnyuk@yandex.ru	f	lev_eat	lev937	ИУ10-15	Багнюк Лев Михайлович	\N	\N	f
5d7c734a-3ae7-4881-b0c8-a80fce8511d2	$2a$08$zTt5QzQBzvMbTfupkC8gyeCJtbozTwWaoicvTjzzJgsuHeLM.uP42	8 916 056 24 61	katya.ertseva@yandex.ru	f	i_1_2_4_q_ok	kateridze_n	ФН2-12Б	Ерцева Екатерина Анатольевна	73ce80d7-2798-442f-918e-bf40511dfb07	2	f
ecb1dddf-0a44-475e-8d80-13673b2cc980	$2a$08$rEV/WZvpdoiYwjFQBe1G4.xLDgBa0P/f71StGZ5kK6OrsrQx9R1xe	89250950577	tkyghl4@gmail.com	f	kaiswify	idtryhokage	ЮР1-14	Молчанова Наталья Олеговна	712edd72-aed6-4ec5-a521-cf4a8bbbf4f7	3	f
cda2b68c-8117-4050-976c-b89b14f51e77	$2a$08$4foxMNZ3BW03H7nl8e7DOusXp6kbGXAMzXtWuORp.KbtyTKP4LFSq	898221821-15	yana.gross.05@mail.ru	f	nuit_etoilee	nuit_etoilee	ЮР1-14	Гросс Яна Александровна	712edd72-aed6-4ec5-a521-cf4a8bbbf4f7	1	f
5f623061-b8a7-4679-9946-b76f5630df18	$2a$08$ZUQ3q5QxvCjOxLTfsctIo.adeZLdfmwxQFcn7/veeO0tnHTRKF5Vm	89854590301	yuliakov2550@gmail.com	f	jvlsi	zapiv0m	Э3-13	Ковалева Юлия Глебовна	485a9afa-6f4c-457b-952a-203d110898b2	2	f
683c2ffb-f567-4832-b944-7096ca55ed92	$2a$08$qro.gKHmAdhC8MUeZFbzM.r9QUPCKTse6JNtEPJGl4FScZwFndmaG	89266250460	mariakuznecova304@gmail.com	f	m_kuz_a	id547190136	ИБМ6-13Б	Кузнецова Мария Алексеевна	ccbdb464-37b4-4fa5-88e0-8a97ea84dceb	3	f
26775a35-5d44-4f4b-a060-6944f5556c57	$2a$08$r9fCjRT0xjiQqdkIlDHvruksYgXEOWnTZPIbsbGp8YUzSVppHt2eq	89687657256	leon209921@yandex.ru	f	leorended	leorended	ИБМ7-12Б	Хабибуллин Леонид Маратович	ccbdb464-37b4-4fa5-88e0-8a97ea84dceb	1	f
16bc730c-c937-4567-9142-b0fc292d5e56	$2a$08$gEWqpykYx89repZPb1y4ueEijdNmDpJwoUnus.J5oNtv3Grx18sJC	8 968 79507-25	lizak.2006@mail.ru	f	ia_durachok	ia_tamara	МТ11-13Б	Козлова Елизавета Алексеевна	\N	\N	f
80fe8459-bfd4-4292-92e7-d308167f9886	$2a$08$b7nGDUaAuXcaZWkh/LWmhOI23PcWjvohhxGCA9b9kiamt/36tFw4a	89057547552	dovbyshliza@gmail.com	f	ElizavetaDD	id533616996	СГН2-11Б	Довбыш Елизавета Дмитриевна	\N	\N	f
f6bce436-05c6-4c94-9f98-df1645408e17	$2a$08$ZHKNzs2SkzXEnEfswJmRnugASopQg3BqDGX.QE99xx2gHp4l8aCvW	89048094666	5levinivan@gmail.com	f	ilevin5	vanya_levin5	РЛ1-13	Левин Иван Владимирович	\N	\N	f
82c74dd9-4cc5-48e7-bd0a-f74004ce3a5a	$2a$08$PmJuiy6LO00NT6zeeP4tbeLrQV9nX.wh43tE4/GMlYKxF41l41yFm	89015557883	salatoyuniga@gmail.com	f	Lysyshash	id815753505	СМ10-11Б	Караванов Александр Михайлович	\N	\N	f
ccee52d4-b38f-491f-9ecc-d28d8b8901e9	$2a$08$FCM1.oq4qsITzr6QNlCoLeJsKwoIhXkykybTyQEBe.IPuaKxAhoL6	8985-201-98-64	jmitrofanova2005@gmail.com	f	yulenokk	hasnuuts	СГН3-12Б	Митрофанова Юлия Александровна	\N	\N	f
9d31137b-f5cf-4fc8-95c8-dcdf14a1973d	$2a$08$6FweSRlSB15z.ezPQCLEwuWpjhwMbEkD8M8pif5c.dL6eSsxfhBQC	89197729700	maksevaviktoria@gmail.com	f	v_u_mi	vickkusha	СМ12-11	Мякшева Виктория Юрьевна	\N	\N	f
7d6f23fc-2967-44ce-a373-a3f26ab68bda	$2a$08$OkczYhyhr6u8SnSldHMq1O82..xc8MEGPr8gTBFyKf4BBjlMbmF8i	89178175999	diana.a.ilina@gmail.com	f	diacould	dianadinadia	БМТ2-12Б	Ильина Диана Артемовна	32cc7a99-1f39-4dbb-a160-e293a14838a1	1	f
90115d90-fc11-433c-95e5-5abde5d3ecfe	$2a$08$6g1ao6zinM6W8I4TA1jpPOrbUm/jkgdNvEym43OMrGs3WVaL.Eky6	89684910028	nikanna05@mail.ru	f	guseniza_vupsen	id538290427	Э9-11Б	Никишина Анна Дмитриевна	6b2934e3-98fb-4a5a-9138-35a62df992b9	2	f
15cce29a-4d8b-450a-94a1-f482d07323d1	$2a$08$LFx6OjKtw23I1sxTUrI./eQpYpoNuRmPPyGKDwXuoMTme/ZH3KkOW	89061443206	dartlednevi@yandex.ru	f	VZ_ID_3994Z23	obi1fet	МТ2-11	Леднев Иван Александрович	40355e30-8138-4e9e-a09c-6d4b9ae9f4aa	1	f
2c391e40-3139-423f-82a1-78888f46cba4	$2a$08$SZ1BoWdwI2sLCSm185MTt.V5276bI60BqemR.mqQnlyf2Yoi9n4Fy	89774195144	lilpo2518929@gmail.com	f	polina_osd	polina_osd	Э5-11Б	Судакова Полина Константиновна	9f8e20ed-c2b2-4f77-9d95-ca22c7ad82d3	3	f
6ba71053-9eb7-452b-8c92-d112c92905fb	$2a$08$NcSAXV9WOT7RdEJDQD8Ci.wLJ.RhwG1pkRSkOG3Tpu3OXMSSGHVfK	89229912292	zykov.s2006@mail.ru	f	sibirega	ruslietuviu	ИУ10-13	Зыков Сергей Павлович	9b7aeb4b-946b-4491-957a-b10688ac71d5	2	f
be3873a0-58b3-4df7-8c49-7aa6efce576c	$2a$08$DImnaWY5yt/KfsOsNq15tuM7diWOUV96sqXPp6kNbnZWsw.V08DeO	89886991478	aliev2002zybair@gmail.com	f	GOSTb_TG	zybairaliev	СМ9-12	Алиев Зубаир Омардибирович	9f8e20ed-c2b2-4f77-9d95-ca22c7ad82d3	1	f
f25b7e58-9e38-4f45-b142-2a4fb7db4edc	$2a$08$nru9fkr/Q/1WLT.LmM.ysexF.SDdgXmOwQYVrghCsDNXvQxEmOaXO	89671273907	dmtsheva07@gmail.com	f	ht_asher	ht_asher	ИУ10-11	Щербина Дмитрий Игоревич	9b7aeb4b-946b-4491-957a-b10688ac71d5	2	f
1efc9e9c-9354-4e0f-9881-a0d0f7ba02e8	$2a$08$H17Aa5sXX8n9HQEEexyVeOVPnWM0q8Poiq69.G6wsOa1Dh0SWtKn.	89164824682	esenovavr@student.bmstu.ru	f	vesenchic	vesenchic	РЛ6-11Б	Есенова Василиса Руслановна	2d40c471-5ebc-4470-a419-97d182c2b096	2	f
bda2b6a9-9644-49ef-84e9-4b120de62b85	$2a$08$nOUxuLjBQFqx7Fcmr3zoMO3j/PVR9t/U8sKFwe6x9iTphI7y8Va5G	8933-339-23-65	twincnaxore@mail.ru	f	twincaxor	twincaxor	Э9-12Б	Дубинина Екатерина Владимировна	\N	\N	f
2fab1e56-3a8d-437e-b20d-2072643d2ef8	$2a$08$ACIFcgcXTEl7veyqgygn/ODYnVJJl2.KUZiBW5dy/DJEEfhDWL2vG	.	.	f	.	.	.	.	\N	\N	f
484f8b49-f02c-4204-8c49-f9897f691d5f	$2a$08$vjla2NUxlDmAiJU6HwUJ0OEwtOWYK1rTkWbTlPwKExrH7/kUv7M5G	8968-409-80-54	shishkina.anastasia05@yandex.ru	f	anastas9999	n.shishkina0	ИУ3-11Б	Шишкина Анастасия Элуардовна	\N	\N	f
84b38673-7e31-4feb-a9b7-2dd7ef01945a	$2a$08$qRaFwUSHymgBu8ME4UPwUOcFl4cJsswqP9wTGwTV3yYBieRGVdrOi	89683572908	abp27@mail.ru	f	moscow_toy	bogdanov_moskow	СМ7-13Б	Богданов Андрей Петрович	\N	\N	f
41d31c87-2dd6-49c9-ab18-df0ad1c29c6f	$2a$08$1FNJY86eU1IcvuHffifKX.axBaQvQRvmLfIhiOKErP2MFhKRpPTM6	89166249568	i@kiralpatov.ru	f	insomnii_aa	insomnii_aa	СМ1-111	Алпатов Кирилл Викторович	\N	\N	f
19f92c4f-40b5-4556-a852-793e9412fcd7	$2a$08$053q8qRPF09xteh5qdb1O.s.JyrjrcZeQsVtz0vYfygfY43pR.AVi	89997212134	89997212134@mail.ru	f	GoroSheK_00	goroshek_goroshek	МТ11-12Б	Федоров Сергей Викторович	\N	\N	f
10a7109a-d683-4259-ab51-7057b6f62635	$2a$08$reDkw0n5knQ2vwI91wEpK.nK/s0b03K1AK8J.zQY7wUGbaIf2XcBu	89165499311	mr.alex1710@yandex.ru	f	Nhoro283	id815297146	РЛ2-13	Макаров Александр Михайлович	\N	\N	f
dfe37018-f1b8-42d3-b73e-7becf427e259	$2a$08$iBSJrOe8F8R.Z5tWjiINzuJY3J19BFZkllOIF7whQtpGQ9AZMKVZa	89257383233	kudriavcevvlad2005@mail.ru	f	White_Bread_1	vladkudryavtsev5	СМ9-12	Кудрявцев Владислав Александрович	\N	\N	f
c7bc9a5e-0e73-4823-8dfb-48bdcc58fb08	$2a$08$vtdJ3H2khiTT1/IJYBZyJuxV1..Ra0VIYYAd/FnQEGubzYMq/oCKW	89999121250	marydub50@gmail.com	f	nastasiiia264	nastasia___egorova	ИБМ2-12Б	Егорова Анастасия Александровна	\N	\N	f
1bac3b6a-b9e9-4d8b-bfe2-b3786266e9e3	$2a$08$DWQp30FJiSzYEoRhpLhh1ec46m9HIdjlsr8hpmxXAMEqk9iwbSUkm	89526653446	strugovsikovartem@gmail.com	f	inzshenegr	inzshenegr	СМ9-12	Струговщиков Артем Сергеевич	\N	\N	f
09479e03-f26e-41f2-b568-90e29f209c5c	$2a$08$RcmU0iJRJ0kusUgZaqaM../CWXPyMdtMMadLy.Oyaw43AWWuz0l2y	89196964501	azalika05@mail.ru	f	azalikkka	azaliya21	ИБМ3-14	Нургалиева Азалия Азатовна	\N	\N	f
df605ba3-db8c-48a0-8323-68fe78df4b06	$2a$08$M04NupPDUFEwpY.L1/ViN.725yaX3UNcaZl9dVmi/Mre6Q2b6J6.6	8 916 07651-31	eshuidina@gmail.com	f	Sayo_una_ra	lousama	МТ7-11	Шуйдина Елизавета Андреевна	\N	\N	f
53d258cb-f02a-4752-a74a-8d7f0083554b	$2a$08$utR9m6dn6.Hk75LqnmY6UuVOnSLHuCKcRxPNmIEwmREDfpkZ9WlNe	8903-720-42-26	serchainikov@gmail.com	f	tomastu	0	СМ7-13Б	Чайников Сергей Вячеславович	\N	\N	f
3e3a8086-23d0-47a8-bb56-63ac4186dcbb	$2a$08$16d2lezlDFXiqAFwqCFZkeiLg.QvyiqgtqzNIb11CK7KXzzLXc9S2	89254175908	evgenji2609@gmail.com	f	nukacake	nukacake	ФН2-12Б	Усов Евгений Максимович	deb66347-3dcb-4a55-8e42-b28de913c9b8	1	f
0ff7642b-df7d-40f7-a04b-c255059aaa98	$2a$08$6GGkKfTzyhVPqrD8kOA0eecCX3DjZCmhYdmyu/cZ5xQlOLUIfWlK6	89263735445	mironenkovam@student.bmstu.ru	f	SEENCHAA	seencha	ИУ5-11Б	Мироненков Арсений Максимович	7b146ffd-3941-4851-ab0a-e557884042d4	1	f
2e946b60-9ad7-44de-aeb2-68439e00cbb4	$2a$08$FMbl4gcdOia0iooOqqnBhuVb5y5.thCEazqShxGrwWSzIr5Bopkne	89510841882	30072005v@mail.ru	f	Qwertyasdfghdgz	id427256522	Э2-11Б	Плохих Виолетта Сергеевна	9e61f1aa-978a-461a-90f6-99808580d083	1	f
c43dcc45-d7dc-4a81-b7cf-f5b0694d4bd5	$2a$08$E1UcsSND3/D50nGE.cPXqOhUTENF6N2drMjWUhQ50Td64mxWSFMTu	89264644645	yershovama@gmail.com	f	atirag_ram	yersh_junior	Э4-12Б	Ершова Маргарита Алексеевна	6fa56c83-f29f-4909-9d63-ed00152c144b	1	f
ef1ae43f-a56d-448c-b8a1-bb97f88e7aa9	$2a$08$7tVPHBPr.uhW1J.bbyhIU.pOalAvM4lU4FOAMTqCJBkbI9Dj9q1ty	8123-456-78-90	legends@bmstu.tu	f	legends_bmstu	legends_bmstu	ОЭ2-11	Иванов Иван Иванович	\N	\N	f
0a34584c-1a5b-41e6-b614-d793b10a135e	$2a$08$uP2Z6Ytda9jpPDcInFIL4ukg4Awfju5EAVzj4z3VFMibKiG2efnAO	89214926967	gulutaelizaveta@mail.ru	f	elizafox17	17eliza	МТ4-11Б	Гулюта Елизавета Сергеевна	\N	\N	f
e0202039-9543-4237-8868-aeafeb07d78b	$2a$08$bUWgeYOcRjz1UhhoM9NlneY4WXyvU77Pnsx6tVi.ZGhp/dQQVrVye	89857487526	anpalma328@gmail.com	f	anpalmaa	anpalmaa	СМ12-11	Мамонов Андрей Павлович	\N	\N	f
64e9fe87-69be-4db5-8117-c6fe52fd1f67	$2a$08$CmDn5Xt9sptrrGeaAyN0TugA/Q11EzmQ0eFR/p0ckTl4eUde7qwgq	89689838622	lashkovadaria1203@gmail.com	f	dariohaa	darioha	СМ3-11	Лашкова Дарья Валерьевна	3155c953-27df-4bf3-93a7-3c3cce34deb3	3	f
e0901d5f-2f0f-40bd-92af-d88f1fa18037	$2a$08$Cz9gX5Yj/kUYspMbQce4zOPlGs8fiahwEUY8RZlEQISp5nvbM3D1q	89836225575	6225575eea@gmail.com	f	avorogetak	avorogetak	Э9-11Б	Егорова Екатерина Андреевна	\N	\N	f
c83124a5-d562-4921-942f-b593c3e0fde9	$2a$08$DUkR346bmrfqSeSVf/W0uOd7ktM.cJPDy7eAhfPg4BNHCk8d8Uvgu	89014266560	sergey-belenkov00@mail.ru	f	hellcrit	sojoyless	МТ3-11	Беленьков Сергей Александрович	\N	\N	f
a177b716-1629-4938-8821-29d89ca2dc4f	$2a$08$Jh4OtDRI.p0Vrjh09b5dwuSwyZJesqhE1MAk1WrAIAY3Dbot42wue	89688761040	kpomu@bk.ru	f	linaaamiii	kafka_uwu	РК6-16Б	Тежикайте Каролина Руслановна	1c6d8c8a-9bda-4ec8-bbf6-8dd0b83b98d6	3	f
2bd1e691-5d14-4f5d-b69c-c0b843eaede4	$2a$08$t.cRiwVvDACjtqu3MoGY7OJ6zzVgwLmng0vb2643mMgG4vODKmAAy	89274200676	sharovaliya14@gmail.com	f	kinoenemo	id454380261	БМТ2-12Б	Шарова Лия Ренатовна	32cc7a99-1f39-4dbb-a160-e293a14838a1	1	f
45e1f1d5-0cb8-4bc4-8ec9-81afa959198f	$2a$08$tpZuB9iJve2.4blokdD9G./Kx.X0CkCdPlx8UoVtiydGD4srICHpm	89252725170	maxim.tomik2@gmail.com	f	Manul707	manul707	РК6-16Б	Томик Максим Олегович	1c6d8c8a-9bda-4ec8-bbf6-8dd0b83b98d6	2	f
ba7e91c9-fc78-4ffd-8edf-b6b2534e2c95	$2a$08$tqyUxFI9b7wlu7Ld95I6OeRyuZK7OgfcI1dRvBoIkAJyl7jEfuQtm	89152764269	telekhav@student.bmstu.ru	f	motorrbikee	motorrbikee	РТ4-11	Телех Артем Викторович	1c6d8c8a-9bda-4ec8-bbf6-8dd0b83b98d6	2	f
a374adc9-5ec6-46c3-bdb0-388a8af4b05c	$2a$08$jJX0vY1i/wNL9oWKYO1/tuLU3KqqWN8sSXS.7D//mETmiPw7dU1E6	89064900993	ps4.stas@yandex.ru	f	moldfgh	stasrck	МТ8-11	Жданов Станислав Сергеевич	02e685fa-19e7-4f7e-8025-8dbfbee483a0	2	f
7dd63709-0bab-49f0-9809-8e9c51898f1a	$2a$08$2v12htqk6rA8vVCyw2RK7.5XGZU6BUtabg62hRBvS2ODM1.diqPYa	8 919 06744-32	rozkovaviktoria11@gmail.com	f	vikkkline	vikkkline	Э7-12	Рожкова Виктория Александровна	17beb218-e4fc-4644-935f-cfbcb462b0ec	2	f
48334108-53d4-4e16-88be-269e0ef81d46	$2a$08$z7QfXdDmdWACwbwtOaizu.QInj6BZAvIeDq77rEmaa8Q5zozWR6vu	89653002549	verlev1111@gmail.com	f	vovaverle	vverle	РК6-12Б	Верле Владимир Сергеевич	572bfabc-f038-4482-97a6-cefd372cd700	1	f
4cf54b4c-61c0-486e-a4e5-0fcf06fe6e1b	$2a$08$VonrebSoUAoKmORM18svTu/b.b0Q.WpZKvqTM1BRe4Np0yzbuX/D2	89002145343	shinkarevapv@student.bmstu.ru	f	polltergayst	po11tergeist	ФН11-13Б	Шинкарева Полина Вячеславовна	bd9782f4-4e81-41b1-b6a6-be7bd5bfde4f	1	f
5c727dd5-5cc0-4c88-8f48-0256b8d05467	$2a$08$Wk3dT421tBGPN5SEl/e6uuwqzfekKJ9eMONnsLK1uuc3vvypaElyK	89774770214	volkovlad342@gmail.com	f	Vlad13300	just_a_men16	СМ10-11Б	Попов Владислав Игоревич	23845812-ffff-45e2-94e8-08b2581f1be2	3	f
81d3dbc4-96c3-400c-bf7f-62fc1a4184b4	$2a$08$LMJRYVq20Zll/e40ArerqelgVnzKHRJAWK2CEMqzlmp3XdG851iT.	89774830458	indirikbahshillaeva@yandex.ru	f	minimogirl_13	minimogirl_13	Э7-11	Бахшиллаева Индира Баходиржановна	\N	\N	f
066bec04-cc50-475f-83fa-1a6e1e0a5c8c	$2a$08$2qvzRdvT35ncb7wwfrpNIOSfCTeGak1lNmqIZ4D/7qUuPEUCkSBDe	89154342108	sharafutdinovazim@gmail.com	f	xuxuxu_xa	hs643	БМТ1-13Б	Шарафутдинов Азим Тимурович	\N	\N	f
de1a01ed-e79e-47f8-a7f0-c805b92596c4	$2a$08$PdFUuDVOsuXJw9xPZIhX7OGRhN1kK4UQLtr2acgyhK42N1yDN7lxm	89969573179	turzin.dmitrei@gmail.com	f	dmit_r18	guessmyname	ИБМ4-13Б	Турзин Дмитрий Александрович	\N	\N	f
ceb9a4da-ee85-4dea-bf6e-1af03c783b3b	$2a$08$by3XrlORTp9uTlYXupnu5.SFdZjKlLgCDjG3sWk9QRAJXB094iBqG	8 951 86354-80	nik20050513@gmail.com	f	NikBesedin	id491572246	Э3-13	Беседин Никита Дмитриевич	\N	\N	f
e1fedb4c-7cd1-4479-873e-52259379ede4	$2a$08$Ujd.ihREaCGHLWOtjXiQ5OdfECpZfQkEwT0dGQ10EH/0PwHLDVBG2	89629966932	alexanderbabenko75@gmail.com	f	miller_79	m1ller79	АК1-12	Бабенко Александр Андреевич	\N	\N	f
9df07869-b373-45b6-b63e-6ed01642958e	$2a$08$epxAd9/nUqb5db54NFw6v./BXkCqvhykt4.q2VE.9ZqKk1nbW.sEC	89161525465	laa23m424@student.bmstu.ru	f	tulencheeek	tuleenchik	СМ12-11	Логинова Анна Алексеевна	\N	\N	f
f3e98fc7-6c3c-4f39-94f7-a4bf2a9d3f65	$2a$08$ZaplwLm6abr6pbb.roGtHux5t5692BMLgOHB/4z2xavC2DMThPwJe	8 977 75256-91	kirillpluhin110@gmail.com	f	kirillplu	id506548348	ИУ10-14	Плюхин Кирилл Сергеевич	\N	\N	f
47757f3b-f65d-45ae-9018-1e078b7f2c52	$2a$08$HELzvnQAdEBekHyWHwChZuW0yFvIrMSwwbTW5kiFeUXsURZ3VLcAq	8 977 57137-12	alexander.pliukhin@mail.ru	f	plulex0	id507639252	ИУ10-14	Плюхин Александр Сергеевич	\N	\N	f
2b43570b-4f24-4269-9bc0-4495bd531363	$2a$08$9XznFc8LJ8Pamnpwxavb3e3uNAT6QRQC/OLMCixX913H8apxkDs0q	89869048476	arseniyaa2005@mail.ru	f	arsecks	flyfuck887	МТ8-11	Андреев Арсений Алексеевич	\N	\N	f
787ba7fe-cd47-4c20-aa04-feb82b0135bb	$2a$08$HzNBDvqaW7DVgArosSa0Kuvh3SBO7yKkjrio2mifr32qxjvRJQVdu	89257024936	zajsevaalina741@gmail.com	f	resses_time	id500243708	СМ2-11	Зайцева Алина Максимовна	\N	\N	f
35dc1e0b-9340-4cb2-8966-64f56d9e39b7	$2a$08$grg7u/FGOgQDE5B20k8irOxSnDNxztbBVtF2aNg/YjT20m71UFe2.	89663353440	romanoganesyan13@gmail.com	f	romaigne	e63amgnesyan	ФН11-14Б	Оганесян Роман Ваганович	\N	\N	f
b0fed02d-6e13-408d-8ce0-f4d0f855cdce	$2a$08$prV3/NoyV8tIfmRjw0ejqesIww.TvizR6dmItUFTOxU//mY9NMG/S	8902-274-16-31	mariailinyh612@gmail.com	f	maria_ilinyh	maria_ilinyh	ИБМ7-12	Ильиных Мария Павловна	\N	\N	f
51dd9df9-ea88-4831-a085-2a9fb4d0003e	$2a$08$GwLgQ7sx1uWjvr3ixAaPO.B9Kg1hA7UBvvN55vMDITRK1Pb7nVDjy	8 999 67237-31	ana_sta_sy@mail.ru	f	liemelly	nas_ta_sy	СМ3-11	Бибикова Анастасия Юрьевна	3155c953-27df-4bf3-93a7-3c3cce34deb3	1	f
aa4d14dc-b01a-478a-b0db-0dfe02ead47b	$2a$08$A9onh1Rue/v0QmlHzzepPebKyBiKJfc3C3uy9lmUb.SoAVg6OokVy	8925-236-44-21	vania.gogolew@gmail.com	f	vangglw	ivan_gogolew	ФН4-12Б	Гоголев Иван Максимович	b9b02aae-0db5-4970-961c-9a5a198f7ee7	3	f
87ecbd7e-4f79-461d-86aa-a35dfa4ef394	$2a$08$nZ4SZYTRHJZ8Ah3uwHt.4usnNb93m8MrjqW6pf6dYtTZGARf9RK.O	8926-259-90-09	morozov_egor616@mail.ru	f	tchaivmore	highqualitybastard	ФН4-12Б	Морозов Егор Дмитриевич	b9b02aae-0db5-4970-961c-9a5a198f7ee7	1	f
863b5f90-814d-4e44-b498-53b02f3cd8c9	$2a$08$sFYTOYtGolX3ZX7sJNP1iurZVg563oOQ1bZ/EhybSdMQbQhZn2AzO	8 985 874 05 36	tonya.ferr@mail.ru	f	sasha12334	id435302079	СМ11-11Б	Михайлова Александра Борисовна	a845c25f-1727-40ef-90fa-8422bd94a20c	1	f
9b5f4e8f-71c2-42b9-8cf3-be67e4462b88	$2a$08$XTnCyf4HisSZ/CtayPWHTe2HarZHTe4XqzzsQVPnIdGYe2xbA64ti	89170462260	raeear111@mail.ru	f	arinooosh	arinooosh	ИБМ6-11Б	Руденко Арина Евгеньевна	a614efba-69c4-43fb-a9a4-d70d8be4c365	1	f
05dbeaa6-869c-472e-ad70-6f3679c414ac	$2a$08$mpgmEhOLsTDx566xdCBh8eHUkQcV64ucYLnr2KQX9ZtKKVZAstmBy	893011059-46	volkova.yulia2005@gmail.com	f	aaaaaAAAwasd	xemug	ИУ10-11	Волкова Юлия Васильевна	\N	\N	f
9de6d2dc-4d1a-4ca6-bfc4-71bd08895516	$2a$08$GJnEODl3G2j6c/vcLT4JReiR7iKxPyT8ninG5LjIvExcnpZuiKYH6	8919-763-50-63	ivanroman2005@yandex.ru	f	Guardian_Svarog	guardian_svarog	СМ4-13	Романов Иван Сергеевич	\N	\N	f
4bec4c14-6c2e-4b5c-bf9e-b2cd4258732b	$2a$08$MJCqInxd.nxbL9FWoauFm.G1qNXF/SdFw65ghdMtGvQyJouCYjtE.	89091500749	aaloginova17@gmail.com	f	lgnvaaa	lgnvaaaa	ЮР-16	Логинова Александра Александровна	\N	\N	f
992b283b-1d49-4a19-b8ec-1daa41d6ad96	$2a$08$SUk3fKdJ9b9Wcji/OylKYeo6YUCg6tozGXCpDIKin1doJB4pVlKhO	89856161788	stasikkalinin@mail.ru	f	kelvinreachs	kelvinreachs	ЮР-11	Калинин Никита Сергеевич	\N	\N	f
09b50002-4181-40c2-af61-4e0b33be04e1	$2a$08$4Q0t9SO1mSpWqUJzZhKyzuqXd3GFV8ZKwTh4SoELwjzr2xppyR81q	891996565-35	vvcharkin@gmail.com	f	Molodoy5003	vladimir5003	ИУ10-14	Чаркин Владимир Константинович	\N	\N	f
cd616b20-32af-44bb-8891-0ab0a4895968	$2a$08$/o/6sODseO6HAZv6BgzABeWAJH17ZZTgNT/OyVDyvwg6zVcjUTRA2	89852865078	v_rovda@mail.ru	f	RovdaVladimir	vrovda	РК6-13Б	Ровда Владимир Евгеньевич	\N	\N	f
4e18ae99-b875-4539-b42d-429a0538a208	$2a$08$uKxaxxGM7JpJs9O4w7.AEeyY07HlM2/QdbfH0LDfCJb7kxRc4gCPy	89260319297	novikov.nikita1609@yandex.ru	f	nik_prvt_aethetic	pumbochka	СМ12-11	Новиков Никита Ярославович	\N	\N	f
eac4983e-1c9e-47f6-87b1-e7be6b1a9a4e	$2a$08$kO7fVhJnWNe74oMtPAV9..BAagUwXdXmjziKscvfR5.JmH0jvNRrK	89522506005	myramy07@yandex.ru	f	alinakkoks	alinakkoks	Э4-12	Коновалова Алина Евгеньевна	\N	\N	f
0681b198-a27d-45bc-9e83-850e6a0d985f	$2a$08$9wZYKf1E92JDO.7thxu/6ey14n/nqEgA1hR7fYWXDJBJq3hPxz23y	8977-557-40-94	dashagasparyan2005@gmail.com	f	darkamx	darkamx	ИУ7-12Б	Гаспарян Дарьяна Кареновна	786d1844-89ec-4d6a-bcc2-ed5a5b83ab95	3	f
edad207c-9d76-4f03-b012-5d49cc5acd76	$2a$08$lLqSeyCAflD.Dv/EckcT.OYpEjAk9fMTUT2oGLUI19yjXJpyVsTJu	89104776141	yarik.shako@gmail.com	f	Yaroslavus5	biba844	ИУ6-12Б	Шако Ярослав Олегович	786d1844-89ec-4d6a-bcc2-ed5a5b83ab95	1	f
25e5633c-aef8-4f49-8fc4-efe1ec037337	$2a$08$u6z3EhFVCUWPQv7H13kyk.BZNLrlg2ci9TyBJrEUir1lCImx1TFmG	89876718875	kabanenko2005@gmail.com	f	NikitosikKitosik	nikitosikkitosik	Э8-12Б	Никита Константинович Кабаненко	6148d3c6-58fd-4ee4-830f-c8643ef8938f	1	f
97c2cfd0-9cd8-4053-a5b2-025731cd2e32	$2a$08$dRRzhQdT6Nc88ATMB.598elfxFrwOYn3MA2DvplJDOPfjDK3DdJxO	89885913618	osykina.ksenia@gmail.com	f	gotierr	gotiier	Л4-13Б	Осыкина Ксения Сергеевна	8f8f8c5c-9df1-43ca-aa12-6eed5edfea17	3	f
3402230a-e8ad-4788-9ac8-447466e480de	$2a$08$KRV57/jGjyZ5B2fx15r6Rel9dd3aqQ6eIgD3hipdVogIF4d.VhNtG	89937996322	diana.shrb@gmail.com	f	dianashrb	dianashrb	Э3-11	Щербакова Диана Алексеевна	485a9afa-6f4c-457b-952a-203d110898b2	1	f
0c413aca-3e5a-410b-9bcb-62d0edef7b03	$2a$08$c7D5XWngLWcIHXrsQvx4rO6FyM0wR1Va9IAjBrZPgH0I3nzY9iQaS	89166110725	nazarchukvanja@gmail.com	f	BalIvanMCK	balikhin3	СМ9-11	Балихин Иван Кириллович	572bfabc-f038-4482-97a6-cefd372cd700	3	f
41c30c25-9738-4499-8ccf-e5873fb69185	$2a$08$fjYita3m3TxugmKx6SZNP.kHfVzzg7e/LHTPkP7aYB5YQ9tRTZPx2	89055579545	natatakalashnikova2005@gmail.com	f	puppythesutul	yadoch	ИБМ4-12Б	Калашникова Наталья Максимовна	572bfabc-f038-4482-97a6-cefd372cd700	2	f
80a5b7dc-e4a6-4dc1-9d3a-5d456b34c557	$2a$08$h5qFI0J9fbUhOGzCUX2xfuz/zdRAqn8rIcJb9fuWhrR4D9YWAewm2	89500554233	suvorkina.vlada@mail.ru	f	aziza271	id370255804	ИБМ7-12Б	Суворкина Влада Евгеньевна	\N	\N	f
a622435a-9a2c-48e4-a108-cbd3b4a35471	$2a$08$P6/lGsMlc9sS5d8RDDIJE.Hnuno.4e2jo1Hm086U/n5yxBl4aDJFW	89295005127	vova.avetyan.0605@mail.ru	f	Mr_RuVa	r_u_v_a	ИБМ6-11Б	Аветян Владимир Эрикович	\N	\N	f
2b436ca9-de73-4892-a960-e61e4b250de6	$2a$08$YWqdha0yCODtHQ.V3bBCMOv/7a8IZT2wb4N6GaV29XFdowsb7V/D2	89092515804	m_solovey7@mail.ru	f	uuwushaaa	uuwushaaa	Л4-11Б	Соловьёва Мария Владимировна	\N	\N	f
597181b4-2283-4e77-939c-b6d98a3f0d16	$2a$08$jVVI5XOBfSl.NMdaQ/I6nOZPIrZ1J51eOuEKhD81o.8j1IMow2NWq	89670632149	dartaid2017@yandex.ru	f	pluxuryskx	pluxuryskx	ИБМ6-14Б	Денисов Кирилл Александрович	\N	\N	f
068c3201-883c-4f31-a552-f4a2bd95a7ef	$2a$08$myxErELKCiSDc3CS7.nND.QYRyZDJJr42Ai5cBRyZdN4Z8YxBxVZS	89039784776	alyawhite@bk.ru	f	alyawhite	xennexi	Л4-11Б	Белова Алевтина Сергеевна	\N	\N	f
4a9dcba3-fd79-4cac-81e1-e0c66f9fec44	$2a$08$WR4Y51tatj2csTGxeqaRo.clbY.5skq7IC5ikKKZDvMz05tcQEHc6	89154471841	chkyu23r222@student.bmstu.ru	f	k1sa1kun	k1da1kun	РК9-11Б	Чебурков Кирилл Юрьевич	\N	\N	f
c15e6b6e-280d-484e-8a38-0865ceb0be1a	$2a$08$0H//kGf0H2ga3JtcY29cZexYVhy/xH9b6Z/eUyYwSHc3pRXqMSP9m	89191060488	l_shvecov@inbox.ru	f	fmtrk	leo_shv	Э4-14	Швецов Леонид Дмитриевич	\N	\N	f
9034a690-42d2-42d6-9736-88d3f447d21e	$2a$08$2GOY6cn0SJ4uA/IrvEWW7uGjfdzoa99tVbOzn7Q9a.FtE9iCXfaNm	8 980 63814-85	orlovkirill2218@mail.ru	f	kira_prodacrion	kiruxa007	Э2-11Б	Орлов Кирилл Александрович	\N	\N	f
a1b7cac6-9963-4753-80a4-bccfae3c9d40	$2a$08$U1KJO6AyuIaaPfClwfr5HuM6WCfhD2k7QRUenBLh5iMgMMA2LcFz2	89652423591	pavelmart3575@gmail.com	f	n3ytron	id363411119	Л4-11Б	Мартынов Павел Романович	\N	\N	f
6235f765-7322-43c2-b810-b84f852bdaa2	$2a$08$COOwdaRk1XHWEA8bOrrRt.gV70GB9uLlnmdtFyjoHZzXrArOV8OAe	89877293448	fedormineev2005@gmail.com	f	minfedik18	minfed18	Э4-12	Минеев Федор Александрович	\N	\N	f
5ce4bf2f-68b6-46b2-9b43-8d03da660406	$2a$08$KFmfajjRJzY4jMCJxSQodefY/v9QMGRnrUbCTYitqGZ7Wtv33jRpy	89684423388	arseniy.pleasse@gmail.com	f	wormgore	arseniytimakov	Э7-12	Тимаков Арсений Валентинович	\N	\N	f
aa32485a-ab6e-4117-82b3-0ca2af0e5a28	$2a$08$.qO2At4l3psuqm.HkGdBy.ZMUdFMQssk6syp1GlVhHhJ4ywcISlry	89506831040	makarlisechko@yandex.ru	f	fuwaast	cherry_donut	Э7-12	Лисечко Макар Артёмович	\N	\N	f
53123599-507d-4777-b139-59ac0f559e3b	$2a$08$7N9LdFYmBPKQ9w2bNySY7unOlEExv/BCK63i/nHI9UlaTseIcsD8O	89163227619	krasava2005000@gmail.com	f	Chevab_Chechnya	id_chevab_chechnya	МТ1-11	Кузьмин Елисей Даниилович	\N	\N	f
4035ad4d-e9db-4f69-bb3c-b436b9ad8fa8	$2a$08$WWCw.v0Zeiz6q6JkrN2fGuzWRFDbTsLujNIiiqx93QlfeFnn7DAZu	89260722409	yura.zubov.02@mail.ru	f	yuraStarbiu	braxlush	МТ5-11	Зубов Юрий Витальевич	\N	\N	f
920773d2-b430-4b98-9cdb-e7f7af0a3d9c	$2a$08$0yjT/cTneEUzWotdrlbJIOMdaz304YvGSJPmB9paT0aHFZePakifG	89205824770	v1ka.chizhik@yandex.ru	f	chizzzhh	shkalvika	МТ11-13Б	Чижик Виктория Глебовна	\N	\N	f
cf8392b2-e8ca-4186-8465-2340543b11b1	$2a$08$E1jPRQO4UW8AnAmc/D4YEeKNhk9ttvFbIxW8xTV4OU1683FzEL7lO	89850110799	stepanovads1@student.bmstu.ru	f	di_step	diashastepanova	ИУ4-11Б	Степанова Диана Станиславовна	83c9e33a-e863-4ef5-9909-2ab2db18b7e4	3	f
d4f9cf5b-72f7-4ecf-b8de-6395bbd21149	$2a$08$YIJYmamEJZGZxG4JBTkKzeflLVdAokrKjKXAx5w2GU4Jb9IR6Rj8y	89917868812	o12olchik@yandex.ru	f	ololollgo	olololgo	РЛ2-11Б	Ковалёва Ольгв Викторовна	543a4e0b-e35d-44f5-85df-43431142aab6	3	f
d17c8b94-abc0-40c8-af70-1f9110cd20bd	$2a$08$KZvCY5MGd2W3lqaa3dflIu0nTAyo2sSl4CkiWNVbrIwHCgZckKu56	89689056250	alena.zhuk.2005@mail.ru	f	puantic	paukzhuk	БМТ1-12Б	Жук Алёна Игоревна	300ef3ae-2e98-43cc-8151-a7f42980de56	1	f
0b1d6910-994e-4475-86fb-cc9c8c597742	$2a$08$/U13PCwzKN5NSI5mGUMJQ.Ug2KlrUZ0l6NwtTEM766IqtO59JYcO2	89277097631	koltakovadiana0@gmail.com	f	didianaa12	dkoltakova12	Л4-13Б	Колтакова Диана Александровна	28fd4455-c957-4065-97f6-847b632daf2b	1	f
8fd782d3-9468-4fd4-9880-87a1b0de0d5c	$2a$08$HpmlsBw6AEK68Pd.NUbw4etGqPgv0pbOQGhi/omdcZnwH6MJTkjB2	89650363663	osuvor3@yandex.ru	f	Nata_hzz	id184849703	РК5-13Б	Суворова Наталья Максимовна	300ef3ae-2e98-43cc-8151-a7f42980de56	1	f
10340800-8ab7-4da3-8b1a-4fbfb59a1c13	$2a$08$pYwDcPLyhyulFlYt/zqZjOEUVGhrbu4bH.t9H8qYRCieZNEbzmOWK	89175538184	lisa.ziberova@mail.ru	f	callmetsundereass	tsundereass	Л4-13Б	Зиберова Елизавета Александровна	28fd4455-c957-4065-97f6-847b632daf2b	1	f
2d317cac-95a6-4851-9c55-7ca546d3cad8	$2a$08$ik5asU1EU4SK7dtRyKsh6.4fU0Ih4EGFH.RyuBdxidcJrH6UTgvx2	89998109178	rebzinka@gmail.com	f	StephenMarkman	stephenmarkman	ИУ6-12Б	Кондратов Степан Юрьевич	\N	\N	f
85de0ea3-2618-4c46-8c88-b54b1fc2ef89	$2a$08$.89LuwBrMX0OiAVAZJ3JceXyUxQ6/zMOnpNKLDUAlgFIrE5SHr8CS	89290822888	akbar_01_05@mail.ru	f	darkl0tus	ne9zaxodi9sojret	МТ6-11	Умурзаков Акбар Мирадилович	\N	\N	f
4efb5d41-b549-4ba1-bf23-946848b91ca1	$2a$08$lZS/tHYL7HEQh4kqejju0uOW0BEJRuPojyhmzuGbc1HkOeC1W4lAu	891527902-52	fedaydemetkot@gmail.com	f	feday234	feday234	Э8-12Б	Борисенко Фёдор Дмитриевич	\N	\N	f
62aea429-a827-4744-92db-3f6f1c8f0e57	$2a$08$2dLQhTBsm2.5eRHG/Cs0LuvUFTJZCmcmEbbTugvHVMPAS9yPu2.02	89154112349	andridudin5@gmail.com	f	cotoucan	cotoucan	БМТ2-12Б	Дюдин Андрей Сергеевич	\N	\N	f
1ff44f08-6f8d-464d-8e9e-751f180574dc	$2a$08$OZtYCwLKGRWe80sQRu4/JeK42brPkdKDI2CdSRKasV8bmX4iHmVW.	89507290455	isheevnikita@rambler.ru	f	knyzlama	knyz.lama	МТ10-12	Ишеев Никита Евгеньевич	\N	\N	f
7c32ba6d-6959-455a-ae4d-126b24f46d4b	$2a$08$um6r9zUDxuLX.SYZ14Z2hOi3XfwcjIPafWrLhCRx7Q0zWaTwKWMGu	8 916 8451451	egorkurochkin21@gmail.com	f	Yeger991	id384549119	Э1-11	Курочкин Егор Ильич	\N	\N	f
78a95119-1e64-474c-862f-49f5ee3b269c	$2a$08$/cXbfeX2Cml2wNOL8SIac..2h0qCT72Ekm9Xpwcq.YeDH4GoG8Iv.	89198791495	dasher.romiro@yandex.ru	f	imdasher	d_kaluga	Э5-11	Калуга Дарья Сергеевна	\N	\N	f
08919015-2f29-4a03-a789-d50f35e19bc8	$2a$08$7pVuZRtnsvEhLndkdu7PmeUe3SzO0e4dkdHFY9TDiDcYB8EYt7rWm	89605569062	mit23560@gmail.com	f	dvmelnikov	mcdimasik	ОЭ2-11	Мельников Дмитрий Вячеславович	\N	\N	f
0625b86e-31d9-4850-96e8-97e144f9f2a6	$2a$08$bw91wvo586pLPSEB7uKPI.TW0Zse0xwbXu/rHmtF/hyl6ohgzmxxu	89307060621	ar.lis@mail.ru	f	a_sh1t	akxeliess	Э8-12Б	Лисина Арина Игоревна	05735fa8-6c86-4424-865c-79df5b9850c2	3	f
459de883-ba46-4947-96b6-632f224a2c21	$2a$08$pmwYyZNAlukLWEpeeIutzegBSqw/PVWq59ZPRS1gkllJd3ReFsvTi	89520255154	s4sha2706@yandex.ru	f	sashshoo	id560792736	МТ5-12	Сергеев Александр Александрович	2b55f5a7-5872-4904-bdea-7d54243826e8	1	f
5f154485-4cba-4090-b9a3-cf842e9a209b	$2a$08$XXvQFtOm5TOBwzGn9MCwr.D6pQxrO7Dx7EHzpQ1/inIqLeMkO1R4O	89226797329	pastin785@mail.ru	f	xedyshka	xedyshka	Э5-11Б	Костюкевич Кирилл Андреевич	9f8e20ed-c2b2-4f77-9d95-ca22c7ad82d3	1	f
82e9f83a-b537-46a7-a370-ab6944ec5924	$2a$08$iV6ansHkblNtHD752qPGp.QmX2RaTPTU28E/JBhqAXZX3RI8MDCJm	8 961 30503-45	nazarenkodi@student.bmstu.ru	f	id1687988937	4sv_toksik	ИУ4-11Б	Назаренко Денис Игоревич	83c9e33a-e863-4ef5-9909-2ab2db18b7e4	2	f
3afc9357-abed-4a7f-bafa-a733f9dfb0fc	$2a$08$5imfoqKhmeyPIWW6mS/I8uwS7MI5957ciyP6blvsjtaLng1JpRF6m	89898303306	belovaanastasis@yandex.ru	f	GennadyCrocodile	gennady.crocodile	СМ12-11	Белова Анастасия Евгеньевна	\N	\N	f
a701e02d-fab1-45ff-b0cc-1ece526e3af2	$2a$08$bmsuGd8/52IFX5SbsKXh7ej32qzdatQqZspDG4imT17f46Napd8ki	89053721962	ellellina@yandex.ru	f	callbyslown	dntunderstand	ИУ2-14	Валеева Эллина Диасовна	\N	\N	f
bdf1c1b2-4e4e-4e3d-9e4c-cda97661e0c5	$2a$08$IaALCACsGsqfxWagcfHaCu8r6onCdNOyfz/dE7vII2Kx2MgkFdrOO	8 916 5706740	marisha.mashinina@yandex.ru	f	goluba_ya_laguna	goluba_ya_laguna	ИУ1-13	Машинина Мария Романовна	\N	\N	f
26fd4c19-2cfd-4196-9f33-fc48e22c1d4c	$2a$08$TNzTnCkPZ.ZG5Ml.LRyZYus7/bpquZiAzCB4Z5BS2aNcCgwXxYtdG	89999229323	gav23l271@student.bmstu.ru	f	allariai	alariaii	РЛ6-11Б	Горшкова Аполлинария Викторовна	\N	\N	f
a88bae33-aad1-4a92-8341-46de0b78172a	$2a$08$6.TyKIqcVLQazFB4z47tMuShbFDkZgzrb6mdZ2C/dYRuYyU5.rqJm	89996070973	varakinkirill3@gmail.com	f	by_Biska	by_biska	ФН1-11Б	Варакин Кирилл Андреевич	6b2934e3-98fb-4a5a-9138-35a62df992b9	3	f
bcda4dd5-438c-4617-b781-9efd27049b59	$2a$08$qQjLUDOBiKUSmJTWjAghse5BhJonbl/LKRN1kBTwpyAdu3kLXX.i.	89966749101	samoylyukvi@student.bmstu.ru	f	VladislavSamoiluk	vladislavsamoiluk	ФН11-13Б	Самойлюк Владислав Игоревич	73ce80d7-2798-442f-918e-bf40511dfb07	1	f
a2c2f92d-abe5-4354-9125-5b71ff844946	$2a$08$dyui5OK1Fx92up8UuAjOcOrHFTaiu5KCT5HI5limwYK4u7scyKfnu	89099548758	ryabkov19@bk.ru	f	DeepStage	deepstagee	Э8-11Б	Рябков Михаил Андреевич	71580588-3eca-4258-817b-a9d3ab09d703	3	f
c2c425bf-8026-4e33-8e75-d3104dcf4d87	$2a$08$BZfeLPhzlzkVc6PtOjHB8.BNIIWW3wTNzkbll3aI4r2DVDFEcp.0u	8917-021-02-52	shchetanovalina@gmail.com	f	wolfer_p	lwolfer	Э4-11Б	Щетанова Полина Денисовна	95fafbf5-1bbd-4f58-9ee3-1e1cf6eff939	3	f
f5b6d319-8878-492e-983d-bf5fc4b07499	$2a$08$wSxEYAluyTkzZD81rbcQc.hVkLgE2Q/EUXf54sRWxae503WjHssOm	89653400520	kirozan2006@gmail.com	f	tipical_kirill	id512076816	ФН1-11Б	Розанов Кирилл Игоревич	5095faa5-51b3-4dcc-90cd-744b3566ccf2	3	f
350729dd-6b9b-4f4c-9356-d4e0366c251e	$2a$08$n7HbBAQqlE7fuQtf6qRQDeb2bqnpf5SxIZkCGrfKWqWp26ML7ySqK	89204195265	alna.borodina.05@mail.ru	f	teatete	borodinaa_a	СМ11-11Б	Алёна Алексеевна Бородина	a845c25f-1727-40ef-90fa-8422bd94a20c	1	f
47c8e3e7-750e-4461-bb38-29800725004b	$2a$08$mr4haWr936wh0jf1TwZwdOONY5v3lJYUq83IPWg2WUHtxNm6sptB.	8 965 26371-64	litovchenkomo@student.bmstu.ru	f	Biffal0	biffal	ФН1-11Б	Литовченко Максим Олегович	5095faa5-51b3-4dcc-90cd-744b3566ccf2	2	f
b12256fa-f0a2-45a4-aa7b-308a510f4aa4	$2a$08$TKTRq71dJCT7ETrweSmTHOn.8QUyMz3opO4eyxgy0ZlXfr4coppOK	89056083655	seredina.2005@inbox.ru	f	varechka_1_1	v.seredinaa	МТ12-11	Середина Варвара Александровна	29103222-c68f-427b-ab8c-2cb674f1fb97	3	f
83f87fb2-211c-4c25-9144-4694b87253a0	$2a$08$Zlh9nJ7NFSJeR/gs4P7Q2O.6RKTSUYzocaHtMiQHWoSZIBxXH5H0K	89154929202	anna.mek2005@gmail.com	f	mikhailova_aniuta	anna_mikhaylova0	МТ12-12	Анна Валерьевна Михайлова	29103222-c68f-427b-ab8c-2cb674f1fb97	1	f
2066934a-cac2-46fa-b036-496ddec0821e	$2a$08$ETSOA0rzm90rDXYajRFHw.wQluQfRiHV/oljtdJMV70ckdbDvAgoe	89379035665	manchedelia06@gmail.com	f	deliaman	delia_man4ik	ЮР-14	Манченко Делия Руслановна	\N	\N	f
adac4ce4-6f61-448a-ac3e-3826d8ac1c22	$2a$08$s5bfRl0JSv5qgVjssYtx9ur4bioRwhW82o/m2qhcb1THEEncsy7PK	89775926755	ivan2005ae0408@mail.ru	f	chikiborboni	chikiborboni	Э1-13	Черницын Иван Евгеньевич	\N	\N	f
a62943ce-86e7-40f8-b0e0-b439ee755366	$2a$08$1JasjYMawiBT3qaQd4fUPeePXb2Ze60MTxZA73Lm4/1mRLpcJuaqO	89255437257	anry2005@icloud.com	f	ANNDRYUSSHA	a.n.drusha	Э4-12	Тюмин Андрей Олегович	\N	\N	f
e14c494a-5f79-4bb3-9177-dfe0874ba13e	$2a$08$I3/6n0.LFzBRUFJA945Na.Ok43Npdzxrn7bld.cTJnwBDG1dH3mtG	8960-589-19-18	anosovalina6319@gmail.com	f	l_h_l_i_l_h_l	a.n.osova	Э7-12	Аносова Алина Дмитриевна	\N	\N	f
652c6ded-567b-4462-8476-ad48c542f603	$2a$08$Kg5Y30lorQdQv7jEeIRcSu6s.IEN9Oyoct/NCzrnnA/OZ0zDB6rXy	89850108585	alinavoronova_05@mail.ru	f	bippiu	ilshk314	ПС4-12	Воронова Алина Андреевна	\N	\N	f
096d7a90-66da-4b11-ae43-aad3b8546b61	$2a$08$DpaRCRLgr4J/5t.H7i.qXOPN.TdbUJv5LubZyA1KUpKyHxIOCuMe6	8 915 14647-78	yegor.naydenov.04@mail.ru	f	egorthefounder	egor1580	ФН2-12Б	Найденов Егор Юрьевич	\N	\N	f
b30d5d48-3e87-49e7-8358-f1e8cd6e76d5	$2a$08$t2HmjrQmdvwPShBOIsExiO8RVN63dXBTfYIfaFymSnN/Fjx.ma3si	89056531577	p.mandrykina@inbox.ru	f	polya_beskraynie	polya_beskraynie	ИБМ2-12Б	Мандрыкина Полина Андреевна	\N	\N	f
6480cece-76be-42c1-b6e0-cbbd24a04f0f	$2a$08$hSan/T7FOOaxo6GZne4uWOHVQAExuq67VmxvJweMLRzuBq2SGDjRS	89103395165	konopelkokatya2006@gmail.com	f	Kathrinekkk	kathrinekkkk	ФН11-11Б	Конопелько Екатерина Сергеевна	\N	\N	f
958d01d5-7dba-4815-8666-af784e67e3fd	$2a$08$s2ue.pzhbePhLEXdY3QDvOf00lsjlUSKdh0UmLr87mh8OvqNdVx8q	89175012977	karevmax2005@mail.ru	f	zampolitkarev	ya_pelmenxl	ЮР-14	Карев Максим Игоревич	\N	\N	f
a0d16b88-a63b-4ae6-a428-bbe224c3005a	$2a$08$C8/3Wi6nrG5gr7U7d/zIOe1iTYl1qWkEefcJ5IC5o/Cvx8JFQIVKi	89032457728	dasha_dmitrovaa1@mail.ru	f	glavnaiai	idlovedodiki	ФН11-11Б	Дмитрова Дарья Михайловна	\N	\N	f
706e3f1c-6d71-425f-854e-509f68530ac8	$2a$08$DXFrc3ryY9gc00DO7yNN9urQIMw4gjbVs7muAsiBCa/4L47ad49LW	89852819933	dmitriy_gudkov_2015@mail.ru	f	sempaiturist	artist_dima	МТ6-12	Гудков Дмитрий Максимович	\N	\N	f
74dfd041-09bc-441c-a253-c3c72282b77b	$2a$08$zy/mUr.HD5.I5SZyzVo7OeOckzpbM/dBGDYZNapdq399qiGbOc72u	8917-487-04-96	gilmaletdinovagyu@student.bmstu.ru	f	guzelchansky	guzelchansky	ФН2-12Б	Гильмалетдинова Гузель Юлаевна	73ce80d7-2798-442f-918e-bf40511dfb07	3	f
ff4f01fb-87c8-469e-85aa-86c1e1d3d25f	$2a$08$FVM/tFtyAQV.lj3AHW06O.qnuPkJ3WX1F06dUZd8VWpWg85A344P6	89053431415	sininkirill978@gmail.com	f	HocHyBebRy	russischscheisskerl	Э8-12Б	Синин Кирилл Сергеевич	05735fa8-6c86-4424-865c-79df5b9850c2	1	f
a49013d5-d1fc-4d96-8405-a5637ac6d298	$2a$08$X1ry135xIoE6t4KwF8bjJ.DY4ylYJ1ZrWPBwuInItjVdIJ5AzPASS	89178657210	ilya.lis22@gmail.com	f	dolph1n123	tvoya_me4ta_zzzz	МТ12-12	Лисихин Илья Андреевич	29103222-c68f-427b-ab8c-2cb674f1fb97	1	f
d7a67c7d-5ad0-47ec-be36-ad2b78385212	$2a$08$Ic8Igd4GTAFTQjc0tmCKJetN.oKzWt3JHOcYi1.Nl6f7cAnPwXQKu	89002709880	r3d049nf@mail.ru	f	chpocarianian	chpocarianian	СМ5-11	Лышнов Иван Владимирович	ccbdb464-37b4-4fa5-88e0-8a97ea84dceb	1	f
d174bc36-f57f-440c-8b66-573c7319851e	$2a$08$aFeg9yc5.NToEjQbyO0j.u/nxSsyM6jybqGHCyt6BH.v7ozKLA2Pu	89374675009	baltyke@mail.ru	f	uauave	uauave	Э8-11Б	Эренценов Балтык Игоревич	71580588-3eca-4258-817b-a9d3ab09d703	1	f
84d3cf3c-6d27-446f-a300-7d276b6e4c9c	$2a$08$jDHL2qgvtjuJrSeN96gHR.6rAbil9n25QU1XfjLnBYnlbvnLFowf.	89156246994	grigorii25frolov@gmail.com	f	gggreggfr	ggreggg	РК6-16Б	Фролов Григорий Николаевич	1c6d8c8a-9bda-4ec8-bbf6-8dd0b83b98d6	1	f
c6b43f49-3e1c-4930-a3be-dae25d045833	$2a$08$NlfU4qEeWX4lAeUaw.66EOSrwLIBg/rCa5N5MCFJkbYMogqzxSEYG	89671172894	sskaverinass@gmail.com	f	kertoflenida	sofa_mur	ИУ5-11Б	Каверина Софья Георгиевна	7b146ffd-3941-4851-ab0a-e557884042d4	1	f
4f481eae-3f8d-4e5f-8d32-287ed321a274	$2a$08$Z5ZlZ9EqYmDypLJUnwxGTew6peRshjjjx/iE3kffx9BbTnIh3T3aG	8 909 95249-51	askerko71@gmail.com	f	Hisa127	hope127	Э10-12Б	Аскерко Надежда Дмитриевна	20cd8a22-243a-42f8-8253-5cadbeeae107	1	f
ce911989-7956-4b5d-9c4f-aac88f24a318	$2a$08$t8NxYkINlNm32iNBKdAUh.Ria3sqevOdPiYtQ9QkrH.d0ONd3/93a	8950-048-59-26	fenochkinamarina@gmail.com	f	mussaurus	id306750893	ИУ2-14	Феночкина Марина Павловна	\N	\N	f
a23816ba-71ad-4eda-a0b4-e28dc42e9856	$2a$08$cVATXSy3xNE.lB1C1JJ7Te03w9NAiOoY1Z.ZxcjWFyF7ClZ5abPKa	89831898483	vvaigel@yandex.ru	f	vvvkulz	vtoryv	Э4-11Б	Виктория Владимировна Вайгель	95fafbf5-1bbd-4f58-9ee3-1e1cf6eff939	1	f
9cc58c54-d05a-4f65-af8a-5b6b8187cfe5	$2a$08$WkW5TNAhFkqEsWIjpzJz0uSb5UyTnas5qivO0TLE3/sWxaAl0Bbei	89164281537	viktoriaborisenko.vb@gmail.com	f	viktorialeevb	ltyviktoria14_	ИБМ3-11Б	Борисенко Виктория Вадимовна	\N	\N	f
da6b0e01-08f1-4cc2-af68-8c5cdc7fbc98	$2a$08$OV/ynrcGfTgXu6m/P0EO7ORcwp62teqfHalGvehniunvlrOdxE1yq	89502542144	dyanchikk@gmail.com	f	dyanch1kk	wowlemdmd	ИУ1-12Б	Мурашкина Диана Валерьевна	\N	\N	f
f924b5d5-7190-40be-92dc-d3f0e04e296d	$2a$08$LQY7mgWcg/T/7u0PaW.RYuSZ7lM6gF.3P2z5I/lAFk10bo.sSu7Fm	89253490636	mshtakhanov@mail.ru	f	leckererpfannkuchen	mishasht	ЮР-11	Штаханов Михаил Юрьевич	\N	\N	f
e29bf816-c21e-44b6-a399-36b58b250730	$2a$08$lExndv0efMll0LYyrvVQ1e8.ewmlaRxGY8bDB9F8aeR10f.QQwjUu	89157208911	orloff201520152015@yandex.ru	f	Pavel_BMSTU	orlov_bmstu	СМ1-11	Орлов Павел Александрович	\N	\N	f
8e878474-7414-4a84-bc9a-cc89dadfdd14	$2a$08$aTOnsgzJdVKJ1HIrKPMynOuywpSVcPawhhrnafxtY5MkKjd1xHdzW	89253602061	pichkurpa@student.bmstu.ru	f	koi_f	id298289820	ФН4-12Б	Пичкур Полина Александиовна	\N	\N	f
e626acdc-7b4a-443f-8200-67d78ae32976	$2a$08$aj.7GPAcHUMjGw5tnFyXSOjd1ihHcgBOHkf7w59zPu0BKEWOBsTOC	8965-304-31-20	orlovs2006@mail.ru	f	yaanaa06	yana.orlovskaya6	БМТ1-12	Орловская Яна Валерьевна	\N	\N	f
f7d014a3-9046-4664-9870-bec450cd9bc8	$2a$08$s2O/3B9DQ53ZOApBg4a04uTbeN/z.C7Nlqj1/O1ICxz2EyO7KhzmG	89255356423	egor.oreshnikov@yandex.ru	f	O_P_E_X_G	1dreamer5	ИУ2-14	Орешников Георгий Андреевич	\N	\N	f
7935ea87-bece-4f67-b45d-2181bb63fdcd	$2a$08$up93z7d7GaiLZtfLAZJ/gOsqpdioKYcCf7.VHTZLmaB50wK1eoaNG	89272343640	kodokuschi@yandex.ru	f	kodokuschi	hrsboyg	МТ10-12	Белов Александр Дмитриевич	\N	\N	f
a232fa3a-4265-47b9-a29a-9cf4e8c06511	$2a$08$7wpKN/qskE8AR6GWi/Rwt.WudTipD9ar6Lbc4FjY6svot54dD8TSC	89266465047	tima.kaplin@mail.ru	f	wildbeart	x_x_x_x_x_x_x_x__x_x_x_x_x_x_x_x	ИУ1-12	Каплин Тимофей Сергеевич	589132a7-7358-41cf-833c-7a5bd9f8e1b9	1	f
70a30c87-ff18-4f40-bd19-c63ece5e5ce3	$2a$08$114f3VnQdJMn9BXeHlE0mulmAKyaF1.0qO0TB01J7pdEjsxN5Xk.y	89518674372	polina.izvekova@yandex.ru	f	Issvekova	Issvekova	ИУ3-13Б	Извекова Полина Евгеньевна	ab1abad2-ef51-48f7-a15c-797d70061adc	2	f
a7ed4bcd-63b5-48b6-92f7-d7704004e697	$2a$08$eNuVFTFSPg4eqHNkGajcf.Pup8k0/4mniLkbXNgUrhJAJ1KV6J2NO	8 910 400 82 10	alya.godina.99@mail.ru	f	akiogawa	akiogawa	Э4-12Б	Година Алина Васильевна	6fa56c83-f29f-4909-9d63-ed00152c144b	1	f
d09b55f8-ed1d-45af-905a-123e679ef89e	$2a$08$zGcxgQAOzCoTaTwZ0gwoWu2yy2yUn5zOKTKDq68QrM5y4Zv2/7d1y	89252423574	butsvasia2005@gmail.com	f	F3n1X_X	fen1x__x	МТ7-11	Буц Василий Александрович	\N	\N	f
145174b9-6cae-48f2-9208-de5feee5592b	$2a$08$ixDrLp6QChh20F.ZVOo5g.Q..PGBWdW/M7TQMe3vkdB8o5fIfW85q	89063072072	bestlana06110@mail.ru	f	svetlanakudr	s.kudriavtseva	ИУ2-14	Кудрявцева Светлана Геннадьевна	\N	\N	f
b5b92a75-d67c-4475-921b-e7164d9458fc	$2a$08$bbaYsrCKvC3vtPuNneVpsObVeHR8.gtj8ewZFfPKGJnZWwIQv.8tK	8 982 27891-05	ziborova.yevgeniya@mail.ru	f	JeniyaZiborova	id608271276	МТ10-11	Зиборова Евгения Дмитриевна	\N	\N	f
53549c80-07dd-434c-a137-de79962e2471	$2a$08$EBzIf.WKQonUtESZh2w/auT3juIIGF0O3kyCOgVI.iQ/YHuMWqZOa	89991399764	kiselevaid@student.bmstu.ru	f	irinkakiseleva	id_8800555353	МТ10-11	Киселева Ирина Дмитриевна	\N	\N	f
1146d11d-bf32-4b58-ad02-a5aff7589073	$2a$08$eSRqBV0bfaXr0TSGEler3OmeeSkS.w1jL2ES3PFQP1teM8VmdEGOm	89153388838	pok2005@yandex.ru	f	Ilya_Pokrovsky	ilya____7	ИУ2-14	Покровский Илья Алексеевич	\N	\N	f
7f73e6c0-b5c1-439a-9ba8-931c1985334e	$2a$08$.01DnZYfnOq.6gN3sXnAA.akMWhvrbKTgbwLw9bwdd0nOkG68J7bq	89645782078	sophia_ivleva_2005@mail.ru	f	sonyaivleva	sonyaivlevaa	Л4-13Б	Ивлева Софья Сергеевна	\N	\N	f
1bdf4a31-8331-4498-b143-810846198db8	$2a$08$LetNS/Q0ZWooqF3olbz2P.2a1DcL.NoI0F2ows8cCsxryiUNv3HsO	89260722409	zubovyuv@student.bmstu.ru	f	yuraStarbiu	braxlush	МТ5-11	Зубов Юрий Витальевич	\N	\N	f
1f596bd7-d6d6-409e-b173-4785a8ff48d4	$2a$08$RmoUwbX1gkzOvAzWYl8GtOPxA4yTkE9wOspsMajKcicOshkwCIwQW	89851926410	soniaiva22052005@gmail.com	f	str0nz00	str0nz00	СМ13-13	Иванова Софья Андреевна	\N	\N	f
d4f332c9-5b37-4274-8a8c-4d59d9c64b72	$2a$08$xWYcwDMOMCjx6Yum4yfq6eRwYUyXCD84OIpUTSDjWUadUrGblo63.	89260122668	turbo@lesnye.ru	f	ghiler_er3	ghiler_er3	ФН1-11Б	Илья Грачев	\N	\N	f
341c48b7-b484-4919-afb4-916cd024d110	$2a$08$Fjoyk9gTqT4.SirOqoQ6uOlXoPC.4NRiU/I7klWaEbQwjARvkFEoe	89784429657	cherbenfo@gamil.com	f	cherbenko	churbenfo	ИБМ3-11	Журбенко Михаил Петрови	\N	\N	f
8211e787-155b-4e5a-b0fd-5bdc9c183266	$2a$08$pWflEV968hC6jaRQKWJUtOMG4kKTGBnZ0s16.xE0dhXoQ8uTVNX1e	89250252814	ilya111550@gmail.com	f	Bragadir	brigadlr	ИУ9-12Б	Чистяков Илья Дмитриевич	\N	\N	f
fdd48fec-6991-4752-99dc-f8bb59130c0a	$2a$08$CBMzO8BvR5vt7Hk6H9HaeOu.1VjTV3Zg4etWB.E4pLMOuNg2ARmMy	89166508876	sovetkin087@gmail.com	f	makar_no_ne_iz_chbd	zip.file_exe	МТ6-12	Советкин Макар Эдуардович	\N	\N	f
7af99109-b08c-45c7-8ffa-2ab440e44013	$2a$08$TixSkug6hkBjNTR4J.SB9OlviVjORH6UEQ4/JGPpAy3BKzIlYmc52	89575517123	qwet123@mail.ru	f	telegram	vkvkvk	РК6-13	Банько Евгений Александрович	\N	\N	f
dfc78805-16a4-465e-9bad-dcf4be19e366	$2a$08$Txw3gw5rQAMHkiXlw6ZXtuXYV26YqHxB4qxCCGf19XyMbHM2YYpi2	89575592684	testtesttess123@mai.ru	f	tgtgtg	vkvkvkvk	ОЭ2-11	Тестов Тест Тестови	\N	\N	f
2354f1e9-b13b-4d2a-86ef-260a7608cedb	$2a$08$A0.2/1GUm/MvT2lb6VJZROMxmU8UBzzNxCCCxJUgPaH.AG0Cw4H1C	89263911043	superilyshka228@gmail.com	f	cmeshar1k	cmeshar1k	РК6-16Б	Мещуров Илья Алексеевич	\N	\N	f
c64e5d66-4c50-44ca-af82-d7df737b0e28	$2a$08$pcSsgGmxsKUlIUY7WfqykukmS58poKel47E6COQPt82hsZBAs/HWa	89538523216	vvponev@gmail.com	f	Poneve25	id502597156	СМ8-11	Поневежская Варвара Вячеславовна	\N	\N	f
b5be0d67-1c16-4063-844f-7f849b83cc85	$2a$08$9T0uFvUIxGcOZn.4FM2UtOFDis4z7wiJfPgcunHG.WycoV73LpnIK	897757580-06	kiralarionova2006@gmail.com	f	Ciliate_Shoe	id455478241	МТ5-11	Ларионова Кира Дмитриевна	307c360f-258f-45b6-b1cc-81c946df749e	3	f
441090f0-d5ec-4dee-94c9-1204b2a03771	$2a$08$TKTiBeSA0ZMthOsfCfKBfe8uvNK3FJrDnnHIcNW/Mr80I1stEMYuS	89152323263	evgnnnnnnn@gmail.com	f	evgrish	valer9vna	СГН2-11Б	Гришина Евгения Валерьевна	2b55f5a7-5872-4904-bdea-7d54243826e8	1	f
3709243d-fb29-4cbf-a6f0-f53fbb4ff6c0	$2a$08$jAL0qnL0YrVRDr43kYn1fuHB.JSDC.ww82s15l8Ae8F3sH5.AHwsy	89261010141	len.baburova@bk.ru	f	LenLevi	id481481887	БМТ1-11Б	Бабурова Елена Руслановна	074448b3-2a26-4b58-b63a-cecbdad6590d	1	f
f443bda7-d860-4334-bbdf-5e1be78f1c17	$2a$08$jbegAJKlGSmvyar1BgCBJO2oiHLJaTQY8lIhZvQwLX8yagvpBd3Bu	8987-359-64-16	aibushev.goscha@yandex.ru	f	Pupkenberg	id328495776	ИУ8-15	Айбушев Георгий Бегенчевич	4b718a67-1bea-4e14-ada5-9b2eaa2731ed	3	f
acb4d98e-363e-40f8-9b22-1627b6261ece	$2a$08$HH1mCpgDEieAcQBVS2Uif.2T42ggT12BWrrPGNQqxJEoa489qjF.C	8 915 4767627	kanevskayay0w@gmail.com	f	kustebatb	scolopendraposhlanahui	БМТ1-13Б	Каневская Юлия Евгеньевна	\N	\N	f
c181ec19-5a7e-4e69-bda9-0aa021dca304	$2a$08$54uje1UwbaVEWyi3NN24CeaKDpCZNh0kQjMLTah3yEGxbGany0wmK	89221183417	dkondrich@bk.ru	f	dekondrich	den4ikon13	АК4-11	Кондрич Денис Леонидович	\N	\N	f
6adbcf28-857d-4c08-9490-91de3741635e	$2a$08$6c4dNOMfS.H0hxMowKVId.hoTku3Z1Etynn2xQ6lMwPYRDy6Zj9gS	89009264675	dreikonline5554545@gmail.com	f	i_5kor	i_5kor	ИУ9-11Б	Скороходов Иван Александрович	\N	\N	f
84ee35e2-0e6c-402b-aa2c-43811db2077d	$2a$08$v.OtythxlNbPYXmgYVNgFOxhzM/kkuyLkPCDGEgfMXrbARSpoh3RW	89017319353	ann.naumova@inbox.ru	f	yatalantt	swndk	ИУ1-12Б	Наумова Анна Андреевна	\N	\N	f
62834d65-3d2e-4872-8fdd-d19e6af1b403	$2a$08$NtUh.QNuqvdfWoxkXKEcJ.Tbg4synjE03SWV2v48.XCxSr8l7eEr2	89773142006	antipixk782@gmail.com	f	AnPixel	anp1xel	РЛ2-14	Илюхин Артём Владимирович	\N	\N	f
0dea7b1e-17b1-4432-8639-c69dad32fb41	$2a$08$3hW1SQvFWlKYWeIa8Ejw3ejams60hEENiffad5qByIBOGgSANeFYa	8985-143-39-85	negotive8@yandex.ru	f	sashka00502	sasha00502	ИУ5-11	Мишакин Александр Олегович	7b146ffd-3941-4851-ab0a-e557884042d4	1	f
bd856d3c-b466-49c5-949c-2525e652f964	$2a$08$b85hn0qVVSrBdSUWxgQ7c.OqQL6Walxzye9zqm1su6OL7e20/y272	89164066207	mygirlst@mail.ru	f	jelerhy	sonyavdev	ИУ3-13Б	Тихонцева София Романовна	ab1abad2-ef51-48f7-a15c-797d70061adc	3	f
f522bebe-a481-4cc0-ab8a-15ee2ac1ff0a	$2a$08$TDWc/3VbdT28bzPTzH5e.eL7HzPAV1G0QmzFU9GeV35NE1pQl8QSG	8927-068-59-05	mes21m423@student.bmstu.ru	f	Bear01062003	instagrameeeeegooooor	СМ9-12	Медведев Егор Сергеевич	a845c25f-1727-40ef-90fa-8422bd94a20c	3	f
8e3dc98c-224c-412f-8571-b740f2671d19	$2a$08$Tl2kxaPNn4VnkqvidleNZ.fBpywu/e0icPKZlUIA7OmFBPC/4rQmW	89261865880	arinasmirnova05@yandex.ru	f	89261865880	arina.smirnova123	СМ1-12	Смирнова Арина Александровна	2acc4cbe-c250-4830-a19e-8b8af134ce2b	1	f
3ac8457c-8ce6-48eb-bcc6-8973ed795dee	$2a$08$ay5pLa9HAOXXUsDFAlvS2eINGw1LLrXesxep4AmXfilu7sbcxYzYS	89161279517	uevstuniceva@gmail.com	f	erforderl1ch	id471569866	ИУ1-12Б	Евстюничева Юлия Александроана	\N	\N	f
72894616-54c5-4f80-84b7-8951c139243c	$2a$08$MDYSFFTjMXNS.A9mPdcqmu.4a7TIBOE2yUmwkTcXPzuIhPpRs3H4q	89013542677	daniildolgov990@gmail.com	f	danchik667	t1ktac	ИУ2-14	Долгов Даниил Дмитриевич	\N	\N	f
d14eb1da-2441-4939-b9d8-47c9c79b6a55	$2a$08$/7s6N4GEHvoO0Cvz3/wBOe4eH0e3cMCNt3hh971A9PD.8ZUU6dlhy	89160932860	sererga115@gmail.com	f	tyapkin_s	tyapkin_s	РК2-11М	Тяпкин Сергей Сергеевич	\N	\N	f
b237c4c7-b549-408e-8bd0-f2a42bfb937e	$2a$08$u0h53Ayg7qa3hcm.Sm1NAOCuIxH5UEcHd6.9Ss9yRe4d38NgNeRpu	8903-644-12-07	tikristina53@gmail.com	f	kris_tin28	kris.tin_28	МТ3-12	Шевчук Кристина Вадимовна	\N	\N	f
c2f239c2-1747-4e6f-83ef-841fa2d088b3	$2a$08$Zqg73NdzLG12Z5rZqKE4vO2.VyAZNXBuYZEZDsprdZD9wxno9nqMa	8 966 347 42 86	sks23u494@student.bmstu.ru	f	zaneruk	zaneruk	ИУ1-12Б	Синюшин Константин Сергеевич	\N	\N	f
f1ec9aa1-1e0b-42ae-9aff-3ea3d2ed2f1e	$2a$08$rqnGHfNsbvRd33/wkIAZ3.HfcopcJsIZjiKb9BWKbSDhy63Z.dxpS	89636098898	ivanova.ae.mail@gmail.com	f	saanechkaaaa	sanchous1243	ИУ3-12	Иванова Александра Евгеньевна	\N	\N	f
c647ffe8-56dc-459b-bd51-b8d56bd92e70	$2a$08$qRtFagxPMi/ZoVp/nvkxLOxyu3R4hvPYYfogKEp0FOslMCjzABL3W	89779736039	nicitow111@gmail.com	f	fuhrerbaumanki	id815649892	ФН11-11Б	Струк Никита Андреевич	\N	\N	f
614a9f38-252a-4e3d-ac84-b3b705c2bc97	$2a$08$T2j36hn5RFzkNBNd7fmNwenRKzg/nnWRY5//jxIfwnhMHY5p4VnAO	89163900271	tramkatya@rambler.ru	f	tramkatya	tramkatya	РК6-15Б	Трамбицкая Екатерина Алексеевна	\N	\N	f
98ed0c2d-296d-4d04-b7b2-1aa653ee59db	$2a$08$pU4SrYk3JWOa53QX2r2TqeQEBQAO5r3tIZgwEh9zD08HV8mQAdEgK	89634668610	mikhail.zaykov.2003@yandex.ru	f	ZamMagiks	legends	Э1-13	Зайков Михаил Андреевич	\N	\N	f
147914ae-243c-4fa5-837f-c033c541af03	$2a$08$cZG3HXamIXOV1m1ItEltmOoxieVbt.fugut5Er7x9h9dQHIvtggeW	89096822824	eduard_filatov_05@list.ru	f	Atakyn	atakync	ЛТ9-14Б	Филатов Эдуард Владимирович	\N	\N	f
db9ed2c7-7097-4e2a-bf9e-194a7214e66c	$2a$08$kRxwLrJaBfKYx8B9jd.2beSyFD0stpDRXHG4.0PsnxzePsh3uwW1y	89162042918	eqiewu@mail.ru	f	equuuuuuuuu	iwaa_channnn	Э4-12	Розанова Светлана Юрьевна	\N	\N	f
ee7c625a-029b-4fa2-8835-b6f9579ca275	$2a$08$W2XUdBWTVeYvfwwn8ETfde.Ho4dTw2msC/SkyHyBZAl4CakzO6Uy.	8915 64596-36	haik.23.09.2005@gmail.com	f	trusiki_v_nosochkah	trusiki_v_nosochkah	Э8-11	Журавлева Дарья Дмитриевна	\N	\N	f
6be26698-dff0-4250-bf91-2bf5bbf7890c	$2a$08$6TiI6P0BlyHNHhBx2RV6mOVQp9b2u.oym6YoDu38r/oxt3ITKtoXq	8 915 14755-72	zarevama@yandex.ru	f	maryzarieva	maryzarieva	СМ13-13	Зарьева Мария Артёмовна	\N	\N	f
98ed3f2c-3fbb-42fa-8ef7-2907cee4db09	$2a$08$Xcz7j9d3cFoXU89D8VupxOsRWsfHd2OIMLVPX3EL2ccRqP6xD2JvG	89204690091	echernavskiy@yandex.ru	f	ErRoRcH	echernav	МТ5-12	Чернавский Евгений Алексеевич	2b55f5a7-5872-4904-bdea-7d54243826e8	1	f
e159d5a0-8ff9-43d0-a6bb-828f279331ed	$2a$08$BwTSuAllIP10hToc9HMT/eVCvAcDKQcRMqAJj0q2YeAZq4aeeOsPi	89859881441	nikitondyakov@gmail.com	f	e5erg	5ergg	Э8-12Б	Дьяков Никита Сергеевич	6148d3c6-58fd-4ee4-830f-c8643ef8938f	3	f
457d36dd-c827-4fde-9d63-e7fb22c39719	$2a$08$/FXOPEmtS5f0LajwpxEv4uN/GE.d4/zWW3w0dxIl9qKXopBKMctLO	8999-059-58-21	arina.shigina@mail.ru	f	silleennce	velevaret	МТ5-11	Шигина Арина Максимовна	307c360f-258f-45b6-b1cc-81c946df749e	1	f
76571251-d22b-4ef2-a4ca-8227b554795d	$2a$08$Usz6YnGl6wfF0itveu4GC.L0PIv57eo7lasMUpSXsYgxvBbY7CThm	89533243106	rtf1234@icloud.com	f	sorrybutididthis	hellforeverybody	МТ11-13Б	Богачева Ксения Андреевна	0c68c5b1-5185-497b-9a60-b1c1e1763b6e	3	f
85a6d73d-b906-4c61-974e-fab93784b1cd	$2a$08$sY69xEcOh8kix2NyI/jbqe2nlpgGL4zgs9xjL6JMSP29v6xvQKABK	8925-026-35-38	nastyaglazunova60@gmail.com	f	glazunova_naa	glazynovaa_naa	МТ4-11Б	Глазунова Анастасия Андреевна	1c6d8c8a-9bda-4ec8-bbf6-8dd0b83b98d6	1	f
4d9fc711-ba78-4d40-b671-e30aee993c77	$2a$08$BwqHUTzNHluGzYOs3buAfuhXaYOybjjx10dk/V6zGzPYzlczgbJkC	8967-265-57-78	gogolevea@student.bmstu.ru	f	durrisdeer	durrisdeer	Э10-12Б	Гоголев Егор Андреевич	20cd8a22-243a-42f8-8253-5cadbeeae107	3	f
566e7043-7419-4d63-bf9e-d2d593f31ebe	$2a$08$EpRzmPwNp8BzsmjWtpeJ1uK1LQqp.wzXu0/0z8SwJd3LROaR2tlJa	89140946702	voroninala@student.bmstu.ru	f	luiza_vv	luizavoronina	Э3-11	Воронина Луиза Алексеевна	485a9afa-6f4c-457b-952a-203d110898b2	1	f
5b22eac7-9e38-4d73-b627-325e2033bcee	$2a$08$TsQnqQQ05g2imak8DGSNQOyvp9qzHC/ZEdD/y.fP/wkFipuOvpED6	89160331701	chirikov_n@list.ru	f	k01yasik1	n_chirikov	ИУ7-16Б	Чириков Николай Владимирович	6e7bc851-c93b-4c94-bb2e-fdb0656c528a	3	f
4859ada4-aa2f-442d-9bc9-0fbd479056cb	$2a$08$UO9KudxrcvkhD3opEq8.gOZYtIxQvgnZ/lvRoP2yp9bC7JobWxOca	89152679919	laymaarhiv@gmail.com	f	bananoviilime	id470323592	МТ11-12Б	Петракова Ангелина Антоновна	0c21a026-7815-4005-ab4f-ac2b5192acb9	3	f
117d10ec-7bf3-4b1f-8cb9-c4b69ec28819	$2a$08$H9dZ.ChRzpHTM/I7I1yhWOfxKONlcXOZL/70h14ffse3oMQyUiOnC	8965-307-25-85	oyuv23u604@student.bmstu.ru	f	yyuulaa	orlovskaya.yulya	ИУ4-11Б	Орловская Юлия Валерьевна	a5ea392d-d4f6-4c87-a04e-002e252cfb6c	2	f
9d21f760-6b3a-4328-b670-28633fe6812b	$2a$08$mz6MLeVKlqWkgpDwE16SJO6nMB4PZASl7X5ES5daAx1/7DLdDceC.	89500934627	maksbarslq@bk.ru	f	MaxBars00	oomaxbarsoo	ФН2-12Б	Барсегян Максим Оганесович	deb66347-3dcb-4a55-8e42-b28de913c9b8	1	f
bab60cb7-29b3-4767-9253-79eeb46f6301	$2a$08$p5VnH1J79a3LHx9YeRvXJOD7dd2UvYVO/PDyOMQndklNheRifGCfy	89653245125	asya1362@mail.ru	f	BlumWinxClud	1pomidorka7	ЮР-14	Салихьянова Асия Фаридовна	712edd72-aed6-4ec5-a521-cf4a8bbbf4f7	1	f
dea2862c-a36d-49a1-a963-c37e5b56c584	$2a$08$NwyWGBqREdHKkMWFQTpu7Oj3ZT0CiHtgGcwQKGv/Oi64B8HfcHHhK	8977-129-52-04	smelovavaleriia@mail.ru	f	s_lerchik	s.lerchik	ЮР-14	Смелова Валерия Романовна	712edd72-aed6-4ec5-a521-cf4a8bbbf4f7	1	f
f850ef10-27a1-4e5b-9bc3-19064aa09112	$2a$08$qdIFTv8U.x9HVaQKm15b6uv4X6OJpdXqIuoTiQAmKf.lo4WbDtJ4.	89001358533	ksdukhopelnikova@mail.ru	f	ksdvsks	ks_dukhopelnikova	Э4-11Б	Духопельникова Ксения Витальевна	95fafbf5-1bbd-4f58-9ee3-1e1cf6eff939	1	f
35fa2bc7-b5a3-4172-a750-641a9dfa10c8	$2a$08$.ySEeqDO59S1xM7ZVvZjsutKEVYV3HxfKSdI31FzKRBtRwCBimFBa	8977-819-07-30	asinka181@gmail.com	f	asya_k21	asyak2005	РЛ2-11Б	Колесникова Анастасия Петровна	543a4e0b-e35d-44f5-85df-43431142aab6	1	f
4e91ee8b-eb1b-4db5-a4a9-72f6c9c8a099	$2a$08$T3QL9IgbZVYcXxQ/wF26N.OpMJUgVkIZM/FsWVDpR.SF3YldaSGT.	8915-338-33-89	i@akuznetsv.ru	f	Kutuzoz	russiaisthebestcountry	Э3-12	Кузнецов Антон Юрьевич	589132a7-7358-41cf-833c-7a5bd9f8e1b9	1	f
b35c69ca-c899-45dc-8b2c-c58429faaa5a	$2a$08$vwaB2e4LlDundt0/5xQPLun3y2GvV.ZfdM6LTmj6vqmhTl/rDjKpS	8 916 83880-11	n1k1r4c1tor_5002@mail.ru	f	True_Dark_Lord	true_dark_lord	ФН2-12Б	Чернов Никита Владимирович	28fd4455-c957-4065-97f6-847b632daf2b	1	f
e387ebe4-26b5-45a5-ae3d-e0b01dc44b8b	$2a$08$kq4z.l63ESpvaGNYcp49y.ZhR7UFMDoNb48QR5hQ3w6nqh7pccXSG	89153801425	velichkovskayass@student.bmstu.ru	f	Quinfers	qqffff	РЛ2-11Б	Величковская Софья Сергеевна	543a4e0b-e35d-44f5-85df-43431142aab6	1	f
259e0287-3ad7-4a71-b153-50926c828524	$2a$08$sn/Dd2DJwTRgrbsR5t8BDuqOTa7IOO3auMRLGokGEnTJ/QTAWPEGq	89290309730	bazhukhinaev@student.bmstu.ru	f	galaxy_v3	galaxy_v3	МТ9-11Б	Бажухина Екатерина Владимировна	d68a442e-2c48-43d4-b184-15596fc420ec	3	f
853daa34-225f-4849-a3d9-603fd3987f5f	$2a$08$k9kvnqHlN06UMzDsVQiko.O2wSs5lUyZeNKNcM9bfDwbVHV8hGNmu	89179373024	tim-tigs@yandex.ru	f	useinac	young_placy	ФН4-12Б	Шайхуллин Тимур Рустамович	ccbdb464-37b4-4fa5-88e0-8a97ea84dceb	1	f
ae2ac781-fdd1-43b1-b22a-266098f1db9b	$2a$08$8O47dvMqxfT5QnJ47j2HeeAYSFV7gA043QKP0.kyMBvyZJnQm1w0y	89125516312	ostapenko.diana.00@mail.ru	f	eeveryonehass	everyonehass	МТ9-11Б	Остапенко Диана Алескандровна	d68a442e-2c48-43d4-b184-15596fc420ec	2	f
9c992c2d-d60f-4c57-b08f-8f3ee19ce341	$2a$08$pWAcmaPXk5dfGBFAdvgwCumJq7Gp/CndPqpKT.HHpCw7DSLSoisoG	89372814562	veronika.fr2016@gmail.com	f	vvvvveronikaa	id307556729	ИУ2-12	Фролова Вероника Романовна	\N	\N	f
e6a771f7-7131-4980-9169-932e5c689271	$2a$08$URsP/4bmavZvimrOSwE3t.l3tA1TqMAlG46YSi0IQuE2be1GQYxF.	8925-381-61-65	serfirefox@yandex.ru	f	mishashtick	mishashtick	РЛ1-12	Шляков Михаил Васильевич	\N	\N	f
bba08d7e-3a04-4005-91c2-59fbfe08033f	$2a$08$YutLnRezAMhqckjLCHYU6u/ei2b4mWSZ0zW8U5VciZRD6UaTFEWh6	89058828565	vanyatep25@gmail.com	f	grinding_nerves	grinding_nerves	ФН3-12Б	Тепляков Иван Александрович	\N	\N	f
b55c0799-0823-4270-aded-5afe7af0b889	$2a$08$6S4sZHeK8lIUO/sdh1uPnOncM9wcV.xYJI0..QZJogvGiAWujPjJq	89153879164	alisamilshtein@yandex.ru	f	Alickich	milshteyn3	Л4-13Б	Зеленова Алиса Антоновна	\N	\N	f
7d5b97be-75ef-4719-b5bb-a3c6c0e449aa	$2a$08$9VR8X1wv44vgwm5PJdIsh.lgAHr/YtqdFAKDXU3zJ3TYAwxVbhHAi	8916-963-00-13	frumartem2005@mail.ru	f	Kozmikq	artemfrum	ФН12-11Б	Фруктов Артем Антонович	\N	\N	f
5b92b48c-894c-4eb1-8273-287a9b2c030e	$2a$08$cOTiM16pu5smbrfxpSHqJezvmifAqZ.JyfBXxszp5hXFyRYMkeMHK	8 926 84740-21	mariashirshova.2006@mail.ru	f	maris_sh	marishirshova	РЛ2-13	Ширшова Мария Владимировна	\N	\N	f
68f149b5-e4d5-4e1a-85cc-c25a52f50a5d	$2a$08$sNO5eQ97JFd3WRsSnYvv1.oAmXpoc97Z3FgKUlOsABa32aeqrNlFG	8985-122-55-81	db.planshet1@gmail.com	f	Nastinka_production	nastyuha.production	ИБМ6-15Б	Кошлакова Анастасия Дмитриевна	\N	\N	f
43308f69-2086-4953-a533-e0429f28ad32	$2a$08$glrGSC5xpt6MCPMNJwbl7.mT.2XJU.SA9dO6jphJtE7IDdHL4osrO	89151048659	lera.vdovina04@mail.ru	f	ktoktok	tototowo	БМТ2-12Б	Вдовина Валерия Сергеевна	\N	\N	f
25baf4bd-a819-4d86-b4ac-88e7bdc3fb46	$2a$08$eAp6sHI0uSe/4FscvAaOqOdKqJ4fkt/CukAHElocOrfTE74BUBDhi	89627442967	besaeva.zhanna@gmail.com	f	waitidk	waitidk0	БМТ2-12Б	Бесаева Жанна Алановна	\N	\N	f
cd64f3b4-5d34-48c6-8932-6e5f710f0bd6	$2a$08$znNCylAN6tASb8C6doU7A.jbvM1EeMsk7qHrnFHQ8rv3mwUMjSjdm	89887976400	zi.shamilova18@mail.ru	f	allmllm	sentiire	ИУ1-12Б	Шамилова Зиният Мурсаловна	\N	\N	f
23854d04-6759-4c0d-af67-f82a85d90e8c	$2a$08$KqAsCOieVTuJVYYR9S2GNOCymgAlSBUsEt1vepm40WpBRfTkNzle6	89034772277	bigunilaevah@gmail.com	f	khaidikbig	khaidik	ИУ7-11Б	Бигунилаева Хадижат Мурадовна	\N	\N	f
ab95256e-9e46-42d9-a613-396b94ed31b2	$2a$08$90mx1EtJvtfFT6bCTnW2s.ot9jxvwkIud0ZkT3Y9gOjJDelKPg0fy	88005553535	meme@bmstu.ru	f	memes_bmstu	mmmavrodi	ИУ12-11М	Токарев Адам Антонович	\N	\N	f
905752a4-1b80-4cd4-bd62-6772322d7c6f	$2a$08$P5YxR/K0McZKzwDIiEB2DOoB5.0xtPwkNJfUeBcjM2fPmBUZ4afKO	89624370769	2006orex@gmail.com	f	mnogotochnik	mnogotochnik	БМТ1-12Б	Шеболдасов Дмитрий	\N	\N	f
f30f851e-760a-46e3-9859-5003075c652a	$2a$08$lFNaL8KdrGc0Qcw18QjQgudWhzx3EIyzY2t6uGmXEc6e0lquAUxqi	89397930505	yura_iz_ulan-ude@mail.ru	f	yura_kuklin	id678844646	Э7-12	Юрий Дмитриевич Куклин	\N	\N	f
92f4bf69-84d5-4a1c-b0c8-0fbb9d9452a2	$2a$08$SGFoRKXttGqHzwS.QJW4MuYc1Zm/3Eygas9mbPgLDSEsqeO1pt5ju	89773823042	vasilisainshakova@mail.ru	f	vassilisaaa	vasilisa_818	ИБМ5-11Б	Иншакова Василиса Сергеевна	\N	\N	f
953a2af7-e09a-4b7d-a200-7e0a76871f45	$2a$08$WQhIBXy1F.29kKpyrOmGoeDGkW4RURuWYbog1/UBj.Y04lmut2DsW	8 916 85060-55	makes.k.s@mail.ru	f	beamz_xxx	beamz20	ИУ8-15	Лозанов Илья Иванович	\N	\N	f
40faec13-2e5f-4a88-a2d8-022f571eb6de	$2a$08$.WTgRqjHhxwQbgB9T3sfxOLrWVs3B0WQUbQ6GvXOFXsRkjilMl6HW	89963253390	renduhmax@gmail.com	f	weqsapo	weqsapo	МТ12-12	Рендухов Максим Дмитриевич	29103222-c68f-427b-ab8c-2cb674f1fb97	1	f
29233546-8e21-4c49-83bb-f493dcfa037f	$2a$08$1oy4LsqZoEgd7bSocE3Yk.zrDW4J6lAE/N6xMKEGVC2PbmmnGKGjW	89101560753	max.popov0110@gmail.com	f	keep_the_v1be	carfb	Э7-12	Попов Максим Евгеньевич	17beb218-e4fc-4644-935f-cfbcb462b0ec	1	f
cd744cd2-7bea-433c-90c6-f4ac63de4475	$2a$08$0NwlvjuzSaTDIFSIelj9/OqSH/v4DE09pNbwlyIvHcP0DAkImd9ya	89773136652	fenin_antony@mail.ru	f	anton_feni	hinnini	Л4-13Б	Фенин Антон Дмитриевич	8f8f8c5c-9df1-43ca-aa12-6eed5edfea17	1	f
3eff2dda-c5ce-4559-8e27-2d9d915070a1	$2a$08$1kI9FtNLJ7XK/s4MWfZMk.nGnzeRXch1Ebx6oeOu4vOcQ.ibjuZKu	89190488873	nozikoffmichael@ro.ru	f	Mishgan_228_smol	nozikov_67	Э7-11	Нозиков Михаил Дмитриевич	16956eba-5185-4ba2-bce2-a8239d6a89a9	1	f
c5dff961-e039-4eda-8b28-905a3bc21d8e	$2a$08$z2zdvMleZwJaX6ML0nX/Z.dO/4Ex/qzYvqCQvanA1/LHs5DPP0.8i	89150367059	r89169404216@yandex.ru	f	General055	general_055	ИБМ6-12Б	Плешков Ростислав Александрович	572bfabc-f038-4482-97a6-cefd372cd700	2	f
bdc6af74-00b3-434b-8883-346184bb5c24	$2a$08$1xQ7Hem7S78vZPBeU2qsOuMht9QyqC0oj5facjGizYTu20ZoUBrB6	8 950 60581-99	ilya.milpops@gmail.com	f	niger_s_getto	niger_s_getto	Э10-11Б	Милов Илья Вадимович	\N	\N	f
eb924338-c28a-413a-8183-4d3e9e8457f9	$2a$08$fqKarbi/yKKoICn.rvpLUeg34EqoDjJUvJxFq6RUqsqG/PptvJPbe	89917246288	svetkako00@mail.ru	f	svetlanakostushko	svetlanakostushko	ИУ8-15	Костюшко Светлана Юрьевна	\N	\N	f
b9c10bc1-5a12-44a9-bbc8-727b67a56cbb	$2a$08$/uliuNQcG4WQnIImFKj0q.4JhE.3vn6b4Z5fzyMdw80D0VxlxoPlm	89251430548	frolia2005@ya.ru	f	Darl1ngzzzzz	darl1ng88	Э7-12	Фролов Игорь Алексеевич	\N	\N	f
8108e4a3-7d7a-4454-88d0-9f9537f21be4	$2a$08$P2KDZjdtv0esWzN28pK.gOrkfRyxLwZWnDDb/N4shBqKk2t22hezu	89853135484	aydargiz08@gmail.com	f	chai00k	chai00k	ФН2-11Б	Гизатуллин Айдар Арсланович	deb66347-3dcb-4a55-8e42-b28de913c9b8	1	f
41671d6c-8191-40e2-ab61-ed184a86539b	$2a$08$eVHKjlqRudmYuiFTS9F8kOuLtB1GSZPtP2fs15BJhSQdPjPZuWvGK	89685465329	marchenko.05@yandex.ru	f	anzhelikapopusk	murlikkrut	ИУ8-15	Марченко Антон Михайлович	2d40c471-5ebc-4470-a419-97d182c2b096	3	f
5442e76d-0aac-4f83-b53d-2c436381c0bf	$2a$08$CAEIMZt374l0zMgZzWjYYu2wNYLZ3kxAvna0CoM2MxrhNiJeYxE4W	891682554-75	anton091005@yandex.ru	f	Toha_kartoha0910	id410116202	ИУ4-11Б	Титов Антон Артемович	83c9e33a-e863-4ef5-9909-2ab2db18b7e4	1	f
3dcab817-e0fd-4c08-81f6-e83f74a46d74	$2a$08$JB4C4F3JpyWTeukVEYE/kuPX8uXsjvoAJEG/c.MUQ.pG8mKyR9QFy	896736507-26	zakirovaltynbek57@gmail.com	f	smnn3000	smnn3000	ИУ4-11Б	Закиров Алтынбек Алмазович	83c9e33a-e863-4ef5-9909-2ab2db18b7e4	1	f
a235e486-fd7c-439c-b093-30efa14d7fc2	$2a$08$M0WZL/c.t3lXng56UHTiFuhtS4t26bOuFKVpMJNZVBIAWWfegtdc6	8903-553-47-74	alibek.chotchaev05@mail.ru	f	Al_BE05	id734179324	Э2-11Б	Чотчаев Алибек Расулович	9e61f1aa-978a-461a-90f6-99808580d083	3	f
0bf149fb-0574-48a8-87f8-334046a240f7	$2a$08$toT2UWP7cnDFFBCTsdlwoevXBlikoLRkI/AuC/gQONjYy0k0a7Xua	89165148083	multifora.menu@gmail.com	f	Diwix6	pl0tva	ИУ4-11Б	Кулинич Дмитрий Олегович	83c9e33a-e863-4ef5-9909-2ab2db18b7e4	1	f
8c025888-5232-43f2-8cae-15a30258616e	$2a$08$AjaW9krpdAQ6qGoNbtDw2.bA3WuFLXKiEtt09KRe483TPfqOEkG7S	89618641171	ask1204@yandex.ru	f	Octopus_dambo	koshherno	РК5-12Б	Куклина Анастасия Сергеевна	ccbdb464-37b4-4fa5-88e0-8a97ea84dceb	1	f
6d872118-e913-435b-9aee-670f3282232b	$2a$08$LHnhAvM.v25v/DB8pEhqUuc8g3sGnzmEOiQ/T9WR0z8DyBfXpGUAi	89653231360	karabanova.alena9@gmail.com	f	Kotality_4	kotaliti	Э4-12Б	Карабанова Алёна Сергеевна	6fa56c83-f29f-4909-9d63-ed00152c144b	2	f
91da8353-c4a6-4880-8fe1-5a04dabfbad8	$2a$08$2WFyCpXX17fSDuzkhO5Fye0vyNW4QD0S6yNyrEWu04K.IKcwSEfim	89991064681	arlanirmuhambetov@mail.ru	f	pelmenioneluv	realmotomoto	Э7-12	Ирмухамбетов Арлан Бакэткелдыевич	\N	\N	f
5193c257-f980-4dbf-a931-1d411b2ea873	$2a$08$7o.44.LmxE5BWJ7R5Mv9hOYTXe.GwJ/voILZcYPIwGIFdcSz7iuMW	89684423388	tema.timakov1968@gmail.com	f	wormgore	arseniytimakov	Э7-12	Тимаков Арсений Валентинович	\N	\N	f
2911a414-eea1-473d-904f-576ec1637158	$2a$08$gt9Fm4155OByyba1OqDbTeWX1wecn07RcTrIaOdwJXbuJXUj9Irg6	89687961168	grischin.egor2001@yandex.ru	f	hitriy_potsyk	grishin.egor	СГН3-11М	Гришин Егор Борисович	\N	\N	f
0da70db0-7968-4994-bd5d-7423be9faa44	$2a$08$8wuhrhq6MFuR9aqXuz6i1ecqEHTt4DNjT9PLSj7aaHM.6VlKPwMte	8 966 35026-52	andrey.mon2005@gmail.com	f	terraningenuity	id607452950	ИУ7-15Б	Монастырский Андрей Алексеевич	\N	\N	f
5f2d5f19-3c68-400b-a28e-a3c91a9e5596	$2a$08$q9ZNnUsJfbW5hHlmGJ4ZBu/mLA3q2yvvfjW9echs2NAguwXQ9NunS	89993785148	dibiyaev009@gmail.com	f	dreamslove000	ne_bouca	ИБМ7-12Б	Дибияев Авраам Русланович	\N	\N	f
f34e6f93-1530-4e37-acad-8edf18c8a01a	$2a$08$gbiEokMoRh418IXuSaicJOznCs/KKBia51wzjKJjFPE.wOhh8UFwu	89151887832	ekaterina.perova.2001@mail.ru	f	urlurlurlur	depressedlysuicidal	Л4-13Б	Перова Екатерина Ивановна	\N	\N	f
959b1f43-1bd5-4573-b08e-fae13a7f6182	$2a$08$t.qbMQ9i26HJD8izcb54cOlirvXyRl8I6gqUGDomzaZKZNRjdpVhC	89104243557	just4n0ther@ya.ru	f	nottabaddguyy	nottabaddguyy	Л4-11Б	Даниил Михайлович Михайлов	\N	\N	f
61471c24-eb20-4fc6-a377-d3923d40d4c8	$2a$08$iU4UCC64wp/gDAZleBm2BOWXmmBhPMSFMaLTanV75nAtOqNdpHT1S	89151196668	lublub.f@gmail.com	f	FLyuba999	grandma_lyube	СМ7-14Б	Федорова Любовь Сергеевна	\N	\N	f
e7f897e8-7cbc-43d6-96fe-cceb27aeb725	$2a$08$PBA5pyTRNS1enAqeL9cycerEGp8fyHep/KQDNldLi1NfwPaJRxS7S	89153143090	yurybudanov2020@gmail.com	f	imya_zahyato	yurets404	СМ6-12	Буданов Юрий Владимирович	\N	\N	f
68bd8a51-267f-4ac6-8d65-e8456a8cebba	$2a$08$ZD.0EDXfIRZqsrpcgiErBOQid9ObyFiRVqeVYRLkwNbU5NDl7P/Za	89775242789	samaraivan2005@yandex.ru	f	samarin_ivan	i.samarin	ИУ8-15	Иван Алексеевич Самарин	\N	\N	f
8332ad8e-f895-494d-afd1-f5779d3f78cd	$2a$08$7F2qH.YQOYNjEFqElVGhGuYhSuliZVK/J8tlb7P.u5KJN0/xJWo6q	8999-862-36-12	vita.kopylova712@gmail.com	f	votetorofffl	votetorofffl	БМТ2-12	Копылова Виталия Александровна	\N	\N	f
610f0b94-b6c2-4c4f-95b9-04ac4af83dc9	$2a$08$Tbm/BG1b767HhO26bnpbe.VquIj5wrnq.w51WI6kt82bXTyQo69Fu	89163803279	nika15082005@gmail.com	f	Jessy_Dess	veronica_des	ФН11-14Б	Янченко Вероника Алексеевна	\N	\N	f
3ab7c152-b7c5-4237-ab2d-6b317f9df007	$2a$08$AFiv7334g/EPbTp8ZD20kO02n9lfsDquWriO3B2Iy1p491CDBrfgG	89653400520	rozanovki@student.bmstu.ru	f	tipical_kirill	id512076816	ФН1-11Б	Розанов Кирилл Игоревич	\N	\N	f
f52c3b76-fe95-4b1b-ae20-c7be1e26e786	$2a$08$Qkz5MG4wFWgd.CwqvAXoiOsh3aRCjp7RA9egK5Fnw7evenQHNPKfy	89776452199	thesaponess@yandex.ru	f	w4nnasleep	wanna.sleep.forever	СМ4-11	Шевцов Владислав Сергеевич	\N	\N	f
d43de75e-d7e6-40c6-b9e2-d2dc26aed331	$2a$08$4f/2RjUICnQT2YfF4Yddo.0n1sMfZubr9W4RIpITqJsaUywHpbUM6	89270957868	nikita.nikishin2014@gmail.com	f	nikita50000	id389378135	СМ1-11	Никишин Никита Юрьевич	\N	\N	f
b340595f-3342-4933-9245-6f00277dc58e	$2a$08$NYd43t1w8Nu0nTD2wzapmujAEUbDTaTO79NPMziZyqbTkkxGRnzKm	89166230794	alexgyz2525@gmail.com	f	tenshiqr	tenshiqr	Э4-12Б	Грязнов Алексей Евгеньевич	6fa56c83-f29f-4909-9d63-ed00152c144b	1	f
76ee3743-c9d2-4e83-97e0-897da10cb3d4	$2a$08$tzIeAJRA5kkNJsmVx2.tX.HRXLZW592tc6PkTlntrEN.sxUWhSIvC	89685808056	khabievadm@student.bmstu.ru	f	hab1batiii	hab1eva	Э2-11Б	Хабиева Дарья Маратовна	9e61f1aa-978a-461a-90f6-99808580d083	1	f
0b25cd45-979c-42f0-b35b-7408986664a7	$2a$08$vGesvortNfapmZr5ezW0EeSkGRHyeY/mbgjV4lzfKbTVWxZGt4uui	89261615353	matveyschenikov@mail.ru	f	Parsifal_2287	yukinomercy_bmstu	МТ5-12	Щеников Матвей Сергеевич	2b55f5a7-5872-4904-bdea-7d54243826e8	3	f
c9f1836e-1868-427f-b6d8-fc2f61155722	$2a$08$IlARZvvMRllqKqi.NVd2POtNVaVm0dkwYy66RoaZNk6ITbBY0jxwe	89252820919	katushakopilova@gmail.com	f	k_copy	id481512260	Э7-12	Копылова екатерина сергеевна	17beb218-e4fc-4644-935f-cfbcb462b0ec	3	f
289ad01c-d7a0-4afa-9a81-f6c4dc8847ad	$2a$08$P/Vx/SvqY2WWDDAu7gpKauXNo8NfbHYgDZeC7aNwBzaaHKgYyDgDq	898571528-15	2aut@mail.ru	f	panule	medol08	Л4-13Б	Медведева Ольга Сергеевна	8f8f8c5c-9df1-43ca-aa12-6eed5edfea17	1	f
33455292-fac0-4965-8b29-f40492823699	$2a$08$FAOfTEpdPjLIxutUcPM6HuEE0qS6l8JkisHBiShL2IxgKR1OZnA/q	89856119393	anvol879@gmail.com	f	mmacardi	mmacardi	СМ13-11Б	Волков Антон Владиславович	2b3b1e0a-1f6f-4f7d-a0da-7877dbf854f5	1	f
1305e247-b222-443d-8759-51bc4cea9eb8	$2a$08$fazHHMzKO6o4TmDhAK66muYgeUIhyp1Pjy9UMTtSH80gELagS/VXC	89605273508	mischenkoivan567@gmail.com	f	m1_van	id599045444	МТ1-12	Мищенко Иван Сергеевич	\N	\N	f
01e59c2d-ee44-44b5-8d47-e8008ade609f	$2a$08$4mWvZqC9ndc1WpPwqYbs/uJ2k4qKNSfduFVVLOQQGQFVulYAmgRa6	8 920 945 69 40	averinaps@student.bmstu.ru	f	polissazzz	polissazzz	СМ7-14Б	Аверина Полина Сергеевна	\N	\N	f
4fd19647-808c-43ed-ae4e-f25dabc9e226	$2a$08$yIhhRIr/U7tBK//7Hbjcoetv24JDjIQJZwtbHK0U4uQ0GGeZP5ZfC	8951-423-16-91	olenev.art@mail.ru	f	Arrtyo_m	artyom_olenev	ИУ6-15Б	Оленев Артём Андреевич	\N	\N	f
a2ff656c-8b03-4579-a732-c8717e64ade1	$2a$08$0ZA0gx1SK6GA3RgCk4Xyru2Vc.2abE86BeQkE5/pxxONmaJgmtdUK	8915 3836093	mayorov_slavik_2005@mail.ru	f	KotPerdolya	slavkabulavka	СМ2-11	Майоров Вячеслав Владимирович	\N	\N	f
ed1f78f8-c9f6-4469-bb1f-d56eb9d4f84b	$2a$08$zLtPoW0.NKjOBNAq/gyNBulaCwKyngYuBq7hD1HjcNbkKCWQ4OlPO	89506808955	zurab.bregadze1005@gmail.com	f	zurabbregadze	zura11	ФН2-12Б	Брегадзе Зураб Гелаевич	deb66347-3dcb-4a55-8e42-b28de913c9b8	3	f
8aabfb91-3778-4912-b733-2843367bac7d	$2a$08$yri4VoWSW8sjabCuvrobROCD0//83gNHjWZpvYYqhvPMaMProEV7e	89858914945	ermak05kira@gmail.com	f	kira3009	kira3009	Э4-12Б	Ермакова Кира Александровна	6fa56c83-f29f-4909-9d63-ed00152c144b	3	f
c6b956ac-41b1-4ca1-950c-ef310837503a	$2a$08$Vv9YnSUHMsD.lYwHOhdpqOlsk0sDW1XDyMDkGPZmI.2/ELjcYRESW	8 929 69510-20	copavel@mail.ru	f	kavevpokylol	kavevpokylol	СМ13-11Б	Копылов Павел Антонович	6b2934e3-98fb-4a5a-9138-35a62df992b9	1	f
f3ea6021-5fe6-4949-9308-fc1e11c58c61	$2a$08$v1qmgoIAW7Y6TIUhDNbFXOk3gsLrv396oFVWR9i8Vp4kAmrcjNpCC	89661201636	mariakuzmina05@mail.ru	f	lnlduck	lonely.duck	ФН11-14Б	Кузьмина Мария Андреевна	73ce80d7-2798-442f-918e-bf40511dfb07	1	f
969f8276-8305-4541-8b48-6b9d0591fa7f	$2a$08$VYfeysR3RcVNrY.yfBTAN.XfbCSBx4qUTJExUqFTDzSruwwYt3Rje	89683355244	privetprivet29674@gmail.com	f	xuxund	cockleto4ka	СМ3-11	Яровой Илья Александрович	3155c953-27df-4bf3-93a7-3c3cce34deb3	2	f
fd893f63-2810-427e-9de9-20f5c2efc70c	$2a$08$nUXC5wsO17Yr5TaecBCa9uDzwNXSkieq15iY4q/GaWA12889s9CPa	89128748055	ktoykina@mail.ru	f	plmsk	ktshilar	Э4-12Б	Тойкина Екатерина Васильевна	6fa56c83-f29f-4909-9d63-ed00152c144b	1	f
67f2d494-3756-4f9e-8866-0ac897489f36	$2a$08$brRy64h.utAsWjZ8PP9mm./.5js68rCypQh9KgJtCpdH16dtTk2jq	89250235744	lvv1401@gmail.com	f	vikki151114	id280231084	Э4-12Б	Лебедевич Виктория Вадимовна	6fa56c83-f29f-4909-9d63-ed00152c144b	1	f
890eb389-24fc-45e3-9f41-c23a0983ab6c	$2a$08$ZCLHqueSlaenLofdr5TVa.r3sKEhbiKHCIk2UC4adj42WBqU3YQ3K	8 977 57137-12	alexander.pliukhin@mail.ru	f	plulex0	id507639252	ИУ10-14	Плюхин Александр Сергеевич	\N	\N	f
3ed25164-78dd-41fc-925d-72679ee88ab4	$2a$08$/pE3uPWywT81P/WtJfQmh.ZXAyREIwZv8rT1xsdxbx.eHIQSctDLu	89092066542	ilapelihov5@gmail.com	f	ili_y_a	id416708602	ИУ3-11Б	Пелихов Илья Игоревич	\N	\N	f
6a5c3e94-e81f-4bd3-a14a-650c46fade44	$2a$08$oxyoKMeb/tNOSnVgGDEUrubBuhVr4nk6/jctaO8LVp8G9GxGcigNK	89258339877	saprinaleksandr15@gmail.com	f	lp_morginal	super_kontik2005	РТ1-11	Саприн Александр Андреевич	\N	\N	f
ad3ceea6-a70b-4433-acac-701a381963e6	$2a$08$rNkbeBX4BuSoMD.s8mjIX.E7.8i2LiW8.thGKoK92yHMHIyJS62d.	89379807068	yarik.chernov.05@mail.ru	f	the_himick	id594970353	Э8-12Б	Чернов Ярослав Дмитриевич	\N	\N	f
2b91f223-30dd-4ed5-9310-31e1d69eb81d	$2a$08$OuJdI3IxsQF9MzrUN4lHhOF37OrnKdfZMh12TVHQh1EAJiZ.IWKCq	89254490115	mr.krasava34@mail.ru	f	HOLDITBABYSOHOT	antonisan	ИУ6-11Б	Ягубов Руслан Марикович	\N	\N	f
cfae0fca-a6ef-4741-b191-2010edaafee0	$2a$08$MDRyN3qd9W7Mxv5YE5owfOaMI7M934hXxPScGsOJQpP/BK7m2YMVG	8 916 45703-62	misteruniversal1606@gmail.com	f	Universal0016	misteruniversal016	ИУ3-12Б	Сосков Сергей Сергеевич	\N	\N	f
579c5a97-1de2-48d3-9325-1a78519d1fea	$2a$08$0BRy8U5cQ9STMKbpYRX4euoXyoTmmSdGDuezoyOk/QlHbTLuFlhQG	89055886045	v2lentin2005@yandex.ru	f	just_VaIya	id567242157	ФН3-11Б	Киселев Валентин Валерьевич	\N	\N	f
83305b4c-fa1b-45fb-a595-cbb198a961fc	$2a$08$fL2QwgemUlZC0d6QvXHUoumdLv45HXchtdj9wr.ilPSYE1GQFbf9O	8 985 11973-72	slava.korovin.05@mail.ru	f	VlAdiSlAvkOro	vladislavalb	РК5-11Б	Коровин Владислав Евгеньевич	\N	\N	f
cb540cf5-efc2-4a4d-9226-75215c764255	$2a$08$Vi3.hM8VSZIGKX49lpxaM.nZ2HgqsFHdwIERGTM8Hm1sh7SMqLuNu	89107458853	faton2281337@gmail.com	f	Ebuchiykolhoznik	ebuchiykolhoznik	СМ1-11	Ставер Федор Николаевич	\N	\N	f
4367eff3-3a08-48a5-9703-3d6679b2b11e	$2a$08$Wo2vYEjR9xHsHRO0bibfZeh7fn2pU0nwwpZ6Ii63qKO/1FNIw0X0O	89118896628	arturgatovsky@gmail.com	f	tapalil	id250037973	СМ10-12	Гатовский Артур Антонович	\N	\N	f
5eb4fdb6-144e-43b4-9faa-d9e26327acf6	$2a$08$x5/EGqHMUF9NFPVCaS8iZOuw5QLD/yHggJwH1ytL/TIwX5iXVYEHa	89263403909	miknik666666@gmail.com	f	Mrs_NiKeR	mrniker	ИУ5-15Б	Герасимович Никита Иванович	\N	\N	f
0be7459b-b007-4e4d-b5b7-707231bc5a7d	$2a$08$uq.l9q0Oqe/LJHX0EFKKFuIpkW40PIme0i/i/bNK5ugFxJOOLTjje	8977-880-43-30	krasov.i2005@gmail.com	f	o0o0we	krasov_i	СМ8-11	Красов Иван Алексеевич	\N	\N	f
30eeb430-51ce-49ee-a015-7b8fad95e8b2	$2a$08$ZePooU1CthpKOHrn8Q/5xuNds/sUlO5DRmjAty3onYDqOLLI4A9FK	89154804484	dima_petreev@mail.ru	f	Charlicrown	charlicrown	РТ4-11	Петреев Дмитрий Денисович	\N	\N	f
c8f5be62-d887-4f74-a63b-176c6865cf24	$2a$08$PUhStUe/CqEraLLanibUyuyJ5fCAW1umBugy8A02nA65JuR3cA7ka	89195667850	mi_k2005@mail.ru	f	MI_K2005	mashakorebo	СМ3-11	Коребо Мария Игоревна	\N	\N	f
0268e91a-cc13-400d-9385-285101b57da6	$2a$08$YvJpQdU1eBKdmMW8dORSPOi.xpoAvmIwWjk2i2ZJmbo.yInQXCc6u	89312447865	an1kina.olha@yandex.ru	f	Olga04anikina	id347778302	РЛ2-11Б	Аникина Ольга Игоревна	\N	\N	f
33dfcba4-d298-4980-b163-467aa48fd2fa	$2a$08$9v.RcEOx6koyxNDizl2TEecJY28aYk/n8.XoUoG22P4k0dNcBEW5O	89096322014	nikolaeffedor@yandex.ru	f	TalonDueCouteau	id563550529	РЛ6-19	Николаев Фёдор Антонович	\N	\N	f
3ca00e59-0c69-4e44-b550-2cae4536417f	$2a$08$GC6rbcL7O51oUQS62ISJKO2/aUA4HIG/XcFTy1RAy01TY/TO2m/FW	89026405482	5b_alexey@mail.ru	f	legents_bmstu	qw	РЛ2-12Б	Емельянов Алексей Михайлович	\N	\N	f
ec9b8aed-aa86-4088-9ef9-c618fe073d65	$2a$08$Y6sc74exUV3VGfYLP8inF.DRULjfKTpI4CX9eJqMqrxGaE4lmKgwG	89158184896	gleyzx@yandex.ru	f	GleyZX	zxgley	БМТ1-11Б	Лукьянчиков Глеб Михайлович	\N	\N	f
c4fc1eba-b266-4b9c-990a-fc9b516f488b	$2a$08$unfXD3poRUl7ProFDjhodurGwdWtieiXBfIyOUxODHDSCwazA7sUy	89997813143	melnikovav65@gmail.com	f	lichinus05	id328809417	РЛ1-13	Мельникова Виктория Александровна	\N	\N	f
00ef4554-dbdb-43a6-b6ec-6ecc997d3709	$2a$08$rem0itgu0QIjk.sDhfqCluOp9dm0Dxo37tAQ2Zer26uaVPqLRLiZq	89951127118	acsis144@gmail.com	f	akc1s	csis14	РЛ2-11Б	Алферьев Александр Ильич	\N	\N	f
740e2ec7-23f9-42c9-85a4-6c0c89a6c76f	$2a$08$T.XwUT6K0Hp0dkeMEBnA3eLPeNbZVx9Enx8RmR/AkaBFIHPHaTLGS	89851200590	vernaya.de@gmail.com	f	yeteer	yeteer	БМТ2-11	Верная Екатерина Ярославовна	\N	\N	f
3cde0a68-8d72-4c54-b629-a84e53207b32	$2a$08$pQUAx3JmOQOVrUwgMeYWlO/hMThY7s3csDzb/dAkHm5IjVp080d3.	89842630151	gogaafan@mail.ru	f	Georgi_Ferrinio	georgii2019	СМ10-11Б	Афанасенко Георгий Игоревич	5095faa5-51b3-4dcc-90cd-744b3566ccf2	1	f
1ccbb879-47e6-442a-b4c3-7ae8f05a8dcb	$2a$08$PgjGiCOP1UJU7h15pHlld.Oyr1lVGJz4taQtN5fzQQLs6W7itg992	89677930333	ksm485@mail.ru	f	KekereKekek	kekerekekek	СМ1-11	Колесников Семен Михайлович	\N	\N	f
4639d652-feef-4ebb-a53d-938652c2f8f1	$2a$08$nws86lTrV1HKSy9Ae5kU5epBnFsZmEwnXC8Bd7RoJG0Lh8GgY2dO6	89995554149	cristalixxx@list.ru	f	89995554149	SatanIsMyGod	Э1-13	Щербаков Григорий Владимирович	\N	\N	f
238e9066-a81a-4331-b40d-c885b818ec6d	$2a$08$kCGgw9daIii74069m2lQSO3PAQrmOu.fK9csT0toKN7ljy/Pb87Hm	89522914563	s.matylenok@gmail.com	f	Cezzarius	amd_phenomii	ИУ1-11	Матылёнок Александр Александрович	\N	\N	f
1df45526-dfbc-46ef-9b10-7aeabb978496	$2a$08$fo3qTb3FaliwsrkcKRi8l.sA9fHWWRNU360ZsjQn6zl0/JofOhftO	89183816113	slender5@yandex.ru	f	jopistok	jopistok	ИУ1-11Б	Ракитин Андрей Петрович	\N	\N	f
d5761890-d625-4816-bf6f-0d5cd2fa347a	$2a$08$GIifuhsDDC7tO2ccx2VdFOLD04VTvaE6ZMtcOXSv1AT4C/ByPTtb2	89827550170	ekatshat@icloud.com	f	sfloyx	sfloyx	Э9-12Б	Шатунова Екатерина Андреевна	7b146ffd-3941-4851-ab0a-e557884042d4	1	f
c8733110-e937-40e0-afb7-45eef919768c	$2a$08$oZoRG8sbzeQik1uYN0JOk.Hz.K2wjNXpVtdyxmxOnrQyS/D2PswqO	89265753536	drrnkn@mail.ru	f	drnkn	drrnkn	РЛ2-12	Доронкина Юлия Олеговна	01dec7e5-69f0-4948-8adf-16def255ab41	2	f
54829277-5252-4900-aa34-3a1c93a81caa	$2a$08$jHat1k7fgOXWtBt6gM307.C8Kb0/KfnOsHdrN64lm8jGZBIuaa63u	8925-701-22-98	kozpavel05@gmail.com	f	pavelkozlov_ae	pavel.kozlov_ae	ИБМ6-14Б	Козлов Павел Андреевич	4fb50154-1111-4f44-badb-e48f83e0d433	3	f
48058ec2-7829-4ccc-8795-634f09a275f8	$2a$08$ckvSX3MtDtWPGZo.zgkuEu1VL6GvQcZGysjc8WsRTjGMBsN/Z9Dxq	89100884171	ryzhov.andryuha23@yandex.ru	f	Andryuha23	dushkinsk	СМ9-11	Рыжов Андрей Дмитриевич	300ef3ae-2e98-43cc-8151-a7f42980de56	1	f
37f3b38d-321c-4788-8ef3-12e2062f128f	$2a$08$9qXJ6TGDPngsVitFLht.rOihbU1St09LJzwfENGtPYMkpOccTrADO	89774771275	anna.pigaryova@yandex.ru	f	sov01n0	atar_v	БМТ1-13Б	Пигарёва Анна Алексеевна	300ef3ae-2e98-43cc-8151-a7f42980de56	1	f
d39b1c15-b1b4-4cd2-a52c-b90a7da3e15a	$2a$08$sXDI7AkVHxzIcpETerxDOeKHMXw36dE4P1wOxRohZYaTbh/DI/rLm	89536298990	matvey3142@mail.ru	f	MSavos132	matveika31	БМТ1-11Б	Савостиков Матвей Иванович	f9edfee0-2a1c-46c9-8518-d5a92b0cff1b	1	f
408e9d88-d2a6-424f-a100-4fccd8d9908d	$2a$08$birNe8siiC8P.kt2wUZ4iOgUnqDxmZ80u6NyofUIHo4uaAN4PcwSK	8 952 50771-40	bariaryjkowa@gmail.com	f	Doshik727	idariario	ИУ6-14Б	Рыжкова Дарья Дмитриевна	\N	\N	f
8c66ddbc-1298-4405-9935-e637f15e8dc9	$2a$08$gsq4DDCHsjxBmXHASXYxdeIIrsDtGTaEz2LxdF9..f5eVp5IoSuUa	89032007260	mishatheg@gmail.com	f	MishaTheGreate	mishatheg	ИУ5-15Б	Шадрин Михаил Константинович	\N	\N	f
e6c12eef-7040-4c68-92b1-5bbabf91e8dd	$2a$08$uGyMj6zlEEUfHITzPhm6g.HnuAhBI4TlijHI1Fc0BMkAcMk/KBpNa	89854365891	geor.avdeeff@yandex.ru	f	GeorgeAvdeev	GeorgeAvdeev	ИБМ7-11Б	Авдеев Георгий Александрович	\N	\N	f
4d18da4e-e2d2-4621-a06f-0e3cb0c7746c	$2a$08$h9iftIBngmGh41/923.vNeW6fn3T9dfFJ9171do6J8pn9ZN6v92T2	89254456262	nikita.petin.04@gmail.com	f	agamemnon_767	nik1510	Э10-12Б	Петин Никита	\N	\N	f
0b019438-7eda-42f9-a03b-1204529f18e2	$2a$08$rf4UxZ/vELuG/Wty.6eY5ORsUeKD2e2NSv.ZqPRj7UNiKFAXOFSUS	89268823582	ctepanoohall@gmail.com	f	soeviy_korol	korol_zamka	ЮР-11	Второв Степан Александрович	\N	\N	f
fd365581-2698-4d1a-9db6-36fc4dc3bfea	$2a$08$d/gqV1EPJvVWqU7s.YovDu/x5mHnzzItHTX6i/uIkBv7gp21oLUkK	89153050206	slr23ae063@student.bmstu.ru	f	plhgx	leonid229	Э3-13	Семенов Леонид Романович	\N	\N	f
881c4d25-9c63-411e-ba3c-244676b1d507	$2a$08$OSwV5oOEOosVZVqcNZ/dMeiNgvE6R7kylcyMSEb1lKmYVGghpI6.m	89530124615	kaz47835@mail.ru	f	micoramma	micoramma	СМ3-11	Казакова Ольга Анатольевна	3155c953-27df-4bf3-93a7-3c3cce34deb3	1	f
68c22a6e-d793-4c6c-b191-a5eb06caad3e	$2a$08$Y8UZ.8xtLs/8pNb4QIM7iemDLevys3trmd8gxNH0Lo2BA5SEiuf1y	8 977 57137-12	alexander.pliukhin@gmail.com	f	plulex0	id507639252	ИУ10-14	Плюхин Александр Сергеевич	\N	\N	f
0331aeff-f438-4577-ba32-c749a074962f	$2a$08$vdGWw6CfarWXqsgKT1doZuLlEmdQ2G.bij7CSEjONS.wNGek3/rEG	88005553535	meme@bmstu.ru	f	memes_bmstu	mmmavrodi	ИУ12-11М	Токарев Адам Антонович	0979f441-ea8c-4322-a670-3bb16cf27f6d	3	f
0aaf7e4f-99e9-49be-a5cd-c8a8a73504ba	$2a$08$6wtmcawqz4XID9xUJI7e1u/hqfgrB2IrANBEEgx/mUNv4XMsogqrO	8977-303-80-77	g.anufriev25@yandex.ru	f	goshaanufriev	goshaspg	ФН2-12Б	Ануфриев Георгий Алексеевич	deb66347-3dcb-4a55-8e42-b28de913c9b8	2	f
985cdefc-aed1-457c-a6c4-b41fe2681cfe	$2a$08$PUODqdcAPKx/XHzjJhas6OsavLHuFqZPbyLrEYOlg6mrCzKa0Sy2O	8986 9881220	prugunov2005@mail.ru	f	prygunov_1907	prygunov_19	РТ1-11	Прыгунов Сергей Александрович	40355e30-8138-4e9e-a09c-6d4b9ae9f4aa	1	f
2471e33c-6e39-431e-9b8c-7155d3c52031	$2a$08$C7w72NNoWIhhcv7j84dCauQR7TO8J3JKIPUl51gVXs3zsbUsTghsy	89524928280	nikitadanilin17052005@gmail.com	f	ni4itos	n.danilkin20	Э10-11Б	Данилин Никита Сергеевич	6b2934e3-98fb-4a5a-9138-35a62df992b9	1	f
28c3221f-e5d3-491b-92ea-859a8d57af2e	$2a$08$Ore6hAs1aStSq3jhL6K2juaAItevcutyu2DlQcLOFTCk810f2bw0a	89604482679	mit23d@gmail.com	f	telegramm	vkvkvck	ИУ4-13Б	Мельнико Дмитрий Вячеславович	\N	\N	f
e3e58a32-1f42-44b3-881f-79611010dc6e	$2a$08$vVB6PZ6sGdo1vSPRBPeO.O/05KzUV.FxGZhIFDsGyrOJuYxhMbi9m	8 926 40050-25	nastuk0509@mail.ru	f	sorika_594	sora9879_deku	РТ1-11	Коваленко Анастасия Вадимовна	40355e30-8138-4e9e-a09c-6d4b9ae9f4aa	1	f
17efec82-60ea-4e53-8625-6cd1a400287a	$2a$08$rYnCW3l5gQghPnK.tg3B5enmV16bTmQBYnNeOo49/JGJe6nbhFJj.	8909937 7299	maslov_srg@mail.ru	f	maslov_srg	s.ks_3_top	ИУ4-12Б	Маслов Сергей Андреевич	6b2934e3-98fb-4a5a-9138-35a62df992b9	1	f
64d29804-7f53-4a17-b4e2-39a23d3dd1b0	$2a$08$y6In6EmLz.53yRcvbi5Ine2xsH2jy3fPPs0s9r59gGcYLOy/61KZ6	89263700740	uhminvladislav@gmail.com	f	vert000	vert000	СМ3-11	Юхмин Владислав Кириллович	3155c953-27df-4bf3-93a7-3c3cce34deb3	1	f
a52a92a4-d604-4f97-ab65-0f64cf745447	$2a$08$XtVCujaMPFCjsBTso7GpBOVQAuCrZuNsRjWXmfaQhIuiwtUK51Ydm	89779669555	2004asat@gmail.com	f	l_kenchik_l	l_kaneki_kenl	ФН11-13Б	Базарбаев Асадилло Анварович	73ce80d7-2798-442f-918e-bf40511dfb07	1	f
6437f508-bc80-401d-87cf-10f191749e88	$2a$08$e3Nxb1408osFMvhMhCWRYOBfxNWyLTYSnYiemsBDCKuX2GcNa1jIq	8901-181-17-16	zver@bmstu.ru	f	lkahses	beb	ИУ3-11Б	Николай Тищенко Матвеевич	\N	\N	f
d6bc8f30-0ac8-4444-8fe8-f2fa4f032d65	$2a$08$GuifsVHqGOMthyf7MF.UTeZKJ22deBmPLF/qvqIJndDHqH6CxTFUW	89511546733	sasha.chernikov10@gmail.com	f	sanechkacher	id454814055	СМ3-11	Черников Александр Викторович	3155c953-27df-4bf3-93a7-3c3cce34deb3	1	f
e8895401-1e87-4713-ae9f-1a224f40c1a3	$2a$08$YnD36d5z7wr0qEtLZYhBz.rtXwecWIFBf.X1TCTtGaIz9TlZURcJO	89160239030	ilmikhivanov@gmail.com	f	IlyaivanovM	id49612033	РТ4-11	Иванов Илья Михайлович	6b2934e3-98fb-4a5a-9138-35a62df992b9	1	f
2c52cfcc-b539-4f54-b9f7-b429088924e8	$2a$08$xBuZBgVC4kWNYlViyZWj.ObSdqCx9ghXycWg.BwOOZSgQwDRbPj4.	89060223937	danis.1304@mail.ru	f	danis1304	nkn11	Э3-11	Бикмучев Данис Ракипович	2b55f5a7-5872-4904-bdea-7d54243826e8	1	f
5a14418a-f233-4c85-ac9c-718bb2edc932	$2a$08$xoK0lV4jOvy9zPh2St.kZecNB4WjRt6XgzlS5NLahRrkIpEKrW/XS	89052875010	vasily.rybin2005@yandex.ru	f	rybin2005	rybin2005	Э7-12	Рыбин Василий Михайлович	2b3b1e0a-1f6f-4f7d-a0da-7877dbf854f5	1	f
180f94bf-1361-434b-8cdc-9ecc9bda35ce	$2a$08$uZyHPzVbcpxLXuKxzjQQ3eD2obMg4KWo4gCDAKJRbYKb6iREHUV4S	89804610062	timashovaevgenia58@gmail.com	f	Tuuchka12	tuuchkaaaa	РЛ2-12	Тимашова Евгения Петровна	01dec7e5-69f0-4948-8adf-16def255ab41	3	f
5679f9e2-8642-4f17-a213-92b7281ff4cf	$2a$08$C.uK/LdK7m3msUZ3uOXnqet/HxOhneQzYy2AbODv369.ugxb0ijIG	89160932860	tyapkin2002@mail.ru	f	tyapkin_s	tyapkin_s	РК2-11М	Тяпкин Сергей Сергеевич	7b1a54ca-dbe6-40f6-bd58-955ed42987f4	1	f
fa91f879-1c9e-4033-a016-ad5a6a2b49f2	$2a$08$RUhKnhth1MhEWWYTAZc0ne0pF60.Wr8DkGb1EuffdFcVeRrwAFn8C	896518353-00	borenko.nastya@mail.ru	f	bad19u141	hii	ИУ7-11Б	Боренко Анастасия Денисовна	9909f062-467c-44d2-a9a3-0e43553d6892	3	f
e755162f-52e8-458e-8ee0-5fa58752966c	$2a$08$fHwwR8WkMLcJ9IB3CcbOMOmS9dgAMZcoFtK68Ku8iZpGSTiBWpyjK	8 918 97507-30	natali.ivanova16090899@gmail.com	f	DaRyAGolovko	darya___golovko	РЛ2-12	Головко Дарья Павловна	01dec7e5-69f0-4948-8adf-16def255ab41	1	f
ebbd2b6e-3d7d-4536-8780-277405dbfecf	$2a$08$NMbnPQViIdtGhmXPgk.rueGuxIeYeUX8p2OlI3c545zh4P7HjSqhS	8 916 41584-91	iannatry@yandex.ru	f	melarinesss	iannatry	РЛ2-12	Иванова Яна Александровна	01dec7e5-69f0-4948-8adf-16def255ab41	1	f
7d56b7cd-9249-4f7f-9037-8db19e34f74e	$2a$08$nO4RWWeh3j6BvLXBaEnJUecTTLujrUTyVMQlZ6lv0v5j/TvStHB8S	89991128124	sanyabatevsky@yandex.ru	f	SanKryasSan	alexbatevsky	ОЭ2-11Б	Пересветов Александр Михайлович	9e61f1aa-978a-461a-90f6-99808580d083	1	f
839dc0b4-0295-49c5-bba8-3f232272d416	$2a$08$lxH4RRu8ohlEoEIuDGCvGOImaKG6h8f.zfo3Det5Z5nzVPddbCuEi	89778463387	ilyamyrom12@mail.ru	f	your_Ilyushka	kvad_rik	РТ1-11	Черных Илья Валентинович	40355e30-8138-4e9e-a09c-6d4b9ae9f4aa	1	f
19d1c2f0-16ce-439b-943f-eddbb58f675a	$2a$08$/MnV6baYKfPykW2lm.o6Y.I4wUlbO0eJkcb6lCz2V9n.CdI2kpIKe	89608946010	ladakochikina@gmail.com	f	oooSHIKAMARUooo	ya_belyashik_l	ФН11-12Б	Кочикина Лада Вячеславовна	b9b02aae-0db5-4970-961c-9a5a198f7ee7	1	f
2fd258ba-2ceb-4938-b694-39899e670ee2	$2a$08$FZa/TXQaBI71r4HxiZjs/OEWYQd5m6.NWkhr.pf3hTwFXw5AFawEa	896613380-08	fresh-human@yandex.ru	f	reroshi	reroshi	ФН4-11Б	Кружкова Татьяна Дмитриевна	b9b02aae-0db5-4970-961c-9a5a198f7ee7	1	f
84cc6730-b35b-43a9-bddb-1b0a5c799aaf	$2a$08$R.Wpk5oj0syzzpKP2E21cetKD.KmQ8uWp0b8Y/5AK/yKGvXm1f6VG	89261205844	boroff.tosha@yandex.ru	f	yaruun	yaruun	ИБМ6-14Б	Касаткин Ярослав Антонович	4fb50154-1111-4f44-badb-e48f83e0d433	1	f
c942f13c-a1d6-49c8-84bc-6a3234b81829	$2a$08$ZstPcO72ike.RWzKHVQo/.E5aozypRELoOi9YUThp0Af9rgyfEpY6	89671545974	lyigdi@gmail.com	f	monsenior_qwerty	palm_alex	Э2-11Б	Пятибратов Александр Владимирович	9e61f1aa-978a-461a-90f6-99808580d083	1	f
0129002c-5ab9-4f3a-89c5-caab2b0025e7	$2a$08$nrDzYwF88lPyRM/8cOxF.OgPfsdqO81ehEdcv6oWEjaFEh3uShUpW	8 916 84972-72	a916849727@gmail.com	f	Lyn1xx	lyn1xx	Э4-11Б	Ясенев Арсений Дмитриевич	95fafbf5-1bbd-4f58-9ee3-1e1cf6eff939	1	f
5c8ea37f-0161-4f8d-9ede-675d74d38aba	$2a$08$t.sRked.1EtNoM1FDb2J5uGfdmKUwnMPZgfXWBk6Eu4Em7druLOse	89109167803	yur.fedor2112@gmail.com	f	maladoy_pozitivniy	krosawczeg	Э4-11Б	Юрченков Федор Денисович	95fafbf5-1bbd-4f58-9ee3-1e1cf6eff939	1	f
59510ca2-dc99-43a3-a804-aa5585212366	$2a$08$ObfRW7Yfr5pUe6ciqKY8duquGhyR1vwn/UaipZCQAS/EtPM.3wtS6	89174755946	prosvirov.valeru@gmail.com	f	orfism1	orfism	Э8-11Б	Просвиров Матвей Николаевич	4fb50154-1111-4f44-badb-e48f83e0d433	1	f
d7136177-4a50-4a4b-a3a1-03c058134a6f	$2a$08$k653C8wRlc5LInNdaI9n2uxPKEeWM3Rwp5u6JJPuaLPC0CH9EgXSe	8 903 52477-75	zatseva26@gmail.com	f	oksik_z_y	oksik_z	ИБМ6-14Б	Зайцева Оксана Юрьевна	\N	\N	f
c01a8c6d-9fad-480e-b4bf-f27daed4646e	$2a$08$37IShYiTmBS5MS3a9JYTSeXEvNbz9imhzUv7cXCMVLYriExN8UaNe	89654182671	nikitka.tyutin@gmail.com	f	TyutinEz	sosatpls	МТ8-11Б	Тютин Никита Кириллович	02e685fa-19e7-4f7e-8025-8dbfbee483a0	1	f
110f5254-c63d-4b79-84b5-f246ff178f0d	$2a$08$ALaTTB.FkWRRw0SyHI2YJe.HEVY1CD5xr3C29xBtvxayjvs.pnnwC	89191858502	lme88@yandex.ru	f	marusya_lev	massssshhhhhha	ФН2-11Б	Левашова Мария Егоровна	f9edfee0-2a1c-46c9-8518-d5a92b0cff1b	1	f
535edae1-a75a-427c-a709-6f67ec7b503d	$2a$08$WFe/uWjt4yehKjOehetcnuiO9fFVuCRdekK7Bj5ds2G5ZpxfBVjR.	8 951 61863-15	maria1akinina@gmail.com	f	akin1m	maria1akinina	РК5-11Б	Акинина Мария Александровна	c8fafbd5-03e9-481f-bea4-93819d84a7bb	3	f
2bcd2cc2-1e43-4dcf-b58e-49c50206fb6c	$2a$08$i7rgGUGbY7Y3cxwHSej60.eP4DCyxf6QlqZa35qnrTiFjLPpsvfkK	89779168972	novichkov.05@yandex.ru	f	Reptiloyd	id597064619	РК5-11Б	Новичков Михаил Сергеевич	c8fafbd5-03e9-481f-bea4-93819d84a7bb	2	f
eb15cc6c-7b04-4f68-991d-a1ea8711c175	$2a$08$llT7jX/T/2xjSnL46XaVZOtsHIVolImYB/lZ8pF/OlTK1kOlR2pDq	89255701058	qazaqky@mail.ru	f	modnik_ivan	modnik_ivan	ФН2-11Б	Миронов Иван Сергеевич	f9edfee0-2a1c-46c9-8518-d5a92b0cff1b	2	f
beeafa30-ebe2-4964-8def-7a5a09a52a9c	$2a$08$x9Q6oexwFuMbCD62ObBsy.h/n9YhoXPSwahlQ59YqeI3WR5lPazfW	8 916 84860-26	alandrus05@yandex.ru	f	alandrus000	alandrus	ИБМ2-11Б	Андрус Александр Алексеевич	\N	\N	f
e8e84ab7-c557-4cdd-a118-f93e55e1b540	$2a$08$5L.a0DltVQZVdeAqLPpkQ.v3oSMMjbxMy2N9srFNQCzkjkOoG957q	89631442714	arsen.ishmuratov.05@bk.ru	f	Directed_Byyy	arsenchusss	ФН2-11Б	Ишмуратов Арсен Русланович	f9edfee0-2a1c-46c9-8518-d5a92b0cff1b	2	f
16025870-6ccf-453b-8a6c-d49857b42477	$2a$08$S5FLRN1JW2vOFXoB.AMMPuRpos/o2vNA7iz/podOoPfSiodqWO/be	8906-5480748	yagodkinay12@yandex.ru	f	julgerman	yagodkinay	ИУ2-12	Ягодкина Юлия Игоревна	074448b3-2a26-4b58-b63a-cecbdad6590d	3	f
e08f06a8-4ed1-40c8-9d26-4eb4e706412d	$2a$08$kjq14LhC0T1d/6OvG9D.nOzRAUGtGXHeQNCSQKypMH5Ca2HiKYmkS	8968-868-20-70	milyaevaav@student.bmstu.ru	f	milyaevaanastasia	a_milyaeva	ИУ2-12	Миляева Анастасия Вячеславовна	074448b3-2a26-4b58-b63a-cecbdad6590d	1	f
09368ce3-2e57-41c5-905a-6b3bdeaba633	$2a$08$ascZtyAdDBgvRNr6Lyf1e.OE/q6GWxgacpJSEcDP1GvUxE7n8LE6q	89194602947	slav_ov05@list.ru	f	slavavov	slav_ov	РЛ6-11	Овчинников Вячеслав Сергеевич	02e685fa-19e7-4f7e-8025-8dbfbee483a0	1	f
34943121-dd9c-4675-a7bb-5d9e0351b6f6	$2a$08$FAG6GwyE.bwyTweb9CiMb.ozIocAJ.MZF0DU.4xBghOE9e7oZK7QS	89651835300	borenko.nastya@gmail.com	f	bad19u141	hii	ИУ7-11	Боренко Анастасия Денисовна	\N	1	f
72cad169-90ef-4e13-9dff-32581350568a	$2a$08$KmrtExMpxQ8IdwNSn9NDae1nA8.7gJA6Set5Vx/fSPz2GSFL9kw7O	896518353-00	anastasya.borenko@yandex.ru	f	bad19u141	hi	ИУ7-11	Боренко Анастасия Денисовна	a0ff82fd-80f4-4fbd-9e4a-9ee65ab57a3f	3	f
7f2d7373-3489-4bbc-baec-1d24a66f4dc4	$2a$08$vh5JXMCHvtaW9M2H6yA7o.aJDwFKjNrUaahuzX9CIIkrXH70Bis0e	89773760126	bogdana06.ch@mail.ru	f	meohani	bo0gdaanaa	ИУ4-11Б	Чечетко Богдана Андреевна	a5ea392d-d4f6-4c87-a04e-002e252cfb6c	1	f
c354eed2-e39b-4595-80d8-f27facb57412	$2a$08$sjBs4RDC.FKO41zLm7yn9eA0097tPdqnk5aY0txany6yglUcFRO36	89993613143	lidiya0chaynicova@gmail.com	f	lidiya_0	qgeqge	ИУ4-11Б	Чайникова Лидия Юрьевна	a5ea392d-d4f6-4c87-a04e-002e252cfb6c	3	f
a22e0451-308b-46fb-a705-acc1148abfa9	$2a$08$bXHTZzWLh775CB.ce5ZOn.AKeGjU9yyWPLu5Nz20ZfNhGRhqwQoh6	89253171808	bazarovamasha05@gmail.com	f	mrllwwwww	mrllww	ИУ4-11Б	Базарова Мария Денисовна	a5ea392d-d4f6-4c87-a04e-002e252cfb6c	1	f
d4d15afe-3873-446d-8c26-4f876cdc4bcc	$2a$08$NGB2rY3BFnT7xZSMqaVBMe4u0zlcjqqrtE0QLf5/oTC9q3xtPcdEa	89776546214	sasha.filimonovich@gmail.com	f	PocketChamber	pocketchamber	ИУ4-11Б	Филимонов Александр Александрович	a5ea392d-d4f6-4c87-a04e-002e252cfb6c	1	f
4499c980-960c-48e5-9ebb-2d59d2ba28ed	$2a$08$z8uK5c7giDk3tmk6xeCVVOa8gt3paPMmtI6ZOaKVRL2Tf7WHrr2Vm	89175012977	karevmax2005@bmstu.ru	f	zampolitkarev	ya_pelmenxl	ЮР-14	Карев Максим Игоревич	\N	\N	f
53cede40-6b23-43ec-871a-f4ae186fdb9c	$2a$08$XqFjVFGRGAVW2dGncOkOo.WWFnvsxmcG7x8TcnrKki2VJ3c4C2SLC	89281728103	tomas.krik2023@gmail.com	f	tk2023gm	id812737849	ФН4-12Б	Чернявский Юрий Геннадьевич	73ce80d7-2798-442f-918e-bf40511dfb07	1	f
233e50c0-05c1-459f-8eff-bc836c3c3a39	$2a$08$CAwYZRdjQ15m7dbHpMwPaOik2uoOVZlZImymQ1z7qXfB52tsQ4b9y	8915-476-76-27	scolopendraposhlanahui@gmail.com	f	kustebatb	scolopendraposhlanahui	БМТ1-13Б	Каневская Юлия Евгеньевна	300ef3ae-2e98-43cc-8151-a7f42980de56	3	f
33ec1b6f-fd30-48a7-9997-e9d01fa9d461	$2a$08$YCVFh66fR532R3erQfrBiuY/qyEKXbjhDN3uNl2Vs6gS08Qic5V1W	89266697749	gvidocrazy228@gmail.com	f	gvido_crazy	gvido_crazy	РТ1-11	Зиновьев Даниил Владиславович	40355e30-8138-4e9e-a09c-6d4b9ae9f4aa	1	f
a91a3a6e-1092-4d80-972d-26024edbb9a6	$2a$08$ZuDgwvHnxHbIC5jHxnlCSeDWfY3Gol.mc.m68vYS9Syb.hdNYJnCu	89153063836	sbaranov2111@gmail.com	f	CrabEvolution	CrabEvolution	БМТ1-11	Баранов Станислав Андреевич	\N	\N	f
b64ddc39-93ff-4aba-ad2e-ff45fd695af0	$2a$08$QDnDAK5xhPzUY9SO5BJ7z.GZWikhu2Ri8kvl/kUk1KQGz1U7Xay.S	89308467748	a.k.12345543210@mail.ru	f	Anastasia0084	id316056941	РЛ2-11Б	Кондратьева Анастасия Олеговна	543a4e0b-e35d-44f5-85df-43431142aab6	1	f
c8ab03fe-88cf-4cbe-8446-a4510a78b937	$2a$08$7V0C8A6uWRdr2A9m4Anjt.sHNE3ib8w31gZ0JzS0Wp6jyxhsQJaR.	8 915 04886-93	danya9365@gmail.com	f	danik_pryanik2005	legendarnii_chelovek	ИУ6-11Б	Мухин Даниил Николаевич	2d40c471-5ebc-4470-a419-97d182c2b096	1	f
7f073cb9-974d-45ad-833e-4cb5e76ad1b2	$2a$08$lB6WwKXSfA5PrrCPlvNz4OEbWsCpqnGVW5HevdivBiivSZkRKj95S	89954247009	vladavladakirmir@mail.ru	f	vladislaviyyy	i_am_vladosia	БМТ1-13Б	Киричкова Владислава Евгеньевна	300ef3ae-2e98-43cc-8151-a7f42980de56	1	f
1f21d264-66c4-4768-82a7-ec7c2f446c5f	$2a$08$vBOwzx7JFCQe3j3xbsKSh.batPZZYhz/lrTtDm6snzviF7mD7PhMm	89273906641	kudimovapolina1@gmail.com	f	kavorry	kudimovapolina	Л4-13Б	Кудимова Полина Андреевна	28fd4455-c957-4065-97f6-847b632daf2b	1	f
0f225150-eb5b-4233-8fd4-94cf1d04cdc3	$2a$08$wVJ6oYW.SfOxsGmp0JFn5O4fXilhZQ36rpDR2tIokhkgnXiEViepq	89023981053	artem0072kondr@gmail.com	f	InStable0072	instable_0	РКТ1-11	Кондратьев Артëм Олегович	543a4e0b-e35d-44f5-85df-43431142aab6	1	f
13865d1c-cbe8-459f-9e1f-f63c52de7378	$2a$08$AwdQfc5/MiVHKMcjcy0DaONZASgNBvoZOKgVLezkxhZI/m201jXli	89652081241	alexbagrin708@gmail.com	f	AlexBag	Vk	Э10-12Б	Багрин Алексей Петрович	\N	\N	f
36ccb13b-ef92-4a8d-9c28-60e0ab9565fe	$2a$08$b2hYTRWBJ13UsRXU1nQ1l.8fuvRu2tm6ikGqbfeJl5mnWk.JAt676	8 920 54123-73	aleksandrasedrina6100@gmail.com	f	sashka170veu	sashka170veu	БМТ1-12Б	Щедрина Александра Владимировна	300ef3ae-2e98-43cc-8151-a7f42980de56	1	f
8eb888e9-c392-46b8-b3ec-3887240384f0	$2a$08$M01sRHB8cBCOhKmy94UUAO0vosdj1T0Pyrc4gzYo9g6e.44e3Gg96	89213638670	puskinandrej25@gmail.com	f	barashek01	barashek_lives_matter	ИУ2-13	Пушкин Андрей Никлаевич	28fd4455-c957-4065-97f6-847b632daf2b	1	f
08aaa3c9-7340-4829-8845-4e86d26904ca	$2a$08$rvSJJRNCQrQMPB6euKhtSepwJranf1sbACfz2gJE2CRptrWrg0e/.	89165800028	ivanytch2@gmail.com	f	DIHamen_2005	id455927095	АК1-11	Хаменушко Денис Иванович	\N	\N	f
497965af-a79d-42fc-b9b2-c0e49d60957e	$2a$08$ZgVJqcrIJM4FoyPZ2qj.7.SsW1rrNG34/QUSJZ0OLxIkP4cgly5qa	89952609678	urmatalmazbekov13@gmail.com	f	Urmat_Almazbekov	id466023508	БМТ1-11Б	Алмазбеков Урмат	\N	\N	f
268b3baf-cec5-492f-9f9f-1408de44054b	$2a$08$Z2yd9RBpWpC.eRhDhZTJPOQA1iJPfCOzTXKeY3BRn9JinsqyDFH62	89091545748	tigr.my@gmail.com	f	tigran_mushegian	mushegyan4	БМТ1-11Б	Мушегян Тигран Мушегович	\N	\N	f
7673146a-c0b5-444f-b654-bda289449438	$2a$08$0S/Y1ukfnQJFUMG66MLUDuxd.KikxPXCfM3Mga6xi.xL84MBwjJxG	89852855801	nikskl999@gmail.com	f	srnm9	9srnm	ИУ7-16Б	Скляр Никита Михайлович	\N	\N	f
eba19f7b-49fb-48ea-ae75-666983c5f2bb	$2a$08$w.aXt8n.s7L4POCTAg5a4OBnKLQO9OIe7ZiZ/14gT2w7jErrwH.nO	89661808844	sert.fou58@gmail.com	f	sertfou	sertfou	АК1-11	Огонькова Мария Сергеевна	\N	\N	f
edb0bb5d-1840-4707-b0b1-ebe1718aeff0	$2a$08$tgpWSPFvGWXD3X3z3Rcfh.3GPWFOy3idaKaahjz88R6tx0JqWflky	890152212-65	borgleb13@yandex.ru	f	borgleb13	borgleb13	СМ1-12	Борисовец Глеб Алексеевич	2acc4cbe-c250-4830-a19e-8b8af134ce2b	1	f
1a14d554-9b56-43a5-94c2-9f5566498c12	$2a$08$dVk4X8lugI440dbigUlVQeojVn9PM/FyOQ2cs5UhKlD5uBjjJZ9vG	89251560470	legoshin640@gmail.com	f	kokorik_marandich	legoshin	СМ1-12	Легошин Егор Александрович	2acc4cbe-c250-4830-a19e-8b8af134ce2b	1	f
d47545ab-6a1b-4123-8968-bff1a26cd046	$2a$08$Yd1wELxlZpmCVvOnoZ9WcuXQ/LzQlGcgqtgpmnUWTI5jj7gFmcRT2	89850913139	stasbaranovskij7@gmail.com	f	XxEgerxX	mister_superstar	РК5-11Б	Алексеев Георгий Константинович	c8fafbd5-03e9-481f-bea4-93819d84a7bb	2	f
64bf3ce9-31c4-4966-bbbf-be83e7ae8569	$2a$08$KoHrWtEicN24prReW5Yhl.gFRnTCmkUWwW4ZR6l3QT4d4HaQazhgu	8965-304-31-20	orlovskayayav@student.bmstu.ru	f	yaanaa06	yana.orlovskaya6	БМТ1-12	Орловская Яна Валерьевна	\N	\N	f
db6d2897-e254-4ce5-bd0c-5458885f7a6d	$2a$08$.W2/v8vzH6DsyPthf6zlj.R3DUXbKw6WxIOPQ.0CrY48okTMmVSKW	89199627676	rustsharipov56@gmail.com	f	Eustos_Hail	id788968569	ФН1-11Б	Шарипов Рустамджон Олимджонович	5095faa5-51b3-4dcc-90cd-744b3566ccf2	1	f
603c4ba5-5edd-42ad-9f4f-1e2f69dd5f40	$2a$08$E5tc8mVxUaY6ku8UjNm4zOrnJ9Ch7YiE5wPIf8BPom1tp3VvHjIym	8 983 56515-97	a_artem2005@mail.ru	f	engineer0044	engineer0044	СМ1-11	Айхлер Артём Александрович	\N	\N	f
9b342fc3-7a94-4402-9fb2-69bb936f28c9	$2a$08$fBXlI85c824/Zd7kwonSaOOtunAREFYoAOl3Xcs8pYbrwjaOY65Xm	89823180582	butkovladislava@gmail.com	f	Nanochelkins	vladusikb	МТ11-11Б	Бутко Владислава Алексеевна	fd8040f4-8e1e-40fe-8012-0398c80c5324	1	f
4d4c3642-3660-41d3-9dd9-ad328bd5235c	$2a$08$9XWSN7AFNPF5s5Pq1CewNe59LSKeKGRwoucjDOxNkLkG3qZDR0FdO	89604354988	ilezgamberdov5@gmail.com	f	g_ilez	ukineee	МТ11-11Б	Илез Гамбердов	fd8040f4-8e1e-40fe-8012-0398c80c5324	1	f
1359cc8b-0707-4d44-9ebe-58d5303dff84	$2a$08$PwQV0tHpjNvs0wi6dqujW.GmsiqrH7OQUhUR8TE/T02qfn2P0crKK	89998623612	prostoogurec9@gmail.com	f	votetorofffl	votetorofffl	БМТ2-12	Копылова Виталия Александровна	32cc7a99-1f39-4dbb-a160-e293a14838a1	1	f
dea8fd44-7fa9-4617-b6e7-ae0b6d959626	$2a$08$8mHFvTyYVRqzs8pPRxV7LeAFYvLvKdi5vJsYO3GxG6NY.pTK90KO2	8 987 09310-85	danilerm2005@mail.ru	f	pmolaich	pmolaich	МТ11-13Б	Ермолаев Данил Сергеевич	0c68c5b1-5185-497b-9a60-b1c1e1763b6e	1	f
0d31cf04-3b99-4e41-928a-04d86c270802	$2a$08$MH56q/KjugUiR7JtiZBAOuNfk2z8nweDUiqEM5X6NJWD4Wt6vi0nu	891754832-56	ivan0712ivanov@yandex.ru	f	kosikov07	id455909886	ПС3-11	Иванов Иван Сергеевич	6b2934e3-98fb-4a5a-9138-35a62df992b9	1	f
842710f5-7e0d-4f69-a19f-4991fd91bb84	$2a$08$kCx2ZvpoUhSMkDzF.d/W0OXRu0Agtp2MEyB4WyS3DhnR52eVGQsI.	89168657749	sharova.mv@mail.ru	f	maren4ik	kvayou	РК6-12Б	Шарова Марина Владимировна	2d40c471-5ebc-4470-a419-97d182c2b096	1	f
033f54f4-c8a9-4557-b9f0-c187046393f9	$2a$08$/Ra3mb8cqKdf86Xds9wPh.KE1LDLhagLx4/VjcDYz6K6oplf3.zla	89689235153	ilya.sorokin200531@gmail.com	f	Ilyas0r0kin	i.sorokin2005	МТ1-11	Сорокин Илья Валерьевич	\N	\N	f
f570a49d-edf1-4756-9ad3-74785041edb0	$2a$08$5broAet7eafy3Ij3cVUv4eh8Jg4COzXOnwpU642BXOYV0qqygKxuC	89169593903	ikolevatykh@gmail.com	f	IgorKolevatykh	Vk	Э10-12Б	Колеватых Игорь Алексеевич	\N	\N	f
e45b9bdb-9ce4-4520-82fe-967d10666a35	$2a$08$jMAMJNpQJlZB7CIOqGX87.1UVH0Mfng6pS5nMJci6QW9CeyJZJDnq	8926-350-23-77	verakoroleva2005@gmail.com	f	greenpeass	green_peass	ФН4-12Б	Королева Вера Вениаминовна	\N	\N	f
0c8f9bd9-155b-4744-81f1-1adeb8bfc64e	$2a$08$nzB7qWYtVpybwoHmmhWv0eDtUGIbhYYyz1Kpfuqrbh5SehPTtLQMG	8 964 83079-92	nmsmirnov2005@mail.ru	f	Sssmirniy	sssmirniy	СМ5-11Б	Смирнов Никита Максимович	a845c25f-1727-40ef-90fa-8422bd94a20c	2	f
3eecdcbd-1825-4f73-a709-31ece1670a09	$2a$08$3cEaYChZ813i7.gvlryI.O9eIAzzO8l7sAw1Wx27BOY01i.F9Umq2	891752135-42	lidiya.vasileva.05@mail.ru	f	llllidiya	lllidiya	СМ1-12	Васильева Лидия Григорьевна	2acc4cbe-c250-4830-a19e-8b8af134ce2b	1	f
cce6dfc5-6579-479b-a499-8bd2a31296b7	$2a$08$zMnQfc1D58.G60Js42tMnuZfUpyU/N513X3LPturKKZPDvtZuXFCC	8 965 26554-19	nadiasycheva1105@yandex.com	f	NadiLisso	nadilisso	СМ7-11Б	Сычева Надежда Дмитриевна	a845c25f-1727-40ef-90fa-8422bd94a20c	1	f
70431c0d-9fcc-450f-aa3a-8c7f3939c9a5	$2a$08$A3BWRM0N8HkTn/wn1CASi.4G.mmfgMKH19SP4tgl.1hj7spZzQHKi	89815599232	kostaudal@ya.ru	f	K0nctan	id807285529	Э8-12Б	Удальцов Константин Игоревич	4fb50154-1111-4f44-badb-e48f83e0d433	1	f
b356a91c-8a83-4122-9193-f477e16b2330	$2a$08$9HY.f2cGzlwpYCcF4w8gnu4mfjYdauik.n.M4kqN7C3QUqaQ7Si9q	8925 1106016	matov-23t254@dummy.example.com	f	Matbl4	old_vobla	МТ8-11Б	Матов Владимир Витальевич	02e685fa-19e7-4f7e-8025-8dbfbee483a0	1	f
d11b09d6-1f11-4512-a305-11ae4f124166	$2a$08$aJ9m8io7TuJb63vm/6h.PuC9DsK1jaiMlwYprH02AWK5aR27kCscG	89969646479	79296686952@mail.ru	f	alasteria	alasteria	СМ1-12	Кожакина Софья Евгеньевна	2acc4cbe-c250-4830-a19e-8b8af134ce2b	1	f
88537553-d042-4a6f-9c0c-4534346116de	$2a$08$VLpd8YlvmMLknTPvKjIVg.ChlRP9wOpXwQwgVc1EPUKXKlOv75xAy	89854697150	notbadbutbad@mail.ru	f	ioymeana	v_palaty	ФН4-11Б	Душабаева Амина Азаматовна	b9b02aae-0db5-4970-961c-9a5a198f7ee7	1	f
0f29d5bc-08e1-4042-baf8-de2d665b67b9	$2a$08$QzGn0gjySvWH.yo3fYpvIONhy5FLLY.C0.Tvy7R5yQcDkQZMq3HYy	89246656364	misha.neti04032005@gmail.com	f	lanrefnl	0iachsv0	ЮР-16	Нетисов Михаил Антонович	94abce80-a6ef-4ec6-8aae-cdfd55e4d47f	1	f
0cfbcc3d-168b-4720-b9b1-82c8f57e2cbd	$2a$08$OiemZV3mb/q8aHt.reJD.ObW7l8rl0u2Qafg.DG8BU3bHEdybg9mi	8916-264-22-80	amira.skakun@yandex.ru	f	murkamrr	mrrmurkka	ИБМ4-12Б	Скакун Амира Алексеевна	d8d79fe2-6702-439e-bdda-ef7f2be4ebae	3	f
ba5f954d-319b-43ed-8988-c8b60c896095	$2a$08$peEr.qydRqWB6vDsD/A8guv3iVKGb.b7UsvRgnW089JHXiNjQX3RC	8 968 58314-58	andreevapm05@gmail.com	f	andreevapm	andreevapm	ИБМ4-12Б	Андреева Полина Максимовна	d8d79fe2-6702-439e-bdda-ef7f2be4ebae	2	f
ec192d80-0c68-45c1-92a6-f31cb17a2adb	$2a$08$sZCQuBH08BGCiFmI9m5wA.xNJSo8b7QR8fW6E3efuNYojE5Vy57Ja	89278773214	zhyravlevaannadi555@gmail.com	f	unclozer	unclozer	РК5-11Б	Журавлева Анна Дмитриевна	c8fafbd5-03e9-481f-bea4-93819d84a7bb	2	f
4a62bd68-d4ec-4b6f-94ef-c0daf4a99cc0	$2a$08$pL31PrwfWsiFoLHgdiQ1ee8mYXWz4PhZL/qaaLd/tII9m0ZGXzLq2	8 916 91372-60	vladzhdanov2000336@gmail.com	f	marmevlad14	marmevlad14	РК5-11Б	Владислав Жданов	c8fafbd5-03e9-481f-bea4-93819d84a7bb	2	f
0d4f1e8b-002d-4170-b083-4b8fe1e1f08c	$2a$08$1lS23Z7GMK1k4ohhNNP0a.0T4z8zl/n7iT9X1RiiBWfPnbE1ZRYce	89779365288	yun.gritsenko@mail.ru	f	yuna_gri	yunona_gritsenko	ЮР-14	Юнона Гриценко	712edd72-aed6-4ec5-a521-cf4a8bbbf4f7	1	f
c39591e9-6125-4494-a930-adddaa1c3efb	$2a$08$OPg3drt3A.1/dmcJT6WxiOdMKwwhoblkmX.k1Rb8NboYsOsZDU9Nm	8910-725-12-02	vlad.vavilchencko21@gmail.com	f	XOCHU_CHAI	vav.vladislav	ЮР-12	Вавильченко Владислав Александрович	94abce80-a6ef-4ec6-8aae-cdfd55e4d47f	3	f
0e78ce6c-26f1-4314-81de-08589d2b1435	$2a$08$Vdd1K8ejIST5i6.iJL4vfedqaz1mzIxN3ITesfkAl3UIBNjR7l4VG	89778819788	artur.abyzov.101@gmail.com	f	the_onixs	the_onixs	ЮР-12	Абызов Артур Александрович	94abce80-a6ef-4ec6-8aae-cdfd55e4d47f	1	f
8a26127e-2e51-4b1b-93c5-9d901fdb07d2	$2a$08$0XgejLJvqVMOnpA118fR/.QrbLeH1DsakyEKrFHXsO8bt85R7mVVS	89154127188	shatalov1907@mail.ru	f	tizwo	tizwo	ЮР-12	Шаталов Дмитрий Сергеевич	94abce80-a6ef-4ec6-8aae-cdfd55e4d47f	1	f
60eb6230-a15c-48d3-a0e2-312f9e7e17e9	$2a$08$qS9KVsrXW4AO2PG.YacFj.pJCM6wagRS4/AfJMVwQ9GibXuMWQPjK	89826610182	zabolockajakate@gmail.com	f	kseroneox	kate_de.swamp	Э4-12Б	Заболоцкая Екатерина Владимировна	6fa56c83-f29f-4909-9d63-ed00152c144b	1	f
01b69d52-1de2-4c4d-aa45-a9a05da5a270	$2a$08$t2uzvaiZxdS2YxKqrLRpAeUxqCKGp9ghE9Dh0Ve4XeygIKH2G95xW	89256010881	vova.zabelin.06@mail.ru	f	V1GINGLE	v1gingle	ЮР-12	Забелин Владимир Александрович	94abce80-a6ef-4ec6-8aae-cdfd55e4d47f	1	f
17cd289e-0537-4009-af8a-eefb7aba7a74	$2a$08$rwBgjyGfzsjHRqx4G/0fqezAhXLSsnaSE2JZ8MJrnpTCOfXrQH05.	89857579622	drozdovnikita2019@mail.ru	f	Nikita_NikD	feed	ЮР-12	Дроздов Никита Алексеевич	\N	\N	f
24e97b69-ca4c-4ff7-9370-4a38c1109f07	$2a$08$eAWwEP3s5LrN0URlQjvkx.2LAZWNaFbHIPqHQno8V.HwOYF7QCrEi	89018042863	alexsob794@gmail.com	f	xenonini	xenonini	СМ1-12	Алексей Дмитриевич Рыженков	2acc4cbe-c250-4830-a19e-8b8af134ce2b	1	f
27e529fe-9dba-4dc8-bb87-35aedf3e2a5c	$2a$08$Pnnyg53RWmHYYEHG8nuP1eWtrpGWIW643k4iRomguwjfpgiVmJJL.	89231442230	bgsaneg777@gmail.com	f	Billyb0bs	lubatovo	СМ7-14Б	Якушев Александр Михайлович	\N	1	f
3ae21f9f-da34-4662-a2ee-ffa17da12591	$2a$08$bjEOKvTqNEdPqm1WUU/BrOqRVzJKCf1dpAxtxqRkHTbqLkWvQf47a	89652252537	baikinanastya@yandex.ru	f	bainastasi	anstya_bai	МТ10-12	Байкина Анастасия Михайловна	6148ed20-a1df-49f4-9e86-58163649a2d0	3	f
a101f262-0f52-45d6-9b46-b13e0fbdf3ff	$2a$08$ibKkKaza32ZCOn0yxeM8se1bF.vmJyY3oHpTielU/mGLTRp4./EmO	89153530023	ekk.kartashov@yandex.ru	f	youngmulaboiiiiii	youngmulaboiiiiii	МТ2-12	Карташов Егор Дмитриевич	6148ed20-a1df-49f4-9e86-58163649a2d0	1	f
e84287ab-e70a-4597-a436-082bdc4baef5	$2a$08$ns69bM9B6Ih1xcZ92Q4FCubIlmDXh5xkuG5IYmFmHhf3hT0R02ARO	89192950250	raaazmv@gmail.com	f	rrrzmvvv	rrrzmvvv	Э2-11Б	Разумеева Анна Алексеевна	6148ed20-a1df-49f4-9e86-58163649a2d0	1	f
e36f0880-2c26-489b-96bf-3454de1c2a7b	$2a$08$McuIHxM6WE9gbWDTuEMHk.0DLZIqPDMc/RjvWIT/zlS1q3LUlHyb6	89161598937	diananime2202@gmail.com	f	Dayashaaaaa	dayashhha	ЮР-12	Козлова Дайана Николаевна	94abce80-a6ef-4ec6-8aae-cdfd55e4d47f	1	f
c7214ed8-d85c-480b-93fd-8250602d1cc7	$2a$08$ES3RGL5i718EeuImNMEB8.qqcIBQRlAEXmZijngYuRJaY54nuRFaK	89161393209	ulianaloseva0709@yandex.ru	f	uliana_loseva	uliana_loseva	ИБМ6-11Б	Лосева Ульяна Андреевна	a614efba-69c4-43fb-a9a4-d70d8be4c365	3	f
cc9c8d37-4432-48bb-9fe2-b31d18f3b0a3	$2a$08$MdcibI4ug1HTBjyP7WIdEOL/NgCqGDUERWlvcSAg6YGrl1h740L7q	89032679494	maximkutuzov2006@gmail.com	f	maximkutuzov2006	dynamo1801	ИБМ6-11Б	Кутузов Максим Владимирович	a614efba-69c4-43fb-a9a4-d70d8be4c365	1	f
b0bbf59a-c089-43d5-a73e-505a5d8f8689	$2a$08$/T6koNsVP.9VB4tIzdudJu.14SK1JCDx/szy.l8vGRAxAbpqZHA1e	89253006420	edkhromykh@gmail.com	f	KkarmaSss	id369071743	Э7-12	Хромых Екатерина Дмитриевна	17beb218-e4fc-4644-935f-cfbcb462b0ec	1	f
de1a06e1-f751-4b2e-a14e-9ec547b246bb	$2a$08$ZB/ZIhDxR7xcgfDCGqFML.MuYjHkd87o3iH35HLxTcXWCdCQzxEY.	89092691670	misha.p2005@mail.ru	f	Icemanbro	sagathtype	ФН3-11Б	Панов Михаил Евгеньевич	b9b02aae-0db5-4970-961c-9a5a198f7ee7	1	f
b526eb1c-1830-4458-bd18-d5aeccc73708	$2a$08$Fm3BgReMMRaozwC1izAOKeQz82DCpK2WPGmDnhydaPud757GCNj9q	89775085898	belyakov20052005@gmail.com	f	o1eg_be1	id607001202	МТ2-12	Беляков Олег Сергеевич	6148ed20-a1df-49f4-9e86-58163649a2d0	1	f
97cee0f2-1463-4303-ae35-f3e988dce1e8	$2a$08$v0dDMiE/NhjyE.G1iDotDOt2IX453g3/kPZZiOJ6.sdfnubg.lDGy	8123-456-78-90	lavashtest@gmail.com	f	qewtyfn	test	ИУ10-14	Корнеплод Виктор Травоядович	\N	\N	f
83627cb8-30a8-40e7-93af-01a2b406e149	$2a$08$QEDcdGFN1gdfYHB5KkfIQeJZ1YLOH6cQ/EbEg8ie17fPi39JoZyzu	8925-381-61-65	m.shlyakov@yandex.ru	f	mishashtick	mishashtik	РЛ1-12	Шляков Михаил Васильевич	6148ed20-a1df-49f4-9e86-58163649a2d0	1	f
daa6f2ad-68cc-44eb-906e-4328df91cc2c	$2a$08$.o.hmmPAUypGbWYiT7t0MekBeu0YcSQwgdfEZ4T/Y2v9V.yof1E62	89687977869	kostya.malnev2005@gmail.com	f	One_rtt	id743612722	МТ11-11Б	Мальнев Константин Сергеевич	fd8040f4-8e1e-40fe-8012-0398c80c5324	1	f
6e9e2717-5bd8-451f-a37a-19928080c0a5	$2a$08$fxXp4.MyhPPXigtIdKSIDOKBvNpgCxC2eo10cKQzEnfRuxZCF8NfG	89265354468	alexgnedovets2005@gmail.com	f	GreedSempai	greed_sempai	ИУ10-16	Гнедовец Алексей Александрович	fd8040f4-8e1e-40fe-8012-0398c80c5324	1	f
3b1517be-08ae-4980-a8e8-f2f34da55113	$2a$08$6mgWTOQQQ1iCppItsAmeo.OtZQbijRA07ZtVigUaI9Tn1lhdgIH9i	89059471177	vika.chervova@inbox.ru	f	ChiktoriaVervova	id268721589	МТ11-11Б	Червова Виктория Евгеньевна	fd8040f4-8e1e-40fe-8012-0398c80c5324	1	f
a23791af-d77d-4067-a46f-d3ae64068ba1	$2a$08$aT6cwYIGsoGvlRCpUhM3e.OpQhE4axSWZ0PJbp9VXVmwE9HMpGONO	898573053-27	sobikfox@gmail.com	f	rrocky_r	rrocky_r	ИУ8-14	Соболева Дарья Евгеньевна	2d40c471-5ebc-4470-a419-97d182c2b096	1	f
f31ec536-11d1-4eb0-a4b5-08fb53e1a358	$2a$08$M91w5SKeL2OFah7QhaA1puDKvAYrfVfVrMyy81qJT9Fmybv.621xu	89260674853	nastyaivanova_2005@mail.ru	f	Nast_yaaiva	id280107147	ИУ8-14	Иванова Анастасия Андреевна	2d40c471-5ebc-4470-a419-97d182c2b096	1	f
b4503646-3f1b-4677-aee1-95f3c9616527	$2a$08$L4hAiX/3Ssig6mh6oL7UEOavN91rlYCDTvgzNiCCn4rq0Bw3uKKI2	89050055755	nurao@yandex.ru	f	comoelsol	id745762483	АК2-11	Нуреева Ангелина Айратовна	28fd4455-c957-4065-97f6-847b632daf2b	1	f
d0766e3c-731b-44ae-81be-700caa12ba22	$2a$08$lh5hbcgZAXaNHTqGfhj/we1QP4oFGcJpL6UR0qHT9N3qwshPUg08O	89214926967	gulutaelizaveta@gmail.com	f	Elizafox17	17eliza	МТ4-11Б	Гулюта Елизавета Сергеевна	1c6d8c8a-9bda-4ec8-bbf6-8dd0b83b98d6	1	f
a270e581-f073-4258-8af0-6bcb5f08a06d	$2a$08$7zfCfS.QIV7NJhxc1A/l3.OMfq2rgsLOvohCh9QFlezXDj40sV4Uq	89777121474	reovinss.inc@gmail.com	f	benzphetamine	vampireeee	ЮР-12	Данилов Левон Артурович	94abce80-a6ef-4ec6-8aae-cdfd55e4d47f	1	f
162bf7fc-b91e-4af4-be3e-59b526a37588	$2a$08$NMQCB7NIpH1hW1qp9EUNHewCC9ySqgmwws61lBAAoqdgbcGjumHBW	89025465242	svyat.gradetskiy@gmail.com	f	svyatik3	svyatoyslava	СМ3-11	Градецкий Святослав Богданович	3155c953-27df-4bf3-93a7-3c3cce34deb3	1	f
d77e5b07-d2ec-49bb-bd77-f8e7b86f83f1	$2a$08$2PhbqBLg7YJv1To0k9XKl.tPVsikVUPrzFzMXaLkw9FFuAbPYA5dy	8985-430-37-63	nikita.krasnov.05@bk.ru	f	nikitos1717	asd171	МТ1-11	Краснов Никита Сергеевич	589132a7-7358-41cf-833c-7a5bd9f8e1b9	3	f
d337ff51-fed7-4a06-a013-4539be52cf9e	$2a$08$kcR13PHes5PH7iAstARxbejvJHi85jx3WmDX2uftIa5FSSZsGDjVW	89125941334	gladkikhivanus@gmail.com	f	abobaeveryday	ivnkin	Э2-12Б	Гладких Иван Андреевич	589132a7-7358-41cf-833c-7a5bd9f8e1b9	1	f
1f91896f-bae4-46ec-9b7c-ad7d5c2d5c3f	$2a$08$Zjwf3sD1/Zw5yPe4CX6V/Od8DUa3ik2DwcbudCR/5jpK6u78TsD9W	89853536912	temkaiva@gmail.com	f	HumpaHumps	ivashin2006	Э10-11Б	Ивашин Артём Андреевич	589132a7-7358-41cf-833c-7a5bd9f8e1b9	1	f
4284c33b-ecbc-42a0-ab6e-407c998fa4ca	$2a$08$GzHQfqZqgHBn1Isz.018leyiCLBjkXCwe4RtyktX9cg.7Y67ZMmqm	89671173539	buld2005.12@gmail.com	f	Ananasik87	ananasikbulik	ЮР-11	Бульдяева Анастасия Алексеевна	\N	\N	f
df951146-c8e8-426f-88c3-f3430e1d5286	$2a$08$2H6Fbu89t/5KNL07ZOgRsupaB1486K0nsJ2638.sccRRRLM6SxlGa	89258339877	saa23k016@student.bmstu.ru	f	lp_morginal	super_kontik2005	РТ1-11	Саприн Александр Андреевич	6b2934e3-98fb-4a5a-9138-35a62df992b9	1	f
e7b904dd-e1b0-416b-8803-bb12c6f1f4d0	$2a$08$kJmy/fCY.FOw1J.lVxVLRuz4VFA0qUwjB/Lm3knK8.IgfW7k73d0G	89101855189	ksu.osipatova@gmail.com	f	mairenissyll	randkea	СМ4-13	Осипатова Ксения Александровна	a845c25f-1727-40ef-90fa-8422bd94a20c	1	f
7e5b7d15-ffea-4375-91d7-f13e045a6e5c	$2a$08$FiJAUkSBu/Icn9sxhy/ameprEy.IVvHjrcnZdvlk/tFGBAtMjUXWS	89127615216	alekslesnt@yandex.ru	f	Aleks5216	id672942116	ИУ4-11Б	Лесников Алексей Сергеевич	811be257-7e9c-4847-91b3-3552791b9d7e	1	f
ecd1955f-0a09-43b3-b072-e104916fa1b7	$2a$08$hxYKFLH7tzll74BbBTpVA.QUYQmADXU7eyNO3e25By6qAG9TDjqQi	89164473691	m.moskovskikh@mail.ru	f	mihail_moskovskih	mihail_moskovskih	ИУ8-13	Московских Михаил Валерьевич	\N	\N	f
3ef7201f-9e42-40b4-b2ab-6d32221e737c	$2a$08$8C4gbrgdlmLxt9begnHYzOiDiIkk3ZpCCqRAbhjAEQif/1Z1m.R7e	89031230703	seya23ts019@student.bmstu.ru	f	difyzz	vk.com	ИУ5Ц-11Б	Спорыньин Егор	811be257-7e9c-4847-91b3-3552791b9d7e	3	f
e23626b1-2428-4bde-826b-561d279c33f5	$2a$08$VOpw7TcrlxZyUoMjdvhJLeH0esIojAqU3jKc0gxWY12H5TspHujB6	89057799339	chebotareva1221@icloud.com	f	Senurapa	ksu.boetseva	ИБМ3-11Б	Чеботарева Ксения Антоновна	811be257-7e9c-4847-91b3-3552791b9d7e	1	f
dbd71b12-82b4-4639-87fa-b1a6ab7c64b9	$2a$08$DCHICRa2AaNoYlo9onz8cO1RM85NHbnXwNaRinNr/vg1m1HcGZ3ji	89160902568	blumkinass@bmstu.ru	f	sorowfulll	sorowfulll	ФН1-11Б	Блюмкина Софья Станиславовна	\N	\N	f
aae3a00b-68da-4232-848b-81f3b44426a0	$2a$08$EzZy5DOGvlsJMTBjzTEfM.C.n/D5eibOUH28XdmtL8OLAjuWx..Nq	89188277229	maggeor2016@gmail.com	f	goshandrujban	palechunichtojaushiitankov	ИУ5-11Б	Магкеев Георгий Артурович	811be257-7e9c-4847-91b3-3552791b9d7e	1	f
48084d24-485b-4da4-beb2-fcb27daa6753	$2a$08$4nlZCuuqjSw7Xa9nigc3PecYE2Epap6uUnoC6d5zTdTGa7kbKGxM.	89200709973	anna.koptelova004@gmail.com	f	aneaeaea	anna_koptelova_2004	ИУ5-11Б	Коптелова Анна Юрьевна	\N	\N	f
ecb4637a-67c8-47f6-a5ab-9b3c1e7dc00d	$2a$08$ydWAN4hkNzHyiyynJIgCn.0RfQMnJBAzBsMdeg4ucSDiKtUt.kux2	8916-945-03-90	arfaikot@gmail.com	f	arfaikot	id427586694	ИУ6-11Б	Грачева Мария Вячеславовна	4fb50154-1111-4f44-badb-e48f83e0d433	1	f
405404ef-c894-4bb9-a614-cf5b1b0c3a98	$2a$08$Z48JzricblJmxhgZqqJegO1f97P85UF7x1FNyqUZ7ynC/lXFdzcc2	89661053081	kleschinskiyis@student.bmstu.ru	f	xmeteoput	x.meteoput	ИУ8-11	Клещинский Иван Сергеевич	811be257-7e9c-4847-91b3-3552791b9d7e	2	f
5e78d081-c2ee-4656-b3cf-7ec0b58a7592	$2a$08$boO9oVzk0X49far.8IJz..jXvHt/GBwbYpkEhv4IumQ5ZMW0Gs/He	8 906 62820-37	senya.pashkovs@gmail.com	f	chudak_71region	workout_gad	МТ5-11	Пашков Арсений Вадимович	307c360f-258f-45b6-b1cc-81c946df749e	1	f
710a5a82-3952-4d57-a841-293216de65d1	$2a$08$eLuQ2qb6vMGlycqyAZCuoO64ts1sFlUTox7bcqtjiLmYwqgMLOBJy	89777072947	bolgovfedor12@gmail.com	f	III_magnetico_III	III_magnetico_III	Э8-12Б	Болгов Фëдор Сергеевич	05735fa8-6c86-4424-865c-79df5b9850c2	1	f
6fdfc14b-bac1-4c52-b65d-0db47db3e748	$2a$08$Qwov.SzngT9BOmEe9XqRp.RP5oBaN592lTa3TUp1gYaBYuabhc75S	89895322655	kyst2282@mail.ru	f	Toks1chn1y	dedins1de993	Э8-12Б	Запорожченко Глеб Дмитриевич	05735fa8-6c86-4424-865c-79df5b9850c2	1	f
ca07a7f4-6ac9-4801-9af3-7bd4301b5e7c	$2a$08$FjhGlL3vTbqtkBVrZMFH4uqskkDtDDJiaV9zbHzEt.O4RJi8LBUTC	89260722408	zubovyuv@student.bmst.ru	f	yuraStarbiu	braxlush	МТ5-11	Зубов Юрий Витальевич	307c360f-258f-45b6-b1cc-81c946df749e	1	f
db782daf-57a3-4eee-902f-f52ac313e3b4	$2a$08$SMW4Vp2ssdkTuqbegFHWGeyShqqh1.zv6.CnxHJKxZl2qPPt7yjv.	89303045241	rymar.eraw@yandex.ru	f	v00001a	xexxxxxx	БМТ2-12Б	Рымарь Варвара Эдуардовна	32cc7a99-1f39-4dbb-a160-e293a14838a1	1	f
52ecf4f1-1dbf-4d39-8640-8bae26337697	$2a$08$Hv8YJGSBz18b8.IIlE8U6uU2rIFaPnMxtLnRB/sbKkt3GNeJMg/ky	89851247721	sanyaburanov@mail.ru	f	sashka_bu	id244182974	БМТ2-12Б	Буранова Александра Игоревна	32cc7a99-1f39-4dbb-a160-e293a14838a1	1	f
96e77f7c-e393-43e2-b163-a3b89596acc4	$2a$08$QAQAkBv/lYpuO9kWbp/dgeKrs4glUaSdwTDdeJW2ZEOvPc4weJhRi	89777025376	nastya-shum-13@yandex.ru	f	nsahsutmysa	shuumskaayaa	БМТ2-12Б	Шумская Анастасия Олеговна	32cc7a99-1f39-4dbb-a160-e293a14838a1	3	f
68cc056f-6cb9-493e-bc5f-a4756677c6c2	$2a$08$j0jbA852IpUfnSRQTtI/3eQMXHgSBiT0mMYSx/RPpf8/zkWM/ZAL6	89962313394	vvh14082005@gmail.com	f	vikakhomenko2005	vikakhomenko	СМ13-13	Хоменко Виктория Владимирлвна	\N	\N	f
f5f97e7b-bd85-4b2e-83d8-44c8d22e3396	$2a$08$TEFWUkpiNM.JsXNyvB9ptu7G2GWM7gVIPpGtBIyA5BfNwI1UiChoi	89138924062	ulyanitskaya.dar@yandex.ru	f	mathsolver2000	mathsolver2000	Э8-12	Ульяницкая Дарья Владимировна	6148d3c6-58fd-4ee4-830f-c8643ef8938f	2	f
7fc1ed73-2c9c-4bef-8a19-1dba970da772	$2a$08$76L7v9glHrTutlP2bBEvvOfUWsUGJFofln.SPbMxuS1jFyYdfXbCu	89140977711	arina.pak.2005@bk.ru	f	arpark3500	ar.park	Э8-12Б	Пак Арина Александровна	6148d3c6-58fd-4ee4-830f-c8643ef8938f	1	f
bb5c6ab6-f2c6-4135-8f11-39a70167ced4	$2a$08$81o7Aw5EEE0NV8mDo6ofYOuymb61qTdOz5Jp7tYZ21Be1NIoTWWXe	8926-194-19-77	shekhovtsovma@student.bmstu.ru	f	Max00000n	id253996777	Э3-13	Шеховцов Максим Александрович	485a9afa-6f4c-457b-952a-203d110898b2	1	f
6d83e736-9462-4924-b81e-e34355b6411f	$2a$08$YIoFJa1CbQSMsAU9HrQlE.iPdb3QLNYOwSLUtu1DSeVQIIcPtQU9e	8916-993-12-12	ira9931212@gmail.com	f	shmelundra	shmelundra	РЛ1-19	Шмелева Ирина Викторовна	\N	\N	f
24da476a-ff68-49b2-94c1-e7065221e0c0	$2a$08$q16qvosRc2zB8IgAZmdNA.toH4O4LeK3BzMyJH9MFAXOPPxH28H46	89851924917	eakulik20042@gmail.com	f	kulichidse	id640308780	РЛ2-12	Екатерина Александровна Куликова	6148d3c6-58fd-4ee4-830f-c8643ef8938f	1	f
96406d08-4dc7-4b63-a9c9-c4fa0a97f93e	$2a$08$0P569xPGyBoS.1iopHLlBO7/m.f2nTCq1mM/jnj5/vlLJ4r8bCAKG	89779389039	k0f3stprofaceit@mail.ru	f	eliryss	elifryss	МТ5-11	Кособочкин Дмитрий Алексеевич	307c360f-258f-45b6-b1cc-81c946df749e	1	f
40a6ba3b-d5fa-4d72-a680-e576367e2afd	$2a$08$.KOzWfrbH4Q.iOTn02xP5OOI.YRJEwYf3dzJr6wyzYevInmlZpEAW	8 908 59392-26	danzik05@gmail.com	f	stre4ok1314	danzanaush	Э5-11	Аюшеев Данзан Вячеславович	\N	\N	f
ae5d6f2a-a65c-40c7-8d7c-1e8b16be1386	$2a$08$d8GnvM9sYr45VMplfC9.T.YCM5UP5RvGkUSKYcCDc3fQ7KzVweJAS	89169897326	danielmurigin@yandex.ru	f	regul_official	regul_official	ИБМ6-11Б	Мурыгин Даниил Александрович	a614efba-69c4-43fb-a9a4-d70d8be4c365	1	f
515eed81-ae59-4f49-aba6-c79e1f7ffd1e	$2a$08$iEne9GtWIJKqPwYg2tL3R.TcL7uxEdbb1M/J4Vs.xR4I8KeNgRjt6	89165922247	lizabbkn@gmail.com	f	lklsss	lcklz	Э8-12Б	Бабкина Елизавета Владимировна	05735fa8-6c86-4424-865c-79df5b9850c2	1	f
f4dc6a08-df4f-4ad8-b39d-96fe64a1a7a3	$2a$08$gTzeDJDxEqraKnde/XCj1OkMecb9qXdk7chX665yQi2tpVu2A2uWS	89531344128	vladfokin646@gmail.com	f	VladFokin	vlad2308f	ИУ3-11Б	Фокин Владислав Дмитриевич	\N	\N	f
aac52384-66a1-4609-a3a9-e214e5a449c4	$2a$08$9yFzbZS8E1i9jXKAk/wMnOc/JSxf5eOJE.yhOyS9sFkwJdMHDmia.	89672630616	tigerkrut03@yandex.ru	f	ULV_SSU	id252182879	Э4-11Б	Улыбин Леонид Вадимович	\N	\N	f
d7205129-5902-459a-9a53-05e7a35dc710	$2a$08$HS2ltvQN1hMwHZWmRdjHPuDvRrNwRXkD4EKFQP0JsATaLMAXa5aZy	8 927 02109-54	ir1nak0z@yandex.ru	f	k_ira_18	k_ira_18	ИУ7-13Б	Козлова Ирина Васильевна	\N	\N	f
923aba22-34da-4ddd-994e-756e5336d8cc	$2a$08$bVz6V2TNr.ndktKfz12vfuscRkSE4yjzqq5XIyy5w1L0r1kTqIZfC	89513125984	daruianna@yandex.ru	f	daruianna	daruianna	РК6-14Б	Козлова Дарья Сергеевна	2b3b1e0a-1f6f-4f7d-a0da-7877dbf854f5	3	f
e3cfa1c7-3f6c-4aee-b601-17ea1002e7dd	$2a$08$nUcEcLlFnSsbfXqJfuvRhu6DN6XuzovwCweZYa4PA1b8wNRGLMAAa	89779434961	vladislaustchengaev@yandex.ru	f	Chengaev_Vladislav	awful.siemenss	ФН11-14Б	Ченгаев Владислав Викторович	8f8f8c5c-9df1-43ca-aa12-6eed5edfea17	1	f
03298fba-6b87-49ce-98c8-3a22f4ecb06b	$2a$08$cmtwXvOuY9wjME.VPrUsPu.nKfJOaggZVDgD0G.3uQFlPIb0rzKXG	89197310579	kegor7481@gmail.com	f	korolev170	baza1808	СМ13-13	Королев Егор Витальевич	\N	\N	f
49b43c17-0876-49cc-9f70-b508940d19b0	$2a$08$IL5xagB280SdAom9mYNJrekzk/x13k9BtqcuHYNrrmH9yU6qx6AZu	8 903 96512-24	lenya230305@gmail.com	f	Leo23nid	leoshidz	РК6-16Б	Рыженко Леонид Сергеевич	1c6d8c8a-9bda-4ec8-bbf6-8dd0b83b98d6	1	f
6b26eb3d-8cf9-46e1-9054-85a422d1f27e	$2a$08$49mELcdCWws2GHG7f1SerOvIPFvY/OBQ2SqAtZAAc1Hk23c7hH4lW	89162549540	n1z1renk0@yandex.ru	f	N1Z1renk0	orbit_bez_sachara	ФН1-11Б	Назаренко Александр Павлович	5095faa5-51b3-4dcc-90cd-744b3566ccf2	1	f
843fea24-8d57-4d1e-9d94-5f670dda6a94	$2a$08$LKVQdrIdEU7ZN2V2rar.oey68Dal9fljbNRZk9HUlI0.BIm1xGO5S	8 913 34122-85	tulush.05@mail.ru	f	Natcak_Dorzhu	CHUVAK54	Э2-11Б	Тулуш Натцак-Доржу	9e61f1aa-978a-461a-90f6-99808580d083	1	f
d590ff70-1eff-4fcb-a476-dfc6aa975150	$2a$08$7NbljzlmKk16GV/sLzhX7.oVqq4i9hPIRXdUIYSqp4osci7yPXfeu	89259800841	benafsha2005@icloud.com	f	beno2005	benafsha05	РК6-14Б	Хаяти Бенафша АбдулРахман	2b3b1e0a-1f6f-4f7d-a0da-7877dbf854f5	1	f
87e8b36e-4163-4512-81a5-496711ce14a8	$2a$08$cZM1RrS6ustOSK16LhGAd.3Rn6.qA/CkorNugUc5JxvYh5Ti7HA6q	89854485552	vera.korotkaya@yandex.ru	f	verendaya	vera.korotkaya	ИУ7-11Б	Короткая Вера	a0ff82fd-80f4-4fbd-9e4a-9ee65ab57a3f	2	f
d7e9db7c-9532-4b53-91b0-d5f1c711ee6e	$2a$08$Pns4mvbt/hWvgMWfLL.QOeuJVTnndetjI7fwiccpiv90md4hrfPuu	89258918330	zor05kin@yandex.ru	f	enotcompot	enotikcompotik	БМТ2-11Б	Зорькин Александр Сергеевич	\N	\N	f
3628ad02-1e6c-4c83-a047-fe3a710b4782	$2a$08$/OxnR6DaWp2ARV/zsC4DAuxAeVtGhWD4HMP8i.omGmAmHj6.aCHru	8964-560-99-94	novikovfv@student.bmstu.ru	f	keevony	keevony	ФН2-11Б	Новиков Фёдор Васильевич	\N	\N	f
62e93553-da87-4114-8525-fe4fc36bb102	$2a$08$aUacJnsw47a1EoNE9rQJcefGchdPB2w9UObM9BOSGMliwiAz/C0RG	89267384453	rogovaaan@gmail.com	f	rogovskayaan	rogovskayaan	МТ2-12	Рогова Анна Андреевна	\N	\N	f
7984912d-6adf-431c-a120-08cc15d3051f	$2a$08$rcGEqhVezwRk9c31eCVRleXzKCZXIMa0KLGcD/tAIfV19qvlH1V6e	89267886710	bmd197@yandex.ru	f	b0nduel	b0nduel	СМ7-12	Михаил Денисович	\N	\N	f
b0b3aad5-1ca0-4e36-90ba-dd702bbecb7c	$2a$08$wbGKJSPddP850nRzOIscjOVl9Rc2FHH7G6/IovbXQa8VJ6hVabkXC	8 927 65180-08	egor.syechov@mail.ru	f	panikmacher0	panikmacher	Э6-11	Сычев Егор Игоревич	9f8e20ed-c2b2-4f77-9d95-ca22c7ad82d3	1	f
d80a3bae-76a8-4b4c-a507-a10f87377077	$2a$08$WSBn9dOpdwbz95mfwgR36.c7HoMwkLt4rCgL0B0iyqjMtaorjvQc2	89166357452	polozkov.fedor@yandex.ru	f	Serv_Ant	fpolozkov	Э4-11	Полозков Фёдор Дмитриевич	9f8e20ed-c2b2-4f77-9d95-ca22c7ad82d3	1	f
634a6660-6cf2-4d72-a2da-bab64bad7f18	$2a$08$Kph19xgIvbPD2tCv2tGgOeLLY8XZ9sRgG4QyBUNoU8N4Z.psgrBmm	89168623792	mishakilchitskiy@gmail.com	f	metadern	michael_kilchitskiy	Э5-11Б	Кильчицкий Михаил Сергеевич	9f8e20ed-c2b2-4f77-9d95-ca22c7ad82d3	1	f
b8f55c9f-4dbf-4882-8cff-9934a0551046	$2a$08$ue1nbFdOldfwrkYmgGrQ4.7qtYnQaOP5q6gh86I9qgl1U13ZH1nky	89854612169	polina.alexeenko@gmail.com	f	natashhia	id661776043	ФН4-12Б	Алексеенко Полина Олеговна	\N	\N	f
8a6d944e-5fc9-4792-af0f-e3c2d718bc7b	$2a$08$3rRzmfxxGw3dvmPh7eZBJOioHel43faod.2IceWjriBdLDoD/Qzhe	89805117117	tsarkov0309@gmail.com	f	C0tangent	cotangent	БМТ2-11	Александр Царьков Русланович	966087bd-2504-4ec6-b863-422b8f9d24cf	3	f
47a23567-e0ff-40c0-b7f0-e5a288cedfaf	$2a$08$amJGk7lVtmmiBKacJtI/cuTzTOM3SX1JpOevX/kJ5/BCAKTfhoaRO	89235272535	kryuchenkovalm@bmstu.ru	f	lidiarrrr	0rrrrrrr	ИУ8-13	Крюченкова Лидия Михайловна	\N	\N	f
b1580352-1bdf-4fb8-98aa-968e2503659d	$2a$08$WqQndZgyoCR0IKn2Os9XQurO4B6oa4RY9tLGnCzxugV8xNlCnXmSa	89611227715	bodrikov11@yandex.ru	f	mafia_3d	mafia_3d	АК1-12	Иван Андреевич Бодриков	966087bd-2504-4ec6-b863-422b8f9d24cf	2	f
4bee6776-890b-4e9b-8b57-ac6240ff0b5c	$2a$08$BhODEMgjkzzJHv4zksg5j.DyYbYgB1JEo4iPcZk8.KzNM/GJJdqMa	8915-293-99-92	sladkiha@mail.ru	f	sladkilina	sladkikha	ИБМ6-11Б	Сладких Алина Игоревна	\N	\N	f
77936243-1fda-4fb9-a8ac-8859f54f6377	$2a$08$uWPZrF9GjNq8njmxg8l5w.57sghhgNDPCoksqJ9XtU8lS/uBYq99e	89162436584	shahlovesrk1@gmail.com	f	hloshaaaa	hloshaaa	РЛ2-12Б	Ризаева Шахло Толгонбаевна	2b3b1e0a-1f6f-4f7d-a0da-7877dbf854f5	1	f
2fa23af9-186d-4ffc-b7b1-618b445809a4	$2a$08$A.Cv3QV.xAAskPfYpS247u40p2CMxxi6cJaUCGf4GY1KeTkEqtOvK	89383037731	valiapod1607@gmail.com	f	gopasyhami	gopasyhami	СМ13-13	Подчередниченко Валентина Викторовна	\N	\N	f
37ba4745-11d0-422c-af36-6add8537c398	$2a$08$9HgiCei6nVU3KE/WDSzh8O20mnhOnqFhiSwl3byia3MuFyS2IweLO	89604884400	kate.an.belonogova@gmail.com	f	belonkate	belonkate	ИБМ6-15Б	Белоногова Екатерина Андреевна	\N	\N	f
06844a01-19d6-405b-9de8-4c38dc4507ab	$2a$08$bS2fLNNV2JPNku7GuZ9mHuXXhPEdSfGXi0iwxt8GsVETikseHxrl2	8950-186-50-31	danyaaaaaa4000@gmail.com	f	miy31	miy31	СМ13-13	Ахметов Данис Рушатович	\N	\N	f
793c5646-b64f-4c87-99f1-265fc8a15054	$2a$08$iptQVDR9bIwrtVyNRGBZgOJOXCbSMiFbPRew8Il5K.OLDfziIHdq.	89263924311	kolobanovevgeny@mail.ru	f	Hesore	hesore	ИУ7-16Б	Колобанов Евгений Максимович	\N	\N	f
c0a2a1c6-a6b7-4c43-8027-28125f4f8a14	$2a$08$.9JeCIUFl.xJVjUayJjoaOz1UA715w.dlj3KLuYT/MGeNxArou5QG	89605447659	mark.girichev@yandex.ru	f	Mark_Gir	mark.girichev	ИУ7-16Б	Гиричев Марк Сергеевич	1533d334-657f-455a-94b2-edb6a31ddc18	3	f
979966cf-3bf7-4e90-9eac-d937f7a40171	$2a$08$1l9/YAu/C8XVbxJTwl3WLuf/m5F.a//EAuIQDu2y8QoTwX/t51U9O	89161243618	otieno.a@yandex.ru	f	nevvroti	nevrotik	ФН11-13Б	Отиено Артём Маликович	bd9782f4-4e81-41b1-b6a6-be7bd5bfde4f	3	f
2fdbfc6b-b255-43d7-baff-faa66f6ec082	$2a$08$zBhgefN8TOtLiR2KiXNThO5tViTwmlcrN8fWUkgTzGytsiO81lGgy	89161259945	mantserovpavel@gmail.ru	f	klykalo	mantserovpavel	АК2-11	Манцеров Павел Константинович	\N	\N	f
dcab6078-ac00-49c8-adc1-2b308f2bfcab	$2a$08$/auT5Efqnk1y5.7YIWUJ2uYeFTEO2BwAwt1ZtXZE.HWLugJTX9N7S	89250114042	jess180215@mail.ru	f	ajssaaaa	j.ssaa	ФН3-12Б	Джалагания Александра Дмитриевна	bd9782f4-4e81-41b1-b6a6-be7bd5bfde4f	1	f
911cd60b-262e-4383-be4c-355fead23393	$2a$08$ocSn/UeMmO1QTTTZI4ByXO3STWD16JtIDRJDJUmpVXdwpT3Fzau1i	89251150222	artem-solovyev-2005@ya.ru	f	artmslv	artesol	РК6-12Б	Артём Андреевич Соловьев	\N	\N	f
42f282cc-ccc8-4709-923c-f1cafe16e9d6	$2a$08$wqRYMB1JjbgHmSvZ9P.oROld40ZKsYiawlU.vfOlx.LcgIQDNZopu	89270764742	artsiomziankou2005@gmail.com	f	Tema228666777	pupil_of_bomonka	ФН11-13Б	Зеньков Артём Александрович	bd9782f4-4e81-41b1-b6a6-be7bd5bfde4f	1	f
953dc7b6-bb31-40ec-a3e3-544b0563afd2	$2a$08$F3minBb7/0A/gGVubisJJ.Q.lVWiZPKUNuSBx5WA3PVAzc53GUn7m	89028899714	bogdan.volkov.05@inbox.ru	f	Hdbrgxhdbe	kuplyu_zapchasty_izh_ju5	СМ9-11	Волков Богдан Павлович	572bfabc-f038-4482-97a6-cefd372cd700	1	f
817feb83-5489-4e85-a66b-62a502d2b9c9	$2a$08$IvgTEMXcZ5cu1U7I8enxhujH.KjinnVshEszXngEKN5JN.ACEpQ0K	89153009841	semenovia3@student.bmstu.ru	f	yo_teg	burilkaa	ФН3-12Б	Семенов Илья Александрович	bd9782f4-4e81-41b1-b6a6-be7bd5bfde4f	1	f
ba6fe735-1762-4e30-a47b-944ba25958bc	$2a$08$H/nHuRvi3BkmAHEQquvaAedEfAS2Qm4.M5B1/NHQCOcvEReeMAZ3S	89653100079	marknaumov@icloud.com	f	nau_mar	nau_mrk	ПС4-12	Наумов Марк Алексеевич	572bfabc-f038-4482-97a6-cefd372cd700	1	f
0f287979-d4d9-4b0a-a0f6-765033d892a6	$2a$08$u.xc39NyUAiMSYsjedAyq.vnz.Nxp93aMdhD26vrDbQ5xX0RiVsh6	89640568797	maga_malsagov_03@mail.ru	f	EmiliaForeEver	EmiliaTrueLove	ИУ7-16Б	Магомед Ахмедович Мальсагов	6e7bc851-c93b-4c94-bb2e-fdb0656c528a	2	f
04571ac4-f07b-488f-a37d-ede6469680c3	$2a$08$DA5n76n2k2WgDuY5tkbm8.NA6EQ7jup5Yeq/BeQ/finkR4NCGp0ra	89115677427	ytka.dubinin@yandex.ru	f	zdes_nebila_egora	andreydybinin	Э2-11Б	Дубинин Андрей Александрович	\N	\N	f
53e4c075-c8a1-4787-8212-f2e5983c00ec	$2a$08$rwa4wcNO/ZfrxkuaUN4vYOKUyTGWcPaP5qtoabwU4jgsRjvTESa8K	89639223987	himik-daniil@yandex.ru	f	Raroge21	abvgde321	ФН3-12Б	Волков Даниил Максимович	bd9782f4-4e81-41b1-b6a6-be7bd5bfde4f	1	f
513d8708-abbb-4cc7-860b-715267269413	$2a$08$OMX.fdmFmO3NDrgRIvQuo.NWaZrIHwKvFwKFtpHTK1lgBDGWZc9KO	89091511460	lev.klementev.00@mail.ru	f	levlok	levlok	РК6-11Б	Клементьев Лев Кириллович	572bfabc-f038-4482-97a6-cefd372cd700	1	f
d0d9465d-d6d2-4b0e-b0be-2ebd8c946c3b	$2a$08$wxn4dzMYKQldHpG3aVucI.YDryGue0TS2BXMRiuCBe5y9Xp62v5AS	89094873286	valera-roma06@mail.ru	f	Lns_df	romchik299792458	АК1-11	Валерий Олегович Романов	899a1e33-3acd-4935-9330-ffa80a966412	3	f
ca0c231c-7a12-4d1e-b5c8-4c16eee04d31	$2a$08$HHz93.yQsFoxCCFuwfmtKu/FetwAFB8F/ZcTmmwXEkpLOiDxYYfya	8965-143-56-92	iljushinavaleriia@yandex.ru	f	winxvv	winxvv	АК1-11	Ильюшина Валерия Сергеевна	899a1e33-3acd-4935-9330-ffa80a966412	1	f
b20286eb-173a-4115-a87d-b2fd23b77920	$2a$08$ZTI6nasOXfHVgTB02Exxyu1sEFiQsHgnsW3VNNcTcedzYxZ/2lWZ6	8 908 218 92 83	sadokhin1606@gmail.com	f	89082189283	listovs	АК1-11	Садохин Антон Владимирович	899a1e33-3acd-4935-9330-ffa80a966412	1	f
65a90aaa-423b-48be-bf64-1ba426991990	$2a$08$C55UcVO6A3e7aJ40Qz2yr.6vj8qS0.6RlJMJZSqcHl5i187OTGY76	897744826-02	warst353@mail.ru	f	TsuMaks	m.tsurka	МТ13-12Б	Цурка Максим Витальевич	\N	\N	f
32dd76f5-9223-42c9-959b-cf4bc46b6195	$2a$08$oaWKKqIeyEZxGLzRo5InpOwqMw/9/mZJQe5FEM2N6DwMWDZ4sF23W	89124658208	mishard19@gmail.com	f	mikharda	kroshkaenot_pro	МТ11-12Б	Ардашев Михаил Васильевич	0c21a026-7815-4005-ab4f-ac2b5192acb9	1	f
ef0bc4d7-4636-4016-9d46-86dce0f0b5e7	$2a$08$HDG8soM.dC9ysJmmWRi1Z.zDLANurHqGmXixW7MPy4L99LesB0j6S	89574468912	qwe123@gmail.com	f	tgttggtgt	vkvkvkkv	ИУ4-12	Мельников Дмитрий Вячеславович	\N	\N	f
59cbb8a7-56ae-46c6-bdc2-05e8df3dd729	$2a$08$NH37iNzGfVVjd5niTDMKNO8oOhpLgJ3c3JVuJp8Cl2HXLVgwiQ40y	89853479890	vita.rybina@list.ru	f	ultravvviolettt	villxw	ЮР-11	Виолетта Рыбина Игоревна	\N	\N	f
10879010-7b26-4649-a5a4-43b9a31c4e16	$2a$08$8T2N4yIaX1X0xHpi1O9WjewEyj6LD46MzmNMO98b7Clu2bpDJLi/O	89310083449	androsova_var@mail.ru	f	vrvrstv	frossch	Э7-11	Андросова Варвара Дмитриевна	16956eba-5185-4ba2-bce2-a8239d6a89a9	1	f
86318f32-0ff5-4ab5-98d7-8ddc6c8f01ff	$2a$08$FcMyR3cYCmOFf1W0/AhU2OMzheakesv4ZvIQ0Fib1MfH6BQFi5EO2	89101019421	sip_keyed0g@icloud.com	f	asapkocky	platonkosarev123	СМ10-11Б	Косарев Платон Александрович	5095faa5-51b3-4dcc-90cd-744b3566ccf2	1	f
f6a5ad0c-34eb-4f4a-902b-c66337ee26f0	$2a$08$bLbADVKDlGxg25HwGUUYxefJ5DeFVFZeCZFYm.a1qliIGGtk6/Hai	8916-146-09-43	yuklapotun@gmail.com	f	kloopiik	kloopiik	МТ10-12	Клапотун Юлия Павловна	2b3b1e0a-1f6f-4f7d-a0da-7877dbf854f5	1	f
e25a659b-aea0-44fa-9372-bd8f0f55522a	$2a$08$eiMB35diWtjXhG2hfjnBqe5uvh4p0/S40JSwW6t1Rr28iLQYMAGOm	89196797530	sasha2005962@mail.ru	f	belobrisiy0	sashaivanov2	Э7-11	Sasha Ivanov	16956eba-5185-4ba2-bce2-a8239d6a89a9	1	f
\.
