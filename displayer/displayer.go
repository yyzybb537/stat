package displayer

import (
	"os"
	"fmt"
	"strings"
)

var displaying int

func Printf(format string, a ... interface{}) {
	fmt.Printf(strings.Repeat("\b", displaying) + strings.Repeat(" ", displaying))
	os.Stdout.Sync()
	info := fmt.Sprintf(format, a...) + " "
	fmt.Printf(strings.Repeat("\b", displaying) + info)
	os.Stdout.Sync()
	displaying = len(info)
}

var suffixs = []string{"", "K", "M", "G", "T", "P", "E"}                           
func Bytes2String(bytes int64) string {
    var i int                                                                      
    var dot int                                                                    
    for i < len(suffixs) {                                                         
        if bytes < 1024 {                                                           
            break                                                                  
        }                                                                          
                                                                                   
        //dot = bytes / 10 % 100                                                    
        dot = int((float64(bytes)/1024.0 - float64(bytes/1024)) * 100)               
        bytes /= 1024                                                               
        i++                                                                        
    }                                                                              
                                                                                   
    if dot > 0 {                                                                   
        return fmt.Sprintf("%d.%02d%s", bytes, dot, suffixs[i])                     
    }                                                                              
                                                                                   
    return fmt.Sprintf("%d%s", bytes, suffixs[i])                                   
}
