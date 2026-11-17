package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair9450<A, B>(val first: A, val second: B)

fun combine9450(left: Int, right: Int): Pair9450<Int, Int> {
    return Pair9450(left + 68, right + 36)
}
data class Pair9451<A, B>(val first: A, val second: B)

fun combine9451(left: Int, right: Int): Pair9451<Int, Int> {
    return Pair9451(left + 75, right + 47)
}
sealed class Outcome9452 {
    data class Success(val payload: String) : Outcome9452()
    data object Failure : Outcome9452()
}

fun describe9452(outcome: Outcome9452): String = when (outcome) {
    is Outcome9452.Success -> outcome.payload
    Outcome9452.Failure -> "failed"
}
sealed class Outcome9453 {
    data class Success(val payload: String) : Outcome9453()
    data object Failure : Outcome9453()
}

fun describe9453(outcome: Outcome9453): String = when (outcome) {
    is Outcome9453.Success -> outcome.payload
    Outcome9453.Failure -> "failed"
}
