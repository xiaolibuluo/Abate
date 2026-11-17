package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7970 {
    data class Success(val payload: String) : Outcome7970()
    data object Failure : Outcome7970()
}

fun describe7970(outcome: Outcome7970): String = when (outcome) {
    is Outcome7970.Success -> outcome.payload
    Outcome7970.Failure -> "failed"
}
@JvmName("step7971")
fun step7971(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
@JvmName("step7972")
fun step7972(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 98
}
data class Pair7973<A, B>(val first: A, val second: B)

fun combine7973(left: Int, right: Int): Pair7973<Int, Int> {
    return Pair7973(left + 8, right + 9)
}
fun <T, R> List<T>.mapChecked7974(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
