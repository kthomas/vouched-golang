package vouched

const graphqlSubmitJobMutation = `mutation {
	submitJob(input: $input) {
	  success
	  type
	  country
	  state
	  confidences {
		id
		backId
		selfie
		idMatch
		faceMatch
	  }
	}
  }`

const graphqlQueryJobs = `query jobs(
    $id: ID
    $type: String
    $status: String
    $withPhotos: Boolean
  ) {
    jobs(
	  id: $id
	  type: $type
      status: $status
	  withPhotos: $withPhotos
    ) {
      total
      totalPages
      pageSize
      page
      items {
		id
		status
		submitted
		request
		result
		errors
      }
    }
  }`
