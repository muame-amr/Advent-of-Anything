from itertools import product
import requests

# URL to send the POST request
url = "https://hhc23-reportinator-dot-holidayhack2023.ue.r.appspot.com/check"

# Request body data in the specified format
data = "input-1=1&input-2=0&input-3=0&input-4=0&input-5=0&input-6=0&input-7=0&input-8=0&input-9=0"

# Cookies to include in the request (replace with your actual cookies)
cookies = {"ReportinatorCookieYum": "[REDACTED]"}

# Number of inputs
num_inputs = 9

# Generate all combinations of 0 and 1 for the given number of inputs
input_combinations = list(product([0, 1], repeat=num_inputs))

# Print the generated combinations
for comb in input_combinations:
    body_str = f"input-1={comb[0]}&input-2={comb[1]}&input-3={comb[2]}&input-4={comb[3]}&input-5={comb[4]}&input-6={comb[5]}&input-7={comb[6]}&input-8={comb[7]}&input-9={comb[8]}"
    # Sending the POST request with the specified data
    response = requests.post(url, data=body_str, cookies=cookies)
    # Printing the response
    if response.status_code != 400:
        print(f"✅ {comb}")
        print(f"Status Code: {response.status_code}")
        print("Response Content:")
        print(response.text)
        break
    else:
        print(f"❌ {body_str}")

response = requests.post(
    url,
    data="input-1=0&input-2=0&input-3=0&input-4=0&input-5=0&input-6=0&input-7=1&input-8=1&input-9=0",
    cookies=cookies,
)
print(f"Status Code: {response.status_code}")
print("Response Content:")
print(response.text)
