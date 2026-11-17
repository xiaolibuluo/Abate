package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23060<A, B>(val first: A, val second: B)

fun combine23060(left: Int, right: Int): Pair23060<Int, Int> {
    return Pair23060(left + 10, right + 47)
}
data class Pair23061<A, B>(val first: A, val second: B)

fun combine23061(left: Int, right: Int): Pair23061<Int, Int> {
    return Pair23061(left + 17, right + 58)
}
data class Pair23062<A, B>(val first: A, val second: B)

fun combine23062(left: Int, right: Int): Pair23062<Int, Int> {
    return Pair23062(left + 24, right + 69)
}
sealed class Outcome23063 {
    data class Success(val payload: String) : Outcome23063()
    data object Failure : Outcome23063()
}

fun describe23063(outcome: Outcome23063): String = when (outcome) {
    is Outcome23063.Success -> outcome.payload
    Outcome23063.Failure -> "failed"
}
sealed class Outcome23064 {
    data class Success(val payload: String) : Outcome23064()
    data object Failure : Outcome23064()
}

fun describe23064(outcome: Outcome23064): String = when (outcome) {
    is Outcome23064.Success -> outcome.payload
    Outcome23064.Failure -> "failed"
}
