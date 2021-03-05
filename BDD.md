# Behavior Driven Development
Copied from https://en.wikipedia.org/wiki/Behavior-driven_development

# Principles of BDD
Test-driven development is a software-development methodology which essentially states that for each unit of software, a software developer must:

* define a test set for the unit first;
* make the tests fail;
* then implement the unit;
* finally verify that the implementation of the unit makes the tests succeed.

This definition is rather non-specific in that it allows tests in terms of high-level software requirements, low-level technical details or anything in between.
One way of looking at BDD therefore, is that it is a continued development of TDD which makes more specific choices than TDD.

Behavior-driven development specifies that tests of any unit of software should be specified in terms of the desired behavior of the unit.
Borrowing from agile software development the "desired behavior" in this case consists of the requirements set by the business — that is, the desired behavior that has business value for whatever entity commissioned the software unit under construction.
Within BDD practice, this is referred to as BDD being an "outside-in" activity.

## Behavioral specifications
Following this fundamental choice, a second choice made by BDD relates to how the desired behavior should be specified.
In this area BDD chooses to use a semi-formal format for behavioral specification which is borrowed from user story specifications from the field of object-oriented analysis and design.
The scenario aspect of this format may be regarded as an application of Hoare logic to behavioral specification of software units using the Domain Language of the situation.

BDD specifies that business analysts and developers should collaborate in this area and should specify behavior in terms of user stories, which are each explicitly written down in a dedicated document.
Each User Story should, in some way, follow the following structure:

### Title
An explicit title.
### Narrative
A short introductory section with the following structure:
* **As a**: the person or role who will benefit from the feature;
* **I want**: the feature;
* **so that**: the benefit or value of the feature.
### Acceptance criteria
A description of each specific scenario of the narrative with the following structure:
* **Given**: the initial context at the beginning of the scenario, in one or more clauses;
* **When**: the event that triggers the scenario;
* **Then**: the expected outcome, in one or more clauses.

BDD does not have any formal requirements for exactly how these user stories must be written down, but it does insist that each team using BDD come up with a simple, standardized format for writing down the user stories which includes the elements listed above.
However, in 2007 Dan North suggested a template for a textual format which has found wide following in different BDD software tools.
A very brief example of this format might look like this:

    Title: Returns and exchanges go to inventory.
    
    As a store owner,
    I want to add items back to inventory when they are returned or exchanged,
    so that I can track inventory.
    
    Scenario 1: Items returned for refund should be added to inventory.
    Given that a customer previously bought a black sweater from me
    and I have three black sweaters in inventory,
    when they return the black sweater for a refund,
    then I should have four black sweaters in inventory.
    
    Scenario 2: Exchanged items should be returned to inventory.
    Given that a customer previously bought a blue garment from me
    and I have two blue garments in inventory
    and three black garments in inventory,
    when they exchange the blue garment for a black garment,
    then I should have three blue garments in inventory
    and two black garments in inventory.

# Story versus specification
A separate subcategory of behavior-driven development is formed by tools that use specifications as an input language rather than user stories.
An example of this style is the RSpec tool that was also originally developed by Dan North.
Specification tools don't use user stories as an input format for test scenarios but rather use functional specifications for units that are being tested.
These specifications often have a more technical nature than user stories and are usually less convenient for communication with business personnel than are user stories.

An example of a specification for a stack might look like this:

    Specification: Stack
    
    When a new stack is created
    Then it is empty
    
    When an element is added to the stack
    Then that element is at the top of the stack
    
    When a stack has N elements
    And element E is on top of the stack
    Then a pop operation returns E
    And the new size of the stack is N-1

Such a specification may exactly specify the behavior of the component being tested, but is less meaningful to a business user.
As a result, specification-based testing is seen in BDD practice as a complement to story-based testing and operates at a lower level.
Specification testing is often seen as a replacement for free-format unit testing.