package vouched

const graphqlSubmitJobMutation = `mutation submitJob($type: String!, $callbackURL: String, $params: JobParams) {
    submitJob(type: $type, callbackURL: $callbackURL, params: $params) {
      id
      status
      request {
        type
        callbackURL
        parameters {
          idPhoto
          userPhoto
          twicPhoto
          carInsurancePhoto
          dotPhoto
          firstName
          lastName
          dob
        }
      }
      result {
        id
        success
        type
        country
        state
        expireDate
        birthDate
        firstName
        lastName
        confidences {
          id
          backId
          selfie
          idMatch
          faceMatch
        }
      }
      errors {
        type
        message
        suggestion
      }
      submitted
    }
  }`

const graphqlQueryJobs = `query jobs(
    $id: ID
    $ids: [ID]
    $type: String
    $status: String
    $to: String
    $from: String
    $withPhotos: Boolean
    $sortOrder: String
    $sortBy: String
    $page: Int
    $pageSize: Int
  ) {
    jobs(
      withPhotos: $withPhotos
      id: $id
      ids: $ids
      status: $status
      type: $type
      to: $to
      from: $from
      sortOrder: $sortOrder
      sortBy: $sortBy
      page: $page
      pageSize: $pageSize
    ) {
      total
      totalPages
      pageSize
      page
      items {
        id
        status
        request {
          type
          callbackURL
          parameters {
            idPhoto
            userPhoto
            twicPhoto
            carInsurancePhoto
            dotPhoto
            firstName
            lastName
            dob
          }
        }
        result {
          id
          success
          type
          country
          state
          expireDate
          birthDate
          firstName
          lastName
          confidences {
            id
            backId
            selfie
            idMatch
            faceMatch
          }
        }
        errors {
          type
          message
          suggestion
        }
        submitted
      }
    }
  }`
