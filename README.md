Assumptions:
 1. Go has been successfully installed and the installation tested.
 2. The "go" version installed on the test system is 1.9.2
 3. The OS installed on the test system is Ubuntu 16.04.
 
Repository code structure contains two directories as below:
1. golangexercises/driver/ is the main package that contains a driver file to test out the functionality provided by the rectangle package.
2. golangexercises/rectangle/ is the packagage that provides the method to calculate intersection between 2 rectangles.
3. golangexercises/DifferentScenarios is a JPG file that illustrates the different scenarios with respect to rectangles and may help understand with the algorithm and test data.
                
Test 1: Execute the default driver program that has predefined test data:
Follow the below steps:
1. Execute "go get -u github.com/vishwanathj/golangexercises/driver"
2. cd to the bin directory and execute ./driver to see 4 lines of output that displays intersecting rectangle of the given 2 rectangles.

Test 2: Execute the driver program by adding additional test data
1. Edit the src/github.com/vishwanathj/golangexercises/driver/driver.go file to add additional data by following instructions in the source file at lines 10 - 15
2. cd to src directory and execute "go build github.com/vishwanathj/golangexercises/driver/"
3. Next execute "go install github.com/vishwanathj/golangexercises/driver"
4. cd to the bin directory and execute ./driver to see calculated intersection rectangles.

Test 3: Execute the tests in rectangle package
1. cd to the src directory
2. Execute the command "go test github.com/vishwanathj/golangexercises/rectangle/rectangle_test.go"