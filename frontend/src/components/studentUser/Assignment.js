import { Component } from 'react';

import Button from '@material-ui/core/Button';

import { Auth } from '../../context/context';

/** 
 * Assignment consists of:
 *  - Question Dropdown
 *  - Question Text
 *  - Submission Card
 *    - File Upload
 *    - Text Input
 *    - Submit Button       
 *
 * Data required:
 *  - Student ID of User
 *  - Module ID
 *
 * Data retrieved:
 *  - Question IDs of questions open for submission
 *  - Question text
 *
 * Data submitted:
 *  - Submission file / submission content
*/
class Assignment extends Component {
  constructor(props) {
    super(props);
    this.submitDocument = this.submitDocument.bind(this);
    this.state = {
      questionIds: [],
      submitted: false
    };
  }

  submitDocument() { }

  render() {
    // const submitButton = <Button variant="contained" size="large" onClick={this.submitDocument}>Submit</Button>;

  }
}

Assignment.contextType = Auth;

export default Assignment;