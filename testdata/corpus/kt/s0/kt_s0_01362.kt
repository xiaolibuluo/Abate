package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked11220(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome11221 {
    data class Success(val payload: String) : Outcome11221()
    data object Failure : Outcome11221()
}

fun describe11221(outcome: Outcome11221): String = when (outcome) {
    is Outcome11221.Success -> outcome.payload
    Outcome11221.Failure -> "failed"
}
