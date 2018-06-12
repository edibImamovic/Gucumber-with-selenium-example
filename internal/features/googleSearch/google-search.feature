@scenario
Feature: Google Search

    @remote
    Scenario: Submit a form
        Given I navigate to "https://google.com/"
        When I write "Google" to the search bar
        And I click  button search
        Then I should see "Google - Google Pretraživanje" in link text
