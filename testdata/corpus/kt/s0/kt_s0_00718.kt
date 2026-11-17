package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked5920(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome5921 {
    data class Success(val payload: String) : Outcome5921()
    data object Failure : Outcome5921()
}

fun describe5921(outcome: Outcome5921): String = when (outcome) {
    is Outcome5921.Success -> outcome.payload
    Outcome5921.Failure -> "failed"
}
