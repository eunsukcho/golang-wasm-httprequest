package main

func main() {
	
	js.Global().Set("ConvertingDate", convertDate())

}

func convertDate() js.Func {
	now := time.Now()
	dataConstructor := js.Global().Get("Date")
	return dateConstructor.New(now.Unix() * 1000)
}