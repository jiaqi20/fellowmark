export interface User {
  ID?: number,
  Name?: string,
  Email?: string,
  Password?: string
}

export interface Enrollment {
  ModuleId?: number,
  Student?: User
}

export interface Assignment {
  ID?: number,
  Name?: string,
  ModuleID?: number,
  GroupSize?: number,
  Deadline?: number,
}

export interface Question {
  ID?: number,
  QuestionNumber?: number,
  QuestionText?: string,
  AssignmentID?: number
}

export interface Pairing {
  ID?: number,
	AssignmentID?: number,
	Student?: User,
	Marker?: User,
	Active?: Boolean
}
