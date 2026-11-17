package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked21960(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome21961 {
    data class Success(val payload: String) : Outcome21961()
    data object Failure : Outcome21961()
}

fun describe21961(outcome: Outcome21961): String = when (outcome) {
    is Outcome21961.Success -> outcome.payload
    Outcome21961.Failure -> "failed"
}
