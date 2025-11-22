/*func AFSearchSubjects(c echo.Context) error {
	search:=c.QueryParam("search")
	 if search==""{
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":"Параметр обьязателен(id,iin, or bin)"
		})
	 }
	 reports, err:= GetSearchSubject(search)
	 if err!=nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	 }
	 if len(reports)==0{
		return c.JSON(http.StatusOK, []map[string]string{
			{"message":"Данные не найдены"},
		})
	 }
	 	return c.JSON(http.StatusOK, reports)
}*/
