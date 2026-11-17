package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair23780<A, B>(val first: A, val second: B)

fun combine23780(left: Int, right: Int): Pair23780<Int, Int> {
    return Pair23780(left + 73, right + 63)
}
sealed class Outcome23781 {
    data class Success(val payload: String) : Outcome23781()
    data object Failure : Outcome23781()
}

fun describe23781(outcome: Outcome23781): String = when (outcome) {
    is Outcome23781.Success -> outcome.payload
    Outcome23781.Failure -> "failed"
}
fun <T> List<T>.secondOrNull23782(): T? = if (size > 1) this[1] else null

fun render23782(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked23783(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome23784 {
    data class Success(val payload: String) : Outcome23784()
    data object Failure : Outcome23784()
}

fun describe23784(outcome: Outcome23784): String = when (outcome) {
    is Outcome23784.Success -> outcome.payload
    Outcome23784.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked23785(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step23786")
fun step23786(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
data class Pair23787<A, B>(val first: A, val second: B)

fun combine23787(left: Int, right: Int): Pair23787<Int, Int> {
    return Pair23787(left + 25, right + 18)
}
fun <T> List<T>.secondOrNull23788(): T? = if (size > 1) this[1] else null

fun render23788(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
