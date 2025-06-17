package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://api.eliteskins.gg/accounts/me"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//	log.Fatal("Erro ao criar a requisição:", err)
	} //

	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2NTVkMGZlMWExNTUxNTFjNTUwNjVjYWEiLCJpYXQiOjE3NDk1NzY2MDYsIm5iZiI6MTc0OTU3NjYwOCwiZXhwIjoxNzQ5NTk0NjA2fQ.XEutrmkTwAJnFVKRraEIRgX0CYqv_D6oyzQmUwl5W-AEZl2inmvRnC6sRDrOLoR_1Vz_21aHiWHYZh7IxOr2--iTnHzxB47BYlc_RAnIe06slSZ8WxRdBavbiGFgDFxbGBMR7ia1F9s9OspBvz_c-Dm9irZ0TQVaeD6hjUPeZF3Z7gYC0ctfYTes-N44IcN0WsdAbAW2vi3ey6PqiEfQeAJ0iVgpS1hlwpCUM-fVIpzIpSXQCz0pG0jJ0ziKRAwr1Sznvh_WQ4Ps3Jn7GRkBuKrIOQHVlDHP211CstIzqFW_-JA85QqThWcAEz-IqMuKR3Vrblpf6WZVOOAGoNXLcw")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//	log.Fatal("Erro ao enviar a requisição:", err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)

}
