package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry7660 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7660 = "registry-69"
    }
}
@JvmName("step7661")
fun step7661(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
data class Pair7662<A, B>(val first: A, val second: B)

fun combine7662(left: Int, right: Int): Pair7662<Int, Int> {
    return Pair7662(left + 83, right + 25)
}
sealed class Outcome7663 {
    data class Success(val payload: String) : Outcome7663()
    data object Failure : Outcome7663()
}

fun describe7663(outcome: Outcome7663): String = when (outcome) {
    is Outcome7663.Success -> outcome.payload
    Outcome7663.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked7664(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull7665(): T? = if (size > 1) this[1] else null

fun render7665(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair7666<A, B>(val first: A, val second: B)

fun combine7666(left: Int, right: Int): Pair7666<Int, Int> {
    return Pair7666(left + 14, right + 69)
}
@JvmName("step7667")
fun step7667(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
