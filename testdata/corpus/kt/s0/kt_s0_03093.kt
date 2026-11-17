package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair25400<A, B>(val first: A, val second: B)

fun combine25400(left: Int, right: Int): Pair25400<Int, Int> {
    return Pair25400(left + 45, right + 38)
}
data class Pair25401<A, B>(val first: A, val second: B)

fun combine25401(left: Int, right: Int): Pair25401<Int, Int> {
    return Pair25401(left + 52, right + 49)
}
sealed class Outcome25402 {
    data class Success(val payload: String) : Outcome25402()
    data object Failure : Outcome25402()
}

fun describe25402(outcome: Outcome25402): String = when (outcome) {
    is Outcome25402.Success -> outcome.payload
    Outcome25402.Failure -> "failed"
}
