package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step8510")
fun step8510(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
@JvmName("step8511")
fun step8511(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
sealed class Outcome8512 {
    data class Success(val payload: String) : Outcome8512()
    data object Failure : Outcome8512()
}

fun describe8512(outcome: Outcome8512): String = when (outcome) {
    is Outcome8512.Success -> outcome.payload
    Outcome8512.Failure -> "failed"
}
data class Pair8513<A, B>(val first: A, val second: B)

fun combine8513(left: Int, right: Int): Pair8513<Int, Int> {
    return Pair8513(left + 31, right + 21)
}
data class Pair8514<A, B>(val first: A, val second: B)

fun combine8514(left: Int, right: Int): Pair8514<Int, Int> {
    return Pair8514(left + 38, right + 32)
}
data class Pair8515<A, B>(val first: A, val second: B)

fun combine8515(left: Int, right: Int): Pair8515<Int, Int> {
    return Pair8515(left + 45, right + 43)
}
