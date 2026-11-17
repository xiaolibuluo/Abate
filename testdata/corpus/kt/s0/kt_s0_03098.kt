package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25440 {
    data class Success(val payload: String) : Outcome25440()
    data object Failure : Outcome25440()
}

fun describe25440(outcome: Outcome25440): String = when (outcome) {
    is Outcome25440.Success -> outcome.payload
    Outcome25440.Failure -> "failed"
}
fun <T> List<T>.secondOrNull25441(): T? = if (size > 1) this[1] else null

fun render25441(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome25442 {
    data class Success(val payload: String) : Outcome25442()
    data object Failure : Outcome25442()
}

fun describe25442(outcome: Outcome25442): String = when (outcome) {
    is Outcome25442.Success -> outcome.payload
    Outcome25442.Failure -> "failed"
}
