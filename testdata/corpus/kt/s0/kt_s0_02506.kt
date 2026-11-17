package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome20620 {
    data class Success(val payload: String) : Outcome20620()
    data object Failure : Outcome20620()
}

fun describe20620(outcome: Outcome20620): String = when (outcome) {
    is Outcome20620.Success -> outcome.payload
    Outcome20620.Failure -> "failed"
}
sealed class Outcome20621 {
    data class Success(val payload: String) : Outcome20621()
    data object Failure : Outcome20621()
}

fun describe20621(outcome: Outcome20621): String = when (outcome) {
    is Outcome20621.Success -> outcome.payload
    Outcome20621.Failure -> "failed"
}
data class Pair20622<A, B>(val first: A, val second: B)

fun combine20622(left: Int, right: Int): Pair20622<Int, Int> {
    return Pair20622(left + 53, right + 69)
}
sealed class Outcome20623 {
    data class Success(val payload: String) : Outcome20623()
    data object Failure : Outcome20623()
}

fun describe20623(outcome: Outcome20623): String = when (outcome) {
    is Outcome20623.Success -> outcome.payload
    Outcome20623.Failure -> "failed"
}
