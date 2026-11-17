package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20650<A, B>(val first: A, val second: B)

fun combine20650(left: Int, right: Int): Pair20650<Int, Int> {
    return Pair20650(left + 78, right + 68)
}
data class Pair20651<A, B>(val first: A, val second: B)

fun combine20651(left: Int, right: Int): Pair20651<Int, Int> {
    return Pair20651(left + 85, right + 18)
}
@JvmName("step20652")
fun step20652(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
sealed class Outcome20653 {
    data class Success(val payload: String) : Outcome20653()
    data object Failure : Outcome20653()
}

fun describe20653(outcome: Outcome20653): String = when (outcome) {
    is Outcome20653.Success -> outcome.payload
    Outcome20653.Failure -> "failed"
}
data class Pair20654<A, B>(val first: A, val second: B)

fun combine20654(left: Int, right: Int): Pair20654<Int, Int> {
    return Pair20654(left + 9, right + 51)
}
