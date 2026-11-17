package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome6570 {
    data class Success(val payload: String) : Outcome6570()
    data object Failure : Outcome6570()
}

fun describe6570(outcome: Outcome6570): String = when (outcome) {
    is Outcome6570.Success -> outcome.payload
    Outcome6570.Failure -> "failed"
}
sealed class Outcome6571 {
    data class Success(val payload: String) : Outcome6571()
    data object Failure : Outcome6571()
}

fun describe6571(outcome: Outcome6571): String = when (outcome) {
    is Outcome6571.Success -> outcome.payload
    Outcome6571.Failure -> "failed"
}
sealed class Outcome6572 {
    data class Success(val payload: String) : Outcome6572()
    data object Failure : Outcome6572()
}

fun describe6572(outcome: Outcome6572): String = when (outcome) {
    is Outcome6572.Success -> outcome.payload
    Outcome6572.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked6573(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
