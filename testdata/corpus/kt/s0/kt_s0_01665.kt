package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair13710<A, B>(val first: A, val second: B)

fun combine13710(left: Int, right: Int): Pair13710<Int, Int> {
    return Pair13710(left + 77, right + 29)
}
sealed class Outcome13711 {
    data class Success(val payload: String) : Outcome13711()
    data object Failure : Outcome13711()
}

fun describe13711(outcome: Outcome13711): String = when (outcome) {
    is Outcome13711.Success -> outcome.payload
    Outcome13711.Failure -> "failed"
}
