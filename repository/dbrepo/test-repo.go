package dbrepo

import (
	"errors"
	"log"
	"time"

	"github.com/bertoxic/bert/models"

)

func (m *testDBRepo) AllUsers() bool {
	return true
}

func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {

	if res.RoomID == 2 {
		return 0, errors.New("cannot return room id 2 in test repo")
	}
	return 1, nil
}

//inserts a room restriction into the database

func (m *testDBRepo) InsertRoomRestriction(md models.RoomRestriction) error {

	if md.RoomID == 1000 {
		return errors.New("some error")
	}
	return nil

}

// returns true if availability for roomID but false if not

func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	if roomID == 100 {
		return false, errors.New("this roomID does not exist")
	}

	return false, nil
}

// SearchAvailabilityForAllRooms returns room id and name fo all rooms available
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	layout := "2006-01-02"
	str := "2049-12-31"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}

	testDateToFail, err := time.Parse(layout, "2060-01-01")
	if err != nil {
		log.Println(err)
	}

	if start == testDateToFail {
		return rooms, errors.New("some error")
	}

	if start.After(t) {
		return rooms, nil
	}

	// otherwise, put an entry into the slice, indicating that some room is
	// available for search dates
	room := models.Room{
		ID: 1,
	}
	rooms = append(rooms, room)

	return rooms, nil
}

func (m *testDBRepo) GetRoomByID(roomID int) (models.Room, error) {
	if roomID == 100 {
		return models.Room{}, errors.New("something")
	}
	var room models.Room

	if roomID > 50 {
		return room, errors.New("non-existent room")
	}

	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}


func (m *testDBRepo)Authenticate(email, testPassword string)(int, string, error){
	if email=="me@here.com" {
		return 1,"",nil
	}
	return 1, "", errors.New("could not authenticate use thisemail 'me@here.com'") 
}

func (m *testDBRepo)AllReservations()([]models.Reservation, error){

	var reservations []models.Reservation

	return reservations,nil
}

func (m *testDBRepo) AllNewReservations()([]models.Reservation, error){


	var reservations []models.Reservation


		return reservations, nil
	}

	func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	
	
		var res models.Reservation

	return res , nil
	}
	
	func (m *testDBRepo) UpdateReservation (u models.Reservation) error {

		return nil
	}

	func (m *testDBRepo) DeleteReservation(id int) error {
	
		return nil
	}
	
	
	//UpdateReservationProcessed updates processed for reservation
	func (m *testDBRepo) UpdateReservationProcessed(id,processed int) error {
	
		return nil
	}

	
func (m *testDBRepo) AllRooms()([]models.Room, error){
	
	
	var rooms []models.Room
	
	return rooms , nil
	
}

func (m *testDBRepo) GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction,error){

	var restrictions []models.RoomRestriction

	return restrictions, nil
}

func (m *testDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {

		return nil

}

func (m *testDBRepo) DeleteBlockBYID(id int) error {

		return nil

}