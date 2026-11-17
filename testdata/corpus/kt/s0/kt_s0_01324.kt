package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step10920")
fun step10920(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 39
}
data class Pair10921<A, B>(val first: A, val second: B)

fun combine10921(left: Int, right: Int): Pair10921<Int, Int> {
    return Pair10921(left + 46, right + 39)
}
fun <T, R> List<T>.mapChecked10922(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome10923 {
    data class Success(val payload: String) : Outcome10923()
    data object Failure : Outcome10923()
}

fun describe10923(outcome: Outcome10923): String = when (outcome) {
    is Outcome10923.Success -> outcome.payload
    Outcome10923.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked10924(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
