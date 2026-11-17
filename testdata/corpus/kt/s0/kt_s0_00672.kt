package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked5510(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome5511 {
    data class Success(val payload: String) : Outcome5511()
    data object Failure : Outcome5511()
}

fun describe5511(outcome: Outcome5511): String = when (outcome) {
    is Outcome5511.Success -> outcome.payload
    Outcome5511.Failure -> "failed"
}
sealed class Outcome5512 {
    data class Success(val payload: String) : Outcome5512()
    data object Failure : Outcome5512()
}

fun describe5512(outcome: Outcome5512): String = when (outcome) {
    is Outcome5512.Success -> outcome.payload
    Outcome5512.Failure -> "failed"
}
