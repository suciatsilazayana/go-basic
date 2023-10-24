package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name  string
	Class string
}

func saveToFile(users []User, filename string) error {
	// Mengkodekan slice users ke JSON
	jsonData, err := json.Marshal(users)
	if err != nil {
		return err
	}

	// Menyimpan JSON ke file
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func loadFromFile(filename string) ([]User, error) {
	// Membaca isi file JSON
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Membuka file JSON
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func main() {
	users := []User{
		{Name: "NooB", Class: "A"},
		{Name: "Java", Class: "B"},
		{Name: "Docker", Class: "C"},
	}

	err := saveToFile(users, "users.json")
	if err != nil {
		fmt.Println("Gagal menyimpan ke file:", err)
	} else {
		fmt.Println("Data pengguna berhasil disimpan ke users.json")
	}
}
