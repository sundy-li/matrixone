
// Code generated by tool; DO NOT EDIT.
package config

import (
	"fmt"
	"sync"
	"github.com/BurntSushi/toml"
)

//all parameters in the system
type AllParameters struct{
	//read and write lock
	rwlock	sync.RWMutex

	/**
	Name:	boolSet1
	Scope:	[global]
	Access:	[file]
	DataType:	bool
	DomainType:	set
	Values:	[true]
	Comment:	boolSet1
	UpdateMode:	dynamic
	*/
	boolSet1    bool

	/**
	Name:	boolSet2
	Scope:	[global]
	Access:	[file]
	DataType:	bool
	DomainType:	set
	Values:	[false]
	Comment:	boolSet2
	UpdateMode:	hotload
	*/
	boolSet2    bool

	/**
	Name:	boolSet3
	Scope:	[global]
	Access:	[file]
	DataType:	bool
	DomainType:	set
	Values:	[]
	Comment:	boolSet3
	UpdateMode:	dynamic
	*/
	boolSet3    bool

	/**
	Name:	stringSet1
	Scope:	[global]
	Access:	[file]
	DataType:	string
	DomainType:	set
	Values:	[ss1 ss2 ss3]
	Comment:	stringSet1
	UpdateMode:	dynamic
	*/
	stringSet1    string

	/**
	Name:	stringSet2
	Scope:	[global]
	Access:	[file]
	DataType:	string
	DomainType:	set
	Values:	[]
	Comment:	stringSet2
	UpdateMode:	dynamic
	*/
	stringSet2    string

	/**
	Name:	int64set1
	Scope:	[global]
	Access:	[file]
	DataType:	int64
	DomainType:	set
	Values:	[1 2 3 4 5 6]
	Comment:	int64Set1
	UpdateMode:	dynamic
	*/
	int64set1    int64

	/**
	Name:	int64set2
	Scope:	[global]
	Access:	[file]
	DataType:	int64
	DomainType:	set
	Values:	[1 3 5 7]
	Comment:	int64Set2
	UpdateMode:	fix
	*/
	int64set2    int64

	/**
	Name:	int64set3
	Scope:	[global]
	Access:	[file]
	DataType:	int64
	DomainType:	set
	Values:	[]
	Comment:	int64Set3
	UpdateMode:	dynamic
	*/
	int64set3    int64

	/**
	Name:	int64Range1
	Scope:	[global]
	Access:	[file]
	DataType:	int64
	DomainType:	range
	Values:	[1000 0 10000]
	Comment:	int64Range1
	UpdateMode:	dynamic
	*/
	int64Range1    int64

	/**
	Name:	float64set1
	Scope:	[global]
	Access:	[file]
	DataType:	float64
	DomainType:	set
	Values:	[1.0 2.0 3.0 4.00 5.0 6.0]
	Comment:	float64Set1
	UpdateMode:	dynamic
	*/
	float64set1    float64

	/**
	Name:	float64set2
	Scope:	[global]
	Access:	[file]
	DataType:	float64
	DomainType:	set
	Values:	[1.001 3.003 5.005 7.007]
	Comment:	float64Set2
	UpdateMode:	fix
	*/
	float64set2    float64

	/**
	Name:	float64set3
	Scope:	[global]
	Access:	[file]
	DataType:	float64
	DomainType:	set
	Values:	[]
	Comment:	float64Set3
	UpdateMode:	dynamic
	*/
	float64set3    float64

	/**
	Name:	float64Range1
	Scope:	[global]
	Access:	[file]
	DataType:	float64
	DomainType:	range
	Values:	[1000.01 0.02 10000.03]
	Comment:	float64Range1
	UpdateMode:	dynamic
	*/
	float64Range1    float64


	//parameter name -> parameter definition string
	name2definition map[string]string
}//end AllParameters

//all parameters can be set in the configuration file.
type configuration struct{
	//read and write lock
	rwlock	sync.RWMutex


	
	/**
	Name:	boolSet1
	Scope:	[global]
	Access:	[file]
	DataType:	bool
	DomainType:	set
	Values:	[true]
	Comment:	boolSet1
	UpdateMode:	dynamic
	*/
	BoolSet1    bool  `toml:"boolSet1"`

	

	
	/**
	Name:	boolSet2
	Scope:	[global]
	Access:	[file]
	DataType:	bool
	DomainType:	set
	Values:	[false]
	Comment:	boolSet2
	UpdateMode:	hotload
	*/
	BoolSet2    bool  `toml:"boolSet2"`

	

	
	/**
	Name:	boolSet3
	Scope:	[global]
	Access:	[file]
	DataType:	bool
	DomainType:	set
	Values:	[]
	Comment:	boolSet3
	UpdateMode:	dynamic
	*/
	BoolSet3    bool  `toml:"boolSet3"`

	

	
	/**
	Name:	stringSet1
	Scope:	[global]
	Access:	[file]
	DataType:	string
	DomainType:	set
	Values:	[ss1 ss2 ss3]
	Comment:	stringSet1
	UpdateMode:	dynamic
	*/
	StringSet1    string  `toml:"stringSet1"`

	

	
	/**
	Name:	stringSet2
	Scope:	[global]
	Access:	[file]
	DataType:	string
	DomainType:	set
	Values:	[]
	Comment:	stringSet2
	UpdateMode:	dynamic
	*/
	StringSet2    string  `toml:"stringSet2"`

	

	
	/**
	Name:	int64set1
	Scope:	[global]
	Access:	[file]
	DataType:	int64
	DomainType:	set
	Values:	[1 2 3 4 5 6]
	Comment:	int64Set1
	UpdateMode:	dynamic
	*/
	Int64set1    int64  `toml:"int64set1"`

	

	

	
	/**
	Name:	int64set3
	Scope:	[global]
	Access:	[file]
	DataType:	int64
	DomainType:	set
	Values:	[]
	Comment:	int64Set3
	UpdateMode:	dynamic
	*/
	Int64set3    int64  `toml:"int64set3"`

	

	
	/**
	Name:	int64Range1
	Scope:	[global]
	Access:	[file]
	DataType:	int64
	DomainType:	range
	Values:	[1000 0 10000]
	Comment:	int64Range1
	UpdateMode:	dynamic
	*/
	Int64Range1    int64  `toml:"int64Range1"`

	

	
	/**
	Name:	float64set1
	Scope:	[global]
	Access:	[file]
	DataType:	float64
	DomainType:	set
	Values:	[1.0 2.0 3.0 4.00 5.0 6.0]
	Comment:	float64Set1
	UpdateMode:	dynamic
	*/
	Float64set1    float64  `toml:"float64set1"`

	

	

	
	/**
	Name:	float64set3
	Scope:	[global]
	Access:	[file]
	DataType:	float64
	DomainType:	set
	Values:	[]
	Comment:	float64Set3
	UpdateMode:	dynamic
	*/
	Float64set3    float64  `toml:"float64set3"`

	

	
	/**
	Name:	float64Range1
	Scope:	[global]
	Access:	[file]
	DataType:	float64
	DomainType:	range
	Values:	[1000.01 0.02 10000.03]
	Comment:	float64Range1
	UpdateMode:	dynamic
	*/
	Float64Range1    float64  `toml:"float64Range1"`

	


	//parameter name -> updated flag
	name2updatedFlags map[string]bool
}//end configuration

/**
prepare something before anything else.
it is unsafe in multi-thread environment.
*/
func (ap *AllParameters) prepareAnything(){
	if ap.name2definition == nil {
		ap.name2definition = make(map[string]string)
	}
}

/**
set parameter and its string of the definition.
*/
func (ap *AllParameters) PrepareDefinition(){
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	ap.prepareAnything()
	
	ap.name2definition["boolSet1"] = "	Name:	boolSet1	Scope:	[global]	Access:	[file]	DataType:	bool	DomainType:	set	Values:	[true]	Comment:	boolSet1	UpdateMode:	dynamic	"
	
	ap.name2definition["boolSet2"] = "	Name:	boolSet2	Scope:	[global]	Access:	[file]	DataType:	bool	DomainType:	set	Values:	[false]	Comment:	boolSet2	UpdateMode:	hotload	"
	
	ap.name2definition["boolSet3"] = "	Name:	boolSet3	Scope:	[global]	Access:	[file]	DataType:	bool	DomainType:	set	Values:	[]	Comment:	boolSet3	UpdateMode:	dynamic	"
	
	ap.name2definition["stringSet1"] = "	Name:	stringSet1	Scope:	[global]	Access:	[file]	DataType:	string	DomainType:	set	Values:	[ss1 ss2 ss3]	Comment:	stringSet1	UpdateMode:	dynamic	"
	
	ap.name2definition["stringSet2"] = "	Name:	stringSet2	Scope:	[global]	Access:	[file]	DataType:	string	DomainType:	set	Values:	[]	Comment:	stringSet2	UpdateMode:	dynamic	"
	
	ap.name2definition["int64set1"] = "	Name:	int64set1	Scope:	[global]	Access:	[file]	DataType:	int64	DomainType:	set	Values:	[1 2 3 4 5 6]	Comment:	int64Set1	UpdateMode:	dynamic	"
	
	ap.name2definition["int64set2"] = "	Name:	int64set2	Scope:	[global]	Access:	[file]	DataType:	int64	DomainType:	set	Values:	[1 3 5 7]	Comment:	int64Set2	UpdateMode:	fix	"
	
	ap.name2definition["int64set3"] = "	Name:	int64set3	Scope:	[global]	Access:	[file]	DataType:	int64	DomainType:	set	Values:	[]	Comment:	int64Set3	UpdateMode:	dynamic	"
	
	ap.name2definition["int64Range1"] = "	Name:	int64Range1	Scope:	[global]	Access:	[file]	DataType:	int64	DomainType:	range	Values:	[1000 0 10000]	Comment:	int64Range1	UpdateMode:	dynamic	"
	
	ap.name2definition["float64set1"] = "	Name:	float64set1	Scope:	[global]	Access:	[file]	DataType:	float64	DomainType:	set	Values:	[1.0 2.0 3.0 4.00 5.0 6.0]	Comment:	float64Set1	UpdateMode:	dynamic	"
	
	ap.name2definition["float64set2"] = "	Name:	float64set2	Scope:	[global]	Access:	[file]	DataType:	float64	DomainType:	set	Values:	[1.001 3.003 5.005 7.007]	Comment:	float64Set2	UpdateMode:	fix	"
	
	ap.name2definition["float64set3"] = "	Name:	float64set3	Scope:	[global]	Access:	[file]	DataType:	float64	DomainType:	set	Values:	[]	Comment:	float64Set3	UpdateMode:	dynamic	"
	
	ap.name2definition["float64Range1"] = "	Name:	float64Range1	Scope:	[global]	Access:	[file]	DataType:	float64	DomainType:	range	Values:	[1000.01 0.02 10000.03]	Comment:	float64Range1	UpdateMode:	dynamic	"
	
}

/**
get the definition of the parameter.
*/
func (ap *AllParameters) GetDefinition(name string)(string,error){
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	ap.prepareAnything()
	if p,ok := ap.name2definition[name];!ok{
		return "",fmt.Errorf("there is no parameter %s",name)
	}else{
		return p,nil
	}
}

/**
check if there is the parameter
*/
func (ap *AllParameters) HasParameter(name string)bool{
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	ap.prepareAnything()
	if _,ok := ap.name2definition[name];!ok{
		return false
	}else{
		return true
	}
}

/**
Load the initial values of all parameters.
*/
func (ap *AllParameters) LoadInitialValues()error{
	ap.PrepareDefinition()
	var err error
	
		
		
			boolSet1choices :=[]bool {
				
				true,
					
			}
			if len(boolSet1choices) != 0{
				if err = ap.setBoolSet1( boolSet1choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","BoolSet1",err)
				}
			}else{
				
					if err = ap.setBoolSet1( false ) ; err != nil{
						return fmt.Errorf("set%s failed.error:%v","BoolSet1",err)
					}	
				
			}
		
	
		
		
			boolSet2choices :=[]bool {
				
				false,
					
			}
			if len(boolSet2choices) != 0{
				if err = ap.setBoolSet2( boolSet2choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","BoolSet2",err)
				}
			}else{
				
					if err = ap.setBoolSet2( false ) ; err != nil{
						return fmt.Errorf("set%s failed.error:%v","BoolSet2",err)
					}	
				
			}
		
	
		
		
			boolSet3choices :=[]bool {
					
			}
			if len(boolSet3choices) != 0{
				if err = ap.setBoolSet3( boolSet3choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","BoolSet3",err)
				}
			}else{
				
					if err = ap.setBoolSet3( false ) ; err != nil{
						return fmt.Errorf("set%s failed.error:%v","BoolSet3",err)
					}	
				
			}
		
	
		
		
			stringSet1choices :=[]string {
				
				"ss1",
				
				"ss2",
				
				"ss3",
					
			}
			if len(stringSet1choices) != 0{
				if err = ap.setStringSet1( stringSet1choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","StringSet1",err)
				}
			}else{
				//empty string
				if err = ap.setStringSet1( "" ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","StringSet1",err)
				}
			}
		
	
		
		
			stringSet2choices :=[]string {
					
			}
			if len(stringSet2choices) != 0{
				if err = ap.setStringSet2( stringSet2choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","StringSet2",err)
				}
			}else{
				//empty string
				if err = ap.setStringSet2( "" ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","StringSet2",err)
				}
			}
		
	
		
		
			int64set1choices :=[]int64 {
				
				1,
				
				2,
				
				3,
				
				4,
				
				5,
				
				6,
					
			}
			if len(int64set1choices) != 0{
				if err = ap.setInt64set1( int64set1choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","Int64set1",err)
				}
			}else{
				
					if err = ap.setInt64set1( 0 ) ; err != nil{
						return fmt.Errorf("set%s failed.error:%v","Int64set1",err)
					}
				
			}
		
	
		
		
			int64set2choices :=[]int64 {
				
				1,
				
				3,
				
				5,
				
				7,
					
			}
			if len(int64set2choices) != 0{
				if err = ap.setInt64set2( int64set2choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","Int64set2",err)
				}
			}else{
				
					if err = ap.setInt64set2( 0 ) ; err != nil{
						return fmt.Errorf("set%s failed.error:%v","Int64set2",err)
					}
				
			}
		
	
		
		
			int64set3choices :=[]int64 {
					
			}
			if len(int64set3choices) != 0{
				if err = ap.setInt64set3( int64set3choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","Int64set3",err)
				}
			}else{
				
					if err = ap.setInt64set3( 0 ) ; err != nil{
						return fmt.Errorf("set%s failed.error:%v","Int64set3",err)
					}
				
			}
		
	
		
		
			int64Range1choices :=[]int64 {
				
				1000,
				
				0,
				
				10000,
					
			}
			if len(int64Range1choices) != 0{
				if err = ap.setInt64Range1( int64Range1choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","Int64Range1",err)
				}
			}else{
				
					if err = ap.setInt64Range1( 0 ) ; err != nil{
						return fmt.Errorf("set%s failed.error:%v","Int64Range1",err)
					}
				
			}
		
	
		
		
			float64set1choices :=[]float64 {
				
				1.0,
				
				2.0,
				
				3.0,
				
				4.00,
				
				5.0,
				
				6.0,
					
			}
			if len(float64set1choices) != 0{
				if err = ap.setFloat64set1( float64set1choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","Float64set1",err)
				}
			}else{
				
					if err = ap.setFloat64set1( 0.0 ) ; err != nil{
						return fmt.Errorf("set%s failed.error:%v","Float64set1",err)
					}
				
			}
		
	
		
		
			float64set2choices :=[]float64 {
				
				1.001,
				
				3.003,
				
				5.005,
				
				7.007,
					
			}
			if len(float64set2choices) != 0{
				if err = ap.setFloat64set2( float64set2choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","Float64set2",err)
				}
			}else{
				
					if err = ap.setFloat64set2( 0.0 ) ; err != nil{
						return fmt.Errorf("set%s failed.error:%v","Float64set2",err)
					}
				
			}
		
	
		
		
			float64set3choices :=[]float64 {
					
			}
			if len(float64set3choices) != 0{
				if err = ap.setFloat64set3( float64set3choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","Float64set3",err)
				}
			}else{
				
					if err = ap.setFloat64set3( 0.0 ) ; err != nil{
						return fmt.Errorf("set%s failed.error:%v","Float64set3",err)
					}
				
			}
		
	
		
		
			float64Range1choices :=[]float64 {
				
				1000.01,
				
				0.02,
				
				10000.03,
					
			}
			if len(float64Range1choices) != 0{
				if err = ap.setFloat64Range1( float64Range1choices[0] ) ; err != nil{
					return fmt.Errorf("set%s failed.error:%v","Float64Range1",err)
				}
			}else{
				
					if err = ap.setFloat64Range1( 0.0 ) ; err != nil{
						return fmt.Errorf("set%s failed.error:%v","Float64Range1",err)
					}
				
			}
		
	
	return nil
}



/**
Get the value of the parameter boolSet1
*/
func (ap * AllParameters ) GetBoolSet1() bool {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.boolSet1
}

/**
Get the value of the parameter boolSet2
*/
func (ap * AllParameters ) GetBoolSet2() bool {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.boolSet2
}

/**
Get the value of the parameter boolSet3
*/
func (ap * AllParameters ) GetBoolSet3() bool {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.boolSet3
}

/**
Get the value of the parameter stringSet1
*/
func (ap * AllParameters ) GetStringSet1() string {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.stringSet1
}

/**
Get the value of the parameter stringSet2
*/
func (ap * AllParameters ) GetStringSet2() string {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.stringSet2
}

/**
Get the value of the parameter int64set1
*/
func (ap * AllParameters ) GetInt64set1() int64 {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.int64set1
}

/**
Get the value of the parameter int64set2
*/
func (ap * AllParameters ) GetInt64set2() int64 {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.int64set2
}

/**
Get the value of the parameter int64set3
*/
func (ap * AllParameters ) GetInt64set3() int64 {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.int64set3
}

/**
Get the value of the parameter int64Range1
*/
func (ap * AllParameters ) GetInt64Range1() int64 {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.int64Range1
}

/**
Get the value of the parameter float64set1
*/
func (ap * AllParameters ) GetFloat64set1() float64 {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.float64set1
}

/**
Get the value of the parameter float64set2
*/
func (ap * AllParameters ) GetFloat64set2() float64 {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.float64set2
}

/**
Get the value of the parameter float64set3
*/
func (ap * AllParameters ) GetFloat64set3() float64 {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.float64set3
}

/**
Get the value of the parameter float64Range1
*/
func (ap * AllParameters ) GetFloat64Range1() float64 {
	ap.rwlock.RLock()
	defer ap.rwlock.RUnlock()
	return ap.float64Range1
}







/**
Set the value of the parameter boolSet1
*/
func (ap * AllParameters ) SetBoolSet1(value bool)error {
	return  ap.setBoolSet1(value)
}



/**
Set the value of the parameter boolSet2
*/
func (ap * AllParameters ) SetBoolSet2(value bool)error {
	return  ap.setBoolSet2(value)
}



/**
Set the value of the parameter boolSet3
*/
func (ap * AllParameters ) SetBoolSet3(value bool)error {
	return  ap.setBoolSet3(value)
}



/**
Set the value of the parameter stringSet1
*/
func (ap * AllParameters ) SetStringSet1(value string)error {
	return  ap.setStringSet1(value)
}



/**
Set the value of the parameter stringSet2
*/
func (ap * AllParameters ) SetStringSet2(value string)error {
	return  ap.setStringSet2(value)
}



/**
Set the value of the parameter int64set1
*/
func (ap * AllParameters ) SetInt64set1(value int64)error {
	return  ap.setInt64set1(value)
}





/**
Set the value of the parameter int64set3
*/
func (ap * AllParameters ) SetInt64set3(value int64)error {
	return  ap.setInt64set3(value)
}



/**
Set the value of the parameter int64Range1
*/
func (ap * AllParameters ) SetInt64Range1(value int64)error {
	return  ap.setInt64Range1(value)
}



/**
Set the value of the parameter float64set1
*/
func (ap * AllParameters ) SetFloat64set1(value float64)error {
	return  ap.setFloat64set1(value)
}





/**
Set the value of the parameter float64set3
*/
func (ap * AllParameters ) SetFloat64set3(value float64)error {
	return  ap.setFloat64set3(value)
}



/**
Set the value of the parameter float64Range1
*/
func (ap * AllParameters ) SetFloat64Range1(value float64)error {
	return  ap.setFloat64Range1(value)
}






/**
Set the value of the parameter boolSet1
*/
func (ap * AllParameters ) setBoolSet1(value bool)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	
		
		
			choices :=[]bool {
				
				true,
					
			}
			if len( choices ) != 0{
				if !isInSliceBool(value, choices){
					return fmt.Errorf("setBoolSet1,the value %t is not in set %v",value,choices)
				}
			}//else means any bool value: true or false
		

	

	ap.boolSet1 = value
	return nil
}

/**
Set the value of the parameter boolSet2
*/
func (ap * AllParameters ) setBoolSet2(value bool)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	
		
		
			choices :=[]bool {
				
				false,
					
			}
			if len( choices ) != 0{
				if !isInSliceBool(value, choices){
					return fmt.Errorf("setBoolSet2,the value %t is not in set %v",value,choices)
				}
			}//else means any bool value: true or false
		

	

	ap.boolSet2 = value
	return nil
}

/**
Set the value of the parameter boolSet3
*/
func (ap * AllParameters ) setBoolSet3(value bool)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	
		
		
			choices :=[]bool {
					
			}
			if len( choices ) != 0{
				if !isInSliceBool(value, choices){
					return fmt.Errorf("setBoolSet3,the value %t is not in set %v",value,choices)
				}
			}//else means any bool value: true or false
		

	

	ap.boolSet3 = value
	return nil
}

/**
Set the value of the parameter stringSet1
*/
func (ap * AllParameters ) setStringSet1(value string)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	

		
			choices :=[]string {
				
				"ss1",
				
				"ss2",
				
				"ss3",
					
			}
			if len( choices ) != 0{
				if !isInSlice(value, choices){
					return fmt.Errorf("setStringSet1,the value %s is not in set %v",value,choices)
				}
			}//else means any string
		

	

	ap.stringSet1 = value
	return nil
}

/**
Set the value of the parameter stringSet2
*/
func (ap * AllParameters ) setStringSet2(value string)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	

		
			choices :=[]string {
					
			}
			if len( choices ) != 0{
				if !isInSlice(value, choices){
					return fmt.Errorf("setStringSet2,the value %s is not in set %v",value,choices)
				}
			}//else means any string
		

	

	ap.stringSet2 = value
	return nil
}

/**
Set the value of the parameter int64set1
*/
func (ap * AllParameters ) setInt64set1(value int64)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	

		
			choices :=[]int64 {
				
				1,
				
				2,
				
				3,
				
				4,
				
				5,
				
				6,
					
			}
			if len( choices ) != 0{
				if !isInSliceInt64(value, choices){
					return fmt.Errorf("setInt64set1,the value %d is not in set %v",value,choices)
				}
			}//else means any int64
		

	

	ap.int64set1 = value
	return nil
}

/**
Set the value of the parameter int64set2
*/
func (ap * AllParameters ) setInt64set2(value int64)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	

		
			choices :=[]int64 {
				
				1,
				
				3,
				
				5,
				
				7,
					
			}
			if len( choices ) != 0{
				if !isInSliceInt64(value, choices){
					return fmt.Errorf("setInt64set2,the value %d is not in set %v",value,choices)
				}
			}//else means any int64
		

	

	ap.int64set2 = value
	return nil
}

/**
Set the value of the parameter int64set3
*/
func (ap * AllParameters ) setInt64set3(value int64)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	

		
			choices :=[]int64 {
					
			}
			if len( choices ) != 0{
				if !isInSliceInt64(value, choices){
					return fmt.Errorf("setInt64set3,the value %d is not in set %v",value,choices)
				}
			}//else means any int64
		

	

	ap.int64set3 = value
	return nil
}

/**
Set the value of the parameter int64Range1
*/
func (ap * AllParameters ) setInt64Range1(value int64)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	

		
			choices :=[]int64 {
				
				1000,
				
				0,
				
				10000,
					
			}
			if !(value >= choices[1] && value <= choices[2]){
				return fmt.Errorf("setInt64Range1,the value %d is not in the range [%d,%d]",value,choices[1],choices[2])
			}
		

	

	ap.int64Range1 = value
	return nil
}

/**
Set the value of the parameter float64set1
*/
func (ap * AllParameters ) setFloat64set1(value float64)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	

		
			choices :=[]float64 {
				
				1.0,
				
				2.0,
				
				3.0,
				
				4.00,
				
				5.0,
				
				6.0,
					
			}
			if len( choices ) != 0{
				if !isInSliceFloat64(value, choices){
					return fmt.Errorf("setFloat64set1,the value %f is not in set %v",value,choices)
				}
			}//else means any float64
		
	

	ap.float64set1 = value
	return nil
}

/**
Set the value of the parameter float64set2
*/
func (ap * AllParameters ) setFloat64set2(value float64)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	

		
			choices :=[]float64 {
				
				1.001,
				
				3.003,
				
				5.005,
				
				7.007,
					
			}
			if len( choices ) != 0{
				if !isInSliceFloat64(value, choices){
					return fmt.Errorf("setFloat64set2,the value %f is not in set %v",value,choices)
				}
			}//else means any float64
		
	

	ap.float64set2 = value
	return nil
}

/**
Set the value of the parameter float64set3
*/
func (ap * AllParameters ) setFloat64set3(value float64)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	

		
			choices :=[]float64 {
					
			}
			if len( choices ) != 0{
				if !isInSliceFloat64(value, choices){
					return fmt.Errorf("setFloat64set3,the value %f is not in set %v",value,choices)
				}
			}//else means any float64
		
	

	ap.float64set3 = value
	return nil
}

/**
Set the value of the parameter float64Range1
*/
func (ap * AllParameters ) setFloat64Range1(value float64)error {
	ap.rwlock.Lock()
	defer ap.rwlock.Unlock()

	

		
			choices :=[]float64 {
				
				1000.01,
				
				0.02,
				
				10000.03,
					
			}
			if !(value >= choices[1] && value <= choices[2]){
				return fmt.Errorf("setFloat64Range1,the value %f is not in the range [%f,%f]",value,choices[1],choices[2])
			}
		
	

	ap.float64Range1 = value
	return nil
}



/**
prepare something before anything else.
it is unsafe in multi-thread environment.
*/
func (config *configuration) prepareAnything(){
	if config.name2updatedFlags == nil {
		config.name2updatedFlags = make(map[string]bool)
	}
}

/**
reset update flags of configuration items
*/
func (config *configuration) resetUpdatedFlags(){
	config.rwlock.Lock()
	defer config.rwlock.Unlock()
	config.prepareAnything()
	
	
		config.name2updatedFlags["boolSet1"] = false
	
	
	
		config.name2updatedFlags["boolSet2"] = false
	
	
	
		config.name2updatedFlags["boolSet3"] = false
	
	
	
		config.name2updatedFlags["stringSet1"] = false
	
	
	
		config.name2updatedFlags["stringSet2"] = false
	
	
	
		config.name2updatedFlags["int64set1"] = false
	
	
	
	
	
		config.name2updatedFlags["int64set3"] = false
	
	
	
		config.name2updatedFlags["int64Range1"] = false
	
	
	
		config.name2updatedFlags["float64set1"] = false
	
	
	
	
	
		config.name2updatedFlags["float64set3"] = false
	
	
	
		config.name2updatedFlags["float64Range1"] = false
	
	
}

/**
set update flag of configuration item
*/
func (config *configuration) setUpdatedFlag(name string,updated bool){
	config.rwlock.Lock()
	defer config.rwlock.Unlock()
	config.prepareAnything()
	config.name2updatedFlags[name] = updated
}

/**
get update flag of configuration item
*/
func (config *configuration) getUpdatedFlag(name string)bool{
	config.rwlock.RLock()
	defer config.rwlock.RUnlock()
	config.prepareAnything()
	return config.name2updatedFlags[name]
}

/**
Load parameters' values in the configuration string.
*/
func (config *configuration) LoadConfigurationFromString(input string) error {
	config.resetUpdatedFlags()

	metadata, err := toml.Decode(input, config);
	if err != nil {
		return err
	}else if failed := metadata.Undecoded() ; len(failed) > 0 {
		var failedItems []string
		for _, item := range failed {
			failedItems = append(failedItems, item.String())
		}
		return fmt.Errorf("decode failed %s. error:%v",failedItems,err)
	}

	for _,k := range metadata.Keys(){
		config.setUpdatedFlag(k[0],true)
	}

	return nil
}

/**
Load parameters' values in the configuration file.
*/
func (config *configuration) LoadConfigurationFromFile(fname string) error {
	config.resetUpdatedFlags()

	metadata, err := toml.DecodeFile(fname, config);
	if err != nil {
		return err
	}else if failed := metadata.Undecoded() ; len(failed) > 0 {
		var failedItems []string
		for _, item := range failed {
			failedItems = append(failedItems, item.String())
		}
		return fmt.Errorf("decode failed %s. error:%v",failedItems,err)
	}

	for _,k := range metadata.Keys(){
		config.setUpdatedFlag(k[0],true)
	}

	return nil
}

/**
Update parameters' values with configuration.
*/
func (ap * AllParameters ) UpdateParametersWithConfiguration(config *configuration)error{
	var err error
	
	
	if config.getUpdatedFlag("boolSet1"){
		if err = ap.setBoolSet1(config.BoolSet1); err != nil{
			return fmt.Errorf("update parameter boolSet1 failed.error:%v",err)
		}
	}
	
	
	
	if config.getUpdatedFlag("boolSet2"){
		if err = ap.setBoolSet2(config.BoolSet2); err != nil{
			return fmt.Errorf("update parameter boolSet2 failed.error:%v",err)
		}
	}
	
	
	
	if config.getUpdatedFlag("boolSet3"){
		if err = ap.setBoolSet3(config.BoolSet3); err != nil{
			return fmt.Errorf("update parameter boolSet3 failed.error:%v",err)
		}
	}
	
	
	
	if config.getUpdatedFlag("stringSet1"){
		if err = ap.setStringSet1(config.StringSet1); err != nil{
			return fmt.Errorf("update parameter stringSet1 failed.error:%v",err)
		}
	}
	
	
	
	if config.getUpdatedFlag("stringSet2"){
		if err = ap.setStringSet2(config.StringSet2); err != nil{
			return fmt.Errorf("update parameter stringSet2 failed.error:%v",err)
		}
	}
	
	
	
	if config.getUpdatedFlag("int64set1"){
		if err = ap.setInt64set1(config.Int64set1); err != nil{
			return fmt.Errorf("update parameter int64set1 failed.error:%v",err)
		}
	}
	
	
	
	
	
	if config.getUpdatedFlag("int64set3"){
		if err = ap.setInt64set3(config.Int64set3); err != nil{
			return fmt.Errorf("update parameter int64set3 failed.error:%v",err)
		}
	}
	
	
	
	if config.getUpdatedFlag("int64Range1"){
		if err = ap.setInt64Range1(config.Int64Range1); err != nil{
			return fmt.Errorf("update parameter int64Range1 failed.error:%v",err)
		}
	}
	
	
	
	if config.getUpdatedFlag("float64set1"){
		if err = ap.setFloat64set1(config.Float64set1); err != nil{
			return fmt.Errorf("update parameter float64set1 failed.error:%v",err)
		}
	}
	
	
	
	
	
	if config.getUpdatedFlag("float64set3"){
		if err = ap.setFloat64set3(config.Float64set3); err != nil{
			return fmt.Errorf("update parameter float64set3 failed.error:%v",err)
		}
	}
	
	
	
	if config.getUpdatedFlag("float64Range1"){
		if err = ap.setFloat64Range1(config.Float64Range1); err != nil{
			return fmt.Errorf("update parameter float64Range1 failed.error:%v",err)
		}
	}
	
	
	return nil
}

/**
Load configuration from file into configuration.
Then update items into AllParameters
*/
func LoadconfigurationFromFile(filename string,params *AllParameters) error{
	config := &configuration{}
	if err := config.LoadConfigurationFromFile(filename); err != nil{
		return err
	}

	if err := params.UpdateParametersWithConfiguration(config); err != nil{
		return err
	}
	return nil
}
