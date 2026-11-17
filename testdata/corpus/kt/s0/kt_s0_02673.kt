package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked22080(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked22081(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step22082")
fun step22082(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
sealed class Outcome22083 {
    data class Success(val payload: String) : Outcome22083()
    data object Failure : Outcome22083()
}

fun describe22083(outcome: Outcome22083): String = when (outcome) {
    is Outcome22083.Success -> outcome.payload
    Outcome22083.Failure -> "failed"
}
