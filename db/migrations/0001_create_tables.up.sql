CREATE TABLE users (
                       id INTEGER PRIMARY KEY AUTOINCREMENT,  -- уникальный идентификатор записи, автоинкрементный
                       user_id INTEGER NOT NULL, -- идентификатор пользователя Telegram
                       balance REAL NOT NULL DEFAULT 0.0 -- баланс пользователя
);


CREATE TABLE products (
                          id INTEGER PRIMARY KEY AUTOINCREMENT, -- уникальный идентификатор
                          description TEXT, -- описание товара
                          photo_url TEXT, -- URL фотографии товара
                          price REAL NOT NULL, -- цена товара
                          quantity INTEGER NOT NULL DEFAULT 0, -- количество доступных единиц товара
                          location_city TEXT, -- город расположения товара
                          location_coordinates TEXT, -- координаты геолокации в формате "lat,long"
                          hidden_photo_url TEXT DEFAULT '', -- скрытая URL фотографии товара по умолчанию пустая строка
                          hidden_description TEXT DEFAULT '', -- скрытое описание товара по умолчанию пустая строка
                          paid_flag BOOLEAN NOT NULL DEFAULT 0 -- флаг оплаты
);

CREATE TABLE locations (
                           id INTEGER PRIMARY KEY AUTOINCREMENT,
                           country TEXT NOT NULL, -- страна
                           city TEXT NOT NULL, -- город
                           city_district TEXT NOT NULL -- район города
);
