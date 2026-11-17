package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16600<A, B>(val first: A, val second: B)

fun combine16600(left: Int, right: Int): Pair16600<Int, Int> {
    return Pair16600(left + 51, right + 39)
}
sealed class Outcome16601 {
    data class Success(val payload: String) : Outcome16601()
    data object Failure : Outcome16601()
}

fun describe16601(outcome: Outcome16601): String = when (outcome) {
    is Outcome16601.Success -> outcome.payload
    Outcome16601.Failure -> "failed"
}
