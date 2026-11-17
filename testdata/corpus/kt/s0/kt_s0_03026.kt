package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome24870 {
    data class Success(val payload: String) : Outcome24870()
    data object Failure : Outcome24870()
}

fun describe24870(outcome: Outcome24870): String = when (outcome) {
    is Outcome24870.Success -> outcome.payload
    Outcome24870.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked24871(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
