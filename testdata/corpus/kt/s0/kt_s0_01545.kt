package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair12790<A, B>(val first: A, val second: B)

fun combine12790(left: Int, right: Int): Pair12790<Int, Int> {
    return Pair12790(left + 45, right + 56)
}
data class Pair12791<A, B>(val first: A, val second: B)

fun combine12791(left: Int, right: Int): Pair12791<Int, Int> {
    return Pair12791(left + 52, right + 67)
}
fun <T, R> List<T>.mapChecked12792(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12793 {
    data class Success(val payload: String) : Outcome12793()
    data object Failure : Outcome12793()
}

fun describe12793(outcome: Outcome12793): String = when (outcome) {
    is Outcome12793.Success -> outcome.payload
    Outcome12793.Failure -> "failed"
}
