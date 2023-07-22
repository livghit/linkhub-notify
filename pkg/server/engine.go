package server

/*
This will be the rendering engine for the application here you will be able to change to the rendering engine you want 
for more information about rendering engines please follow fibers documents https://docs.gofiber.io/template/

here we will have a variable that will controll the templating engine it will be private and have a public setter
exp 

by default it will be mustache , but you will be able to change it with the ChangeTemplateEngine function

var templateEngine := mustache.New('path to the views' , mustache)

we will also have for each templating engine a variable called like the engine taht will retunr a string './enginename'
ex.
var (
mustache := './mustache'
django := './django'
....
)
*/



