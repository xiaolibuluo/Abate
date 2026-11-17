package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24960 {
    data class Success(val payload: String) : Outcome24960()
    data object Failure : Outcome24960()
}

fun describe24960(outcome: Outcome24960): String = when (outcome) {
    is Outcome24960.Success -> outcome.payload
    Outcome24960.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked24961(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked24962(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
