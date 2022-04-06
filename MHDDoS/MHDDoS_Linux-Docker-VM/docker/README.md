## 🪖 Docker (Windows, Mac, Linux)

# [auto_mhddos](https://github.com/theorlovsky/auto_mhddos) – інструмент для автоматизації роботи [mhddos_proxy](https://github.com/porthole-ascend-cinnamon/mhddos_proxy)

Інструмент по типу "запусти та забудь". Він сам періодично буде
підтягувати [цілі](https://raw.githubusercontent.com/Aruiem234/auto_mhddos/main/runner_targets)
від [Українського Жнеця](https://t.me/ukrainian_reaper_ddos) та атакувати їх.

## Встановлення

Встановіть Docker з [офіційного сайта](https://docs.docker.com/get-docker/) або будь-яким іншим методом.

ВАЖЛИВО: в залежності від обраного метода, вам може знадобитися запускати команди докера від імені root
користувача (`sudo docker run ...`).

## Застосування

### Запустити у фоні

```shell
docker run -dit --name auto_mhddos --restart unless-stopped --pull always ghcr.io/theorlovsky/auto_mhddos:latest
```

### Подивитися логи

```shell
docker logs -f auto_mhddos
```

### Зупинити

```shell
docker rm -f auto_mhddos
```

### Налаштування

Додатково можна передати наступні параметри під час запуску контейнера:

#### `--parallel 3`, `--parallel all`

За замовчуванням: `1`

Скільки цілей атакувати одночасно. Вкажіть `all` для атаки по усім активним цілям.

#### `--restart-interval 1h`

За замовчуванням: `30m`

Мінімально: `15m`

Як часто оновлювати цілі. Можна вказати у хвилинах (`m`), годинах (`h`) або днях (`d`).

#### `--debug false`

За замовчуванням: `true`

Виводити логи чи ні.

#### Також підтримуються усі параметри з [документації mhddos_proxy](https://github.com/porthole-ascend-cinnamon/mhddos_proxy#usage) (окрім цілей та `-c`)
  
  
