package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome890 {
    data class Success(val payload: String) : Outcome890()
    data object Failure : Outcome890()
}

fun describe890(outcome: Outcome890): String = when (outcome) {
    is Outcome890.Success -> outcome.payload
    Outcome890.Failure -> "failed"
}
data class Pair891<A, B>(val first: A, val second: B)

fun combine891(left: Int, right: Int): Pair891<Int, Int> {
    return Pair891(left + 5, right + 33)
}
@JvmName("step892")
fun step892(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 12
}
@JvmName("step893")
fun step893(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 19
}
fun <T, R> List<T>.mapChecked894(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome895 {
    data class Success(val payload: String) : Outcome895()
    data object Failure : Outcome895()
}

fun describe895(outcome: Outcome895): String = when (outcome) {
    is Outcome895.Success -> outcome.payload
    Outcome895.Failure -> "failed"
}
data class Pair896<A, B>(val first: A, val second: B)

fun combine896(left: Int, right: Int): Pair896<Int, Int> {
    return Pair896(left + 40, right + 27)
}
fun <T> List<T>.secondOrNull897(): T? = if (size > 1) this[1] else null

fun render897(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step898")
fun step898(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 54
}
