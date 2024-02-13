# School Prototype App

The School Prototype App is a straightforward school management application offering essential authentication and authorization features for users across different roles. As of now, the app facilitates user creation, login, and grants access to authenticated users' information.

## Features

- **User Creation**: Provides an unprotected endpoint for the system's new user registration.
- **Login**: Authenticates users and generates a JWT for authenticated sessions.
- **User Information Access**: A protected endpoint delivering information about the authenticated user, based on the supplied JWT token.

## Endpoints

### Create User

- **URL**: `/users`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "username": "bob",
    "password": "password123"
  }
  ```
### Login
 - **URL**: `/login`
 - **Method**: `POST`
 - **Request Body**:
 ```json
  {
    "username": "bob",
    "password": "password123"
  }
 ```

### Access User Information
 - **URL**: `/auth/me`
 - **Method**: `POST`
 - **Headers**: `Authorization: Bearer <token>`

## Password Security
The application currently employs SHA-256 for hashing passwords. Although SHA-256 is secure for many uses, it's not the best choice for password hashing due to its susceptibility to brute force and rainbow table attacks.

### Suggested Improvement
For enhanced security of passwords, it's advisable to use algorithms specifically intended for password hashing, like bcrypt, Argon2, or PBKDF2. These algorithms are designed to be computationally demanding and automatically incorporate a salt, making brute force attacks substantially more difficult.

## Implementing Permissions
### Next steps
The app is set to support various user roles, including STUDENT, TEACHER, and ADMIN, each with distinct permissions for nuanced access control based on the user's role:
 - **STUDENT**: Restricted access tailored to student-related information and functionalities.
 - **TEACHER**: Authorized to manage educational content and access student data.
 - **ADMIN**: Granted full system access, encompassing user management and application settings.
The development of a comprehensive permissions framework is essential to ensure that users access only those features and information pertinent to their role within the application.