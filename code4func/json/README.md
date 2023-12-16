# JSON encoding and decoding

how to convert from JSON raw data (strings or bytes) into Go types like structs, arrays, and slices, 
as well as unstructured data like maps and empty interfaces

### Unmarshaling (Parsing) Raw JSON Strings

We can convert JSON strings into bytes and unmarshal the data into a variables address:

```go
import "encoding/json"
//...

// ... 
myJsonString := `{"some":"json"}`

// `&myStoredVariable` is the address of the variable we want to store our
// parsed data in
json.Unmarshal([]byte(myJsonString), &myStoredVariable)
//...
```
There are two types of data you will encounter when working with JSON:

- Structured data
- Unstructured data
### Structured Data (Decoding JSON Into Structs)

```json
{
  "species": "pigeon",
  "description": "likes to perch on rocks"
}
```
create a struct that mirrors the data you want to parse. 
In our case, we will create a bird struct which has a Species and Description attribute:
```go
type Bird struct {
  Species string
  Description string
}
```
And `unmarshal` it as follows:

```go
birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
var bird Bird	
json.Unmarshal([]byte(birdJson), &bird)
fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
//Species: pigeon, Description: likes to perch on rocks
```

Playground:
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string
	Description string
}

func main() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
	//Species: pigeon, Description: likes to perch on rocks
}
```
Output: 
```go
Species: pigeon, Description: likes to perch on rocks
```
### JSON Arrays
Let’s look at how we can decode an array of objects, like below:
```json
[
  {
    "species": "pigeon",
    "description": "likes to perch on rocks"
  },
  {
    "species":"eagle",
    "description":"bird of prey"
  }
]
```

```go
birdJson := `[{"species":"pigeon","description":"likes to perch on rocks"},{"spices":"eagle","description":"bird of prey"}]`
var birds []Bird
json.Unmarshal([]byte(birdJson), &birds)
fmt.Printf("Birds : %+v", birds)
```

Playground:
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string
	Description string
}

func main() {
	birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf("Birds : %+v", birds)
	//Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]
}
```
Output:
```go
Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]
```

### Nested Objects
Now, consider the case when you have a property called Dimensions, that measures the Height and Length of the bird in question:

```json
{
  "species": "pigeon",
  "description": "likes to perch on rocks"
  "dimensions": {
    "height": 24,
    "width": 10
  }
}
```
we need to mirror the structure of the object in our Go code. To add a nested dimensions object, lets create a dimensions struct :
```go
type Dimensions struct{
	Height int
	Width int
}
```
Now, the `Bird` struct will include a `Dimensions` field:
```go
type Bird struct{
	Spicies string
	Description string
	Dimensions Dimensions
}
```
We can unmarshal this data using the same method as before:

```go
birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
var birds Bird
json.Unmarshal([]byte(birdJson), &birds)
fmt.Printf(bird)
// {pigeon likes to perch on rocks {24 10}}
```
Playground:
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

func main() {
	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)
	// {pigeon likes to perch on rocks {24 10}}
}
```
Output:
```go
{pigeon likes to perch on rocks {24 10}}
```
### Primitive Types
We mostly deal with complex objects or arrays when working with JSON, but data like 3, 3.1412 and "birds" are also valid JSON strings.

```go
numberJson := "3"
floatJson := "3.1412"
stringJson := `"bird"`

var n int
var pi float64
var str string

json.Unmarshal([]byte(numberJson), &n)
fmt.Println(n)
// 3

json.Unmarshal([]byte(floatJson), &pi)
fmt.Println(pi)
// 3.1412

json.Unmarshal([]byte(stringJson), &str)
fmt.Println(str)
// bird
```
### Time Values

### JSON Struct Tags - Custom Field Names
Although sometimes, we want a different attribute name than the one provided in your JSON data. For example, consider the below data:

```json
{
  "birdType": "pigeon",
  "what it does": "likes to perch on rocks"
}
```
To solve this, we can use struct field tags:
```go
type Bird struct{
	Species string `json:"birdType"`
	Description string `json:"what it does"`
}
```
Now, we can explicitly tell our code which JSON property to map to which attribute.

```go
birdJson := `{"birdType": "pigeon","what it does": "likes to prech on rocks"}`
var bird Bird
json.Unmarshal([]byte(birdJson), &bird)
fmt.Println(bird)
```
Playground:
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
  Species string `json:"birdType"`
  Description string `json:"what it does"`
}

func main() {
	birdJson := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)
	// {pigeon likes to perch on rocks}
}
```
### Decoding JSON to Maps - Unstructured Data

If you don’t know the structure of your JSON properties beforehand, you cannot use structs to unmarshal your data.
Instead you can use maps. Consider some JSON of the form:
```json
{
  "birds": {
    "pigeon":"likes to perch on rocks",
    "eagle":"bird of prey"
  },
  "animals": "none"
}

```
There is no struct we can build to represent the above data for all cases since the keys corresponding to the birds can change, which will change the structure.
To deal with this case we create a map of strings to empty interfaces:
```go
birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
var result map[string]any
json.Unmarshal([]byte(birdJson), &result)

// The object stored in the "birds" key is also stored as 
// a map[string]any type, and its type is asserted from
// the `any` type
birds := result["birds"].(map[string]any)

for key, value := range birds {
  // Each value is an `any` type, that is type asserted as a string
  fmt.Println(key, value.(string))
}

```

### Validating JSON Data

```json
{
  "birds": {
    "pigeon":"likes to perch on rocks",
    "eagle":"bird of prey"
```
If we try to unmarshal this, our code will panic:

```go
birdJson := `{"birds": {"pigeon":"likes to perch on rocks","eagle":"bird of prey"`
var result map[string]any
json.Unmarshal([]byte(birdJson), &result)
```
Playground:

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func main() {
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"`
	var result map[string]any
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]any type, and its type is asserted from
	// the `any` type
	birds := result["birds"].(map[string]any)

	for key, value := range birds {
		// Each value is an `any` type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}
	//pigeon likes to perch on rocks
	//eagle bird of prey
}
```


### Ref:
1. https://www.sohamkamani.com/golang/json/#best-practices-structs-vs-maps
2. 





