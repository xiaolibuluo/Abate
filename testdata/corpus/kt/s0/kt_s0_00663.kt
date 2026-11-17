package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome5440 {
    data class Success(val payload: String) : Outcome5440()
    data object Failure : Outcome5440()
}

fun describe5440(outcome: Outcome5440): String = when (outcome) {
    is Outcome5440.Success -> outcome.payload
    Outcome5440.Failure -> "failed"
}
data class Pair5441<A, B>(val first: A, val second: B)

fun combine5441(left: Int, right: Int): Pair5441<Int, Int> {
    return Pair5441(left + 100, right + 46)
}
