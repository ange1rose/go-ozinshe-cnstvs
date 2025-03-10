## Ozinshe

Личная онлайн-платформа, предоставляющая информацию о фильмах. Система позволяет оценивать фильмы, помечать фильмы 
просмотренными и формировать очередь просмотра.

### Функциональные требования

* Создавать, редактировать, удалять фильмы и их данные: включает название, описание, режиссера, год производства, жанр, ссылку на трейлер и постер;
* Иметь возможность сортировать и фильтровать фильмы по критериям;
* Иметь возможность оценивать фильмы;
* Иметь возможность формировать очередь просмотров;
* Иметь возможность пометить фильм просмотренным;
* Создавать, редактировать, удалять жанры;
* Создавать, редактировать, сбрасывать пароль, удалять пользователей;
* Пользователь должен авторизоваться в системе по имейлу и паролю для входа

### Нефункциональные требования

* Наличие документированного API по спецификации OpenAPI
* Возможность контейнеризировать приложение в Docker

## Как запустить UI?

Чтобы запустить UI Озенше локально тебе нужен будет [git](https://git-scm.com/) и [docker](https://www.docker.com/).  

### Запуск UI без авторизации

```
docker run --name ozinshe-ui -p "8080:3000" -e VITE_API_URL="http://localhost:8081" -e VITE_FEATURE_AUTH="false" -e VITE_SIMPLIFIED_MOVIE="true" -d kchsherbakov/ozinshe-ui:latest
```

### Запуск UI с обязательной авторизацией

```
docker run --name ozinshe-ui -p "8080:3000" -e VITE_API_URL="http://localhost:8081" -e VITE_FEATURE_AUTH="true" -e VITE_SIMPLIFIED_MOVIE="false" -d kchsherbakov/ozinshe-ui:latest
```

Для логина используй:
* логин: `admin@admin.com`
* пароль: `admin`