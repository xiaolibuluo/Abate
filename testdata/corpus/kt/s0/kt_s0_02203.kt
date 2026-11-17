package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked18140(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome18141 {
    data class Success(val payload: String) : Outcome18141()
    data object Failure : Outcome18141()
}

fun describe18141(outcome: Outcome18141): String = when (outcome) {
    is Outcome18141.Success -> outcome.payload
    Outcome18141.Failure -> "failed"
}
sealed class Outcome18142 {
    data class Success(val payload: String) : Outcome18142()
    data object Failure : Outcome18142()
}

fun describe18142(outcome: Outcome18142): String = when (outcome) {
    is Outcome18142.Success -> outcome.payload
    Outcome18142.Failure -> "failed"
}
