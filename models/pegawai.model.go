package models

import (
	"context"
	"github.com/rafimuhammad01/learn-go-echo-rest/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Pegawai struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Nama    string `bson:"nama" json:"nama"`
	Alamat  string `bson:"alamat" json:"alamat"`
	Telepon string `bson:"telepon" json:"telepon"`
}




//GET
func FetchAllPegawai() (Response, error) {

	
	var arrObj []Pegawai
	var res Response

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	con, err := db.Connect()
	if err != nil {
		return res, err
	}
	//defer db.Disconnect()

	cursor, err := con.Collection("pegawai").Find(ctx, bson.M{})
	if err != nil {
		return res, err
	}

	for cursor.Next(ctx) {
		var obj Pegawai
		err := cursor.Decode(&obj)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)

	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrObj
	return res, nil


}

//POST
func StorePegawai(nama, alamat, telepon string) (Response, error) {
	var res Response
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)


	con, err := db.Connect()
	if err != nil {
		return res, err
	}
	//defer db.Disconnect()

	pegawaiCollection := con.Collection("pegawai")

	pegawaiResult, err := pegawaiCollection.InsertOne(ctx,Pegawai{
			ID : primitive.NewObjectID(),
			Nama:    nama,
			Alamat:  alamat,
			Telepon: telepon,
		},
	)


	if err != nil {
		return res, err
	}

	id := pegawaiResult.InsertedID

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]interface{}{"id": id}
	return res, nil
}

//UPDATE
func UpdatePegawai(_id primitive.ObjectID, newNama, newAlamat, newTelepon string) (Response, error) {
	var res Response
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	con, err := db.Connect()

	if err != nil {
		return res, err
	}

	selector := bson.M{"_id" : _id}
	changes := Pegawai{
		ID:      _id,
		Nama:    newNama,
		Alamat:  newAlamat,
		Telepon: newTelepon,
	}


	pegawaiCollection := con.Collection("pegawai")


	_, err = pegawaiCollection.UpdateOne(ctx,selector,bson.M{"$set": changes})
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = changes

	return res, nil

}

//Delete
func DeletePegawai(_id primitive.ObjectID) (Response, error) {
	var res Response
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	con, err := db.Connect()

	pegawaiCollection := con.Collection("pegawai")

	pegawaiResult, err := pegawaiCollection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]interface{}{"deleted_count" : pegawaiResult.DeletedCount}

	return res, nil

}

