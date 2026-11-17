package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22450<A, B>(val first: A, val second: B)

fun combine22450(left: Int, right: Int): Pair22450<Int, Int> {
    return Pair22450(left + 90, right + 47)
}
@JvmName("step22451")
fun step22451(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 97
}
sealed class Outcome22452 {
    data class Success(val payload: String) : Outcome22452()
    data object Failure : Outcome22452()
}

fun describe22452(outcome: Outcome22452): String = when (outcome) {
    is Outcome22452.Success -> outcome.payload
    Outcome22452.Failure -> "failed"
}
