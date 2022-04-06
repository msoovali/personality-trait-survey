# Personality trait survey
Figure out which personality trait You are!
## How to get it up and running?
1. Use provided dockerfiles (webapp is separate angular application)
2. If needed change initial data in conf/data.json
## TODO roadmap for perfection :sunglasses:
1. Write unit tests for trait and survey services. All test files with atleast 80% coverage. Current coverage can be seen in coverage/coverage.html
2. Implement jest unit test runner for webapp and write unit tests
3. Add swagger documentation. Currently rest requests are documented in requests.rest
4. Implement redis or database support aside to application memory storage. Make it configurable
5. Improve webapp's UI