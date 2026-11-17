package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step18850")
fun step18850(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 66
}
sealed class Outcome18851 {
    data class Success(val payload: String) : Outcome18851()
    data object Failure : Outcome18851()
}

fun describe18851(outcome: Outcome18851): String = when (outcome) {
    is Outcome18851.Success -> outcome.payload
    Outcome18851.Failure -> "failed"
}
data class Pair18852<A, B>(val first: A, val second: B)

fun combine18852(left: Int, right: Int): Pair18852<Int, Int> {
    return Pair18852(left + 80, right + 50)
}
object Registry18853 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18853 = "registry-87"
    }
}
fun <T> List<T>.secondOrNull18854(): T? = if (size > 1) this[1] else null

fun render18854(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked18855(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step18856")
fun step18856(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
fun <T, R> List<T>.mapChecked18857(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
