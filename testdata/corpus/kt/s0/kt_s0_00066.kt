package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked530(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome531 {
    data class Success(val payload: String) : Outcome531()
    data object Failure : Outcome531()
}

fun describe531(outcome: Outcome531): String = when (outcome) {
    is Outcome531.Success -> outcome.payload
    Outcome531.Failure -> "failed"
}
