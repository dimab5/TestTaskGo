# Приложение EWallet

Приложение EWallet предназначено для реализации системы обработки транзакций платежной системы. Приложение реализовано в виде HTTP-сервера, предоставляющего REST API. Сервер должен реализовать четыре метода с следующей логикой:

### 1. Создание кошелька

- **Эндпоинт:** `POST /api/v1/wallet`
- **Параметры запроса:** Отсутствуют
- **Ответ:** JSON-объект, содержащий состояние созданного кошелька с параметрами:
  - `id`: Строковый идентификатор кошелька (генерируется сервером)
  - `balance`: Вещественное число, представляющее баланс кошелька
- Созданный кошелек должен иметь баланс 100.0 у.е.

### 2. Перевод средств

- **Эндпоинт:** `POST /api/v1/wallet/{walletId}/send`
- **Параметры запроса:**
  - `walletId`: Строковый идентификатор кошелька (указан в пути запроса)
  - JSON-объект в теле запроса с параметрами:
    - `to`: ID кошелька, на который нужно перевести деньги
    - `amount`: Сумма перевода
- **Статус ответа:**
  - 200, если перевод успешен
  - 404, если исходящий кошелек не найден
  - 400, если целевой кошелек не найден или на исходящем кошельке недостаточно средств

### 3. Получение истории транзакций

- **Эндпоинт:** `GET /api/v1/wallet/{walletId}/history`
- **Параметры запроса:**
  - `walletId`: Строковый идентификатор кошелька (указан в пути запроса)
- **Ответ с статусом:**
  - 200, если кошелек найден. Ответ должен содержать массив JSON-объектов с входящими и исходящими транзакциями кошелька. Каждый объект содержит параметры:
    - `time`: Дата и время перевода в формате RFC 3339
    - `from`: ID исходящего кошелька
    - `to`: ID входящего кошелька
    - `amount`: Сумма перевода (вещественное число)
  - 404, если указанный кошелек не найден

### 4. Получение текущего состояния кошелька

- **Эндпоинт:** `GET /api/v1/wallet/{walletId}`
- **Параметры запроса:**
  - `walletId`: Строковый идентификатор кошелька (указан в пути запроса)
- **Ответ с статусом:**
  - 200, если кошелек найден. Ответ должен содержать JSON-объект с текущим состоянием кошелька. Объект содержит параметры:
    - `id`: Строковый идентификатор кошелька (генерируется сервером)
    - `balance`: Вещественное число, представляющее баланс кошелька
  - 404, если кошелек не найден

## Требования к реализации

### Безопасность:

- В приложении не должно быть уязвимостей, позволяющих произвольно менять данные в базе.

### Персистентность:

- Данные и изменения не должны "теряться" при перезапуске приложения.

### Стек технологий:

- Язык реализации – Go 1.21
- База данных – PostgreSQL, SQLite или MongoDB

### Дополнительно (плюсом будет):

- Dockerfile для сборки контейнера с приложением.
- Хранение исходного кода в системе контроля версий (например, Git) с публикацией на GitHub. В решении предоставить ссылку на репозиторий.
