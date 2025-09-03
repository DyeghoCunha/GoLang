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
		Image: "https://steamcommunity-a.akamaihd.net/economy/image/-9a81dlWLwJ2UUGcVs_nsVtzdOEdtWwKGZZLQHTxDZ7I56KU0Zwwo4NUX4oFJZEHLbXH5ApeO4YmlhxYQknCRvCo04DEVlxkKgpoo6m1FBRp3_bGcjhQ09-jq5WYh8j3Jq_ummJW4NE_3rqSoI2njQSwr0Y6Yj3xIdKXI1c3YVDU_lHvwOnngMXtv5jPziZn6z5iuyi8MxH2DQ/",
		Name:  "USP-S | Córtex",
		ID:    "685c405ff1d688f58946c7f7",
	}

	giveaway := Giveaway{
		Tier:        "Common",
		TicketValue: 1,
		ItemOne:     item,
		EndsAt:      time.Date(2025, time.July, 12, 23, 59, 59, 0, time.UTC),
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
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2NTVkMGZlMWExNTUxNTFjNTUwNjVjYWEiLCJpYXQiOjE3NTE1NDM2OTEsIm5iZiI6MTc1MTU0MzY5MywiZXhwIjoxNzUxNTYxNjkxfQ.UGN2ntmACyBonFqrXDIzqlsSUqQmThE9tHz5CC6g0AnOL6MGC1CRymdjAnrLeGinNEEhYlK5zzmvnCZ64tH2o-3Qi_A1tQ7deibjcDsyoo4aDu0Qz6WEAXa2j0X_v0s8Juvw9k6qUlXyca0_k8hKGYQXhCSuiDvs-EuUV2BF8jk6lnr2rx1k1OP3UH9smV2Z203utsPgHGWpQWvDLemBCJChrY_TB2SBon5xwF0xQpUFH5BiyXny-ahL9h5h-XFEWrU2mmTHeSsen_qsJ_0RANKyxBrVi0n4ujy828jnuLTK4TX0TlWVpZh7FexX5P85jB8VvOsPKm1FHV8jRhFXrQ")

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
