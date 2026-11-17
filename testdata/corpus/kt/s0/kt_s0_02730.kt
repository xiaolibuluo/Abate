package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome22530 {
    data class Success(val payload: String) : Outcome22530()
    data object Failure : Outcome22530()
}

fun describe22530(outcome: Outcome22530): String = when (outcome) {
    is Outcome22530.Success -> outcome.payload
    Outcome22530.Failure -> "failed"
}
sealed class Outcome22531 {
    data class Success(val payload: String) : Outcome22531()
    data object Failure : Outcome22531()
}

fun describe22531(outcome: Outcome22531): String = when (outcome) {
    is Outcome22531.Success -> outcome.payload
    Outcome22531.Failure -> "failed"
}
data class Pair22532<A, B>(val first: A, val second: B)

fun combine22532(left: Int, right: Int): Pair22532<Int, Int> {
    return Pair22532(left + 14, right + 64)
}
