# nusmods-test
Offline test harness for NUSMods API, made for fun over the weekend.

Since this was only made in two days, only the basic, happy scenarios are included.
- Day 1: collect all endpoints, generate all JSONs, create verification script 
- Day 2: create config package and convert to subtests that can be noticed in Allure
- Future: convert to more compact test suites to cover more non-happy scenarios

## Setup
```bash
git clone git@github.com:RussellDash332/nusmods-test.git
cd nusmods-test
go mod tidy
```

## Generating testcase data
The convenient way is to simply use Python to create multiple JSON files. This is an ad-hoc solution because it's pretty much independent to the testing itself.

```bash
python3 gen.py
```

## Testing API endpoint(s)
```bash
# Test everything
# -count=1 to prevent caching
go test -v ./tests/... -count=1

# Test moduleInfo
go test -v ./tests/moduleInfo/... -count=1

# Test moduleInfo and venues
go test -v ./tests/moduleInfo/... ./tests/venues/... -count=1
```

## Generating a test report
To generate a nice and tidy test report, Allure is used along with `go-junit-report`.

### Setup
```bash
mkdir report && touch report/results.xml
```

### Installing Allure
```bash
cd ~ # Goes to home directory
wget https://github.com/allure-framework/allure2/releases/download/2.18.1/allure-2.18.1.tgz
tar -xvf allure-2.18.1.tgz
sudo cp -r allure-2.18.1/* /usr
```

### Installing go-junit-report
```bash
cd $GOPATH/bin # make sure GOPATH exists on your go env
go install github.com/jstemmer/go-junit-report@latest
```

### Creating a XML report
```bash
# Include all tests in report
go test -v ./tests/... -count=1 | go-junit-report > report/results.xml
```

### Generate Allure report
```bash
allure serve ./report
```

## Cool tips
It seems nice to have the Go tests colored so let's install something else.

```bash
cd $GOPATH/bin # make sure GOPATH exists on your go env
go install github.com/rakyll/gotest@latest
```

Afterwards, change every single `go test` syntax with `gotest`, for example
```bash
gotest -v ./tests/... -count=1
```

## Remarks
This not-so-hackathon is made to understand more about offline test harness during my internship. I might add more tests or assertions if I have more time.

Some tests might be failing due to data refresh, so `python3 gen.py` is the way to go if it's the case.