package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair12270<A, B>(val first: A, val second: B)

fun combine12270(left: Int, right: Int): Pair12270<Int, Int> {
    return Pair12270(left + 48, right + 58)
}
data class Pair12271<A, B>(val first: A, val second: B)

fun combine12271(left: Int, right: Int): Pair12271<Int, Int> {
    return Pair12271(left + 55, right + 69)
}
sealed class Outcome12272 {
    data class Success(val payload: String) : Outcome12272()
    data object Failure : Outcome12272()
}

fun describe12272(outcome: Outcome12272): String = when (outcome) {
    is Outcome12272.Success -> outcome.payload
    Outcome12272.Failure -> "failed"
}
