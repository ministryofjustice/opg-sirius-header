describe('header spec', () => {

  beforeEach(() => {
    cy.visit('../index.html')
  })

  it('has all of the expected information within the header', () => {
    cy.get('.moj-header__logo').should('be.visible');
    cy.get('.govuk-grid-column-two-thirds > .moj-header__link').should('contain.text', 'Sirius');
    cy.get('.govuk-grid-column-two-thirds').should('contain.text', 'Sirius');
    cy.get('.govuk-grid-column-two-thirds').should('contain.text', 'Power of Attorney');
    cy.get('.govuk-grid-column-two-thirds').should('contain.text', 'Supervision');
    cy.get('.govuk-grid-column-two-thirds').should('contain.text', 'Admin');
    cy.get('.govuk-grid-column-two-thirds').should('contain.text', 'Sign out');
  })

  it('has all the expected links within the secondary nav list', () => {
    cy.get('.moj-primary-navigation__list').should('contain.text', 'Create client');
    cy.get('.moj-primary-navigation__list').should('contain.text', 'Workflow');
    cy.get('.moj-primary-navigation__list').should('contain.text', 'Finance');
  })
})