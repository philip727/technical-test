import requests

#r = requests.post("http://127.0.0.1:8080/login", json ={
#    "username": "JohnDoe599",
#    "password": "JohnDoe",
#})

# print(r.Text)

# Define the GraphQL query
#body = """
#{
#    employees {
#        id
#        firstName
#        lastName
#        email
#        dateOfBirth
#        departmentId
#        position
#    }
#}
#"""
#
#
#r = requests.post("http://127.0.0.1:8080/employee", 
#    data=body
#)

body ="""
{
        employees(sort: ID_DESC, amount: 3) {
            id
            firstName
            lastName
            }
        }
"""

r = requests.post("http://127.0.0.1:8080/employee", 
    data=body
)
print(r.json())
