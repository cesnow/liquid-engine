package Settings

type GinConf struct {
	RunMode      string `envField:"gin:RunMode" default:"debug"`
	HttpPort     int    `envField:"gin:HttpPort" default:"8080"`
	ReadTimeout  int    `envField:"gin:ReadTimeout" default:"10"`
	WriteTimeout int    `envField:"gin.WriteTimeout" default:"10"`
}

//func (ginConf *GinConf) ApplyConfig(engine *Engine) {
//	ele := reflect.ValueOf(ginConf).Elem()
//	typeOfT := ele.Type()
//	for i := 0; i < ele.NumField(); i++ {
//		f := ele.Field(i)
//
//		if !f.CanSet() {
//			continue
//		}
//
//		fName := typeOfT.Field(i).Name
//		fTag := typeOfT.Field(i).Tag
//
//		fmt.Println(fTag.Get("default"))
//		//pv := reflect.New(f.Type().Elem())
//		field := ele.FieldByName(fName)
//		field.SetString(reflect.New(f.Type()))
//	}
//}
