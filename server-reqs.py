import requests

#r = requests.post("http://127.0.0.1:8080/login", json ={
#    "username": "JohnDoe599",
#    "password": "JohnDoe",
#})


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

#body ="""
#{
#        employees(sort: ID_DESC, amount: 3) {
#            id
#            firstName
#            lastName
#            }
#        }
#"""

#body="""
#mutation {
#        deleteEmployee(id: 2) 
#        }
#"""

#body="""
#mutation {
#        updateEmployee(input: {
#            id: 1
#            firstName: "Philip"
#            })
#        }
#"""

body="""
mutation {
        createEmployee(input: {
            firstName: "Joe"
            lastName: "Banana"
            username: "TheBanana112"
            password: "hahaha"
            email: "BananaJoe@gmail.com"
            dateOfBirth: "2004-06-15"
            departmentId: 11
            position: SENIOR
            })
        }
"""
r = requests.post("http://127.0.0.1:8080/employee", 
    data=body
)

print(r.text)
