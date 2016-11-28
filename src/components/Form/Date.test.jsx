import React from 'react'
import { shallow } from 'enzyme'
import { Date } from './Date'

describe('The date component', () => {
  it('renders appropriately with an error', () => {
    const expected = {
      name: 'input-error',
      label: 'Date input error',
      help: 'Helpful error message',
      error: true,
      focus: false,
      valid: false
    }
    const component = shallow(<Date name={expected.name} label={expected.label} help={expected.help} error={expected.error} focus={expected.focus} valid={expected.valid} />)
    expect(component.find('label.usa-input-error-label').length).toEqual(3)
    expect(component.find('input#' + expected.name + '-month').length).toEqual(1)
    expect(component.find('span.usa-input-error-message').length).toEqual(3)
    expect(component.find('span.hidden').length).toEqual(0)
  })

  it('renders appropriately with focus', () => {
    const expected = {
      name: 'input-focus',
      label: 'Date input focused',
      help: 'Helpful error message',
      error: false,
      focus: true,
      valid: false
    }
    const component = shallow(<Date name={expected.name} label={expected.label} help={expected.help} error={expected.error} focus={expected.focus} valid={expected.valid} />)
    expect(component.find('label').length).toEqual(3)
    expect(component.find('input#' + expected.name + '-month').length).toEqual(1)
    expect(component.find('input#' + expected.name + '-month').hasClass('usa-input-focus')).toEqual(true)
    expect(component.find('span.hidden').length).toEqual(3)
  })

  it('renders appropriately with validity checks', () => {
    const expected = {
      name: 'input-success',
      label: 'Date input success',
      help: 'Helpful error message',
      error: false,
      focus: false,
      valid: true
    }
    const component = shallow(<Date name={expected.name} label={expected.label} help={expected.help} error={expected.error} focus={expected.focus} valid={expected.valid} />)
    expect(component.find('label').length).toEqual(3)
    expect(component.find('input#' + expected.name + '-month').length).toEqual(1)
    expect(component.find('input#' + expected.name + '-month').hasClass('usa-input-success')).toEqual(true)
    expect(component.find('span.hidden').length).toEqual(3)
  })

  it('renders sane defaults', () => {
    const expected = {
      name: 'input-type-text',
      label: 'Date input label',
      help: 'Helpful error message',
      error: false,
      focus: false,
      valid: false
    }
    const component = shallow(<Date name={expected.name} label={expected.label} help={expected.help} error={expected.error} focus={expected.focus} valid={expected.valid} />)
    expect(component.find('label').length).toEqual(3)
    expect(component.find('input#' + expected.name + '-month').length).toEqual(1)
    expect(component.find('span.hidden').length).toEqual(3)
  })
})
