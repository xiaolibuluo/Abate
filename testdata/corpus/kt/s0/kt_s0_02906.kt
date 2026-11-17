package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23980 {
    data class Success(val payload: String) : Outcome23980()
    data object Failure : Outcome23980()
}

fun describe23980(outcome: Outcome23980): String = when (outcome) {
    is Outcome23980.Success -> outcome.payload
    Outcome23980.Failure -> "failed"
}
data class Pair23981<A, B>(val first: A, val second: B)

fun combine23981(left: Int, right: Int): Pair23981<Int, Int> {
    return Pair23981(left + 49, right + 31)
}
sealed class Outcome23982 {
    data class Success(val payload: String) : Outcome23982()
    data object Failure : Outcome23982()
}

fun describe23982(outcome: Outcome23982): String = when (outcome) {
    is Outcome23982.Success -> outcome.payload
    Outcome23982.Failure -> "failed"
}
