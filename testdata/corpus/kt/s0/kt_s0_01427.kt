package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome11750 {
    data class Success(val payload: String) : Outcome11750()
    data object Failure : Outcome11750()
}

fun describe11750(outcome: Outcome11750): String = when (outcome) {
    is Outcome11750.Success -> outcome.payload
    Outcome11750.Failure -> "failed"
}
data class Pair11751<A, B>(val first: A, val second: B)

fun combine11751(left: Int, right: Int): Pair11751<Int, Int> {
    return Pair11751(left + 58, right + 10)
}
