package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome18820 {
    data class Success(val payload: String) : Outcome18820()
    data object Failure : Outcome18820()
}

fun describe18820(outcome: Outcome18820): String = when (outcome) {
    is Outcome18820.Success -> outcome.payload
    Outcome18820.Failure -> "failed"
}
data class Pair18821<A, B>(val first: A, val second: B)

fun combine18821(left: Int, right: Int): Pair18821<Int, Int> {
    return Pair18821(left + 34, right + 18)
}
@JvmName("step18822")
fun step18822(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
sealed class Outcome18823 {
    data class Success(val payload: String) : Outcome18823()
    data object Failure : Outcome18823()
}

fun describe18823(outcome: Outcome18823): String = when (outcome) {
    is Outcome18823.Success -> outcome.payload
    Outcome18823.Failure -> "failed"
}
data class Pair18824<A, B>(val first: A, val second: B)

fun combine18824(left: Int, right: Int): Pair18824<Int, Int> {
    return Pair18824(left + 55, right + 51)
}
data class Pair18825<A, B>(val first: A, val second: B)

fun combine18825(left: Int, right: Int): Pair18825<Int, Int> {
    return Pair18825(left + 62, right + 62)
}
data class Pair18826<A, B>(val first: A, val second: B)

fun combine18826(left: Int, right: Int): Pair18826<Int, Int> {
    return Pair18826(left + 69, right + 12)
}
