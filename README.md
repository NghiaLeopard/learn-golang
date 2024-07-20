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

-- pointer

var ptn *int : declare pointer to an integer

user :=23

ptn = &user : ptn is pointer to address variable user

*ptn : is value of ptn pointer to address