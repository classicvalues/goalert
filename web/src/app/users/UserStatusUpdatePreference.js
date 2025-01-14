import React from 'react'
import { gql } from '@apollo/client'
import p from 'prop-types'
import Query from '../util/Query'
import { Mutation } from '@apollo/client/react/components'
import UserContactMethodSelect from './UserContactMethodSelect'

const query = gql`
  query statusUpdate($id: ID!) {
    user(id: $id) {
      id
      statusUpdateContactMethodID
    }
  }
`
const mutation = gql`
  mutation ($id: ID!, $cmID: ID!) {
    updateUser(input: { id: $id, statusUpdateContactMethodID: $cmID })
  }
`

const disableVal = 'disable'

export default function UserStatusUpdatePreference({ userID }) {
  function renderControl(cmID, updateCM) {
    return (
      <UserContactMethodSelect
        userID={userID}
        label='Alert Status Updates'
        helperText='Update me when my alerts are acknowledged or closed'
        name='alert-status-contact-method'
        value={cmID || disableVal}
        onChange={updateCM}
        extraItems={[{ label: 'Disabled', value: disableVal }]}
      />
    )
  }

  function renderMutation(user) {
    const setCM = (commit) => (e) => {
      const cmID = e.target.value === disableVal ? '' : e.target.value
      commit({
        variables: {
          id: userID,
          cmID,
        },
      })
    }
    return (
      <Mutation mutation={mutation}>
        {(commit) =>
          renderControl(user.statusUpdateContactMethodID, setCM(commit))
        }
      </Mutation>
    )
  }

  return (
    <Query
      query={query}
      variables={{ id: userID }}
      render={({ data }) => renderMutation(data.user)}
    />
  )
}

UserStatusUpdatePreference.propTypes = {
  userID: p.string.isRequired,
}
