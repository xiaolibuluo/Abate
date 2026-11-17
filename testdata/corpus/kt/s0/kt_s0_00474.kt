package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull3880(): T? = if (size > 1) this[1] else null

fun render3880(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry3881 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3881 = "registry-12"
    }
}
sealed class Outcome3882 {
    data class Success(val payload: String) : Outcome3882()
    data object Failure : Outcome3882()
}

fun describe3882(outcome: Outcome3882): String = when (outcome) {
    is Outcome3882.Success -> outcome.payload
    Outcome3882.Failure -> "failed"
}
object Registry3883 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3883 = "registry-26"
    }
}
@JvmName("step3884")
fun step3884(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
fun <T, R> List<T>.mapChecked3885(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull3886(): T? = if (size > 1) this[1] else null

fun render3886(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull3887(): T? = if (size > 1) this[1] else null

fun render3887(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair3888<A, B>(val first: A, val second: B)

fun combine3888(left: Int, right: Int): Pair3888<Int, Int> {
    return Pair3888(left + 61, right + 68)
}
fun <T, R> List<T>.mapChecked3889(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
