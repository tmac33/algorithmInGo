package main

import (
	"fmt"
	"strings"
	"structnew/pubsub"
	"time"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello,world!")
	p.Publish("hello,golang!")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang", msg)
		}
	}()

	//运行一定时间后退出
	time.Sleep(3 * time.Second)
}

// //生产者：生成factor整数倍的序列
// func Producer(factor int, out chan<- int) {
// 	for i := 0; ; i++ {
// 		out <- i * factor
// 	}
// }

// //消费者
// func Consumer(in <-chan int) {
// 	for v := range in {
// 		fmt.Println(v)
// 	}
// }

// func main() {
// 	ch:=make(channle int, 64)

// 	go Producer(3, ch) //生成3的倍数序列
// 	go Producer(5, ch)

// 	go Consumer(ch)

// 	//ctrl+c 退出
// 	sig:=make(chan os.Signal, 1)
// 	signal.Notify(sig, syscall.SIGINT,syscall.SIGTERM)
// 	fmt.Printf("quit (%v)\n", <-sig)
// }

// func main() {
// 	col1 := "red"
// 	col2 := col1
// 	fmt.Println(col1)
// 	fmt.Println(col2)
// 	col1 = "blue"
// 	fmt.Println(col1)
// 	fmt.Println(col2)
// }

// import (

// 	"fmt"
// )
// type Student struct{
// 	id      int
// 	name    string
// 	address string
// 	age     int
// }

// func structTest01()  {
// 	//使用new创建一个Student对象，结果为指针
// 	var s *Student=new(Student)
// 	s.id = 101
// 	s.name = "Mikle"
// 	s.address = "红旗南路"
// 	s.age = 18

// 	fmt.Printf("id:%d\n", s.id)
// 	fmt.Printf("name:%s\n", s.name)
// 	fmt.Printf("address:%s\n", s.address)
// 	fmt.Printf("age:%d\n", s.age)
// 	fmt.Println(s)
// }

// func structTest02(){
// 	var s *Student=&Student{102,"john","nanjing",19}
// }

// func main(){
// 	map1:=make(map[int]int)
// 	for i := 0; i < 3; i++ {
// 		map1[i]=i+1
// 	}

// 	fmt.Println(map1)
// 	ch:=make(chan int)

// 	for k,v:=range map1{

// 		go func (key,value int)  {

// 			fmt.Println("one way:",key,":",value)
// 			ch<-
// 		}(k,v)
// 		go func() {
// 			time.Sleep(time.Second)
// 			fmt.Println("anpther way:",k,":",v)
// 			ch<-
// 		}()
// 	}
// }

// func main(){

// 	strs:=[]string{"flower","flow","flight","nvrnmvlr"}
// 	fmt.Println(cap(strs))

// 	ans:=real(strs)
// 	fmt.Println(ans)

// }

// func real(strs []string) string{
// short:=ShortStrings(str)
// var dhg string
// for i,r:=range short{
// 	for j := 0; j < len(str); j++ {
// 		fmt.Println(str[j][i])
// 		fmt.Println(r)
// 		if str[j][i]!=byte(r){

// 			dhg=str[j][:i]
// 		}
// 	}
// }
// return dhg

// 	short := shortest(strs)
// 	var dhg string

// 	for i, r := range short {
// 		for j := 0; j < len(strs); j++ {
// 			if strs[j][i] != byte(r) {
// 				dhg= strs[j][:i]
// 			}
// 		}
// 	}
// 	return dhg
// }

//找到最短字符
// func ShortStrings(str []string) string{
// 	if len(str)==0{
// 		return ""
// 	}
// 	res:=str[0]
// 	for _,df:=range str{
// 		if len(res)>len(df){
// 			res=df
// 		}
// 	}
// 	return res
// }

// func shortest(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}

// 	res := strs[0]
// 	for _, s := range strs {
// 		if len(res) > len(s) {
// 			res = s
// 		}
// 	}

// 	return res
// }

// type Fetcher interface{
// 	Fetch(url string)（body string,urls []string,err error)
// }

// func main(){
// 	Crawl("https://golang.org",4,fetcher)
// }

// func Crawl(url string,depth int,fetcher Fetcher)  {
// 	if depth<0 {
// 		return
// 	}
// 	body,urls,err:=fetcher.Fetch(url)
// 	if err!=nil{
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("found: %s %q\n",url,body)
// 	for _,u:=range urls{
// 		Crawl(u,depth-1,fetcher)
// 	}
// 	return
// }

// func Fetcher()  {

// }

// type Fetcher interface {
// 	// Fetch returns the body of URL and
// 	// a slice of URLs found on that page.
// 	Fetch(url string) (body string, urls []string, err error)
// }

// // Crawl uses fetcher to recursively crawl
// // pages starting with url, to a maximum of depth.
// func Crawl(url string, depth int, fetcher Fetcher) {
// 	// TODO: Fetch URLs in parallel.
// 	// TODO: Don't fetch the same URL twice.
// 	// This implementation doesn't do either:
// 	if depth <= 0 {
// 		return
// 	}
// 	body, urls, err := fetcher.Fetch(url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("found: %s %q\n", url, body)
// 	for _, u := range urls {
// 		Crawl(u, depth-1, fetcher)
// 	}
// 	return
// }

// func main() {
// 	Crawl("https://golang.org/", 4, fetcher)
// }

// // fakeFetcher is Fetcher that returns canned results.
// type fakeFetcher map[string]*fakeResult

// type fakeResult struct {
// 	body string
// 	urls []string
// }

// func (f fakeFetcher) Fetch(url string) (string, []string, error) {
// 	if res, ok := f[url]; ok {
// 		return res.body, res.urls, nil
// 	}
// 	return "", nil, fmt.Errorf("not found: %s", url)
// }

// // fetcher is a populated fakeFetcher.
// var fetcher = fakeFetcher{
// 	"https://golang.org/": &fakeResult{
// 		"The Go Programming Language",
// 		[]string{
// 			"https://golang.org/pkg/",
// 			"https://golang.org/cmd/",
// 		},
// 	},
// 	"https://golang.org/pkg/": &fakeResult{
// 		"Packages",
// 		[]string{
// 			"https://golang.org/",
// 			"https://golang.org/cmd/",
// 			"https://golang.org/pkg/fmt/",
// 			"https://golang.org/pkg/os/",
// 		},
// 	},
// 	"https://golang.org/pkg/fmt/": &fakeResult{
// 		"Package fmt",
// 		[]string{
// 			"https://golang.org/",
// 			"https://golang.org/pkg/",
// 		},
// 	},
// 	"https://golang.org/pkg/os/": &fakeResult{
// 		"Package os",
// 		[]string{
// 			"https://golang.org/",
// 			"https://golang.org/pkg/",
// 		},
// 	},
// }

// func main()  {
// 	var nums1=[]int{0,1,2,2,3,0,4,2}
// 	var val=2
// 	fmt.Println(removeElement(nums1,val))
// }

// func removeElement(nums []int, val int) int {
// 	newNum:=make([]int,0)
// 	for k,v:=range nums{
// 		if v==val {
// 			newNum=append(nums[:k],nums[(k+1):]...)
// 			continue
// 		}
// 	}
// 	// for i := 0; i < len(nums); i++ {
// 	// 	if nums[i]==val{
// 	// 		nums=append(nums[:i],nums[(i+1):]...)
// 	// 		continue
// 	// 	}
// 	// }
// 	fmt.Println(nums)
// 	return len(nums)
// }

// func main(){
// 	a:=4
// 	b:=5
// 	fmt.Println(abc(a,b))
// }

// type IntHeap []int

// func (h IntHeap) Len() int{
// 	return len(h)
// }

// func main(){
// 	ring:=ring.New(3)
// 	for i:=1;i<=3;i++{
// 		ring.Value=i
// 		ring=ring.Next()
// 	}
// 	s:=0
// 	ring.Do(func(p interface{}){

// 	})
// }

// import "sync"

// var(
// 	cache=make(map[string]*CacheTable)
// 	mutex sync.RWMutex
// )

// func Cache(table string) *CacheTable {

// }

// import "fmt"

// var sli=[]int{3,2,2,3}

// func main(){
// 	fmt.Println(removeElement(sli,3))
// }

// func removeElement(nums []int, val int) int {
// 	i:=0
// 	var newnum []int
// 	for _,value:=range nums{
// 		if value!=val{
// 			i++
// 			newnum=append(newnum,value)
// 		}
// 	}
// 	fmt.Println(newnum)
// 	return i
// }

// func main() {

// J:
//     for j := 0; j < 5; j++ {
//         for i := 0; i < 10; i++ {
//             if i > 6 {
//                 break J //现在终止的是j 循环，而不是i的那个
//             }
//             fmt.Println(i)
//         }
//     }
// }

// import "fmt"

// func main(){
// 	nums:=[]int{1,3,5,6}
// 	sz:=searchInsert(nums,1)
// 	fmt.Println(sz)
// }

// func searchInsert(nums []int, target int) int {
// 	var f int
//     for k,v:=range nums{
// 		if v<target{
// 			continue
// 		}
// 		if target==v {
// 			f=k
// 		}else if v>target{
// 			f=k
// 		}

// 	}
// 	return f
// }
