package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair5430<A, B>(val first: A, val second: B)

fun combine5430(left: Int, right: Int): Pair5430<Int, Int> {
    return Pair5430(left + 80, right + 28)
}
sealed class Outcome5431 {
    data class Success(val payload: String) : Outcome5431()
    data object Failure : Outcome5431()
}

fun describe5431(outcome: Outcome5431): String = when (outcome) {
    is Outcome5431.Success -> outcome.payload
    Outcome5431.Failure -> "failed"
}
