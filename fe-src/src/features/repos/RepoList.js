import React from 'react'
import { useSelector, shallowEqual } from 'react-redux'

const selectRepoIds = (state) => state.repos.map((repo) => repo.id)

const RepoList = () => {
  const repoIds = useSelector(selectRepoIds, shallowEqual)

  const renderedListItems = repoIds.map((repoId) => {
    return  <div>
      repoId: {repoId}
    </div>
  })

  return <ul className="repo-list">{renderedListItems}</ul>
}

export default RepoList
