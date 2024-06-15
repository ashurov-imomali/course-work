create table roles (
id bigserial primary key,
name text not null,
description text not null
);

create table permissions (
id bigserial primary key,
name text not null,
description text not null
);

create table priorities (
id bigserial primary key,
name text not null,
description text not null
);

create table priority2permissions (
priority_id bigint references priorities,
permission_id bigint references permissions
);

create table permissions2role (
permission_id bigint references permissions,
role_id bigint references roles
);

create table departments (
id bigserial primary key,
name text not null,
location text not null,
contact_phone text not null
);

create table employees (
id bigserial primary key,
first_name text not null,
surname text not null,
father_name text not null,
date_of_birth date not null,
gender boolean not null,
citizenship text not null,
passport_series text not null,
passport_number bigint not null,
date_of_issue date not null,
place_of_issue text not null,
residential_address text not null,
role_id bigint references roles,
email text,
phone_number text,
hire_date date,
salary numeric,
department_id bigint references departments
);

create table clients (
id bigserial primary key,
priority_id bigint references priorities,
name text not null,
surname text not null,
father_name text not null,
date_of_birth date not null,
gender boolean not null,
citizenship text not null,
passport_series text not null,
passport_number bigint not null,
date_of_issue text not null,
place_of_issue text not null,
residential_address text not null,
phone text not null,
email text not null,
last_payment_id bigint,
login text not null,
password text not null
);

create table rules (
id bigserial primary key,
commission float not null,
term integer not null,
active boolean not null default true,
min numeric not null default 0 check ( min >= 0 ),
max numeric not null default 1000000000 check ( max > 0 )
);

create table types (
id bigserial primary key,
name text not null,
description text not null,
active boolean not null default true
);

create table services (
id bigserial primary key,
name text null,
description text null,
type_id bigint references types,
rule_id bigint references rules,
active boolean not null default true,
created_at timestamptz default current_timestamp,
created_by bigint references employees,
updated_at timestamptz default current_time,
updated_by bigint references employees,
deleted_at bigint,
deleted_by bigint references employees
);



CREATE TABLE account_types(
id serial primary key,
name text not null,
description text
);

create table accounts(
id serial primary key,
number bigint not null,
type_id bigint references account_types,
client_id bigint references clients,
service_id bigint references services,
amount numeric not null,
active boolean not null default true,
created_at timestamptz default current_timestamp,
updated_at timestamptz default current_timestamp,
closed_at timestamptz
);

create table payments (
id serial primary key,
client_id bigint references clients,
account_id bigint references accounts,
amount numeric check ( amount > 0 ),
payment_time timestamptz default current_timestamp,
status text not null default 'waiting'
);

INSERT INTO "roles" ("id", "name", "descriptions") VALUES
(1, 'Администратор', 'Имеет полный доступ ко всем функциям и настройкам системы'),
(2, 'Менеджер', 'Управляет учетными записями сотрудников и клиентами банка'),
(3, 'Кассир', 'Обрабатывает транзакции и взаимодействует с клиентами на кассе'),
(4, 'Аналитик', 'Просматривает и анализирует финансовые отчеты и статистику'),
(5, 'Служба безопасности', 'Отвечает за безопасность системы и управление доступом пользователей'),
(6, 'Клиентский консультант', 'Консультирует клиентов и предоставляет им информацию о продуктах и услугах банка'),
(7, 'Оператор колл-центра', 'Обрабатывает звонки клиентов и решает их вопросы'),
(8, 'Системный администратор', 'Управляет техническими аспектами системы, включая базу данных и серверы'),
(9, 'Финансовый директор', 'Контролирует финансовую деятельность банка и утверждает крупные транзакции');

INSERT INTO "permissions" ("id", "name", "description") VALUES
(1, 'Чтение', 'Пользователь имеет доступ для просмотра информации о клиентах'),
(2, 'Запись', 'Пользователь имеет доступ для добавления и изменения информации о клиентах'),
(3, 'Удаление', 'Пользователь имеет доступ для удаления информации о клиентах'),
(4, 'Просмотр транзакций', 'Пользователь имеет доступ для просмотра банковских транзакций'),
(5, 'Изменение транзакций', 'Пользователь имеет доступ для редактирования информации о транзакциях'),
(6, 'Управление пользователями', 'Пользователь имеет доступ для создания и управления учетными записями сотрудников'),
(7, 'Просмотр отчетов', 'Пользователь имеет доступ для просмотра финансовых отчетов и статистики'),
(8, 'Управление отчетами', 'Пользователь имеет доступ для создания и изменения финансовых отчетов'),
(9, 'Администрирование', 'Пользователь имеет полный доступ ко всем функциям и настройкам системы');

INSERT INTO "roles" ("id", "name", "descriptions") VALUES
(1, 'Администратор', 'Имеет полный доступ ко всем функциям и настройкам системы'),
(2, 'Менеджер', 'Управляет учетными записями сотрудников и клиентами банка'),
(3, 'Кассир', 'Обрабатывает транзакции и взаимодействует с клиентами на кассе'),
(4, 'Аналитик', 'Просматривает и анализирует финансовые отчеты и статистику'),
(5, 'Служба безопасности', 'Отвечает за безопасность системы и управление доступом пользователей'),
(6, 'Клиентский консультант', 'Консультирует клиентов и предоставляет им информацию о продуктах и услугах банка'),
(7, 'Оператор колл-центра', 'Обрабатывает звонки клиентов и решает их вопросы'),
(8, 'Системный администратор', 'Управляет техническими аспектами системы, включая базу данных и серверы'),
(9, 'Финансовый директор', 'Контролирует финансовую деятельность банка и утверждает крупные транзакции');

INSERT INTO "permissions2role" ("permission_id", "role_id") VALUES
(1, 1), -- Администратор имеет доступ к чтению
(2, 1), -- Администратор имеет доступ к записи
(3, 1), -- Администратор имеет доступ к удалению
(4, 1), -- Администратор имеет доступ к просмотру транзакций
(5, 1), -- Администратор имеет доступ к изменению транзакций
(6, 1), -- Администратор имеет доступ к управлению пользователями
(7, 1), -- Администратор имеет доступ к просмотру отчетов
(8, 1), -- Администратор имеет доступ к управлению отчетами
(9, 1), -- Администратор имеет полный доступ

(1, 2), -- Менеджер имеет доступ к чтению
(2, 2), -- Менеджер имеет доступ к записи
(3, 2), -- Менеджер имеет доступ к удалению
(4, 2), -- Менеджер имеет доступ к просмотру транзакций
(5, 2), -- Менеджер имеет доступ к изменению транзакций
(6, 2), -- Менеджер имеет доступ к управлению пользователями
(7, 2), -- Менеджер имеет доступ к просмотру отчетов

(1, 3), -- Кассир имеет доступ к чтению
(4, 3), -- Кассир имеет доступ к просмотру транзакций
(5, 3), -- Кассир имеет доступ к изменению транзакций

(1, 4), -- Аналитик имеет доступ к чтению
(7, 4), -- Аналитик имеет доступ к просмотру отчетов

(1, 5), -- Служба безопасности имеет доступ к чтению
(6, 5), -- Служба безопасности имеет доступ к управлению пользователями

(1, 6), -- Клиентский консультант имеет доступ к чтению
(4, 6), -- Клиентский консультант имеет доступ к просмотру транзакций

(1, 7), -- Оператор колл-центра имеет доступ к чтению

(1, 8), -- Системный администратор имеет доступ к чтению
(2, 8), -- Системный администратор имеет доступ к записи
(6, 8), -- Системный администратор имеет доступ к управлению пользователями

(1, 9), -- Финансовый директор имеет доступ к чтению
(2, 9), -- Финансовый директор имеет доступ к записи
(4, 9), -- Финансовый директор имеет доступ к просмотру транзакций
(5, 9), -- Финансовый директор имеет доступ к изменению транзакций
(7, 9), -- Финансовый директор имеет доступ к просмотру отчетов
(8, 9); -- Финансовый директор имеет доступ к управлению отчетами

-- приоритеты для клиентов
INSERT INTO "priorities" ("name", "descriptions") VALUES
('обычный', 'Обычный клиент'),
('премиум', 'Премиум клиент'),
('VIP', 'VIP клиент');


-- доступы для обычных клиентов
INSERT INTO permissions (name, description) VALUES
('Просмотр истории транзакций', 'Просмотр истории транзакций'),
('Запрос банковской выписки', 'Запрос банковской выписки'),
('Изменение контактной информации', 'Изменение контактной информации');

-- доступы для премиум клиентов
INSERT INTO permissions (name, description) VALUES
('Перевод денег с комиссией 0%', 'Перевод денег с комиссией 0%'),
('Приоритетное обслуживание в банковских отделениях', 'Приоритетное обслуживание в банковских отделениях'),
('Уведомления о транзакциях по SMS и электронной почте', 'Уведомления о транзакциях по SMS и электронной почте');

-- доступы для VIP клиентов
INSERT INTO permissions (name, description) VALUES
('Доступ к эксклюзивным инвестиционным продуктам', 'Доступ к эксклюзивным инвестиционным продуктам'),
('Личный банковский менеджер', 'Личный банковский менеджер'),
('Бесплатные консультации по финансовому планированию', 'Бесплатные консультации по финансовому планированию');


INSERT INTO priority2permissions ("permission_id", "priority_id") VALUES
(14, 1), (15, 1), (16, 1), -- обычные клиенты
(14, 2), (15, 2), (16, 2), (17, 2), (18, 2), (19, 2), -- премиум клиенты
(14, 3), (15, 3), (16, 3), (17, 3), (18, 3), (19, 3),(20, 3), (21, 3), (22, 3); -- VIP клиенты


INSERT INTO "rules" ("commission", "term", "active")
VALUES
(0.05, 12, true), -- Комиссия 5% за год для кредитов сроком на 12 месяцев
(0.07, 24, true), -- Комиссия 7% за год для кредитов сроком на 24 месяца
(0.03, 6, true), -- Комиссия 3% за год для кредитов сроком на 6 месяцев
(0.01, 3, true), -- Комиссия 1% за год для кредитов сроком на 3 месяца
(0.025, 12, true), -- Комиссия 2.5% за год для вкладов сроком на 12 месяцев
(0.035, 24, true), -- Комиссия 3.5% за год для вкладов сроком на 24 месяца
(0.02, 6, true), -- Комиссия 2% за год для вкладов сроком на 6 месяцев
(0.01, 3, true); -- Комиссия 1% за год для вкладов сроком на 3 месяца

INSERT INTO "types" ("name", "description", "active")
VALUES
('Кредит', 'Финансовая услуга, при которой банк выдает деньги заемщику с обязательством возврата по определенным условиям (срок, проценты и т.д.)', true),
('Депозит', 'Финансовая услуга, при которой клиент оставляет деньги на определенный срок в банке с целью получения дохода в виде процентов', true);


INSERT INTO account_types (name, description)
VALUES
    ('Депозитный', 'Счет для хранения и накопления денежных средств'),
    ('Кредитный', 'Счет для использования заемных средств'),
    ('Расчетный', 'Счет для совершения текущих платежей');