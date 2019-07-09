package utility
//RenderTemplate ....
func RenderTemplate(w http.ResponseWriter,templatefile string,templatedata interface{}){
	t,err:=template.Parsefiles(templatefile)
	if err!=nil{
		log.fatal("error encountered:",err)
	}
	t.Execute(w,templatedata)

}
//RenderUnsafeTemplate ........
func RenderUnsafeTemplate(w http.ResponseWriter,templatefile string,templatedata interface{}){
	t,err:=template.Parsefiles(templatefile)
	if err!=nil{
		log.fatal("error encountered:",err)
	}
	t.Execute(w,templatedata)
	

}