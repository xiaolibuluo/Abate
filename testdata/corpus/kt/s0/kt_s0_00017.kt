package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome140 {
    data class Success(val payload: String) : Outcome140()
    data object Failure : Outcome140()
}

fun describe140(outcome: Outcome140): String = when (outcome) {
    is Outcome140.Success -> outcome.payload
    Outcome140.Failure -> "failed"
}
data class Pair141<A, B>(val first: A, val second: B)

fun combine141(left: Int, right: Int): Pair141<Int, Int> {
    return Pair141(left + 97, right + 57)
}
fun <T, R> List<T>.mapChecked142(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome143 {
    data class Success(val payload: String) : Outcome143()
    data object Failure : Outcome143()
}

fun describe143(outcome: Outcome143): String = when (outcome) {
    is Outcome143.Success -> outcome.payload
    Outcome143.Failure -> "failed"
}
