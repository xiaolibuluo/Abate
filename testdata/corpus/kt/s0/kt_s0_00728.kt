package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step5990")
fun step5990(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 32
}
fun <T, R> List<T>.mapChecked5991(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked5992(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome5993 {
    data class Success(val payload: String) : Outcome5993()
    data object Failure : Outcome5993()
}

fun describe5993(outcome: Outcome5993): String = when (outcome) {
    is Outcome5993.Success -> outcome.payload
    Outcome5993.Failure -> "failed"
}
