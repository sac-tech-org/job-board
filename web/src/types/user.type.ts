export interface Email {
  address: string;
  verified: boolean;
}

export interface User {
  email: Email;
  firstName: string;
  lastName: string;
  username: string;
}
