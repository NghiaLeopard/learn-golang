# learn-golang

break, continue is mean in for , while ,switch

--

[]byte is not press type float, int ,...

Because:
b := []byte{72, 101, 108, 108, 111} // Các byte tương ứng với các chữ cái "Hello"
s := string(b) // s sẽ là chuỗi "Hello"

then string is , do it like:

result = string(data []byte)

use library strconv => strconv.ParseFloat(result)

-- module,package
convention: 1 directory is 1 package

Files have same package used together in directory

-- POINTER

var ptn \*int : declare pointer to an integer

user :=23

ptn = &user : ptn is pointer to address variable user

\*ptn : is value of ptn pointer to address

return 1 value , it also copy value ,return that value copy

if you use pointer then it do not copy value

+| Go là ngôn ngữ "pass-by-value", nghĩa là khi bạn truyền một đối số cho hàm hoặc trả về một giá trị từ hàm, Go sẽ tạo một bản sao của giá trị đó và sử dụng bản sao này trong quá trình xử lý.

Truyền đối số (arguments): Khi bạn truyền một giá trị vào hàm, một bản sao của giá trị đó được tạo và sử dụng bên trong hàm. Bất kỳ thay đổi nào bạn thực hiện đối với giá trị này bên trong hàm sẽ không ảnh hưởng đến giá trị gốc bên ngoài hàm.

Trả về giá trị (return values): Tương tự, khi một hàm trả về một giá trị, một bản sao của giá trị đó được tạo và trả về cho nơi gọi hàm. Bất kỳ thay đổi nào bạn thực hiện đối với giá trị trả về này sẽ không ảnh hưởng đến giá trị gốc bên trong hàm.

+| Receiver đóng vai trò như một tham số đặc biệt của method, cho phép method truy cập và thao tác trên các trường (fields) hoặc giá trị của kiểu dữ liệu mà nó được gắn vào.

+| import "fmt"

func square(x int) int {
x \*= x // Bình phương giá trị x bên trong hàm
return x
}

func main() {
num := 5
result := square(num)
fmt.Println("Giá trị gốc:", num) // In ra 5 (giá trị gốc không thay đổi)
fmt.Println("Kết quả bình phương:", result) // In ra 25
}

return x is value copy . result receive value x . I don't change value result then x also change , should return value must copy to affect value in func

-- STRUCT

var appUser *user.User
var appUser1 *user.User

    appUser = &user.User{
    	FirstName: "Dai Nghia",
    }

    appUser1 = &user.User{
    	FirstName: "Dai Nghia 1",
    }

    appUser.FirstName = "Dai Beo"
    appUser1.FirstName = "Dai Beo 1"

    fmt.Print(appUser)
    fmt.Print(appUser1)

+| struct is blueprint

init appUser = &user.User{} : create 1 pointer to add struct

struct is type value , not reference
then when use it , func or assign 1 variable difference will copy value
if adjust value in struct then must use pointer

depend case you use .

Add method struct: func (u User)

func (u \*User): you adjust value of pointer to that address struct

+| Refer: should use pointer struct to avoid memory of value copy
// func printData(u \*user) {
// fmt.Println(u.firstName,u.lastName,u.birthDate)
// }

-BUFIO

bufio.NewReader:

Tạo ra một bộ đệm để lưu trữ dữ liệu đọc từ nguồn.
Giúp cải thiện hiệu suất đọc dữ liệu.
ReadString('\n'):

Đọc dữ liệu từ bộ đệm cho đến khi gặp ký tự xuống dòng (\n).
Quan trọng: Nó KHÔNG bỏ đi phần dữ liệu còn lại. Phần dữ liệu chưa đọc (sau ký tự \n) vẫn còn trong bộ đệm và có thể được đọc tiếp bằng các lệnh đọc khác.

type Note struct {
Title string `json:"title"`
Content string `json:"content"`
CreatedAt time.Time `json:"created_at"`
}

like talked : want use data difference package

func (note Note) Save() error {
fileName := strings.ReplaceAll(note.Title," ","\_")
fileName = strings.ToLower(fileName) + ".json"

    json,err := json.Marshal(note)

    if err != nil {
    	return err
    }

    return os.WriteFile(fileName,json,0644)

}

tag struct : override name in json with tag struct ``

Convert json data: json.Marshal

-- INTERFACE

Interface like a contract

type saver interface {
Save() error
}

if interface is just have one method then name interface combine to name method ,add er to last name.

interface is required input and not important input
It is care about that data have user method in that interface

-   EMBED INTERFACE

\*type saver interface {
Save() error
}

\*type outputtable interface {
saver
Print()
}

interface{} = any

-SWITCH TYPE

when combine use any type

func printSomething(value interface{}) {
switch value.(type) {
case int:
fmt.Printf("Integer: %d\n",value)
case float64:
fmt.Printf("Integer: %f\n",value)
case string:
fmt.Printf("string: %s\n",value)
}
}

-CHECK TYPE VALUE
inputVal,ok := value.(int)

--\*Generics

func add[T int | float64 | string](a, b T) T {
return a + b
}

--\*Array

Declare
arr := [4]float64{10.10,10.2,12.2,14.1}

--\* Slice arr:

Trong Go, slice được xây dựng dựa trên mảng. Hãy tưởng tượng slice như một cửa sổ nhìn vào một phần của mảng gốc. Cửa sổ này có thể được điều chỉnh (cắt)

slice like 1 pointer

Can change value in main array

Slice don't copy array

Thay đổi kích thước: Slice có thể tự động tăng kích thước khi cần. Khi bạn thêm phần tử vượt quá chiều dài hiện tại nhưng vẫn nằm trong dung lượng hiện tại, slice sẽ sử dụng không gian trống có sẵn trong mảng gốc. Tuy nhiên, nếu bạn thêm phần tử vượt quá dung lượng, Go sẽ tự động tạo một mảng mới lớn hơn, sao chép các phần tử hiện có sang mảng mới và cập nhật slice để trỏ đến mảng mới này.

sliceArr := arr[1:3]

start with index 1 and just get value of index: 1,2
Then exclude index 3

sliceArr := arr[:3]
sliceArr := arr[1:]

\*restSlice := []string{"learn go","learn backend"}

result := appand(sliceArr,restSlice...)

--\* MAP

<!-- Declare websites : create map -->

website := map[string]string{"Google": "https://google.com", "aws": "https://aws.com"}

    fmt.Println(website)


    fmt.Println(website["Google"])

<!-- add Map -->

website["linkedIn"] = "https://linkedIn.com"

fmt.Println(website)

<!-- Delete Map -->

delete(website,"Google")

fmt.Println(website)

\* Special make function

<!-- Slice -->

mySlice := make([]string, 5) // Tạo slice có độ dài 5
fmt.Println(mySlice) // ["" "" "" "" ""]
fmt.Println(len(mySlice)) // 5
fmt.Println(cap(mySlice)) // 5

mySlice := make([]string, 3, 5) // Tạo slice có độ dài 3, sức chứa 5
fmt.Println(mySlice) // ["" "" ""]
fmt.Println(len(mySlice)) // 3
fmt.Println(cap(mySlice)) // 5

<!-- Map -->

mapUrl := make(map[string]string,5)

cap(mapUrl) = 5

Make used to don't pre-allocate

--\*Type alias
Trong Go (Golang), type alias là một cách để tạo ra một tên mới (biệt danh) cho một kiểu dữ liệu đã tồn tại. Điều này không tạo ra một kiểu mới hoàn toàn, mà chỉ đơn giản là cung cấp một cách khác để tham chiếu đến kiểu gốc.

type floatMap map[string]float64

type alias: floatMap -> reference to address original Map

\*loop array,slice,map

With Map : key,value
With Slice: index,value

for key,value := range website {
fmt.Println("Key: ",key)
fmt.Println("Value: ",value)
}
