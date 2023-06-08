describe("Sirius Header homepage", () => {
  beforeEach(() => {
    cy.setCookie("Other", "other");
    cy.setCookie("XSRF-TOKEN", "abcde");
    cy.visit("/supervision/");
  });

  it('finds the content "Hello world!"', () => {
    cy.contains('Hello world!')
  })
});