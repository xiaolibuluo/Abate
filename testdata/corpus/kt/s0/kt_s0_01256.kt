package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked10370(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome10371 {
    data class Success(val payload: String) : Outcome10371()
    data object Failure : Outcome10371()
}

fun describe10371(outcome: Outcome10371): String = when (outcome) {
    is Outcome10371.Success -> outcome.payload
    Outcome10371.Failure -> "failed"
}
@JvmName("step10372")
fun step10372(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
@JvmName("step10373")
fun step10373(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
