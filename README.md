# Реализованные функции
* Регистраия
* Создание поста(доступно только для авторизированных пользователей). Пользователь, написавший пост, может запретить оставлять под ним комментарии.
* Получения списка постов(можно задать количество и смещение).
* Получение поста и комментариев под ним(можно задать количество и смещение). Комментарии передаются в иерархическом виде, позволяя вложенность без ограничений, как на Хабре и Редите.
* Написание комментария(доступно только для авторизированных пользователей). Размер комментария ограничен, значение ограничения задаётся в конфиге при запуске.
* Подписаться на пост. В этом случае подписавшемуся человеку будут асинхронно приходить все новые комментарии, появляющиеся под постом.
# Использованные Фреймворки
* В приложении используется авторизация при помощи токенов. Для реализаии был использован фреймворк - https://github.com/golang-jwt/jwt/
* Для работы с GraphQL использовался фреймворк - gqlgen.
* В качестве базы данных использовалась PostgreSQL. При этом также была реализована возможность хранить всю информацию в памяти. Выбрать тип используемого хранилища можно при запуске в конфиге. 0 - PostgreSQL, 1 - в памяти.

# Команда для запуска приложения и бд
docker compose up -d

# Возможные команды
## Регистрация 
```
mutation{
  Register(RegisterData:{
    login: "sasha"
  }){
      token
  }
}
```
## Создание поста
```
mutation{
  CreatePost(IdentificationData:{
    login: "sasha"
    token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTY4ODQ4NTcsIkxvZ2luIjoic2FzaGEifQ.WY1nkz1YdEt7I92RwOxYxxLDR_wW5ng2iM1N4Wk-lvs"  }
    PostData:"POSTEREL"
    IsCommentEnbale: false){
      result
  }
}
```
## Получение списка постов
```
query{
  Posts(Limit: 3, Offset: 0){
    ID
    subject
    author
    timeAdd
    isCommentEnable
    
  }
}
```
## Добавление комментариев
```
mutation{
  AddComment(IdentificationData:{
    login: "sasha"
    token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTY4MDA5NjAsIkxvZ2luIjoic2FzaGEifQ.J-NPh1rt0sakjTvdkccOzX56HHLIFhpqAiEys746QYU"  }
  Comment:{
    CommentData: "blabla1"
    PostID: "2"
  }), {
      result
  }
}
```
## Получение комментариев к посту
```
query{
  PostAndComment(PostID: 2,limit: 3){
    Post{
    ID
    subject
    author
    timeAdd
    isCommentEnable 
    }

    comments{
      CommentData
      ParentID
      PostID
      CommentID
      NestingLevel
    }
  }
}
```


## Подписка на пост
```
subscription{
  GetCommentsFromPost(IdentificationData:{
    login: "sasha"
    token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTY4MDA5NjAsIkxvZ2luIjoic2FzaGEifQ.J-NPh1rt0sakjTvdkccOzX56HHLIFhpqAiEys746QYU"  }
    PostID: "2"
  ), {
      CommentID
      CommentData
  }
}
```
