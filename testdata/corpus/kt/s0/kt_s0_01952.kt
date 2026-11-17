package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked16110(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome16111 {
    data class Success(val payload: String) : Outcome16111()
    data object Failure : Outcome16111()
}

fun describe16111(outcome: Outcome16111): String = when (outcome) {
    is Outcome16111.Success -> outcome.payload
    Outcome16111.Failure -> "failed"
}
