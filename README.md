# BaseBE
Base a simple backend with Gorm and Gin (DB: Mysql )

I. Create a Mysql DB
- Create a docker-compose.yml init DB
II. Init list API 
[GIN-debug] GET    /ping                     --> github.com/macduyhai/BaseBE/controllers.(*Controller).Ping-fm (4 handlers)
[GIN-debug] POST   /login                    --> github.com/macduyhai/BaseBE/controllers.(*Controller).Login-fm (4 handlers)
[GIN-debug] POST   /api/v1/user/add          --> github.com/macduyhai/BaseBE/controllers.(*Controller).Create-fm (5 handlers)
[GIN-debug] PUT    /api/v1/user/edit         --> github.com/macduyhai/BaseBE/controllers.(*Controller).Edit-fm (5 handlers)
[GIN-debug] DELETE /api/v1/user/delete       --> github.com/macduyhai/BaseBE/controllers.(*Controller).Delete-fm (5 handlers)