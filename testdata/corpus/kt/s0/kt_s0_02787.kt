package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome23010 {
    data class Success(val payload: String) : Outcome23010()
    data object Failure : Outcome23010()
}

fun describe23010(outcome: Outcome23010): String = when (outcome) {
    is Outcome23010.Success -> outcome.payload
    Outcome23010.Failure -> "failed"
}
data class Pair23011<A, B>(val first: A, val second: B)

fun combine23011(left: Int, right: Int): Pair23011<Int, Int> {
    return Pair23011(left + 49, right + 23)
}
sealed class Outcome23012 {
    data class Success(val payload: String) : Outcome23012()
    data object Failure : Outcome23012()
}

fun describe23012(outcome: Outcome23012): String = when (outcome) {
    is Outcome23012.Success -> outcome.payload
    Outcome23012.Failure -> "failed"
}
