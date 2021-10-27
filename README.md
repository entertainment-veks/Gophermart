# Gophermart
Demo project which implements accumulative loyalty system
- - -
#### POST /api/user/register — user registration;
#### POST /api/user/login — user authentication;
#### POST /api/user/orders — user getting order code for payment;
#### GET /api/user/orders — получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях;
#### GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя;
#### POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
#### GET /api/user/balance/withdrawals — получение информации о выводе средств с накопительного счёта пользователем.
