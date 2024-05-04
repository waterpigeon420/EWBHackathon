package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/api/option"

	"github.com/google/generative-ai-go/genai"
)

type Content struct {
	Parts []string `json:Parts`
	Role  string   `json:Role`
}
type Candidates struct {
	Content *Content `json:Content`
}
type ContentResponse struct {
	Candidates *[]Candidates `json:Candidates`
}
type Person struct {
	Calories      string `json:"calories"`
	Diet          string `json:"diet"`
	Carbohydrates string `json:"carbs"`
	Protein       string `json:"protein"`
	Fat           string `json:"fat"`
}

// example request
//
//	"http://localhost:8080/" -Method POST -ContentType "application/json" -Body '{"calories": "2100", "diet": "vegetarian", "carbs": "50", "protein": "25", "fat": "25"}'
func handleRequest(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyBWkfrGEtBFBF1lx1V_4hcwPsI2bwWV3lc"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Make sure you only accept POST methods, if necessary
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON from the request body
	var p Person
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text("You are an agent adept at giving nutritional information. I want to eat "+p.Calories+" calories a day, I am a "+p.Diet+", I want "+p.Carbohydrates+" percent of my diet to be carbs "+p.Protein+" percent to be protein and "+p.Fat+" percent to be fat. I eat 3 meals a day and my metabolic rate is 7100 KJ/day. Suggest groceries I should buy as a shopping list for one day and exact amount for each ingredient in grams. Make the Breakfast the lightest meal by the amount of calories."))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	marshalResponse, _ := json.MarshalIndent(resp, "", "  ")
	//fmt.Println(string(marshalResponse))
	var generateResponse ContentResponse
	if err := json.Unmarshal(marshalResponse, &generateResponse); err != nil {
		log.Fatal(err)
	}
	var response string
	for _, cad := range *generateResponse.Candidates {
		if cad.Content != nil {
			for _, part := range cad.Content.Parts {
				response += part
			}
		}
	}
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
