package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7470 {
    data class Success(val payload: String) : Outcome7470()
    data object Failure : Outcome7470()
}

fun describe7470(outcome: Outcome7470): String = when (outcome) {
    is Outcome7470.Success -> outcome.payload
    Outcome7470.Failure -> "failed"
}
sealed class Outcome7471 {
    data class Success(val payload: String) : Outcome7471()
    data object Failure : Outcome7471()
}

fun describe7471(outcome: Outcome7471): String = when (outcome) {
    is Outcome7471.Success -> outcome.payload
    Outcome7471.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked7472(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
