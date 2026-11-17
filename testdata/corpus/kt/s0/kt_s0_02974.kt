package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24470 {
    data class Success(val payload: String) : Outcome24470()
    data object Failure : Outcome24470()
}

fun describe24470(outcome: Outcome24470): String = when (outcome) {
    is Outcome24470.Success -> outcome.payload
    Outcome24470.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked24471(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome24472 {
    data class Success(val payload: String) : Outcome24472()
    data object Failure : Outcome24472()
}

fun describe24472(outcome: Outcome24472): String = when (outcome) {
    is Outcome24472.Success -> outcome.payload
    Outcome24472.Failure -> "failed"
}
