package com.broodclash

data class Question(
	val id: Long,
	val text: String,
	val answers: List<Answer>
) {
	init {
		assert(text.isNotEmpty())
		assert(answers.size in 1..8)
	}
}

data class Answer(
	val id: Long,
	val text: String,
	val points: Int,
	val revealed: Boolean
) {
	init {
		assert(text.isNotEmpty())
	}
}

