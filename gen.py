# Testcase data generator
import requests, os, json, time

# Get modules preference from config
modules = list(map(lambda x: x.strip(), open('config/modules.txt').readlines()))
years = list(map(lambda x: f'{x}-{x+1}', range(2018, 2022))) # NUSMods API available from 2018-2019

# To prevent QPS issue
old_get = requests.get
def new_get(url, *args):
    print(f'Getting {url}...')
    time.sleep(0.1)
    return old_get(url, *args)
requests.get = new_get

# Generate moduleInfo endpoint
def genModuleInfo():
    print()
    for year in years:
        req = requests.get(f'https://api.nusmods.com/v2/{year}/moduleInfo.json')
        assert req.status_code == 200, f'Status code must be 200 but got {req.status_code} instead'
        data = req.json()

        try:    os.mkdir('tests/moduleInfo/testdata')
        except: pass

        with open(f'tests/moduleInfo/testdata/{year}.json', 'w+') as f:
            json.dump(data, f, indent=4)

# Generate moduleList endpoint
def genModuleList():
    print()
    for year in years:
        req = requests.get(f'https://api.nusmods.com/v2/{year}/moduleList.json')
        assert req.status_code == 200, f'Status code must be 200 but got {req.status_code} instead'
        data = req.json()

        try:    os.mkdir('tests/moduleList/testdata')
        except: pass

        with open(f'tests/moduleList/testdata/{year}.json', 'w+') as f:
            json.dump(data, f, indent=4)

# Generate modules endpoint
def genModules():
    print()
    for year in years:
        for module in modules:
            req = requests.get(f'https://api.nusmods.com/v2/{year}/modules/{module}.json')
            assert req.status_code == 200, f'Status code must be 200 but got {req.status_code} instead'
            data = req.json()

            try:    os.mkdir('tests/modules/testdata')
            except: pass
            try:    os.mkdir(f'tests/modules/testdata/{year}')
            except: pass

            with open(f'tests/modules/testdata/{year}/{module}.json', 'w+') as f:
                json.dump(data, f, indent=4)

# Generate venueInformation endpoint
def genVenueInformation():
    print()
    for year in years:
        for semester in range(1, 5):
            req = requests.get(f'https://api.nusmods.com/v2/{year}/semesters/{semester}/venueInformation.json')
            assert req.status_code == 200, f'Status code must be 200 but got {req.status_code} instead'
            data = req.json()

            try:    os.mkdir('tests/venueInformation/testdata')
            except: pass
            try:    os.mkdir(f'tests/venueInformation/testdata/{year}')
            except: pass

            with open(f'tests/venueInformation/testdata/{year}/{semester}.json', 'w+') as f:
                json.dump(data, f, indent=4)

# Generate venues endpoint
def genVenues():
    print()
    for year in years:
        for semester in range(1, 5):
            req = requests.get(f'https://api.nusmods.com/v2/{year}/semesters/{semester}/venues.json')
            assert req.status_code == 200, f'Status code must be 200 but got {req.status_code} instead'
            data = req.json()
            
            try:    os.mkdir('tests/venues/testdata')
            except: pass
            try:    os.mkdir(f'tests/venues/testdata/{year}')
            except: pass

            with open(f'tests/venues/testdata/{year}/{semester}.json', 'w+') as f:
                json.dump(data, f, indent=4)

genModuleInfo()
genModuleList()
genModules()
genVenueInformation()
genVenues()
