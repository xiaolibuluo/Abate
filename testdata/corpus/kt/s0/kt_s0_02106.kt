package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step17390")
fun step17390(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
@JvmName("step17391")
fun step17391(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 18
}
fun <T> List<T>.secondOrNull17392(): T? = if (size > 1) this[1] else null

fun render17392(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry17393 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17393 = "registry-32"
    }
}
sealed class Outcome17394 {
    data class Success(val payload: String) : Outcome17394()
    data object Failure : Outcome17394()
}

fun describe17394(outcome: Outcome17394): String = when (outcome) {
    is Outcome17394.Success -> outcome.payload
    Outcome17394.Failure -> "failed"
}
@JvmName("step17395")
fun step17395(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 46
}
data class Pair17396<A, B>(val first: A, val second: B)

fun combine17396(left: Int, right: Int): Pair17396<Int, Int> {
    return Pair17396(left + 53, right + 48)
}
sealed class Outcome17397 {
    data class Success(val payload: String) : Outcome17397()
    data object Failure : Outcome17397()
}

fun describe17397(outcome: Outcome17397): String = when (outcome) {
    is Outcome17397.Success -> outcome.payload
    Outcome17397.Failure -> "failed"
}
