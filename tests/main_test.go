package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"portto-assignment/internal/handlers"
	"portto-assignment/internal/routes"
	"portto-assignment/internal/services"
	"portto-assignment/tests/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestEndpoints(t *testing.T) {
	buildTestService()

	t.Run("POST /v1/meme-coin/create", testCreateMemeCoinEndpoint)
	t.Run("PATCH /v1/meme-coin/:id", testUpdateMemeCoinEndpoint)
	t.Run("GET /v1/meme-coin/:id", testGetMemeCoinEndpoint)
	t.Run("DELETE /v1/meme-coin/:id", testDeleteMemeCoinEndpoint)
	t.Run("POST /v1/meme-coin/:id/pock", testPockMemeCoinEndpoint)
}

func testCreateMemeCoinEndpoint(t *testing.T) {
	// Case 1: "name" is not in the request body
	noNameInRequestCaseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/v1/meme-coin/create", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(noNameInRequestCaseRecorder, req)

	resJSONstr := noNameInRequestCaseRecorder.Body.String()
	resJSON := map[string]any{}
	json.Unmarshal([]byte(resJSONstr), &resJSON)
	assert.Equal(t, http.StatusBadRequest, noNameInRequestCaseRecorder.Code)
	assert.Equal(t, "Invalid request body", resJSON["message"])
	assert.NotEqual(t, "", resJSON["error"])

	// Case 2: "description" is not in the request body
	noDescriptionInRequestCaseRecorder := httptest.NewRecorder()
	requestBody := map[string]string{
		"name": "name",
	}
	requestBodyJSON, _ := json.Marshal(requestBody)
	timeBefore := time.Now()
	http.Header.Add(req.Header, "Content-Type", "application/json")
	req, err = http.NewRequest("POST", "/v1/meme-coin/create", bytes.NewReader(requestBodyJSON))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(noDescriptionInRequestCaseRecorder, req)
	timeAfter := time.Now()
	resJSONstr = noDescriptionInRequestCaseRecorder.Body.String()
	resJSON = map[string]any{}
	json.Unmarshal([]byte(resJSONstr), &resJSON)

	assert.Equal(t, http.StatusOK, noDescriptionInRequestCaseRecorder.Code)
	assert.Equal(t, "name", resJSON["name"])
	assert.Equal(t, "", resJSON["description"])
	assert.Greater(t, int(resJSON["id"].(float64)), 0)
	assert.Greater(t, int(resJSON["popularity_score"].(float64)), 0)
	assert.Less(t, int(resJSON["popularity_score"].(float64)), 100)
	assert.GreaterOrEqual(t, resJSON["created_at"].(string), timeBefore.Format(time.RFC3339Nano))
	assert.LessOrEqual(t, resJSON["created_at"].(string), timeAfter.Format(time.RFC3339Nano))

	// Case 3: "name" and "description" are in the request body
	bothNameAndDescriptionInRequestCaseRecorder := httptest.NewRecorder()
	requestBody = map[string]string{
		"name":        "name",
		"description": "description",
	}
	requestBodyJSON, _ = json.Marshal(requestBody)
	timeBefore = time.Now()
	http.Header.Add(req.Header, "Content-Type", "application/json")
	req, err = http.NewRequest("POST", "/v1/meme-coin/create", bytes.NewReader(requestBodyJSON))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(bothNameAndDescriptionInRequestCaseRecorder, req)
	timeAfter = time.Now()
	resJSONstr = bothNameAndDescriptionInRequestCaseRecorder.Body.String()
	resJSON = map[string]any{}
	json.Unmarshal([]byte(resJSONstr), &resJSON)

	assert.Equal(t, http.StatusOK, bothNameAndDescriptionInRequestCaseRecorder.Code)
	assert.Equal(t, "name", resJSON["name"])
	assert.Equal(t, "description", resJSON["description"])
	assert.Greater(t, int(resJSON["id"].(float64)), 0)
	assert.Greater(t, int(resJSON["popularity_score"].(float64)), 0)
	assert.Less(t, int(resJSON["popularity_score"].(float64)), 100)
	assert.GreaterOrEqual(t, resJSON["created_at"].(string), timeBefore.Format(time.RFC3339Nano))
	assert.LessOrEqual(t, resJSON["created_at"].(string), timeAfter.Format(time.RFC3339Nano))
}

func testUpdateMemeCoinEndpoint(t *testing.T) {
	// Setup path
	memeCoinId := 1
	updatePath := "/v1/meme-coin/" + strconv.Itoa(memeCoinId)

	// Case 1: "id" is not in the request body
	noIDInRequestCaseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("PATCH", "/v1/meme-coin/", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(noIDInRequestCaseRecorder, req)

	responseStr := noIDInRequestCaseRecorder.Body.String()

	assert.Equal(t, http.StatusNotFound, noIDInRequestCaseRecorder.Code)
	assert.Equal(t, "404 page not found", responseStr)

	// Case 2: "id" is in the url but id is not numeric
	nonNumericIDCaseRecorder := httptest.NewRecorder()
	req, err = http.NewRequest("PATCH", "/v1/meme-coin/abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(nonNumericIDCaseRecorder, req)

	resJSONstr := nonNumericIDCaseRecorder.Body.String()
	resJSON := map[string]any{}
	json.Unmarshal([]byte(resJSONstr), &resJSON)
	assert.Equal(t, http.StatusBadRequest, nonNumericIDCaseRecorder.Code)
	assert.Equal(t, "Invalid MemeCoin ID", resJSON["message"])
	assert.Equal(t, "Wrong ID format", resJSON["error"])

	// Case 3: "id" is in the url but "description" is not in the request body
	descriptionNotInBodyCaseRecorder := httptest.NewRecorder()
	req, err = http.NewRequest("PATCH", updatePath, nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(descriptionNotInBodyCaseRecorder, req)

	resJSONstr = descriptionNotInBodyCaseRecorder.Body.String()
	resJSON = map[string]any{}
	json.Unmarshal([]byte(resJSONstr), &resJSON)
	assert.Equal(t, http.StatusBadRequest, descriptionNotInBodyCaseRecorder.Code)
	assert.Equal(t, "Invalid request body", resJSON["message"])
	assert.NotEqual(t, "", resJSON["error"])

	// Case 4: "id" is in the url and "description" is in the request body
	idAndDescriptionInRequestCaseRecorder := httptest.NewRecorder()
	requestBody := map[string]string{
		"description": "description updated",
	}
	requestBodyJSON, _ := json.Marshal(requestBody)
	timeBefore := time.Now()
	http.Header.Add(req.Header, "Content-Type", "application/json")
	req, err = http.NewRequest("PATCH", updatePath, bytes.NewReader(requestBodyJSON))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(idAndDescriptionInRequestCaseRecorder, req)
	timeAfter := time.Now()
	resJSONstr = idAndDescriptionInRequestCaseRecorder.Body.String()
	resJSON = map[string]any{}

	json.Unmarshal([]byte(resJSONstr), &resJSON)
	assert.Equal(t, http.StatusOK, idAndDescriptionInRequestCaseRecorder.Code)
	assert.Equal(t, "FakeCoin", resJSON["name"])
	assert.Equal(t, requestBody["description"], resJSON["description"])
	assert.Equal(t, int(resJSON["id"].(float64)), memeCoinId)
	assert.Greater(t, int(resJSON["popularity_score"].(float64)), 0)
	assert.Less(t, int(resJSON["popularity_score"].(float64)), 100)
	assert.GreaterOrEqual(t, resJSON["created_at"].(string), timeBefore.Format(time.RFC3339Nano))
	assert.LessOrEqual(t, resJSON["created_at"].(string), timeAfter.Format(time.RFC3339Nano))
}

func testGetMemeCoinEndpoint(t *testing.T) {
	// Setup path
	memeCoinId := 1
	getPath := "/v1/meme-coin/" + strconv.Itoa(memeCoinId)

	// Case 1: "id" is not in the request body
	noIDInRequestCaseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/v1/meme-coin/", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(noIDInRequestCaseRecorder, req)

	responseStr := noIDInRequestCaseRecorder.Body.String()

	assert.Equal(t, http.StatusNotFound, noIDInRequestCaseRecorder.Code)
	assert.Equal(t, "404 page not found", responseStr)

	// Case 2: "id" is in the url but id is not numeric
	nonNumericIDCaseRecorder := httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/v1/meme-coin/abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(nonNumericIDCaseRecorder, req)

	resJSONstr := nonNumericIDCaseRecorder.Body.String()
	resJSON := map[string]any{}
	json.Unmarshal([]byte(resJSONstr), &resJSON)
	assert.Equal(t, http.StatusBadRequest, nonNumericIDCaseRecorder.Code)
	assert.Equal(t, "Invalid MemeCoin ID", resJSON["message"])
	assert.Equal(t, "Wrong ID format", resJSON["error"])

	// Case 3: "id" is in the url
	idInRequestCaseRecorder := httptest.NewRecorder()
	req, err = http.NewRequest("GET", getPath, nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(idInRequestCaseRecorder, req)

	resJSONstr = idInRequestCaseRecorder.Body.String()
	resJSON = map[string]any{}
	json.Unmarshal([]byte(resJSONstr), &resJSON)
	assert.Equal(t, http.StatusOK, idInRequestCaseRecorder.Code)
	assert.Equal(t, "FakeCoin", resJSON["name"])
	assert.Equal(t, "A fake meme coin", resJSON["description"])
	assert.Equal(t, int(resJSON["id"].(float64)), memeCoinId)
	assert.Greater(t, int(resJSON["popularity_score"].(float64)), 0)
	assert.Less(t, int(resJSON["popularity_score"].(float64)), 100)
}

func testDeleteMemeCoinEndpoint(t *testing.T) {
	// Setup path
	memeCoinId := 1
	deletePath := "/v1/meme-coin/" + strconv.Itoa(memeCoinId)

	// Case 1: "id" is not in the request body
	noIDInRequestCaseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/v1/meme-coin/", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(noIDInRequestCaseRecorder, req)

	responseStr := noIDInRequestCaseRecorder.Body.String()

	assert.Equal(t, http.StatusNotFound, noIDInRequestCaseRecorder.Code)
	assert.Equal(t, "404 page not found", responseStr)

	// Case 2: "id" is in the url but id is not numeric
	nonNumericIDCaseRecorder := httptest.NewRecorder()
	req, err = http.NewRequest("DELETE", "/v1/meme-coin/abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(nonNumericIDCaseRecorder, req)

	resJSONstr := nonNumericIDCaseRecorder.Body.String()
	resJSON := map[string]any{}
	json.Unmarshal([]byte(resJSONstr), &resJSON)
	assert.Equal(t, http.StatusBadRequest, nonNumericIDCaseRecorder.Code)
	assert.Equal(t, "Invalid MemeCoin ID", resJSON["message"])
	assert.Equal(t, "Wrong ID format", resJSON["error"])

	// Case 3: "id" is in the url
	idInRequestCaseRecorder := httptest.NewRecorder()
	req, err = http.NewRequest("DELETE", deletePath, nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(idInRequestCaseRecorder, req)

	resJSONstr = idInRequestCaseRecorder.Body.String()
	resJSON = map[string]any{}
	json.Unmarshal([]byte(resJSONstr), &resJSON)
	assert.Equal(t, http.StatusOK, idInRequestCaseRecorder.Code)
	assert.Equal(t, "FakeCoin", resJSON["name"])
	assert.Equal(t, "A fake meme coin", resJSON["description"])
	assert.Equal(t, int(resJSON["id"].(float64)), memeCoinId)
	assert.Greater(t, int(resJSON["popularity_score"].(float64)), 0)
	assert.Less(t, int(resJSON["popularity_score"].(float64)), 100)
}

func testPockMemeCoinEndpoint(t *testing.T) {
	// Setup path
	memeCoinId := 1
	pockPath := "/v1/meme-coin/" + strconv.Itoa(memeCoinId) + "/poke"

	// Case 1: "id" is not in the request body
	noIDInRequestCaseRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/v1/meme-coin//poke", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(noIDInRequestCaseRecorder, req)

	resJSONstr := noIDInRequestCaseRecorder.Body.String()
	resJSON := map[string]any{}
	json.Unmarshal([]byte(resJSONstr), &resJSON)
	assert.Equal(t, http.StatusBadRequest, noIDInRequestCaseRecorder.Code)
	assert.Equal(t, "Invalid MemeCoin ID", resJSON["message"])
	assert.Equal(t, "Wrong ID format", resJSON["error"])

	// Case 2: "id" is in the url but id is not numeric
	nonNumericIDCaseRecorder := httptest.NewRecorder()
	req, err = http.NewRequest("POST", "/v1/meme-coin/abc/poke", nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(nonNumericIDCaseRecorder, req)

	resJSONstr = nonNumericIDCaseRecorder.Body.String()
	resJSON = map[string]any{}
	json.Unmarshal([]byte(resJSONstr), &resJSON)
	assert.Equal(t, http.StatusBadRequest, nonNumericIDCaseRecorder.Code)
	assert.Equal(t, "Invalid MemeCoin ID", resJSON["message"])
	assert.Equal(t, "Wrong ID format", resJSON["error"])

	// Case 3: "id" is in the url
	idInRequestCaseRecorder := httptest.NewRecorder()
	req, err = http.NewRequest("POST", pockPath, nil)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(idInRequestCaseRecorder, req)

	resJSONstr = idInRequestCaseRecorder.Body.String()
	resJSON = map[string]any{}
	json.Unmarshal([]byte(resJSONstr), &resJSON)
	assert.Equal(t, http.StatusOK, idInRequestCaseRecorder.Code)
	assert.Equal(t, "FakeCoin", resJSON["name"])
	assert.Equal(t, "A fake meme coin", resJSON["description"])
	assert.Equal(t, int(resJSON["id"].(float64)), memeCoinId)
	assert.Greater(t, int(resJSON["popularity_score"].(float64)), 0)
	assert.Less(t, int(resJSON["popularity_score"].(float64)), 100)
}

func buildTestService() {
	// Mock repositories
	mockMemeCoinRepository := &mocks.MockMemeCoinRepository{}

	memeCoinService := services.NewMemeCoinService(mockMemeCoinRepository, nil)
	memeCoinHandler := handlers.NewMemeCoinHandler(memeCoinService)

	// Setup routes
	router = routes.NewRouter(memeCoinHandler)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
}
