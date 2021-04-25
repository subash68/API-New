# APIS LIST

## Onboard Service apis

> _Note_: Base url : http://localhost:8080

> On failure Check Statuscode in header not in the response body

> On invalid token no response will have "invalid token" check statuscode in header for proper error handling

- /o

  - /signUp/Corporate

    - Method: _POST_
    - Content-Type: _multipart/form-data_
    <details>
    <summary>
    <i>Paramenters</i>
    </summary>
    <p>

    ```JSON
        {
            "corporateName":"String, Required",
            "CIN": "String , Required",
            "corporateHQAddressLine1": "String , Not-required",
            "corporateHQAddressLine2": "String , Not-required",
            "corporateHQAddressLine3": "String , Not-required",
            "corporateHQAddressCountry": "String , Not-required",
            "corporateHQAddressState": "String , Not-required",
            "corporateHQAddressCity": "String , Not-required",
            "corporateHQAddressDistrict": "String , Not-required",
            "corporateHQAddressZipCode": "String , Not-required",
            "corporateHQAddressPhone": "String ,13 digits, Not-required",
            "corporateHQAddressEmail": "Email , Not-required",
            "corporateLocalBranchAddressLine1": "String , Not-required",
            "corporateLocalBranchAddressLine2": "String , Not-required",
            "corporateLocalBranchAddressLine3": "String , Not-required",
            "corporateLocalBranchAddressCountry": "String , Not-required",
            "corporateLocalBranchAddressState": "String , Not-required",
            "corporateLocalBranchAddressCity": "String , Not-required",
            "corporateLocalBranchAddressDistrict": "String , Not-required",
            "corporateLocalBranchAddressZipCode": "String , Not-required",
            "corporateLocalBranchAddressPhone": "String ,13 digits, Not-required",
            "corporateLocalBranchAddressEmail": "Email , Not-required",
            "primaryContactFirstName": "String , Required",
            "primaryContactMiddleName": "String , Not-required",
            "primaryContactLastName": "String ,  Required",
            "primaryContactDesignation": "String , Required",
            "primaryContactPhone": "String ,13 digits, Required",
            "primaryContactEmail": "Email , Required",
            "secondaryContactFirstName": "String , Not-required",
            "secondaryContactMiddleName": "String , Not-required",
            "secondaryContactLastName": "String , Not-required",
            "secondaryContactDesignation": "String , Not-required",
            "secondaryContactPhone": "String ,13 digits, Not-required",
            "secondaryContactEmail": "Email , Not-required",
            "corporateType": "String , Required",
            "corporateCategory": "String , Required",
            "corporateIndustry": "String , Not-required",
            "companyProfile": "String , Not-required",
            "attachment": "String , Not-required",
            "yearOfEstablishment": "INT , Required",
            "password": "String ,Required",

        }
    ```

    </p>
    </details>

    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "accountStatus": "Created",
            "message": "OTP sent to Mobile and Email for verification",
            "platformUID": "String, StakeholderID",
            "email": "Registered Email",
            "phoneNumber": "Registered phone number",
            "stakeholder": "Corporate"
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "SignupInformation",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "User details not Found",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /signUp/University

    - Method: _POST_
    - Content-Type: _multipart/form-data_
    <details>
    <summary>
    <i>Paramenters</i>
    </summary>
    <p>

    ```JSON
        {
            "universityName":"String, Not-required",
            "universityHQAddressLine1":"String, Not-required",
            "universityHQAddressLine2":"String, Not-required",
            "universityHQAddressLine3":"String, Not-required",
            "universityHQAddressCountry":"String, Not-required",
            "universityHQAddressState":"String, Not-required",
            "universityHQAddressCity":"String, Not-required",
            "universityHQAddressDistrict":"String, Not-required",
            "universityHQAddressZipcode":"String, Not-required",
            "universityHQAddressPhone":"String, Not-required",
            "universityHQAddressemail":"String, Not-required",
            "universityLocalBranchAddressLine1":"String, Not-required",
            "universityLocalBranchAddressLine2":"String, Not-required",
            "universityLocalBranchAddressLine3":"String, Not-required",
            "universityLocalBranchAddressCountry":"String, Not-required",
            "universityLocalBranchAddressState":"String, Not-required",
            "universityLocalBranchAddressCity":"String, Not-required",
            "universityLocalBranchAddressDistrict":"String, Not-required",
            "universityLocalBranchAddressZipcode":"String, Not-required",
            "universityLocalBranchAddressPhone":"String, Not-required",
            "universityLocalBranchAddressemail":"String, Not-required",
            "primaryContactFirstName":"String, Required",
            "primaryContactMiddleName":"String, Required",
            "primaryContactLastName":"String, Required",
            "primaryContactDesignation":"String, Required",
            "primaryContactPhone":"String ,13 digits, Required ",
            "primaryContactEmail":"Email, required",
            "secondaryContactFirstName":"String, Not-required",
            "secondaryContactMiddleName":"String, Not-required",
            "secondaryContactLastName":"String, Not-required",
            "secondaryContactDesignation":"String, Not-required",
            "secondaryContactPhone":"String, Not-required",
            "secondaryContactEmail":"String, Not-required",
            "universitySector":"Strin , required",
            "yearOfEstablishment":"String, required",
            "universityProfile":"String, Not-required",
            "attachment":"File, Not-required",
            "password":"String, Required min=8,max=15",

        }
    ```

    </p>
    </details>

    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "accountStatus": "Created",
            "message": "OTP sent to Mobile and Email for verification",
            "platformUID": "String, StakeholderID",
            "email": "Registered Email",
            "phoneNumber": "Registered phone number",
            "stakeholder": "University"
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "SignupInformation",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "User details not Found",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /signUp/Student

    - Method: _POST_
    - Content-Type: _multipart/form-data_
    <details>
    <summary>
    <i>Paramenters in FormData</i>
    </summary>
    <p>

    ```JSON
        {
        "firstName":"String, Required",
        "middleName":"String, Not-Required",
        "lastName":"String, Required",
        "email":"Email, Required",
        "phone":"String,13 digits, Required",
        "collegeID":"String, Not-Required",
        "gender":"String, Not-Required",
        "dateOfBirth":"String, Not-Required",
        "aadharNumber":"String, Not-Required",
        "permanentAddressLine1":"String, Not-Required",
        "permanentAddressLine2":"String, Not-Required",
        "permanentAddressLine3":"String, Not-Required",
        "permanentAddressCountry":"String, Not-Required",
        "permanentAddressState":"String, Not-Required",
        "permanentAddressCity":"String, Not-Required",
        "permanentAddressDistrict":"String, Not-Required",
        "permanentAddressZipcode":"String, Not-Required",
        "permanentAddressPhone":"String, Not-Required",
        "permanentAddressemail":"String, Not-Required",
        "presentAddressLine1":"String, Not-Required",
        "presentAddressLine2":"String, Not-Required",
        "presentAddressLine3":"String, Not-Required",
        "presentAddressCountry":"String, Not-Required",
        "presentAddressState":"String, Not-Required",
        "presentAddressCity":"String, Not-Required",
        "presentAddressDistrict":"String, Not-Required",
        "presentAddressZipcode":"String, Not-Required",
        "presentAddressPhone":"String, Not-Required",
        "presentAddressemail":"String, Not-Required",
        "fathersFirstName":"String, Not-Required",
        "fathersMiddleName":"String, Not-Required",
        "fathersLastName":"String, Not-Required",
        "secondaryContactFirstName":"String, Not-Required",
        "secondaryContactMiddleName":"String, Not-Required",
        "secondaryContactLastName":"String, Not-Required",
        "secondaryContactDesignation":"String, Not-Required",
        "secondaryContactPhone":"String, Not-Required",
        "secondaryContactEmail":"String, Not-Required",
        "studentProfile":"String, Not-Required",
        "attachment":"File, Not-Required",
        "password":"String, Required, Min 8 max 15",

        }
    ```

    </p>
    </details>

    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "accountStatus": "Created",
            "message": "OTP sent to Mobile and Email for verification",
            "platformUID": "String, StakeholderID",
            "email": "Registered Email",
            "phoneNumber": "Registered phone number",
            "stakeholder": "Student"
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "SignupInformation",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "User details not Found",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /verifyMobile

    - Method: _*POST*_
    - Content-Type: _multipart/form-data_
    <details>
    <summary>
    <i>Paramenters in FormData</i>
    </summary>
    <p>

    ```JSON
        {
        "stakeholder":"Enum[Corporate,University,Student], required",
        "platformUID":"String, Required",
        "otp":"string, required",
        "phone":"Phone, 13 digits, required",
        }
    ```

    </p>
    </details>

    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "message": "OTP verification successful"
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "SignupInformation",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "User details not Found",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /verifyEmail

    - Method: _*POST*_
    - Content-Type: _multipart/form-data_
    <details>
    <summary>
    <i>Paramenters in FormData</i>
    </summary>
    <p>

    ```JSON
        {
        "stakeholder":"Enum[Corporate,University,Student], required",
        "platformUID":"String, Required",
        "otp":"string, required",
        "email":"email, required",
        }
    ```

    </p>
    </details>

    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "message": "OTP verification successful"
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "Target form",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "Custom message",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /resendOtp

    - Method: _*POST*_
    - Content-Type: _multipart/form-data_
    <details>
    <summary>
    <i>Paramenters in FormData</i>
    </summary>
    <p>

    ```JSON
        {
        "stakeholder":"Enum[Corporate,University,Student], required",
        "platformUID":"String, Required",
        "otpType":"Enum[Email,Phone], required",
        }
    ```

    </p>
    </details>

    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "message": "String "
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "Target form",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "Custom message",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /login

    - Method: _*POST*_
    - Content-Type: _multipart/form-data_
    <details>
    <summary>
    <i>Paramenters in FormData</i>
    </summary>
    <p>

    ```JSON
        {
        "stakeholder":"Enum[Corporate,University,Student], required",
        "userID":"String, Required",
        "password":"String, required",
        }
    ```

    </p>

    </details>
    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "token": "String",
            "redirectURL": "Enum for redirecting to pages, [/dashboard,/payment]"
        }
    ```

    </p>
    </details>
    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "LoginInformation",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "User details not Found",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /sendOTP

    - Method: _*POST*_
    - Content-Type: _multipart/form-data_
    <details>
    <summary>
    <i>Paramenters in FormData</i>
    </summary>
    <p>

    ```JSON
        {
        "stakeholder":"Enum[Corporate,University,Student], required",
        "vrfBy":"Enum [ Phone ,Email]",
        "phone":"String , 13 digits, requird only of vrby is Phone",
        "email":"Email, requird only of vrby is Email",
        "platformUUID":"String, required",
        }
    ```

    </p>
    </details>

    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "message": "String"
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "Target form",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "Custom message",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /resetPassword

    - Method: _*POST*_
    - Content-Type: _multipart/form-data_
    <details>
    <summary>
    <i>Paramenters in FormData</i>
    </summary>
    <p>

    ```JSON
        {
        "stakeholder":"Enum[Corporate,University,Student], required",
        "platformUID":"String, Required",
        "otp":"String, Required",
        "vrfBy":"Enum[Phone,Email], Required",
        "phone":"String, Required",
        "email":"String, Required",
        "platformUUID":"String, Required",
        "newPassword":"String, Required",
        }
    ```

    </p>
    </details>

    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "message": "Password Changed Succesfully",
            "Redirectpath": "/login"
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "Target form",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "Custom message",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /logout

    - Method: _*POST*_
    - Content-Type: _multipart/form-data_
    - No Parameters required
    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "message": "Token deleted"
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "Target form",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "Custom message",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

## Profile Service

- /u

  - /profile

    - Method: _*GET*_
    - No Parameters required
    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "": "Based on stakeholder login, Thi method will return all the details of the stakeholder which captured at sign up with below additional details",
            "dateOfJoining": "Date",
            "accountStatus": "ACTIVE",
            "primaryPhoneVerified": "boolean",
            "primaryEmailVerified": "boolean"

        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "Target form",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "Custom message",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /profile

    - Method: _*PATCH*_
    - Content-Type: _multipart/form-data_
    <details>
    <summary>
    <i>Paramenters in FormData</i>
    </summary>
    <p>

    ```JSON
        {
        "Details":"Key and value of the stakeholders profile update details"
        }
    ```

    </p>
    </details>
    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "message": "String "
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "Target form",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "Custom message",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /profilePic

    - Method: _*POST*_
    - Content-Type: _multipart/form-data_
    <details>
    <summary>
    <i>Paramenters in FormData</i>
    </summary>
    <p>

    ```JSON
        {
        "profilePic":"file, required"
        }
    ```

    </p>
    </details>
    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "message": "String "
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "Target form",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "Custom message",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /profilePic

    - Method: _*GET*_
    - No parameters required
    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "profilePic": "base 64 String "
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "Target form",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "Custom message",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

  - /lut/?lutList=_[]strings_

    - Method: _*GET*_
    - Content-Type: _multipart/form-data_
    <details>
    <summary>
    <i>Query Paramenters in URL</i>
    </summary>
    <p>

    ```JSON
        {
        "lutList":"[]Strings,Required valid strings -> [corporateType,corporateCategory,corporateIndustry,universityCategory,skills,programs,departments]"
        }
    ```

    </p>
    </details>

    <details>
    <summary>
    <i>Success Response</i>
    </summary>
    <p>

    ```JSON
        {
            "corporateTypes": [
                {
                    "codeDescription": "string",
                    "code": "string",
                    "charCode": "String"
                }],
            "corporateCategory": [
                {
                    "codeDescription": "Branch Office",
                    "code": "string",
                    "charCode": "string"
                }],
            "universityCategory": [
                {
                    "codeDescription": "Branch Office",
                    "code": "string",
                    "charCode": "string"
                }],
            "skills": [
                {
                    "SkillID": "string",
                    "skill": "string"
                }],
            "programs": [
                {
                    "ProgramID": "string",
                    "program": "string"
                }],
            "departments": [
                {
                    "departmentID": "string",
                    "ProgramID": "string",
                    "department": "string"
                }]
        }

    ```

    </p>
    </details>

    </details>
    <details>
    <summary>
    <i>Failure Response</i>
    </summary>
    <p>

    ```JSON
        {
            "code": "S1LGN001",
            "message": "Input Validation Error",
            "target": "Target form",
            "errors": [
                {
                    "code": "S1LGN001",
                    "message": "Custom message",
                    "target": "All"
                }
            ]
        }
    ```

    </p>
    </details>

## Publish Service

- /p

  - /crp

    - /hiringCriteria

      - Method: _*POST*_
      - Content-Type: _multipart/form-data_
      <details>
      <summary>
      <i>Paramenters in FormData</i>
      </summary>
      <p>

      ```JSON
          [
              {
                "hcName":"String, Required",
                "programID":"String, Required",
                "departmentID":"String, Required",
                "cutOffCategory":"Enum[CGPA,%], Required",
                "cutOff":"Float, Required",
                "eduGapsSchoolAllowed":"Boolean, Required",
                "eduGaps11N12Allowed":"Boolean, Required",
                "eduGapsGradAllowed":"Boolean, Required",
                "eduGapsPGAllowed":"Boolean, Required",
                "allowActiveBacklogs":"Int, Required",
                "yearOfPassing":"Int, Required",
                "remarks":"String, Not-Required",
            }
          ]
      ```

        </p>
        </details>

    - /hiringCriteria/getByID/:hcID

      - Method: _*GET*_
      - No parameters required

    - /hiringCriteria/all

      - Method: _*GET*_
      - No parameters required

    - /hiringCriteria/:hcID

      - Method: _*PATCH*_
      - Content-Type: _multipart/form-data_
      <details>
      <summary>
      <i>Paramenters in FormData</i>
      </summary>
      <p>

      ```JSON
          {
          "Details":"Key:value of Hiring criteria details to update"
          }
      ```

        </p>
        </details>

    - /hiringCriteria/:hcID

      - Method: _*DELETE*_
      - No parameters required

    - /createJob

      - Method: _*POST*_
      - Content-Type: _multipart/form-data_
      <details>
      <summary>
      <i>Paramenters in FormData</i>
      </summary>
      <p>

      ```JSON
          {
            "jobName":"String, Required",
            "hiringCriteriaID":"String, Not-Required",
            "jobs":[
                {
                    "skillID":"String, Not-Required",
                    "skill":"String, Not-Required",
                    "noOfPositions":"Int, Not-Required",
                    "location":"String, Not-Required",
                    "salaryType":"String, Not-Required",
                    "dateOfHiring":"Date, Required",
                    "status":"String, Not-Required",
                    "remarks":"String, Not-Required",
                    "attachment":"File, Not-Required",
                }
            ]
          }
      ```

        </p>
        </details>

    - /createJob/getByID/:jobID

      - Method: _*GET*_
      - No parameters required

    - /createJob/all

      - Method: _*GET*_
      - No parameters required

    - /createJob/job/:jobID

      - Method: _*PATCH*_
      - Content-Type: _multipart/form-data_
      <details>
      <summary>
      <i>Paramenters in FormData</i>
      </summary>
      <p>

      ```JSON
          {
            "jobName":"String, Required",
            "hiringCriteriaID":"String, Not-Required",
            "jobs":[
                {
                    "skillID":"String, Not-Required",
                    "skill":"String, Not-Required",
                    "noOfPositions":"Int, Not-Required",
                    "location":"String, Not-Required",
                    "salaryType":"String, Not-Required",
                    "dateOfHiring":"Date, Required",
                    "status":"String, Not-Required",
                    "remarks":"String, Not-Required",
                    "attachment":"File, Not-Required",
                }
            ]
          }
      ```

        </p>
        </details>

    - /createJob/mapHC/:jobID

      - Method: _*PATCH*_
      - Content-Type: _multipart/form-data_
      <details>
      <summary>
      <i>Paramenters in FormData</i>
      </summary>
      <p>

      ```JSON
          {
          "hiringCriteriaID":"String, Required"
          }
      ```

        </p>
        </details>

    - /createJob/:jobID

      - Method: _*DELETE*_
      - No parameters required

    - /publishJob

      - Method: _*POST*_
      - Content-Type: _multipart/form-data_
      <details>
      <summary>
      <i>Paramenters in FormData</i>
      </summary>
      <p>

      ```JSON
          {
          "publishJobs":[{
              "jobID":"String, Required"
          }]
          }
      ```

        </p>
        </details>

    - /publishJob/getByID/:pjID

      - Method: _*GET*_
      - No parameters required

    - /publishJob/all

      - Method: _*GET*_
      - No parameters required

    - /publishJob/:pjID

      - Method: _*PATCH*_
      - Content-Type: _multipart/form-data_
      <details>
      <summary>
      <i>Paramenters in FormData</i>
      </summary>
      <p>

      ```JSON
          {
          "jobID":"String, Required"
          }
      ```

        </p>
        </details>

    - /publishJob/:pjID

      - Method: _*DELETE*_
      - No parameters required

  - /unv

    - /proposal

      - Method: _*POST*_ - Content-Type: _multipart/form-data_
      <details>
      <summary>
      <i>Paramenters in FormData</i>
      </summary>
      <p>

      ```JSON
          {
              "programs":[
                {
                  "programID":"String, required",
                  "programType":"String, required",
                  "programName":"String, required",
                  "startDate":"Date, required",
                  "endDate":"Date, required",
                  "enablingFlag":"String, required",
                }
              ],
              "branches":[
                {
                  "programID":"String, required",
                  "branchName":"String, required",
                  "branchID":"String, required",
                  "branchName":"String, required",
                  "startDate":"String, required",
                  "endDate":"String, required",
                  "enablingFlag":"String, required",
                  "noOfPassingStudents":"String, required",
                  "monthYearOfPassing":"String, required",
                }
              ],
              "accredations":[
                {
                  "accredationID":"String, required",
                  "accredationName":"String, required",
                  "accredationType":"String, required",
                  "accredationDescription":"String, required",
                  "issuingAuthority":"String, required",
                  "accredationFile":"String, required",
                  "startDate":"String, required",
                  "endDate":"String, required",
                  "enablingFlag":"String, required",
                }
              ],
              "rankings":[
                {
                  "rank":"String, required",
                  "issuingAuthority":"String, required"
                }
              ],
              "specialOfferings":[
                {
                  "specialOfferingType":"String, required",
                  "specialOfferingName":"String, required",
                  "specialOfferingDescription":"String, required",
                  "internallyManagedFlag":"String, required",
                  "outsourcedVendorName":"String, required",
                  "outsourcedVendorContact":"String, required",
                  "outsourcedVendorStakeholderID":"String, required",
                  "specialOfferingFile":"String, required",
                  "startDate":"String, required",
                  "endDate":"String, required",
                  "enablingFlag":"String, required",
                }
              ],
              "tieups":[
                {
                  "typeupType":"String,Required",
                  "tieupName":"String,Required",
                  "tieupDescription":"String,Required",
                  "tieupWithName":"String,Required",
                  "tieupWithContact":"String,Required",
                  "tieupWithStakeholderID":"String,Required",
                  "tieupfile":"String,Required",
                  "startDate":"String,Required",
                  "endDate":"String,Required",
                  "enablingFlag":"String,Required",
                }
              ],
              "coes":[
                {
                  "coeID":"string, required",
                  "coeType":"string, required",
                  "coeName":"string, required",
                  "coeDescription":"string, required",
                  "internallyManagedFlag":"string, required",
                  "outsourcedVendorName":"string, required",
                  "outsourcedVendorContact":"string, required",
                  "outsourcedVendorStakeholderID":"string, required",
                  "coeFile":"string, required",
                  "startDate":"string, required",
                  "endDate":"string, required",
                  "enablingFlag":"string, required",
               }
              ]
          }

      ```

      </p>
      </details>

    - /proposal

      - Method: _*GET*_
      - No parameters required
