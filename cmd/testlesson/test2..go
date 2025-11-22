// var in_id =flag.String("id", "id for request")
// var ibso =flag.bool("ibso", false, "sending data to Ibso yes or no?")
// type RowData struct{
// 	lim string
// 	treg string
// 	time int
// 	face string
// }
// func AF_JSON_SUBJECT(in_json1 string)(err error){
// 	q:=	`INSERT INTO public.JSON_SUBJECT(
// 	JSON_, status )VALUES($1,$2)`
// 	_,err=pg.DB().Exec(q,in, _json1,1)
// 	if err!=nil{
// 		log.Printf("AFInsert(%s) fiasco:%s", q,err)
// 	}
// 	return
// }
// func AF_GET_STACK_ID_SUBJECT()(pers []string, err error){
// 	q:=`SELECT DISTINCT id from subject where status=1 order by 1 asc`
// 	err=pg.DB().SELECT(&pers,q)
// 	if err!=nil{
// 		log.Printf("AF_GET_STACK_ID_SUBJECT(%s) fiasco:%s", q,err)
// 	}
// 	return
// }
// func AF_p_parse_subject()(err error){
// 	q:=`call public.p_run_parse_subject();`
// 	_, err!=nil{
// 		log.Printf("target_subject.go %s fiasco:%s", err)
// 	}
// 	return
// }

// func main(){
// 	var (body_ora string
// 	err error)

// flag.Parse()
// comps.LoadConfiguration("../config.json")
// authData,_:= model.AFGetAuthStruct()

// setAfterTokenCreation:=int(time.Since(authData.CreatedAt).Seconds())
// if secAfterTokenCreation>=authData.RefreshExpiresIn{
// 	log.Printf("faces.go >token expired: %s -do auth", authData.CreatedAt)
// 	authData,_,err=model.AfAuth(`faces token expired`)
// 	if err!=nil{
// 		log.Printf("faces.go fiasco(%s) fiasco:%s", err)
// 		os.Exit(1)
// 	}
// }
// model.AFGlobalAuthData =authData
// if len(*in_id) > 0 il{
// V_person := *in_id
// log.Printf("AFRequestPlain .. %s\n", v_person)
// body1, status, err := model.AFRequestPlain("GET", comps.Config.AF.APIURL+`/v1/lists/subject/`+V_person,`AFPersonInfo`, authData.AccessToken)

// if err != nil {
// fmt.Fprintf(os.Stderr, "AFRequestPlain fiasco: %s\n", err)
// os.Exit(1)
// }
// if status == 401 && status == 403 && status == 404 {
// fmt.Fprintln(os.Stderr, "AFRequestPlain returned %d", status)
// os.Exit(1)
// }
// f len(body1) > 0{
// body_ora = fmt.Sprintf("{\"total\":1,\"limit\":1,\"offset\":0,\"items\":[ %s ])", body1)
// err = model.OraSendToIBSO( `F_SETCLARR`, string(body_ora))
// if err != nil {
// log.Printf("target_subject >AF_JSON_LIST_SUBJECT OraSendToIBSO fiasco: %s", err)
// }
// AF_JSON_SUBJECT(string(body1))
// //AF_p_run_parse_subject()
// log.Printf("target_subject >AF_JSON_SUBJECT %s: %d",v_person, len(body1))
// }else{
// log.Printf("target_subject > %s : %d", v_person,len(body1))
// }

// log.Printf("target_subject >AF_p_run_parse_subject")
// time.Sleep(1 * time.Second)
// AF_p_run_parse_subject()
// }

// }
	
	