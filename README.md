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

- EMBED INTERFACE

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
