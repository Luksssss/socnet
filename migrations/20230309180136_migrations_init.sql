-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE city
(
    -- первичный ключ
    id serial primary key,
    -- город
    name text NOT NULL
);
ALTER TABLE "city" ADD CONSTRAINT "name_size" CHECK (LENGTH("name") <= 128) NOT VALID;
ALTER TABLE "city" VALIDATE CONSTRAINT name_size;

CREATE TABLE gender
(
    -- первичный ключ
    id smallserial primary key,
    -- пол
    name text NOT NULL
);
ALTER TABLE "gender" ADD CONSTRAINT "name_size" CHECK (LENGTH("name") <= 30) NOT VALID;
ALTER TABLE "gender" VALIDATE CONSTRAINT name_size;

CREATE TABLE users
(
    -- первичный ключ
    id bigserial primary key,
    -- имя
    first_name text NOT NULL,
    -- фамилия
    second_name text NOT NULL,
    -- дата рождения
    date_birth  date NOT NULL,
    -- город
    id_city integer references city(id),
    -- пол
    id_gender integer references gender(id),
    -- пароль
    pass text NOT NULL
);
ALTER TABLE "users" ADD CONSTRAINT "first_name_size" CHECK (LENGTH("first_name") <= 30) NOT VALID;
ALTER TABLE "users" VALIDATE CONSTRAINT first_name_size;

ALTER TABLE "users" ADD CONSTRAINT "second_name_size" CHECK (LENGTH("second_name") <= 30) NOT VALID;
ALTER TABLE "users" VALIDATE CONSTRAINT second_name_size;

ALTER TABLE "users" ADD CONSTRAINT "pass_size" CHECK (LENGTH("pass") <= 64) NOT VALID;
ALTER TABLE "users" VALIDATE CONSTRAINT pass_size;

INSERT INTO city (name) VALUES
    ('Абакан'),
    ('Азов'),
    ('Александров'),
    ('Алексин'),
    ('Альметьевск'),
    ('Анапа'),
    ('Ангарск'),
    ('Анжеро-Судженск'),
    ('Апатиты'),
    ('Арзамас'),
    ('Армавир'),
    ('Арсеньев'),
    ('Артем'),
    ('Архангельск'),
    ('Асбест'),
    ('Астрахань'),
    ('Ачинск'),
    ('Балаково'),
    ('Балахна'),
    ('Балашиха'),
    ('Балашов'),
    ('Барнаул'),
    ('Батайск'),
    ('Белгород'),
    ('Белебей'),
    ('Белово'),
    ('Белогорск (Амурская область)'),
    ('Белорецк'),
    ('Белореченск'),
    ('Бердск'),
    ('Березники'),
    ('Березовский (Свердловская область)'),
    ('Бийск'),
    ('Биробиджан'),
    ('Благовещенск (Амурская область)'),
    ('Бор'),
    ('Борисоглебск'),
    ('Боровичи'),
    ('Братск'),
    ('Брянск'),
    ('Бугульма'),
    ('Буденновск'),
    ('Бузулук'),
    ('Буйнакск'),
    ('Великие Луки'),
    ('Великий Новгород'),
    ('Верхняя Пышма'),
    ('Видное'),
    ('Владивосток'),
    ('Владикавказ'),
    ('Владимир'),
    ('Волгоград'),
    ('Волгодонск'),
    ('Волжск'),
    ('Волжский'),
    ('Вологда'),
    ('Вольск'),
    ('Воркута'),
    ('Воронеж'),
    ('Воскресенск'),
    ('Воткинск'),
    ('Всеволожск'),
    ('Выборг'),
    ('Выкса'),
    ('Вязьма'),
    ('Гатчина'),
    ('Геленджик'),
    ('Георгиевск'),
    ('Глазов'),
    ('Горно-Алтайск'),
    ('Грозный'),
    ('Губкин'),
    ('Гудермес'),
    ('Гуково'),
    ('Гусь-Хрустальный'),
    ('Дербент'),
    ('Дзержинск'),
    ('Димитровград'),
    ('Дмитров'),
    ('Долгопрудный'),
    ('Домодедово'),
    ('Донской'),
    ('Дубна'),
    ('Евпатория'),
    ('Егорьевск'),
    ('Ейск'),
    ('Екатеринбург'),
    ('Елабуга'),
    ('Елец'),
    ('Ессентуки'),
    ('Железногорск (Красноярский край)'),
    ('Железногорск (Курская область)'),
    ('Жигулевск'),
    ('Жуковский'),
    ('Заречный'),
    ('Зеленогорск'),
    ('Зеленодольск'),
    ('Златоуст'),
    ('Иваново'),
    ('Ивантеевка'),
    ('Ижевск'),
    ('Избербаш'),
    ('Иркутск'),
    ('Искитим'),
    ('Ишим'),
    ('Ишимбай'),
    ('Йошкар-Ола'),
    ('Казань'),
    ('Калининград'),
    ('Калуга'),
    ('Каменск-Уральский'),
    ('Каменск-Шахтинский'),
    ('Камышин'),
    ('Канск'),
    ('Каспийск'),
    ('Кемерово'),
    ('Керчь'),
    ('Кинешма'),
    ('Кириши'),
    ('Киров (Кировская область)'),
    ('Кирово-Чепецк'),
    ('Киселевск'),
    ('Кисловодск'),
    ('Клин'),
    ('Клинцы'),
    ('Ковров'),
    ('Когалым'),
    ('Коломна'),
    ('Комсомольск-на-Амуре'),
    ('Копейск'),
    ('Королев'),
    ('Кострома'),
    ('Котлас'),
    ('Красногорск'),
    ('Краснодар'),
    ('Краснокаменск'),
    ('Краснокамск'),
    ('Краснотурьинск'),
    ('Красноярск'),
    ('Кропоткин'),
    ('Крымск'),
    ('Кстово'),
    ('Кузнецк'),
    ('Кумертау'),
    ('Кунгур'),
    ('Курган'),
    ('Курск'),
    ('Кызыл'),
    ('Лабинск'),
    ('Лениногорск'),
    ('Ленинск-Кузнецкий'),
    ('Лесосибирск'),
    ('Липецк'),
    ('Лиски'),
    ('Лобня'),
    ('Лысьва'),
    ('Лыткарино'),
    ('Люберцы'),
    ('Магадан'),
    ('Магнитогорск'),
    ('Майкоп'),
    ('Махачкала'),
    ('Междуреченск'),
    ('Мелеуз'),
    ('Миасс'),
    ('Минеральные Воды'),
    ('Минусинск'),
    ('Михайловка'),
    ('Михайловск (Ставропольский край)'),
    ('Мичуринск'),
    ('Москва'),
    ('Мурманск'),
    ('Муром'),
    ('Мытищи'),
    ('Набережные Челны'),
    ('Назарово'),
    ('Назрань'),
    ('Нальчик'),
    ('Наро-Фоминск'),
    ('Находка'),
    ('Невинномысск'),
    ('Нерюнгри'),
    ('Нефтекамск'),
    ('Нефтеюганск'),
    ('Нижневартовск'),
    ('Нижнекамск'),
    ('Нижний Новгород'),
    ('Нижний Тагил'),
    ('Новоалтайск'),
    ('Новокузнецк'),
    ('Новокуйбышевск'),
    ('Новомосковск'),
    ('Новороссийск'),
    ('Новосибирск'),
    ('Новотроицк'),
    ('Новоуральск'),
    ('Новочебоксарск'),
    ('Новочеркасск'),
    ('Новошахтинск'),
    ('Новый Уренгой'),
    ('Ногинск'),
    ('Норильск'),
    ('Ноябрьск'),
    ('Нягань'),
    ('Обнинск'),
    ('Одинцово'),
    ('Озерск (Челябинская область)'),
    ('Октябрьский'),
    ('Омск'),
    ('Орел'),
    ('Оренбург'),
    ('Орехово-Зуево'),
    ('Орск'),
    ('Павлово'),
    ('Павловский Посад'),
    ('Пенза'),
    ('Первоуральск'),
    ('Пермь'),
    ('Петрозаводск'),
    ('Петропавловск-Камчатский'),
    ('Подольск'),
    ('Полевской'),
    ('Прокопьевск'),
    ('Прохладный'),
    ('Псков'),
    ('Пушкино'),
    ('Пятигорск'),
    ('Раменское'),
    ('Ревда'),
    ('Реутов'),
    ('Ржев'),
    ('Рославль'),
    ('Россошь'),
    ('Ростов-на-Дону'),
    ('Рубцовск'),
    ('Рыбинск'),
    ('Рязань'),
    ('Салават'),
    ('Сальск'),
    ('Самара'),
    ('Санкт-Петербург'),
    ('Саранск'),
    ('Сарапул'),
    ('Саратов'),
    ('Саров'),
    ('Свободный'),
    ('Севастополь'),
    ('Северодвинск'),
    ('Северск'),
    ('Сергиев Посад'),
    ('Серов'),
    ('Серпухов'),
    ('Сертолово'),
    ('Сибай'),
    ('Симферополь'),
    ('Славянск-на-Кубани'),
    ('Смоленск'),
    ('Соликамск'),
    ('Солнечногорск'),
    ('Сосновый Бор'),
    ('Сочи'),
    ('Ставрополь'),
    ('Старый Оскол'),
    ('Стерлитамак'),
    ('Ступино'),
    ('Сургут'),
    ('Сызрань'),
    ('Сыктывкар'),
    ('Таганрог'),
    ('Тамбов'),
    ('Тверь'),
    ('Тимашевск'),
    ('Тихвин'),
    ('Тихорецк'),
    ('Тобольск'),
    ('Тольятти'),
    ('Томск'),
    ('Троицк'),
    ('Туапсе'),
    ('Туймазы'),
    ('Тула'),
    ('Тюмень'),
    ('Узловая'),
    ('Улан-Удэ'),
    ('Ульяновск'),
    ('Урус-Мартан'),
    ('Усолье-Сибирское'),
    ('Уссурийск'),
    ('Усть-Илимск'),
    ('Уфа'),
    ('Ухта'),
    ('Феодосия'),
    ('Фрязино'),
    ('Хабаровск'),
    ('Ханты-Мансийск'),
    ('Хасавюрт'),
    ('Химки'),
    ('Чайковский'),
    ('Чапаевск'),
    ('Чебоксары'),
    ('Челябинск'),
    ('Черемхово'),
    ('Череповец'),
    ('Черкесск'),
    ('Черногорск'),
    ('Чехов'),
    ('Чистополь'),
    ('Чита'),
    ('Шадринск'),
    ('Шали'),
    ('Шахты'),
    ('Шуя'),
    ('Щекино'),
    ('Щелково'),
    ('Электросталь'),
    ('Элиста'),
    ('Энгельс'),
    ('Южно-Сахалинск'),
    ('Юрга'),
    ('Якутск'),
    ('Ялта'),
    ('Ярославль');

CREATE INDEX full_name_idx ON users (lower(first_name) text_pattern_ops, lower(second_name) text_pattern_ops);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
