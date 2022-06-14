# Тестовое задание: Система обработки событий
Необходимо реализовать систему приёма и обработки событий, поступающих от распределённых агентов.
     
### Use Cases
1. Регистрация приходящих событий от агентов
   - конфиденциальная часть события шифруется
   - событие записывается в хранилище

2. Разбор событий
   - Аналитик может просматривать события (включая секретную информацию)
   - Аналитик может отмечать подозрительные события как "инцидент"
   
3. Администрирование системы
   - Администратор может подтвердить регистрацию нового пользователя
   - Администратор может контролировать состояние системы, просматривая список событий без секретной информации 
   
   
### Требования 
1. Система принимает события от агентов по протоколу TCP и сохраняет в персистентное хранилище
   - События от агентов представляют собой поток line-delimited JSON структур аналогичных
     ```go
        type Event struct {
        	EventID    int
        	Created    time.Time
        	SystemName string
        	Message    string
        }
     ```
   - Поле `Message` в получаемых событиях представляет собой конфиденциальную информацию и всегда должно
     храниться в зашифрованном виде. Шифрование должно производиться при помощи приложенной проприетарной
     библиотеки **CloseSSL**
2. Система предоставляет API для доступа к событиям зарегистрированным пользователям:
   - Роль пользователя указывается при регистрации, регистрация должна быть подтверждена
   - Пользователям с ролью "Аналитик" доступен просмотр событий (включая конфиденциальные данные)
     и возможность пометить  одно или несколько событий как "инцидент"
   - Пользователям с ролью "Администратор" доступен просмотр публичных данных событий и возможность
     подтвердить регистрацию пользователя
   
  
### Что сделать

В рамках тестового задания необходимо разработать Backend описанной системы: 
приём и сохранение событий, API (и логику) для веб-интерфейса. 

Неописанные явно архитектурные требования остаются на усмотрение разработчика.
В особо емких местах допускаются TODO с описанием того, как должен работать функционал

### Будет плюсом
- Постраничное отображение событий, сортировка и фильтрация событий на стороне сервера
- Журналирование действий пользователей 
- Тесты, документация API

### Полезные ссылки
- https://linux.die.net/man/3/dlopen
- https://catonmat.net/simple-ld-preload-tutorial

### Приложенные материалы
- `closessl/closessl.h`, `closessl/libclosessl.so` и `closessl/libclosessl.dylib` -- заголовочный
   файл и динамические библиотеки для Linux и MacOS соответственно
- `agent.elf` и `agent.macho` -- тестовые агенты для отладки системы, скомпилированные для 
   Linux и MacOS соответственно. Могут быть использованы для генерации поступающих от агентов
   событий. Использование: `agent.elf --help`