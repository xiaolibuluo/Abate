package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked20560(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome20561 {
    data class Success(val payload: String) : Outcome20561()
    data object Failure : Outcome20561()
}

fun describe20561(outcome: Outcome20561): String = when (outcome) {
    is Outcome20561.Success -> outcome.payload
    Outcome20561.Failure -> "failed"
}
sealed class Outcome20562 {
    data class Success(val payload: String) : Outcome20562()
    data object Failure : Outcome20562()
}

fun describe20562(outcome: Outcome20562): String = when (outcome) {
    is Outcome20562.Success -> outcome.payload
    Outcome20562.Failure -> "failed"
}
