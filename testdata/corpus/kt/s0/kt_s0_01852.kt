package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked15230(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked15231(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome15232 {
    data class Success(val payload: String) : Outcome15232()
    data object Failure : Outcome15232()
}

fun describe15232(outcome: Outcome15232): String = when (outcome) {
    is Outcome15232.Success -> outcome.payload
    Outcome15232.Failure -> "failed"
}
