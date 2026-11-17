package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome26460 {
    data class Success(val payload: String) : Outcome26460()
    data object Failure : Outcome26460()
}

fun describe26460(outcome: Outcome26460): String = when (outcome) {
    is Outcome26460.Success -> outcome.payload
    Outcome26460.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked26461(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked26462(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair26463<A, B>(val first: A, val second: B)

fun combine26463(left: Int, right: Int): Pair26463<Int, Int> {
    return Pair26463(left + 86, right + 20)
}
