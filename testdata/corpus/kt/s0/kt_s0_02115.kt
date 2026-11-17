package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked17480(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair17481<A, B>(val first: A, val second: B)

fun combine17481(left: Int, right: Int): Pair17481<Int, Int> {
    return Pair17481(left + 38, right + 56)
}
data class Pair17482<A, B>(val first: A, val second: B)

fun combine17482(left: Int, right: Int): Pair17482<Int, Int> {
    return Pair17482(left + 45, right + 67)
}
sealed class Outcome17483 {
    data class Success(val payload: String) : Outcome17483()
    data object Failure : Outcome17483()
}

fun describe17483(outcome: Outcome17483): String = when (outcome) {
    is Outcome17483.Success -> outcome.payload
    Outcome17483.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17484(): T? = if (size > 1) this[1] else null

fun render17484(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked17485(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step17486")
fun step17486(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 73
}
object Registry17487 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17487 = "registry-80"
    }
}
object Registry17488 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17488 = "registry-87"
    }
}
object Registry17489 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17489 = "registry-94"
    }
}
