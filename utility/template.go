package utility
//RenderTemplate ....
func RenderTemplate(w http.ResponseWriter,templatefile string,templatedata *signup){
	t,err:=template.Parsefiles(templatefile)
	if err!=nil{
		log.fatal("error encountered:",err)
	}
	t.Execute(w,templatedata)

}
//RenderUnsafeTemplate ........
func RenderUnsafeTemplate(w http.ResponseWriter,templatefile string,templatedata *signup){
	t,err:=template.Parsefiles(templatefile)
	if err!=nil{
		log.fatal("error encountered:",err)
	}
	w.Header().Set("X-XSS-Protection","0")
	t.Execute(w,templatedata)
	

}