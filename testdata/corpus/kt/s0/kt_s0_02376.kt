package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome19530 {
    data class Success(val payload: String) : Outcome19530()
    data object Failure : Outcome19530()
}

fun describe19530(outcome: Outcome19530): String = when (outcome) {
    is Outcome19530.Success -> outcome.payload
    Outcome19530.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked19531(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
