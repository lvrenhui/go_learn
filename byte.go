package main

import "fmt"

func main() {
	// time.AfterFunc(time.Second, test)
	// time.Sleep(time.Second * 2)
	println("u16:" + "\u4e16")
	println("u32:" + "\U00004e16")
	println("16zhi:" + "\xe4\xb8\x96")
	println("8zhi:" + "\344\270\226")

	word := "Hello,世界!"
	for i, s := range word {
		fmt.Printf("%d-%q-%b-%d-%x\n", i, s, s, s, s)
	}
	println(len("世"))
	for i := 0; i < len("世"); i++ {
		fmt.Printf("%d-%x-%o-%b\n", i, "世"[i], "世"[i], "世"[i])
	}
	bs := []byte("世")

	fmt.Printf("%b\n", bs)

	/*
		世
		10进制: 19990
		16进制: 4e16
		2进制：[100][111000]010110
		utf8: 111[00100] 10[111000] 10[010110]

		查看16进制：xxd w.txt
		查看2进制：xxd -b w.txt

		utf8是变长字节，111表示需要3个字节
		utf16是固定2个字节
		utf32是固定4个字节

	*/

}
