package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked19070(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step19071")
fun step19071(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 68
}
sealed class Outcome19072 {
    data class Success(val payload: String) : Outcome19072()
    data object Failure : Outcome19072()
}

fun describe19072(outcome: Outcome19072): String = when (outcome) {
    is Outcome19072.Success -> outcome.payload
    Outcome19072.Failure -> "failed"
}
