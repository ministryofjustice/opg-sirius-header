describe('header spec', () => {

  beforeEach(() => {
    cy.visit('../index.html')
  })

  it('has all of the expected information within the header', () => {
    cy.get('.moj-header__logo').should('be.visible');
    cy.get('.moj-header__logo > .moj-header__link').should('contain.text', 'OPG');
    cy.contains(".moj-header__link", "Sirius");
    cy.get('.govuk-grid-column-two-thirds').should('contain.text', 'Sirius');
  });

  const expectedTitle = ["Power of Attorney", "Supervision", "Admin", "Sign out"];
  const expectedUrl = ["/lpa", "/Supervision", "/Admin", "/Logout"];

  it("has working nav links within header banner", () => {
    cy.get(".moj-header__navigation-list")
        .children()
        .each(($el, index) => {
          cy.wrap($el).should("contain", expectedTitle[index]);
          const $expectedLinkName = expectedUrl[index].toLowerCase();
          cy.wrap($el)
              .find("a")
              .should("have.attr", "href")
              .and("contain", `${$expectedLinkName}`);
        });
  });

  it('has all the expected links within the secondary nav list', () => {

    const expectedTitle = ["Create client", "Workflow", "Finance"];
    const expectedUrl = ["/supervision/#/clients/search-for-client", "/supervision/workflow", "/supervision/#/finance-hub/reporting"];

    cy.get('.moj-primary-navigation__list')
        .children()
        .each(($el, index) => {
          cy.wrap($el).should("contain", expectedTitle[index]);
          const $expectedLinkName = expectedUrl[index].toLowerCase();
          cy.wrap($el)
              .find("a")
              .should("have.attr", "href")
              .and("contain", `${$expectedLinkName}`);
        });
    });
});