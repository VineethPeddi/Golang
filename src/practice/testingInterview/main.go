package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)


type Employee struct {
	Name string `json:"employee_name"`
	Salary int	`json:"employee_salary"`
}

type APIResponse struct{
	Data []Employee `json:"data"`
}

func main() {
	url := "https://dummy.restapiexample.com/api/v1/employees"
	responseData,err:=getResponseFromAPI(url)
	if err!=nil{
		fmt.Println("Error in fetching response from API",err)
		os.Exit(1)
	}
	sort.Slice(responseData.Data, func(i, j int) bool {
		salary1:=responseData.Data[i].Salary
		salary2:=responseData.Data[j].Salary
		if salary1==salary2{
			return responseData.Data[i].Name>responseData.Data[j].Name
		}
		return salary1>salary2
	})

	fmt.Printf("sorted data: %+v \n",responseData)

}

func getResponseFromAPI(url string) (*APIResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	responseData:=new(APIResponse)
	if err=json.Unmarshal(dataBytes,responseData);err!=nil{
		return nil,err
	}
	return responseData,nil
}

func getSumOfAllElementsExceptCurrentElement(a []int) []int {
	b := make([]int, 0)
	sum := 0
	for _, val := range a {
		sum += val
	}
	for _, val := range a {
		b = append(b, sum-val)
	}
	return b
}
