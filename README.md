# jwt-golang


# POST http://localhost:8000/auth/user/register

body :
{
"name":"pendekar koding",
"email":"pendekar@gmail.com",
"username":"pendekar",
"password":"12345678"
}

# POST http://localhost:8000/auth/token

{
"email":"pendekar@gmail.com",
"password":"12345678"
}