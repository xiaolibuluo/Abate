package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked5320(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step5321")
fun step5321(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
fun <T> List<T>.secondOrNull5322(): T? = if (size > 1) this[1] else null

fun render5322(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step5323")
fun step5323(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 55
}
fun <T> List<T>.secondOrNull5324(): T? = if (size > 1) this[1] else null

fun render5324(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome5325 {
    data class Success(val payload: String) : Outcome5325()
    data object Failure : Outcome5325()
}

fun describe5325(outcome: Outcome5325): String = when (outcome) {
    is Outcome5325.Success -> outcome.payload
    Outcome5325.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked5326(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome5327 {
    data class Success(val payload: String) : Outcome5327()
    data object Failure : Outcome5327()
}

fun describe5327(outcome: Outcome5327): String = when (outcome) {
    is Outcome5327.Success -> outcome.payload
    Outcome5327.Failure -> "failed"
}
data class Pair5328<A, B>(val first: A, val second: B)

fun combine5328(left: Int, right: Int): Pair5328<Int, Int> {
    return Pair5328(left + 90, right + 39)
}
