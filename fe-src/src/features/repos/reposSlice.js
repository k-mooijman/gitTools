const initialState = []

function nextRepoId(repos) {
  const maxId = repos.reduce((maxId, repo) => Math.max(repo.id, maxId), -1)
  return maxId + 1
}

export default function reposReducer(state = initialState, action) {
  switch (action.type) {
    case 'repos/repoAdded': {
      console.log("--->")

      // Can return just the new repos array - no extra object around it
      return [
        ...state,
        {
          id: nextRepoId(state),
          text: action.payload,
          completed: false,
        },
      ]
    }
    case 'repos/repoToggled': {
      return state.map((repo) => {
        if (repo.id !== action.payload) {
          return repo
        }

        return {
          ...repo,
          completed: !repo.completed,
        }
      })
    }
    case 'repos/colorSelected': {
      const { color, repoId } = action.payload
      return state.map((repo) => {
        if (repo.id !== repoId) {
          return repo
        }

        return {
          ...repo,
          color,
        }
      })
    }
    case 'repos/repoDeleted': {
      return state.filter((repo) => repo.id !== action.payload)
    }
    case 'repos/allCompleted': {
      return state.map((repo) => {
        return { ...repo, completed: true }
      })
    }
    case 'repos/completedCleared': {
      return state.filter((repo) => !repo.completed)
    }
    default:
      return state
  }
}
