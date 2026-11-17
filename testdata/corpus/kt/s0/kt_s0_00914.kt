package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull7500(): T? = if (size > 1) this[1] else null

fun render7500(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull7501(): T? = if (size > 1) this[1] else null

fun render7501(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome7502 {
    data class Success(val payload: String) : Outcome7502()
    data object Failure : Outcome7502()
}

fun describe7502(outcome: Outcome7502): String = when (outcome) {
    is Outcome7502.Success -> outcome.payload
    Outcome7502.Failure -> "failed"
}
sealed class Outcome7503 {
    data class Success(val payload: String) : Outcome7503()
    data object Failure : Outcome7503()
}

fun describe7503(outcome: Outcome7503): String = when (outcome) {
    is Outcome7503.Success -> outcome.payload
    Outcome7503.Failure -> "failed"
}
