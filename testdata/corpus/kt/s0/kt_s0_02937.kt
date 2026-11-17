package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step24220")
fun step24220(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
fun <T, R> List<T>.mapChecked24221(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair24222<A, B>(val first: A, val second: B)

fun combine24222(left: Int, right: Int): Pair24222<Int, Int> {
    return Pair24222(left + 77, right + 27)
}
data class Pair24223<A, B>(val first: A, val second: B)

fun combine24223(left: Int, right: Int): Pair24223<Int, Int> {
    return Pair24223(left + 84, right + 38)
}
fun <T> List<T>.secondOrNull24224(): T? = if (size > 1) this[1] else null

fun render24224(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome24225 {
    data class Success(val payload: String) : Outcome24225()
    data object Failure : Outcome24225()
}

fun describe24225(outcome: Outcome24225): String = when (outcome) {
    is Outcome24225.Success -> outcome.payload
    Outcome24225.Failure -> "failed"
}
@JvmName("step24226")
fun step24226(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 8
}
fun <T> List<T>.secondOrNull24227(): T? = if (size > 1) this[1] else null

fun render24227(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
