import "cypress-axe";

describe("header spec", () => {
  it("has all of the expected information within the header", () => {
    cy.visit("index.html");
    cy.contains(".govuk-header__link--homepage", "OPG");
    cy.contains(".govuk-header__service-name", "Sirius");
  });

  const expectedTitle = [
    "Power of Attorney",
    "Supervision",
    "Admin",
    "Sign out",
  ];
  const expectedUrl = ["/lpa", "/Supervision", "/Admin", "/Logout"];

  it("has working nav links within header banner", () => {
    cy.visit("index.html");
    cy.get("#header-navigation")
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

  it("has all the expected links within the secondary nav list", () => {
    cy.visit("index.html");
    const expectedTitle = ["Create client", "Workflow", "Guidance", "Finance"];
    const expectedUrl = [
      "/supervision/#/clients/search-for-client",
      "/supervision/workflow",
      "https://wordpress.sirius.opg.service.justice.gov.uk",
      "/supervision/#/finance-hub/reporting",
    ];
    cy.get("#header-navigation")
      .get(".moj-primary-navigation__list")
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

  it("highlights the current page nav", () => {
    cy.visit("/supervision/workflow");
    cy.get("#header-navigation")
      .contains("Supervision")
      .should("not.be.visible");
    cy.get(".moj-primary-navigation__list > .moj-primary-navigation__item")
      .contains("Workflow")
      .should("have.attr", "aria-current", "page");

    cy.visit("/supervision");
    cy.get("#header-navigation")
      .contains("Supervision")
      .should("not.be.visible");
    cy.get(".moj-primary-navigation__list > .moj-primary-navigation__item")
      .contains("Workflow")
      .should("not.have.attr", "aria-current");
  });

  it("hides the finance tab when not set to show", () => {
    // no show-finance passed in (default)
    cy.visit("/supervision/workflow");
    cy.get(".moj-primary-navigation__list")
      .contains("Finance")
      .should("be.hidden");

    // show-finance is "false"
    cy.visit("/supervision");
    cy.get(".moj-primary-navigation__list")
      .contains("Finance")
      .should("be.hidden");
  });

  it("meets accessibility standards", () => {
    cy.visit("/index.html");
    cy.injectAxe();
    cy.checkA11y();
  });
});
