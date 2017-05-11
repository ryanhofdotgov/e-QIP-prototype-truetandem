import ForeignDirectInterestValidator from './foreigndirectinterest'
import { validBranch } from './helpers'

export default class ForeignDirectActivityValidator {
  constructor (state, props = {}) {
    this.hasInterests = props.HasInterests || ''
    this.list = props.List || []
  }

  isValid () {
    if (!validBranch(this.hasInterests)) {
      return false
    }
    if (this.hasInterests === 'No') {
      return true
    }

    if (this.hasInterests === 'Yes' && !this.list.length) {
      return false
    }

    return this.list.every((item) => {
      return new ForeignDirectInterestValidator(null, item.DirectInterest).isValid()
    })
  }
}