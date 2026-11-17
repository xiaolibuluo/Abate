package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked23470(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step23471")
fun step23471(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 65
}
sealed class Outcome23472 {
    data class Success(val payload: String) : Outcome23472()
    data object Failure : Outcome23472()
}

fun describe23472(outcome: Outcome23472): String = when (outcome) {
    is Outcome23472.Success -> outcome.payload
    Outcome23472.Failure -> "failed"
}
