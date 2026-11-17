package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked25360(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome25361 {
    data class Success(val payload: String) : Outcome25361()
    data object Failure : Outcome25361()
}

fun describe25361(outcome: Outcome25361): String = when (outcome) {
    is Outcome25361.Success -> outcome.payload
    Outcome25361.Failure -> "failed"
}
