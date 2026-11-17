package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked12770(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step12771")
fun step12771(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
data class Pair12772<A, B>(val first: A, val second: B)

fun combine12772(left: Int, right: Int): Pair12772<Int, Int> {
    return Pair12772(left + 33, right + 64)
}
sealed class Outcome12773 {
    data class Success(val payload: String) : Outcome12773()
    data object Failure : Outcome12773()
}

fun describe12773(outcome: Outcome12773): String = when (outcome) {
    is Outcome12773.Success -> outcome.payload
    Outcome12773.Failure -> "failed"
}
