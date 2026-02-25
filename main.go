package main

import (
  "fmt"
  "flag"
  "strings"
  "strconv"
  "errors"
  "os"
  "net/http"
  "time"
  )


func logo(){
  fmt.Println(`
     dMMMMMP     dMP     .aMMMb    .aMMMMP     dMMMMMP     dMMMMb 
    dMP         dMP     dMP"dMP   dMP"        dMP         dMP.dMP 
   dMMMP       dMP     dMMMMMP   dMP MMP"    dMMMP       dMMMMK"  
  dMP         dMP     dMP dMP   dMP.dMP     dMP         dMP"AMF   
 dMP         dMMMMMP dMP dMP    VMMMP"     dMMMMMP     dMP dMP    
                                                                  
  `)
  
}

//chek funtion lis of code
func chek(slice []int,code int) bool{
  for _,v:= range slice{
    if v == code{
      return true
    }
  }
  return false
}

func attack(ip *string,code []int,listFile *string){
  //open file
  file,err:= os.ReadFile(*listFile)
  if err != nil{
    fmt.Println("error not file Found")
    return
  }
  content:= string(file)
  lines:= strings.Split(content,"\n")
  timeEror:= 0 
  logo()
  for _,line:= range(lines){
    if line != "" && line != " " && line !="\n"{
      //requste network
      url:= *ip + "/" + strings.TrimSpace(line)
      time.Sleep(3 * time.Second)
      resp,err:= http.Get(url)
      if err != nil{
        fmt.Println("\033[0;32m[\033[0m-\033[0;32m]\033[0;35mERROR:\033[0m Some Error Detcatetd")
        fmt.Println("\033[2mEROR :", err)
        timeEror += 1
        if timeEror > 6{
          
          fmt.Println("\n\n\033[1;32m[\033[0m-\033[1;32m]\033[0;35mError:\033[0m Some Mistakes Detcatet.")
          break;
        }
        continue
      }
      timeEror = 0
      defer resp.Body.Close()
      state:= resp.StatusCode
      isIn:= chek(code,state)
      if isIn {
        fmt.Printf("\033[1;32m[\033[0m+\033[1;32m]\033[0;35mCODE:\033[0m %d \033[1;33m,\033[0;35mPATH:\033[0m %s\n",state,url)
      }
    }
  }
  fmt.Println("──────────────────────")
  fmt.Printf("\n\n\033[0;34mUse Word List:\033[0m %s\n\033[0;32mUse Address:\033[0m %s\n\n",*listFile,*ip)
  fmt.Println("──────────────────────")
}

func codesGenrate(code *string) ([]int,error){
  codes:= strings.Split(*code, ",")
  statusCode:= []int{}
  //cinvot string slice number slice
  for _,item := range(codes){
    number,err:= strconv.Atoi(item)
    if err != nil{
      return nil,errors.New("Not is Number list")
    }
    statusCode = append(statusCode, number)
  }
  //fmt.Println(statusCode)
  return statusCode,nil
}

func main() {
  // get cli agrment
  ip:= flag.String("ip","http://exaple.com","Ip/Url Of Scan Web")
  code:= flag.String("staus", "200,403,500","Catch Status Codes")
  listFile:= flag.String("file","exaple.txt", "to butrofose attack words list")
  flag.Parse()
  //genrate code list 
  codes,err:= codesGenrate(code)
  if err != nil{
    fmt.Println("\033[0;31mERROR: \033[0m",err)
    fmt.Println("\033[2mPlize numbers only add  to  status  Cli agreement ")
    return
  }
  fmt.Println("\033[0;32m[\033[0m*\033[0;32m]\033[0;35mFlager Is Starting...\033[0m\n")
  attack(ip,codes,listFile)
} 