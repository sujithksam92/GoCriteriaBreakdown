package main

import (
	"fmt"
	"math"
)

func participantBreakup(totalCount int) [][]int{
	//No. of participants that can be passed in one request is limited to 50
	//And then computing the number of breakups of all participants possible
	limit:=50
	sizeOfParticipants:=int(math.Ceil(float64(totalCount)/float64(limit)))
	
	//2d Slice initialization to hold the ranges of participant breakups
	//Ex: for 80 participants the output would be [[1 50] [51 80]]
	participants := make([][]int, sizeOfParticipants)
	for item := range participants {
 	   participants[item] = make([]int, 2)
	}


	i,a,index:=50,1,0	
	for {
		//tempArray here would hold the array of upper and lower bounds that would be inserted into the main 'participants' array
		tempArray := make([]int, 2)

		//the If part here would handle the last iteration in the loop where we would have reached the last round and can break out
		//the Else would handle all the other iterations and add upper and lower bounds to the main 'participants' array
		if i>=totalCount{
			r:=i-totalCount
			tempArray[0]=a
			tempArray[1]=i-r
			participants[index]=tempArray
			break
		} else{
			tempArray[0]=a
			tempArray[1]=i				
			a=i+1
			i=i+limit
			participants[index]=tempArray
			index=index+1
		}
	}
	
	return participants
}

func criteriaBreakup(totalCount int) [][]int{
	//No. of criteria that can be passed in one request is limited to 25
	//And then computing the number of breakups of all criteria
	limit:=25
	sizeOfCriteria:=int(math.Ceil(float64(totalCount)/float64(limit)))
	
	//2d Slice initialization to hold the ranges of criteria breakups
	//Ex: for 45 criteria the output would be [[1 25] [26 45]]
	criteria := make([][]int, sizeOfCriteria)
	for item := range criteria {
 	   criteria[item] = make([]int, 2)
	}


	i,a,index:=25,1,0	
	for {
		//tempArray here would hold the array of upper and lower bounds that would be inserted into the main 'criteria' array
		tempArray := make([]int, 2)

		//the If part here would handle the last iteration in the loop where we would have reached the last round and can break out
		//the Else would handle all the other iterations and add upper and lower bounds to the main 'criteria' array
		if i>=totalCount{
			r:=i-totalCount
			tempArray[0]=a
			tempArray[1]=i-r
			criteria[index]=tempArray
			break
		} else{
			tempArray[0]=a
			tempArray[1]=i				
			a=i+1
			i=i+limit
			criteria[index]=tempArray
			index=index+1
		}
	}
	
	return criteria
}


func main() {
	result1 := participantBreakup(80)
	//fmt.Println("Participants outside - ",result1)
	
	result2 := criteriaBreakup(45)
	//fmt.Println("Criteria outside - ",result2)
	
	for _,k := range(result1){
		for _,l := range(result2){
			fmt.Println(k,"*",l)
		}
	}
	
}
