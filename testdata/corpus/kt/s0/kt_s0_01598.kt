package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked13200(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome13201 {
    data class Success(val payload: String) : Outcome13201()
    data object Failure : Outcome13201()
}

fun describe13201(outcome: Outcome13201): String = when (outcome) {
    is Outcome13201.Success -> outcome.payload
    Outcome13201.Failure -> "failed"
}
