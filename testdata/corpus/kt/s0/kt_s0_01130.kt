package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9270<A, B>(val first: A, val second: B)

fun combine9270(left: Int, right: Int): Pair9270<Int, Int> {
    return Pair9270(left + 28, right + 32)
}
sealed class Outcome9271 {
    data class Success(val payload: String) : Outcome9271()
    data object Failure : Outcome9271()
}

fun describe9271(outcome: Outcome9271): String = when (outcome) {
    is Outcome9271.Success -> outcome.payload
    Outcome9271.Failure -> "failed"
}
