package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step10340")
fun step10340(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 61
}
sealed class Outcome10341 {
    data class Success(val payload: String) : Outcome10341()
    data object Failure : Outcome10341()
}

fun describe10341(outcome: Outcome10341): String = when (outcome) {
    is Outcome10341.Success -> outcome.payload
    Outcome10341.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked10342(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
