package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry24680 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24680 = "registry-79"
    }
}
fun <T> List<T>.secondOrNull24681(): T? = if (size > 1) this[1] else null

fun render24681(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step24682")
fun step24682(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
fun <T> List<T>.secondOrNull24683(): T? = if (size > 1) this[1] else null

fun render24683(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked24684(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair24685<A, B>(val first: A, val second: B)

fun combine24685(left: Int, right: Int): Pair24685<Int, Int> {
    return Pair24685(left + 17, right + 16)
}
sealed class Outcome24686 {
    data class Success(val payload: String) : Outcome24686()
    data object Failure : Outcome24686()
}

fun describe24686(outcome: Outcome24686): String = when (outcome) {
    is Outcome24686.Success -> outcome.payload
    Outcome24686.Failure -> "failed"
}
