package main

import "fmt"

func main(){
    unsort :=[]int {9,12,3,6,23,34,45,12}
    var flag int
    alen := len(unsort)

    for i:=0;i<alen -1 ;i++ {
        flag = 0 
        for j:=0 ;j<alen - i -1 ;j++ {
            if unsort[j] > unsort[j+1] {
                unsort[j],unsort[j+1] = unsort[j+1] ,unsort[j]
                flag = 1
            }
        }
        if(flag == 0){
            break ;
        }
    }

    fmt.Println(unsort);
}
