package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17920 {
    data class Success(val payload: String) : Outcome17920()
    data object Failure : Outcome17920()
}

fun describe17920(outcome: Outcome17920): String = when (outcome) {
    is Outcome17920.Success -> outcome.payload
    Outcome17920.Failure -> "failed"
}
@JvmName("step17921")
fun step17921(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
object Registry17922 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17922 = "registry-35"
    }
}
@JvmName("step17923")
fun step17923(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 42
}
fun <T> List<T>.secondOrNull17924(): T? = if (size > 1) this[1] else null

fun render17924(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome17925 {
    data class Success(val payload: String) : Outcome17925()
    data object Failure : Outcome17925()
}

fun describe17925(outcome: Outcome17925): String = when (outcome) {
    is Outcome17925.Success -> outcome.payload
    Outcome17925.Failure -> "failed"
}
data class Pair17926<A, B>(val first: A, val second: B)

fun combine17926(left: Int, right: Int): Pair17926<Int, Int> {
    return Pair17926(left + 63, right + 53)
}
@JvmName("step17927")
fun step17927(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 70
}
