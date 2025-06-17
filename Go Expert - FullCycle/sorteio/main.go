package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Item struct {
	Image string `json:"image"`
	Name  string `json:"name"`
	ID    string `json:"id"`
}

type Giveaway struct {
	Tier        string    `json:"tier"`
	TicketValue int       `json:"ticketValue"`
	ItemOne     Item      `json:"itemOne"`
	EndsAt      time.Time `json:"endsAt"`
}

type RequestWrapper struct {
	Giveaway Giveaway `json:"giveaway"`
}

func main() {
	item := Item{
		Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpovbSsLQJf0PLGeC597c2JloyekvPLP7LWnn8fsZR02O-ZpY3xjAbn_0ZvZG36d4KRdAY7Y1_ZrFm7ye7ngJO0u5icwWwj5Hetfi-2VA/",
		Name:  "Faca de Sobrevivência (★) | Aço Azul",
		ID:    "682f8c03855432713ed93971",
	}

	giveaway := Giveaway{
		Tier:        "Silver",
		TicketValue: 500,
		ItemOne:     item,
		EndsAt:      time.Date(2025, time.July, 1, 23, 59, 59, 0, time.UTC),
	}

	requestBody := RequestWrapper{Giveaway: giveaway}

	body, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatal("Erro ao codificar JSON:", err)
	}

	url := "https://api.eliteskins.gg/admin/new/giveaway"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("Erro ao criar a requisição:", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2NTVkMGZlMWExNTUxNTFjNTUwNjVjYWEiLCJpYXQiOjE3NDk1NzY2MDYsIm5iZiI6MTc0OTU3NjYwOCwiZXhwIjoxNzQ5NTk0NjA2fQ.XEutrmkTwAJnFVKRraEIRgX0CYqv_D6oyzQmUwl5W-AEZl2inmvRnC6sRDrOLoR_1Vz_21aHiWHYZh7IxOr2--iTnHzxB47BYlc_RAnIe06slSZ8WxRdBavbiGFgDFxbGBMR7ia1F9s9OspBvz_c-Dm9irZ0TQVaeD6hjUPeZF3Z7gYC0ctfYTes-N44IcN0WsdAbAW2vi3ey6PqiEfQeAJ0iVgpS1hlwpCUM-fVIpzIpSXQCz0pG0jJ0ziKRAwr1Sznvh_WQ4Ps3Jn7GRkBuKrIOQHVlDHP211CstIzqFW_-JA85QqThWcAEz-IqMuKR3Vrblpf6WZVOOAGoNXLcw")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Erro ao enviar a requisição:", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Erro ao ler a resposta:", err)
	}
	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		fmt.Println("Requisição enviada com sucesso! Status:", resp.Status)
		fmt.Println("Resposta do Servidor:", string(respBody))
	} else {
		fmt.Printf("Falha na requisição. Status: %s\n", resp.Status)
		fmt.Println("Resposta do Servidor:", string(respBody))
	}

}
