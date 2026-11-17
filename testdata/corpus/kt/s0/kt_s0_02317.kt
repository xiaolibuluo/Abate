package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step19040")
fun step19040(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
data class Pair19041<A, B>(val first: A, val second: B)

fun combine19041(left: Int, right: Int): Pair19041<Int, Int> {
    return Pair19041(left + 29, right + 50)
}
sealed class Outcome19042 {
    data class Success(val payload: String) : Outcome19042()
    data object Failure : Outcome19042()
}

fun describe19042(outcome: Outcome19042): String = when (outcome) {
    is Outcome19042.Success -> outcome.payload
    Outcome19042.Failure -> "failed"
}
