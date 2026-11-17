package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome14740 {
    data class Success(val payload: String) : Outcome14740()
    data object Failure : Outcome14740()
}

fun describe14740(outcome: Outcome14740): String = when (outcome) {
    is Outcome14740.Success -> outcome.payload
    Outcome14740.Failure -> "failed"
}
data class Pair14741<A, B>(val first: A, val second: B)

fun combine14741(left: Int, right: Int): Pair14741<Int, Int> {
    return Pair14741(left + 65, right + 29)
}
sealed class Outcome14742 {
    data class Success(val payload: String) : Outcome14742()
    data object Failure : Outcome14742()
}

fun describe14742(outcome: Outcome14742): String = when (outcome) {
    is Outcome14742.Success -> outcome.payload
    Outcome14742.Failure -> "failed"
}
