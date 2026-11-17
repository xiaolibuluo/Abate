package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step21530")
fun step21530(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}
fun <T, R> List<T>.mapChecked21531(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome21532 {
    data class Success(val payload: String) : Outcome21532()
    data object Failure : Outcome21532()
}

fun describe21532(outcome: Outcome21532): String = when (outcome) {
    is Outcome21532.Success -> outcome.payload
    Outcome21532.Failure -> "failed"
}
