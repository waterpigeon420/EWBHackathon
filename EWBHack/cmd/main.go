package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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
	Gender        string `json:"gender"`
	BMI           string `json:"BMI"`
	ETotal        string `json:"Etotal"`
	Age           string `json:"age"`
	Allergy       string `json:"allergy"`
}

type Ingredients map[string]string

// RecipeDetails maps ingredients to their quantities
type RecipeDetails map[string]string

// Recipes maps recipe names to their details
type Recipes map[string]RecipeDetails

// example request
//
//	"http://localhost:8080/" -Method POST -ContentType "application/json" -Body '{"calories": "2100", "diet": "vegetarian", "carbs": "50", "protein": "25", "fat": "25"}'
func handleIngredientRequest(w http.ResponseWriter, r *http.Request) {
	/* Log */
	log.Println("*** Calling Gemini ***")

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
	text := "You are an agent that is adept at giving nutritional information and recipe names, with ingredients and their amounts but not the recipe cooking instructions. I am a" + p.Gender + " and my age is " + p.Age + ". I am allergic to " + p.Allergy + " so please do not suggest any recipes or ingredients that contain this food. My BMI is " + p.BMI + ". I want to eat " + p.Calories + " calories a day, I exert " + p.ETotal + " calories a day. I am a " + p.Diet + ", I want " + p.Carbohydrates + " percent of my diet to be carbs " + p.Protein + " percent to be protein and " + p.Fat + " percent to be fat. I eat 3 meals a day and my metabolic rate is 7100 KJ/day. Make the Breakfast the lightest meal by the amount of calories. Give recipe names and ingredients only, do not provide cooking instructions for the recipes, make sure each recipe has calories listed next to it, do not include which meal of the day it is, just recipe name and calories. Make sure the output always contains ingredients for each recipe. List ingredients under recipes as bullet points, also give a separate list of all the ingredients you mention below the title Ingredients (dont mention the food groups, directly list all the ingredients and their respective amounts separated by a comma e.g ingredient,amount). Always start your output with **Recipes** before writing recipes and finally **Ingredients** before writing ingredients. All bullet points, including recipe names, MUST use - . Recipe names should have ** around them."
	log.Println(text)
	resp, err := model.GenerateContent(ctx, genai.Text(text))
	log.Println(resp)
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

func parseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("*** Parsing ingredients ***")

	if r.Method != "POST" {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	text := string(body)
	log.Println(text)
	ingredients, err := parseIngredients(text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(ingredients)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredients)
}

func parseIngredients(data string) (Ingredients, error) {
	startIndex := strings.Index(data, "**Ingredients**")
	if startIndex == -1 {
		return nil, fmt.Errorf("ingredients section not found")
	}

	// Find the next section to determine the end of the ingredients list
	endIndex := strings.Index(data[startIndex+len("**Ingredients**"):], "**")
	if endIndex == -1 { // if no more sections, assume end of data
		endIndex = len(data)
	} else {
		endIndex += startIndex + len("**Ingredients**")
	}

	ingredientsList := data[startIndex+len("**Ingredients**") : endIndex]
	items := strings.Split(ingredientsList, "\n")
	ingredients := make(Ingredients)

	for _, item := range items {
		item = strings.TrimSpace(item)
		if item != "" {
			// Remove any leading bullet points and trim the item
			item = strings.TrimLeft(item, "*- ")
			parts := strings.SplitN(item, ",", 2)
			if len(parts) == 2 {
				ingredient := strings.TrimSpace(parts[0])
				quantity := strings.TrimSpace(parts[1])
				ingredients[quantity] = ingredient
			}
		}
	}
	log.Println(ingredients);
	return ingredients, nil
}

func recipesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("*** Parsing recipes ***")

	if r.Method != "POST" {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	text := string(body)
	log.Println(text)
	recipes, err := parseRecipes(text)
	log.Println(recipes);
	if err != nil {
		http.Error(w, "Error parsing recipes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func parseRecipes(data string) (Recipes, error) {
	recipes := make(Recipes)
	lines := strings.Split(data, "\n")
	var currentRecipe string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "- **") { // Check if the line is a recipe title
			// Extract the recipe title from the line
			titleEnd := strings.Index(line, "** ")
			if titleEnd != -1 {
				currentRecipe = line[4:titleEnd] // Remove the "- **" and trailing characters
				recipes[currentRecipe] = RecipeDetails{}
			}
		} else if currentRecipe != "" && strings.HasPrefix(line, "-") { // Check if the line is an ingredient line
			line = strings.TrimPrefix(line, "- ") // Remove the leading hyphen and space
			parts := strings.SplitN(line, ",", 2) // Split line to extract ingredient and quantity
			if len(parts) == 2 {
				ingredient := strings.TrimSpace(parts[0])
				quantity := strings.TrimSpace(parts[1])
				recipes[currentRecipe][quantity] = ingredient
			}
		}
	}

	return recipes, nil
}

func sustainabilityHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("*** Calling Gemini again ***")
	if r.Method != "POST" {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	recipes, err := parseRecipes(string(body))
	if err != nil {
		http.Error(w, "Error parsing recipes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var ingred []string
	var recipeList []string
	// Loop through the map
	for recipe, recipeDet := range recipes {
		var result string
		recipeList = append(recipeList, recipe)
		for quantity, ingredients := range recipeDet {
			result += quantity + ":" + ingredients + ", "
			ingred = append(ingred, result)
		}
	}

	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyBWkfrGEtBFBF1lx1V_4hcwPsI2bwWV3lc"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	var finalResp string
	model := client.GenerativeModel("gemini-pro")
	for i, _ := range recipeList {
		resp, err := model.GenerateContent(ctx, genai.Text("You are an agent that is able to only give a single score for environmental sustainability for, usually vegetarian recipes, fruits will have a high sustainability while non-vegetarian ones will have low sustanability, if the dish includes both vegetarian and non-vegetarian ingredients, the non-veg ingredient's low sustainability score should weigh more in your final score, but dont make it extremely low, increase the weight of vegetairan ingredients , in your output a score of 1 is the lowest sustainability while 10 is the highest. Output only a sustainability score for "+recipeList[i]+" with ingredients "+ingred[i]+" . Do not provide justification for the calculations or methods, just a score based on the sustainability of each ingredient that is provided for the recipe, the score should be out of 10, e.g some_score/10."))
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
		finalResp += response
	}

	w.Write([]byte(finalResp))
}

func main() {
	log.Println("Running Golang server")
	http.HandleFunc("/geminiResp", handleIngredientRequest)
	http.HandleFunc("/ingredients", parseHandler)
	http.HandleFunc("/recipes", recipesHandler)
	http.HandleFunc("/sustainability", sustainabilityHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
