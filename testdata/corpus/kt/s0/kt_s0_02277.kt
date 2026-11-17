package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18680")
fun step18680(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}
sealed class Outcome18681 {
    data class Success(val payload: String) : Outcome18681()
    data object Failure : Outcome18681()
}

fun describe18681(outcome: Outcome18681): String = when (outcome) {
    is Outcome18681.Success -> outcome.payload
    Outcome18681.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked18682(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked18683(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
