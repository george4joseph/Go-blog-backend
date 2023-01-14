# Go-blog-backend
Backend for a blog application developed in Go Gin and ent


```
User APIs

GET "/blogs"
POST "/users"
DELETE "/user/:id"
GET "/user/:id/blogs
POST "/user/:id/create"
PATCH "/user/:id_user/edit_blog/:id"
POST "/user/:id/delete_blog"

Admin APIs
GET "/admin/:id_admin/users"
POST "/create_admin"
POST "/admin/:id_admin/assign_admin/:id"
POST "/admin/:id_admin/create_user/:id"
DELETE "/admin/:id_admin/delete_user/:id"
DELETE "/admin/:id_admin/delete_blog/:id"
PATCH "/admin/:id_admin/edit_blog/:id"

```
