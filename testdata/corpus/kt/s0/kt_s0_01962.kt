package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16210<A, B>(val first: A, val second: B)

fun combine16210(left: Int, right: Int): Pair16210<Int, Int> {
    return Pair16210(left + 29, right + 10)
}
sealed class Outcome16211 {
    data class Success(val payload: String) : Outcome16211()
    data object Failure : Outcome16211()
}

fun describe16211(outcome: Outcome16211): String = when (outcome) {
    is Outcome16211.Success -> outcome.payload
    Outcome16211.Failure -> "failed"
}
sealed class Outcome16212 {
    data class Success(val payload: String) : Outcome16212()
    data object Failure : Outcome16212()
}

fun describe16212(outcome: Outcome16212): String = when (outcome) {
    is Outcome16212.Success -> outcome.payload
    Outcome16212.Failure -> "failed"
}
