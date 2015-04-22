package main

type DoitServerError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
