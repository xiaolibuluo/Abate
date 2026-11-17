package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked22430(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome22431 {
    data class Success(val payload: String) : Outcome22431()
    data object Failure : Outcome22431()
}

fun describe22431(outcome: Outcome22431): String = when (outcome) {
    is Outcome22431.Success -> outcome.payload
    Outcome22431.Failure -> "failed"
}
