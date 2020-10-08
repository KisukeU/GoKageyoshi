package main

import ("fmt"
		"flag"
		"os"
		"os/exec"
		"bufio"
		"strings"
	)

func main() {
	fmt.Println("~~~~~~~~~~~~~~STARTING~~~~~~~~~~~~~~") 
	fmt.Println("Tool to Check AXFR vulnerability") 
	fmt.Println("**************************************") 
	fmt.Println("|Flags:") 
	fmt.Println("|-iL -- file that contains subdomains (each subdomain per line)") 
	fmt.Println("|-ns -- nameservers to request (nslookup -type=ns domain.com)") 
	fmt.Println("|Usage example: tryMenu -iL fileWithSubdomains -ns nameServer1,nameServer2") 
	fmt.Println("**************************************") 

	var fileNameOptionValue, nameServersOptionValue string
	flag.StringVar(&fileNameOptionValue, "iL","","Usage")
	flag.StringVar(&nameServersOptionValue, "ns","","Usage")
	flag.Parse()

	var nameServers = strings.Split(nameServersOptionValue,",")

	file, err := os.Open(fileNameOptionValue)
	if err != nil {
   		 panic(err)
		}
	defer file.Close()
	
	
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		//fmt.Println("check subdomain",scanner.Text)
		var currentSubDomain string
		currentSubDomain = scanner.Text()
		for i := 0; i < len(nameServers); i++ {
			fmt.Println("________________________________________________")
			fmt.Println("[.]try: dig +multi AXFR@",nameServers[i]+" "+currentSubDomain)
			argumentsToDigArray := strings.Split("+multi AXFR @"+nameServers[i]+" "+currentSubDomain, " ")
			argumentsToDig := argumentsToDigArray[0:len(argumentsToDigArray)]
			cmd := exec.Command("dig",argumentsToDig...)
    		cmd.Stderr = os.Stderr
    		cmd.Stdout = os.Stdout
    		cmd.Stdin = os.Stdin
    		if err := cmd.Run(); err != nil {
        		fmt.Println(err.Error())
        		os.Exit(1)
    		}
		}        
    }
	fmt.Println("~~~~~~~~~~~~~~Thanks!~~~~~~~~~~~~~~") 
}
