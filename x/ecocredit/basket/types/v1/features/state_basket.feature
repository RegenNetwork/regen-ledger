Feature: Basket

  Scenario: a valid basket
    Given the basket
    """
    {
      "id": 1,
      "basket_denom": "eco.uC.NCT",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When the basket is validated
    Then expect no error

  Scenario: a valid basket with date criteria
    Given the basket
    """
    {
      "id": 1,
      "basket_denom": "eco.uC.NCT",
      "name": "NCT",
      "credit_type_abbrev": "C",
      "date_criteria": {
        "years_in_the_past": 10
      },
      "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When the basket is validated
    Then expect no error

  Scenario: an error is returned if id is empty
    Given the basket
    """
    {}
    """
    When the basket is validated
    Then expect the error "id cannot be zero: parse error"

  Scenario: an error is returned if basket denom is empty
    Given the basket
    """
    {
      "id": 1
    }
    """
    When the basket is validated
    Then expect the error "basket denom: empty string is not allowed: parse error"

  Scenario: an error is returned if basket denom is not formatted
    Given the basket
    """
    {
      "id": 1,
      "basket_denom": "foo"
    }
    """
    When the basket is validated
    Then expect the error "basket denom: expected format eco.<exponent-prefix><credit-type-abbrev>.<name>: parse error"

  Scenario: an error is returned if name is empty
    Given the basket
    """
    {
      "id": 1,
      "basket_denom": "eco.uC.NCT"
    }
    """
    When the basket is validated
    Then expect the error "name: empty string is not allowed: parse error"

  Scenario: an error is returned if name is not formatted
    Given the basket
    """
    {
      "id": 1,
      "basket_denom": "eco.uC.NCT",
      "name": "1"
    }
    """
    When the basket is validated
    Then expect the error "name: must start with an alphabetic character, and be between 3 and 8 alphanumeric characters long: parse error"

  Scenario: an error is returned if credit type is empty
    Given the basket
    """
    {
      "id": 1,
      "basket_denom": "eco.uC.NCT",
      "name": "NCT"
    }
    """
    When the basket is validated
    Then expect the error "credit type abbrev: empty string is not allowed: parse error"

  Scenario: an error is returned if credit type is not formatted
    Given the basket
    """
    {
      "id": 1,
      "basket_denom": "eco.uC.NCT",
      "name": "NCT",
      "credit_type_abbrev": "1"
    }
    """
    When the basket is validated
    Then expect the error "credit type abbrev: must be 1-3 uppercase alphabetic characters: parse error"

  Scenario: an error is returned if curator is empty
    Given the basket
    """
    {
      "id": 1,
      "basket_denom": "eco.uC.NCT",
      "name": "NCT",
      "credit_type_abbrev": "C"
    }
    """
    When the basket is validated
    Then expect the error "curator: empty address string is not allowed: parse error"
