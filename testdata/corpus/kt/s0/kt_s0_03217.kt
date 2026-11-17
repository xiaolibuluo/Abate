package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome26440 {
    data class Success(val payload: String) : Outcome26440()
    data object Failure : Outcome26440()
}

fun describe26440(outcome: Outcome26440): String = when (outcome) {
    is Outcome26440.Success -> outcome.payload
    Outcome26440.Failure -> "failed"
}
@JvmName("step26441")
fun step26441(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 46
}
fun <T, R> List<T>.mapChecked26442(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair26443<A, B>(val first: A, val second: B)

fun combine26443(left: Int, right: Int): Pair26443<Int, Int> {
    return Pair26443(left + 60, right + 67)
}
