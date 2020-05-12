package main

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func GetAllNews() (ReturnType, error) {
	var data ReturnType
	var news []News 
	resp, err := http.Get("http://a6c0e756.ngrok.io/akmola")
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		return data, err1
	}
	json.Unmarshal([]byte(body), &news)

	data.NewsList = news

	///News enddd
	
	////Data about world cases
	var worldData Data
	r, err2 := http.Get("https://corona.lmao.ninja/all")
	if err2 != nil {
		return data, err2
	}
	defer r.Body.Close()
	worldbody, err3 := ioutil.ReadAll(r.Body)
	if err3 != nil {
		return data, err3
	}
	json.Unmarshal([]byte(worldbody), &worldData)

	data.WorldData = worldData

	////
	return data, nil
}