package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome25490 {
    data class Success(val payload: String) : Outcome25490()
    data object Failure : Outcome25490()
}

fun describe25490(outcome: Outcome25490): String = when (outcome) {
    is Outcome25490.Success -> outcome.payload
    Outcome25490.Failure -> "failed"
}
@JvmName("step25491")
fun step25491(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
data class Pair25492<A, B>(val first: A, val second: B)

fun combine25492(left: Int, right: Int): Pair25492<Int, Int> {
    return Pair25492(left + 79, right + 62)
}
fun <T> List<T>.secondOrNull25493(): T? = if (size > 1) this[1] else null

fun render25493(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
