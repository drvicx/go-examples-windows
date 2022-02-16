# go-examples-windows
My various examples of GoLang Windows applications


**=INFO**
- dependencies:<br>
  Go SDK 1.17.3 <br>
  Bootstrap 5.0.2 <br>
  PopperJS 2.9.2  <br>

- IDE:<br>
  GoLand 2021.2.4<br>
  <br>

**=CHANGE LOG**<br>
*новые записи в начале <br>

01: 20220216_1215:
<pre>
- проект очищен для новых экспериментов;
- создана новая структура проекта (по легенде это будет простое веб-приложение на GO,
  т.к. его легко будет впоследствии запускать на Windows):

  /src
   - app
     main.go
     index.tmpl

     - assets
       - css
         main.css
         headers.css
         navbar-top-fixed.css
       - js
         main.js

- к html-шаблону "index.html" подключен CSS-фреймворк Bootstrap и реализована простая страница;
- подключена JavaScript библиотека PopperJS которая необходима для работы некоторых шаблонов Bootstrap; 
- изменения отправлены в "main" ветку репозитория (начальная точка разработки - версия 0.0);

- сборка и запуск приложения:
  cd .\src\app
  go build
  .\app.exe
  http://localhost:8077/

- приложению присвоен тег версии v1.0.0
</pre>
<br>

**=APP-PREVIEW**

- App v1.0.0 -- и адаптивное Меню.. и это было хорошо <br>
  ![preview](_preview/preview_v00_20211220_0320_2-adaptiveMenu.png?raw=true)

- App v1.0.0 -- затем Он создал Дизайн..
  ![preview](_preview/preview_v00_20211220_0320_1-mainView.png?raw=true)

- App v0.0.0 -- в Начале было Ничто.. <br>
  ![preview](_preview/preview_v00_20211220_0040_blank.png?raw=true)
