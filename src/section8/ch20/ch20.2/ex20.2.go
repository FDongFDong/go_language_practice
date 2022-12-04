package main

// import "ch20/fedex"
import "ch20/koreaPost"

type Sender interface {
	Send(parcel string)
}
func SendBook(name string , sender Sender) {
	sender.Send(name)
}


func main(){
// var sender Sender = &fedex.FedexSender{}
var sender Sender = &koreaPost.PostSender{}
	// sender := &fedex.FedexSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)

	// poseSender := &koreaPost.PostSender{}
	// SendBokk("어린 놈",poseSender)
}