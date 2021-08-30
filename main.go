package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Prithvipal/proto-buffer/src/enumexample/enumexamplepb"
	"github.com/Prithvipal/proto-buffer/src/simple/simplepb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	// sm := dosimlpe()
	// readAndWriteDemo(sm)
	// jsonDemo(sm)

	doenum()
}

func doenum() {
	em := enumexamplepb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumexamplepb.DayOfTheWeek_THURSDAY,
	}
	em.DayOfTheWeek = enumexamplepb.DayOfTheWeek_MONDAY
	fmt.Println(em)
}

func jsonDemo(pb proto.Message) {
	res := toJSON(pb)
	fmt.Println("========")
	fmt.Println(res)
	var sm simplepb.SimpleMessage
	fromJSON(res, &sm)
	fmt.Println("*********")
	fmt.Printf("%+v", sm)
}

func toJSON(pb proto.Message) string {
	jStr, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatalln("can't marshall to json", err)
		return ""
	}
	return string(jStr)
}

func fromJSON(in string, pb proto.Message) proto.Message {
	err := protojson.Unmarshal([]byte(in), pb)
	if err != nil {
		log.Fatal("can't unmarshall from json", err)
		return nil
	}
	return pb
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := simplepb.SimpleMessage{}
	readFromFile("simple.bin", &sm2)
	fmt.Println("READ OPEARATION SUCCESSFUL")
	fmt.Printf("%+v", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("Can't serialize")
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatal("Can't write to file")
		return err

	}
	fmt.Println("Data has been written")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal("Someting went wrong while reading file")
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatal("Can't unmarshal data")
		return err
	}

	return nil

}
func dosimlpe() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SimpleList: []int32{1, 4, 7, 9},
	}

	fmt.Printf("%+v\n", sm)
	sm.Name = "I renamed you"
	fmt.Printf("%+v\n", sm)
	return &sm
}
