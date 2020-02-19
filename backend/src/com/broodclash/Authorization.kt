package com.broodclash

import io.ktor.application.ApplicationCall
import io.ktor.application.call
import io.ktor.auth.*
import io.ktor.request.ApplicationRequest
import io.ktor.request.authorization
import io.ktor.util.KtorExperimentalAPI

data class TokenCredential(val token: String) : Credential

class TokenAuthenticationProvider internal constructor(
	configuration: Configuration
) : AuthenticationProvider(configuration) {
	internal val authenticationFunction = configuration.authenticationFunction

	class Configuration internal constructor(name: String?) : AuthenticationProvider.Configuration(name) {
		internal var authenticationFunction: AuthenticationFunction<TokenCredential> = {
			throw NotImplementedError(
				"Token auth validate function is not specified. Use token { validate { ... } } to fix."
			)
		}

		fun validate(body: suspend ApplicationCall.(TokenCredential) -> Principal?) {
			authenticationFunction = body
		}
	}
}

@KtorExperimentalAPI
fun Authentication.Configuration.token(
	name: String? = null,
	configure: TokenAuthenticationProvider.Configuration.() -> Unit
) {
	val provider = TokenAuthenticationProvider(TokenAuthenticationProvider.Configuration(name).apply(configure))
	val authenticate = provider.authenticationFunction

	provider.pipeline.intercept(AuthenticationPipeline.RequestAuthentication) { context ->
		val credentials = call.request.tokenAuthenticationCredentials()
		val principal = credentials?.let { authenticate(call, it) }

		if (credentials == null || principal == null) {
			throw AuthorizationException()
		}
		context.principal(principal)
	}

	register(provider)
}

@KtorExperimentalAPI
fun ApplicationRequest.tokenAuthenticationCredentials(): TokenCredential? {
	val auth = authorization()
	return if (auth != null) TokenCredential(auth) else null
}
