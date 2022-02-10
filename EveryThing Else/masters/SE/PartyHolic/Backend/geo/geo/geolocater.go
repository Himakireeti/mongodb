package geo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Body struct {
	Items []struct {
		Title                  string `json:"title"`
		ID                     string `json:"id"`
		PoliticalView          string `json:"politicalView"`
		ResultType             string `json:"resultType"`
		HouseNumberType        string `json:"houseNumberType"`
		AddressBlockType       string `json:"addressBlockType"`
		LocalityType           string `json:"localityType"`
		AdministrativeAreaType string `json:"administrativeAreaType"`
		Address                struct {
			Label       string `json:"label"`
			CountryCode string `json:"countryCode"`
			CountryName string `json:"countryName"`
			StateCode   string `json:"stateCode"`
			State       string `json:"state"`
			CountyCode  string `json:"countyCode"`
			County      string `json:"county"`
			City        string `json:"city"`
			District    string `json:"district"`
			Subdistrict string `json:"subdistrict"`
			Street      string `json:"street"`
			Block       string `json:"block"`
			Subblock    string `json:"subblock"`
			PostalCode  string `json:"postalCode"`
			HouseNumber string `json:"houseNumber"`
			Building    string `json:"building"`
		} `json:"address"`
		Position struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"position"`
		Access []struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"access"`
		Distance float64 `json:"distance"`
		MapView  struct {
			West  float64 `json:"west"`
			South float64 `json:"south"`
			East  float64 `json:"east"`
			North float64 `json:"north"`
		} `json:"mapView"`
		Categories []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Primary bool   `json:"primary"`
		} `json:"categories"`
		FoodTypes []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Primary bool   `json:"primary"`
		} `json:"foodTypes"`
		HouseNumberFallback bool `json:"houseNumberFallback"`
		TimeZone            struct {
			Name      string `json:"name"`
			UtcOffset string `json:"utcOffset"`
		} `json:"timeZone"`
		Scoring struct {
			QueryScore float64 `json:"queryScore"`
			FieldScore struct {
				Country      float64   `json:"country"`
				CountryCode  float64   `json:"countryCode"`
				State        float64   `json:"state"`
				StateCode    float64   `json:"stateCode"`
				County       float64   `json:"county"`
				CountyCode   float64   `json:"countyCode"`
				City         float64   `json:"city"`
				District     float64   `json:"district"`
				Subdistrict  float64   `json:"subdistrict"`
				Streets      []float64 `json:"streets"`
				Block        float64   `json:"block"`
				Subblock     float64   `json:"subblock"`
				HouseNumber  float64   `json:"houseNumber"`
				PostalCode   float64   `json:"postalCode"`
				Building     float64   `json:"building"`
				Unit         float64   `json:"unit"`
				PlaceName    float64   `json:"placeName"`
				OntologyName float64   `json:"ontologyName"`
			} `json:"fieldScore"`
		} `json:"scoring"`
		Parsing struct {
			PlaceName []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"placeName"`
			Country []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"country"`
			State []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"state"`
			County []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"county"`
			City []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"city"`
			District []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"district"`
			Subdistrict []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"subdistrict"`
			Street []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"street"`
			Block []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"block"`
			Subblock []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"subblock"`
			HouseNumber []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"houseNumber"`
			PostalCode []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"postalCode"`
			Building []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"building"`
			SecondaryUnits []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"secondaryUnits"`
			OntologyName []struct {
				Start float64 `json:"start"`
				End   float64 `json:"end"`
				Value string  `json:"value"`
				Qq    string  `json:"qq"`
			} `json:"ontologyName"`
		} `json:"parsing"`
		StreetInfo []struct {
			BaseName           string `json:"baseName"`
			StreetType         string `json:"streetType"`
			StreetTypePrecedes bool   `json:"streetTypePrecedes"`
			StreetTypeAttached bool   `json:"streetTypeAttached"`
			Prefix             string `json:"prefix"`
			Suffix             string `json:"suffix"`
			Direction          string `json:"direction"`
			Language           string `json:"language"`
		} `json:"streetInfo"`
		CountryInfo struct {
			Alpha2 string `json:"alpha2"`
			Alpha3 string `json:"alpha3"`
		} `json:"countryInfo"`
	} `json:"items"`
}

func GeoAddress(address string) []float64 {
	endpoint, _ := url.Parse("https://geocode.search.hereapi.com/v1/geocode")
	queryParams := endpoint.Query()

	queryParams.Set("apiKey", "UukN87uD7WorbP8yiYnTUKN8gHsxtX_OWPxjlfW22Ns")

	queryParams.Set("q", address)

	endpoint.RawQuery = queryParams.Encode()
	response, err := http.Get(endpoint.String())
	op := make([]float64, 2)

	if err != nil {
		fmt.Printf("Error")
	} else {
		defer response.Body.Close()

		bodyBytes, _ := io.ReadAll(response.Body)

		var responseBody Body
		err := json.Unmarshal(bodyBytes, &responseBody)

		if err != nil {
			fmt.Println(err)
		} else {

			op[0] = responseBody.Items[0].Position.Lat
			op[1] = responseBody.Items[0].Position.Lng

		}

	}
	return op
}
