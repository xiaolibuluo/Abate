package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked8390(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome8391 {
    data class Success(val payload: String) : Outcome8391()
    data object Failure : Outcome8391()
}

fun describe8391(outcome: Outcome8391): String = when (outcome) {
    is Outcome8391.Success -> outcome.payload
    Outcome8391.Failure -> "failed"
}
sealed class Outcome8392 {
    data class Success(val payload: String) : Outcome8392()
    data object Failure : Outcome8392()
}

fun describe8392(outcome: Outcome8392): String = when (outcome) {
    is Outcome8392.Success -> outcome.payload
    Outcome8392.Failure -> "failed"
}
