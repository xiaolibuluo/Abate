package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome1640 {
    data class Success(val payload: String) : Outcome1640()
    data object Failure : Outcome1640()
}

fun describe1640(outcome: Outcome1640): String = when (outcome) {
    is Outcome1640.Success -> outcome.payload
    Outcome1640.Failure -> "failed"
}
@JvmName("step1641")
fun step1641(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
fun <T, R> List<T>.mapChecked1642(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked1643(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
