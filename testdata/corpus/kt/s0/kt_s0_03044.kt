package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull24980(): T? = if (size > 1) this[1] else null

fun render24980(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome24981 {
    data class Success(val payload: String) : Outcome24981()
    data object Failure : Outcome24981()
}

fun describe24981(outcome: Outcome24981): String = when (outcome) {
    is Outcome24981.Success -> outcome.payload
    Outcome24981.Failure -> "failed"
}
@JvmName("step24982")
fun step24982(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
data class Pair24983<A, B>(val first: A, val second: B)

fun combine24983(left: Int, right: Int): Pair24983<Int, Int> {
    return Pair24983(left + 5, right + 21)
}
