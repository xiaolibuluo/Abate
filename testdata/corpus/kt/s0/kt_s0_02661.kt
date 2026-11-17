package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked21990(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome21991 {
    data class Success(val payload: String) : Outcome21991()
    data object Failure : Outcome21991()
}

fun describe21991(outcome: Outcome21991): String = when (outcome) {
    is Outcome21991.Success -> outcome.payload
    Outcome21991.Failure -> "failed"
}
