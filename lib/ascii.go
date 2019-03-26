package lib

import (
	"fmt"
)

// PrintWindows ..
func PrintWindows(out *Output) {

	str := fmt.Sprintf(
		`W     W III N   N DDD   OOO  W     W  SSS %20s%s@%s
W     W  I  NN  N D  D O   O W     W S    %20s%.2fG 
W  W  W  I  N N N D  D O   O W  W  W  SSS  %20s%.2fG 
 W W W   I  N  NN D  D O   O  W W W      S %17s%s
  W W   III N   N DDD   OOO    W W   SSSS  
`, "Hostname: ", out.Username, out.Hostname, "Free Mem: ", out.Freeram, "Total Mem: ", out.Totalram, "Uptime: ", out.Uptime)

	fmt.Println(str)
}

// PrintLinux ..
func PrintLinux(out *Output) {

	fmt.Printf(
		`L    III N   N U   U X   X %20s%s@%s
L     I  NN  N U   U  X X  %20s%.2fG
L     I  N N N U   U   X   %21s%.2fG
L     I  N  NN U   U  X X  %18s%s
LLLL III N   N  UUU  X   X
`, "Hostname: ", out.Username, out.Hostname, "Free Mem: ", out.Freeram, "Total Mem: ", out.Totalram, "Uptime: ", out.Uptime)
}

// PrintOSX ..
func PrintOSX(out *Output) {

	fmt.Printf(
		` OOO   SSS  X   X %20s%s@%s
O   O S      X X  %20s%.2fG
O   O  SSS    X   %21s%.2fG
O   O     S  X X  %18s%s
 OOO  SSSS  X   X 
`, "Hostname: ", out.Username, out.Hostname, "Free Mem: ", out.Freeram, "Total Mem: ", out.Totalram, "Uptime: ", out.Uptime)
}

func Print(out *Output) {
	fmt.Printf("%s%s@%s\n%s%.2f\n%s%.2f,%s%s", "Hostname: ", out.Username, out.Hostname, "Free Mem: ", out.Freeram, "Total Mem: ", out.Totalram, "Uptime: ", out.Uptime)
}
