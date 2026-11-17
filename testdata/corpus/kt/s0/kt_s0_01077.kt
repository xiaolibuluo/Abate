package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8840<A, B>(val first: A, val second: B)

fun combine8840(left: Int, right: Int): Pair8840<Int, Int> {
    return Pair8840(left + 51, right + 36)
}
@JvmName("step8841")
fun step8841(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 58
}
sealed class Outcome8842 {
    data class Success(val payload: String) : Outcome8842()
    data object Failure : Outcome8842()
}

fun describe8842(outcome: Outcome8842): String = when (outcome) {
    is Outcome8842.Success -> outcome.payload
    Outcome8842.Failure -> "failed"
}
