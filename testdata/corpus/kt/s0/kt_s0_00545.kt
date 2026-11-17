package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked4420(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome4421 {
    data class Success(val payload: String) : Outcome4421()
    data object Failure : Outcome4421()
}

fun describe4421(outcome: Outcome4421): String = when (outcome) {
    is Outcome4421.Success -> outcome.payload
    Outcome4421.Failure -> "failed"
}
