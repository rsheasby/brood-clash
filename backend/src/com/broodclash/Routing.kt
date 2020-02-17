package com.broodclash

import io.ktor.auth.authenticate
import io.ktor.locations.Location
import io.ktor.locations.get
import io.ktor.locations.post
import io.ktor.routing.Routing

@Location("/authTest")
class AuthTest

@Location("/questions")
class Questions {
	@Location("/{questionId}")
	class Question(val questionId: Long) {
		@Location("")
		class Get

		@Location("/answers")
		class Answers(val parent: Question) {
			@Location("/{answerId}")
			class Answer(val parent: Answers, val answerId: Long) {
				@Location("/reveal")
				class Reveal(val parent: Answer)
			}
		}
	}

	@Location("")
	class Create(val questions: Iterable<Question>)
}

@Location("/presenter/websocket")
class Websocket

fun Routing.configureApi() {
	authenticate(AUTH_NAME) {
		get<AuthTest> {
			throw NotImplementedError()
		}

		get<Questions> {
			throw NotImplementedError()
		}

		post<Questions.Create> {
			throw NotImplementedError()
		}

		get<Questions.Question> {
			throw NotImplementedError()
		}

		// TODO: I know this is horrible.
		post<Questions.Question.Answers.Answer.Reveal> {
			throw NotImplementedError()
		}
	}

	get<Websocket> {
		throw NotImplementedError()
	}
}
