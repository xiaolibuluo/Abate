package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked2950(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome2951 {
    data class Success(val payload: String) : Outcome2951()
    data object Failure : Outcome2951()
}

fun describe2951(outcome: Outcome2951): String = when (outcome) {
    is Outcome2951.Success -> outcome.payload
    Outcome2951.Failure -> "failed"
}
